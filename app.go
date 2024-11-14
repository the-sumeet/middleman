package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/elazarl/goproxy"
	"github.com/google/uuid"
)

type App struct {
	ctx            context.Context
	proxy          *goproxy.ProxyHttpServer
	proxyStartStop chan bool
	proxyRunning   bool
	webStartStop   chan bool
	webRunning     bool
	database       Database
	config         Config
	logger         *slog.Logger
}

func NewApp() *App {
	config := getConfig()
	logFile := getLogFilePath()
	logWriter, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	multiWriter := io.MultiWriter(os.Stdout, logWriter)
	if err != nil {
		panic(err)
	}

	proxy := goproxy.NewProxyHttpServer()
	proxy.Tr.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
		CipherSuites: []uint16{
			tls.TLS_AES_128_GCM_SHA256,
			tls.TLS_AES_256_GCM_SHA384,
			tls.TLS_CHACHA20_POLY1305_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
		},
	}
	app := &App{
		proxy:          proxy,
		database:       NewSqliteDatabase(config.RuleDbPath, config.RequestDbPath),
		proxyStartStop: make(chan bool),
		webStartStop:   make(chan bool),
		config:         config,
		logger:         slog.New(slog.NewJSONHandler(multiWriter, nil)),
	}

	return app
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	if !certKeyPresent() {
		genCert()
	}
}

func (a *App) responseToLog(r *http.Response, state *State, logId any) {

	// Response body
	responseBytes, err := io.ReadAll(r.Body)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Error reading response body: %s", err))
	}
	r.Body.Close() //  must close
	r.Body = io.NopCloser(bytes.NewBuffer(responseBytes))

	requestLog := HttpRequestLog{
		Timestamp:       time.Now(),
		Scheme:          r.Request.URL.Scheme,
		Method:          r.Request.Method,
		Host:            r.Request.Host,
		Path:            r.Request.URL.Path,
		RequestHeaders:  r.Request.Header,
		ResponseHeaders: r.Header,
		Status:          r.StatusCode,
		RequestBody:     state.RequestBody,
		ResponseBody:    string(responseBytes),
		// Rules info
		Cancelled:              state.IsCancelled,
		Redirected:             state.IsRedirected,
		RequestHeaderModified:  state.IsRequestHeaderModified,
		ResponseHeaderModified: state.IsResponseHeaderModified,
		RequestBodyModified:    state.IsRequestBodyModified,
		ResponseBodyModified:   state.IsResponseBodyModified,
		Delayed:                state.DelayedBy,
	}

	a.database.AddRequest(state.requestId, &requestLog)
}

func (a *App) getOnRequest() func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {

	return func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {

		// State for this request
		var requestId any
		requestId, err := uuid.NewUUID()
		if err != nil {
			a.logger.Error(fmt.Sprintf("Error generating UUID: %s", err))
			// If could not generate UUID, use timestamp.
			requestId = string(time.Now().UnixMilli())
		}
		ctx.UserData = &State{requestId: requestId}
		state := ctx.UserData.(*State)

		// Cancels
		cancels, err := a.database.GetManyRules(CANCEL)
		if err != nil {
			a.logger.Error(fmt.Sprintf("Error getting cancels: %s", err))
			log.Fatal("Error getting cancels: ", err)
		} else {
			for _, cancel := range cancels {
				if !cancel.Enabled {
					continue
				}
				// ToDo: Fix this
				if matches(cancel, r) {
					a.logger.Info("Cancel rule matched", getRequestLogValues(r, "rule", CANCEL)...)
					ctx.UserData.(*State).IsCancelled = true
					resp := &http.Response{
						Request:    r,
						StatusCode: 418,
						Body:       io.NopCloser(strings.NewReader("Request cancelled by Middleman")),
						Header:     make(http.Header),
					}

					a.responseToLog(resp, state, requestId)

					// Store request body in state
					responseBytes, err := io.ReadAll(r.Body)
					if err != nil {
						a.logger.Error(fmt.Sprintf("Error reading request body: %s", err))
					}
					r.Body.Close() //  must close
					r.Body = io.NopCloser(bytes.NewBuffer(responseBytes))
					state.RequestBody = string(responseBytes)

					return nil, resp
				}
			}
		}

		// If not cancelled.
		if !ctx.UserData.(*State).IsCancelled {

			// Redirects
			redirects, err := a.database.GetManyRules(REDIRECT)
			if err != nil {
				a.logger.Error(fmt.Sprintf("Error getting %s: %s", REDIRECT, err))
				log.Fatalf("Error getting %s: %s", REDIRECT, err)
			} else {
				for _, redirect := range redirects {
					if !redirect.Enabled {
						continue
					}
					if matches(redirect, r) {
						a.logger.Info("Redirect rule matched", getRequestLogValues(r, "rule", REDIRECT)...)
						ctx.UserData.(*State).IsRedirected = true
						resp := &http.Response{
							Request:    r,
							StatusCode: 307,
							Body:       io.NopCloser(strings.NewReader("Request cancelled by Middleman")),
							Header: http.Header{
								"Location": []string{redirect.RedirectTo},
							},
						}

						a.responseToLog(resp, state, requestId)
						return nil, resp
					}
				}
			}

			// Check for modify request body
			modRequestBodies, err := a.database.GetManyRules(MODIFY_REQUEST_BODY)
			if err != nil {
				a.logger.Error(fmt.Sprintf("Error getting %s: %s", MODIFY_REQUEST_BODY, err))
				log.Fatalf("Error getting %s: %s", MODIFY_REQUEST_BODY, err)
			} else {
				for _, modifyBody := range modRequestBodies {
					if !modifyBody.Enabled {
						continue
					}
					if matches(modifyBody, r) {
						a.logger.Info("ModifyRequestBody rule matched", getRequestLogValues(r, "rule", MODIFY_REQUEST_BODY)...)
						ctx.UserData.(*State).IsRequestBodyModified = true

						r.Body = io.NopCloser(bytes.NewReader([]byte(modifyBody.RequestBody)))
					}
				}
			}

			// Change request headers
			modifyHeaders, err := a.database.GetManyRules(MODIFY_REQUEST_HEADERS)
			if err != nil {
				a.logger.Error(fmt.Sprintf("Error getting %s: %s", MODIFY_REQUEST_HEADERS, err))
			} else {
				for _, modifyHeader := range modifyHeaders {

					if !modifyHeader.Enabled {
						continue
					}
					if matches(modifyHeader, r) {
						a.logger.Info("ModifyRequestHeader rule matched", getRequestLogValues(r, "rule", MODIFY_REQUEST_HEADERS)...)
						ctx.UserData.(*State).IsRequestHeaderModified = true

						for _, v := range modifyHeader.HeaderMods {
							if v.Action == "add" {
								a.logger.Info("Adding request header", "name", v.Name, "value", v.Value, "action", v.Action)
								r.Header.Add(v.Name, v.Value)
							} else if v.Action == "remove" {
								a.logger.Info("Removing request header: ", "name", v.Name, "value", v.Value)
								r.Header.Del(v.Name)
							} else if v.Action == "override" {
								a.logger.Info("Overriding request header: ", "name", v.Name, "value", v.Value)
								r.Header.Set(v.Name, v.Value)
							}
						}
					}
				}
			}
		}
		return r, nil
	}
}

