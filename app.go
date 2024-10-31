package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/elazarl/goproxy"
)

const WebServerPath = "/middleman"

type App struct {
	ctx              context.Context
	proxy            *goproxy.ProxyHttpServer
	proxyStartStop   chan bool
	webStartStop     chan bool
	database         Database
	config           Config
	logger           *slog.Logger
	httpRequests     []HttpRequestLog
	httpRequestsLock sync.Mutex
	webServerPath    string
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

	app := &App{
		proxy:          proxy,
		database:       NewSqliteDatabase(getDatabasePath()),
		proxyStartStop: make(chan bool),
		webStartStop:   make(chan bool),
		config:         config,
		logger:         slog.New(slog.NewJSONHandler(multiWriter, nil)),
		webServerPath:  WebServerPath,
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

func (a *App) requestToLog(r *http.Request, logId int) {

	a.httpRequests[logId].RequestHeaders = r.Header

}

func (a *App) responseToLog(r *http.Response, logId int) {
	a.httpRequests[logId].ResponseHeaders = r.Header
}

func (a *App) getOnRequest() func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	return func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {

		// Add request to logs
		a.httpRequestsLock.Lock()
		a.httpRequests = append(a.httpRequests, HttpRequestLog{
			Timestamp: time.Now(),
			Scheme:    r.URL.Scheme,
			Method:    r.Method,
			Host:      r.Host,
			Path:      r.URL.Path,
		})
		requestLogId := len(a.httpRequests) - 1
		a.httpRequestsLock.Unlock()

		// State for this request
		ctx.UserData = &State{requestId: requestLogId}

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
					a.httpRequests[requestLogId].Cancelled = true
					resp := &http.Response{
						Request:    r,
						StatusCode: 418,
						Body:       io.NopCloser(strings.NewReader("Request cancelled by Middleman")),
						Header:     make(http.Header),
					}

					a.requestToLog(r, requestLogId)
					a.responseToLog(resp, requestLogId)

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
						ctx.UserData.(*State).IsRedirected = true
						a.httpRequests[requestLogId].Redirected = true
						resp := &http.Response{
							Request:    r,
							StatusCode: 307,
							Body:       io.NopCloser(strings.NewReader("Request cancelled by Middleman")),
							Header: http.Header{
								"Location": []string{redirect.RedirectTo},
							},
						}

						a.requestToLog(r, requestLogId)
						a.responseToLog(resp, requestLogId)

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
						a.httpRequests[requestLogId].RequestBodyModified = true
						a.logger.Info("ModifyRequestBody  rule matched", getRequestLogValues(r, "rule", MODIFY_REQUEST_BODY)...)
						r.Body = io.NopCloser(bytes.NewReader([]byte(modifyBody.RequestBody)))
					}
				}
			}

			// Change request headers
			modifyHeaders, err := a.database.GetManyRules(MODIFY_HEADERS)
			if err != nil {
				a.logger.Error(fmt.Sprintf("Error getting %s: %s", MODIFY_HEADERS, err))
			} else {
				for _, modifyHeader := range modifyHeaders {

					if !modifyHeader.Enabled {
						continue
					}
					if matches(modifyHeader, r) {
						for _, v := range modifyHeader.RequestHeaderMods {
							if !v.IsRequest {
								continue
							}
							if v.Action == "add" {
								a.logger.Info("Adding request header", "name", v.Name, "value", v.Value, "action", v.Action)
								r.Header.Add(v.Name, v.Value)
								a.httpRequests[requestLogId].RequestHeaderModified = true
							} else if v.Action == "remove" {
								a.logger.Info("Removing request header: ", "name", v.Name, "value", v.Value)
								r.Header.Del(v.Name)
								a.httpRequests[requestLogId].RequestHeaderModified = true
							} else if v.Action == "override" {
								a.logger.Info("Overriding request header: ", "name", v.Name, "value", v.Value)
								r.Header.Set(v.Name, v.Value)
								a.httpRequests[requestLogId].RequestHeaderModified = true
							}
						}
					}
				}
			}
		}

		// ToDo: Save request body to log
		a.requestToLog(r, requestLogId)

		return r, nil
	}
}

func (a *App) getOnResponse() func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	return func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {

		if !ctx.UserData.(*State).IsCancelled {

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
						resp.Body = io.NopCloser(bytes.NewReader([]byte(modifyBody.ResponseBody)))
						a.httpRequests[ctx.UserData.(*State).requestId].ResponseBodyModified = true

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
					time.Sleep(time.Duration(delay.DelaySec) * time.Second)
					a.httpRequests[ctx.UserData.(*State).requestId].Delayed = delay.DelaySec
				}
			}
		}

		// Change response headers
		modifyHeaders, err := a.database.GetManyRules(MODIFY_HEADERS)
		if err != nil {
			a.logger.Error(fmt.Sprintf("Error getting %s: %s", MODIFY_HEADERS, err))
		} else {
			for _, modifyHeader := range modifyHeaders {
				if !modifyHeader.Enabled {
					continue
				}
				if matches(modifyHeader, resp.Request) {
					for _, v := range modifyHeader.ResponseHeaderMods {
						if v.IsRequest {
							continue
						}
						if v.Action == "add" {
							a.logger.Info("Adding response header", "name", v.Name, "value", v.Value)
							resp.Header.Add(v.Name, v.Value)
							a.httpRequests[ctx.UserData.(*State).requestId].ResponseHeaderModified = true
						} else if v.Action == "remove" {
							a.logger.Info("Removing response header: ", "name", v.Name, "value", v.Value)
							resp.Header.Del(v.Name)
							a.httpRequests[ctx.UserData.(*State).requestId].ResponseHeaderModified = true
						} else if v.Action == "override" {
							a.logger.Info("Overriding response header: ", "name", v.Name, "value", v.Value)
							resp.Header.Set(v.Name, v.Value)
							a.httpRequests[ctx.UserData.(*State).requestId].ResponseHeaderModified = true
						}
					}
				}
			}
		}
		a.responseToLog(resp, ctx.UserData.(*State).requestId)
		return resp
	}
}

func (a *App) middlemanWeb(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	var tmplFile = "/Users/sumeetmathpati/Projects/sumeet/middleman/webserver/index.html"
	tmpl := template.Must(template.ParseFiles(tmplFile))
	tmpl.Execute(w, struct{ WebServerPath string }{WebServerPath: a.webServerPath})
}

func (a *App) downloadCert(w http.ResponseWriter, r *http.Request) {
	certPath, _ := getCertKeyPath()
	w.Header().Set("Content-Disposition", "attachment; filename=\"Middleman.crt\"")
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, certPath)
}