func (a *App) getOnResponse() func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	return func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {

		state := ctx.UserData.(*State)

		if !state.IsCancelled {

			// Modify response body
			modResBodies, err := a.database.GetManyRules(MODIFY_RESPONSE_BODY)
			if err != nil {
				a.logger.Error(fmt.Sprintf("Error getting %s: %s", MODIFY_RESPONSE_BODY, err))
			} else {
				for _, modifyBody := range modResBodies {

					if !modifyBody.Enabled {
						continue
					}
					if matches(modifyBody, resp.Request) {
						a.logger.Info("ModifyResponseBody rule matched", getRequestLogValues(resp.Request, "rule", MODIFY_RESPONSE_BODY)...)
						ctx.UserData.(*State).IsResponseBodyModified = true

						resp.Body = io.NopCloser(bytes.NewReader([]byte(modifyBody.ResponseBody)))

					}
				}
			}
		}

		// Add delay
		delays, err := a.database.GetManyRules(DELAY)
		if err != nil {
			a.logger.Error(fmt.Sprintf("Error getting %s: %s", DELAY, err))
		} else {
			for _, delay := range delays {
				if !delay.Enabled {
					continue
				}
				if matches(delay, resp.Request) {
					a.logger.Info("Delay rule matched", getRequestLogValues(resp.Request, "rule", DELAY)...)
					ctx.UserData.(*State).DelayedBy = delay.DelaySec

					time.Sleep(time.Duration(delay.DelaySec) * time.Second)
				}
			}
		}

		// Change response headers
		modifyHeaders, err := a.database.GetManyRules(MODIFY_RESPONSE_HEADERS)
		if err != nil {
			a.logger.Error(fmt.Sprintf("Error getting %s: %s", MODIFY_RESPONSE_HEADERS, err))
		} else {
			for _, modifyHeader := range modifyHeaders {
				if !modifyHeader.Enabled {
					continue
				}
				if matches(modifyHeader, resp.Request) {
					a.logger.Info("ModifyResponseHeader rule matched", getRequestLogValues(resp.Request, "rule", MODIFY_RESPONSE_HEADERS)...)
					ctx.UserData.(*State).IsResponseHeaderModified = true

					for _, v := range modifyHeader.HeaderMods {
						if v.Action == "add" {
							a.logger.Info("Adding response header", "name", v.Name, "value", v.Value)
							resp.Header.Add(v.Name, v.Value)
						} else if v.Action == "remove" {
							a.logger.Info("Removing response header: ", "name", v.Name, "value", v.Value)
							resp.Header.Del(v.Name)
						} else if v.Action == "override" {
							a.logger.Info("Overriding response header: ", "name", v.Name, "value", v.Value)
							resp.Header.Set(v.Name, v.Value)
						}
					}
				}
			}
		}
		a.responseToLog(resp, state, state.requestId)
		return resp
	}
}

func (a *App) middlemanWeb(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	var tmplFile = "/Users/sumeetmathpati/Projects/sumeet/middleman/webserver/index.html"
	tmpl := template.Must(template.ParseFiles(tmplFile))
	tmpl.Execute(w, struct{ WebServerPath string }{WebServerPath: a.config.WebServerPath})
}

func (a *App) downloadCert(w http.ResponseWriter, r *http.Request) {
	certPath, _ := getCertKeyPath()
	w.Header().Set("Content-Disposition", "attachment; filename=\"Middleman.crt\"")
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, certPath)
}
