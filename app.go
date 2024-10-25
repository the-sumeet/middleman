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

	databasePath := getDatabasePath()
	if _, err := os.Stat(databasePath); os.IsNotExist(err) {
		// Create the database file
		file, err := os.Create(databasePath)
		file.Write([]byte("{}"))
		if err != nil {
			panic(err)
		}
	}

	database := FileDatabase{
		filePath: databasePath,
	}
	database.load()

	logFile := getLogFilePath()
	logWriter, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	proxy := goproxy.NewProxyHttpServer()

	// Get hostname
	hostname, err := os.Hostname()
	if err != nil {
		hostname = ""
	}
	hostname = strings.ReplaceAll(hostname, " ", "")
	webserverPath := "/middleman"
	if hostname != "" {
		webserverPath += "-on-" + hostname
	}

	app := &App{
		proxy:          proxy,
		database:       &database,
		proxyStartStop: make(chan bool),
		config:         config,
		logger:         slog.New(slog.NewJSONHandler(logWriter, nil)),
		webServerPath:  webserverPath,
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
		cancels, err := a.database.GetMany(CANCEL)
		if err != nil {
			a.logger.Error(fmt.Sprintf("Error getting cancels: %s", err))
			log.Fatal("Error getting cancels: ", err)
		} else {
			for _, v := range cancels {
				cancel := v.(Cancel)
				if !cancel.Enabled {
					continue
				}
				if matches(Request(cancel), r) {
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
			redirects, err := a.database.GetMany(REDIRECT)
			if err != nil {
				a.logger.Error(fmt.Sprintf("Error getting %s: %s", REDIRECT, err))
				log.Fatalf("Error getting %s: %s", REDIRECT, err)
			} else {
				for _, v := range redirects {
					redirect := v.(Redirect)
					if !redirect.Enabled {
						continue
					}
					if matches(redirect.Request, r) {
						ctx.UserData.(*State).IsRedirected = true
						a.httpRequests[requestLogId].Redirected = true
						resp := &http.Response{
							Request:    r,
							StatusCode: 307,
							Body:       io.NopCloser(strings.NewReader("Request cancelled by Middleman")),
							Header: http.Header{
								"Location": []string{redirect.ToValue},
							},
						}

						a.requestToLog(r, requestLogId)
						a.responseToLog(resp, requestLogId)

						return nil, resp
					}
				}
			}

			// Check for modify request body
			modRequestBodies, err := a.database.GetMany(MODIFY_REQUEST_BODY)
			if err != nil {
				a.logger.Error(fmt.Sprintf("Error getting %s: %s", MODIFY_REQUEST_BODY, err))
				log.Fatalf("Error getting %s: %s", MODIFY_REQUEST_BODY, err)
			} else {
				for _, v := range modRequestBodies {
					modResBody := v.(ModifyRequestBody)
					if !modResBody.Enabled {
						continue
					}
					if matches(modResBody.Request, r) {
						a.httpRequests[requestLogId].RequestBodyModified = true
						a.logger.Info("ModifyRequestBody  rule matched", getRequestLogValues(r, "rule", MODIFY_REQUEST_BODY)...)
						r.Body = io.NopCloser(bytes.NewReader([]byte(modResBody.Body)))
					}
				}
			}

			// Change request headers
			modifyHeaders, err := a.database.GetMany(MODIFY_HEADERS)
			if err != nil {
				a.logger.Error(fmt.Sprintf("Error getting %s: %s", MODIFY_HEADERS, err))
			} else {
				for _, v := range modifyHeaders {
					modifyHeader := v.(ModifyHeader)
					if !modifyHeader.Enabled {
						continue
					}
					if matches(modifyHeader.Request, r) {
						for _, v := range modifyHeader.Mods {
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
			modResBodies, err := a.database.GetMany(MODIFY_RESPONSE_BODY)
			if err != nil {
				a.logger.Error(fmt.Sprintf("Error getting %s: %s", MODIFY_RESPONSE_BODY, err))
			} else {
				for _, v := range modResBodies {
					modResBody := v.(ModifyResponseBody)
					if !modResBody.Enabled {
						continue
					}
					if matches(modResBody.Request, resp.Request) {
						resp.Body = io.NopCloser(bytes.NewReader([]byte(modResBody.Body)))
						a.httpRequests[ctx.UserData.(*State).requestId].ResponseBodyModified = true

					}
				}
			}
		}

		// Add delay
		delays, err := a.database.GetMany("delay")
		if err != nil {
			a.logger.Error(fmt.Sprintf("Error getting %s: %s", DELAY, err))
		} else {
			for _, v := range delays {
				delay := v.(Delay)
				if !delay.Enabled {
					continue
				}
				if matches(delay.Request, resp.Request) {
					time.Sleep(time.Duration(delay.DelaySec) * time.Second)
					a.httpRequests[ctx.UserData.(*State).requestId].Delayed = delay.DelaySec
				}
			}
		}

		// Change response headers
		modifyHeaders, err := a.database.GetMany(MODIFY_HEADERS)
		if err != nil {
			a.logger.Error(fmt.Sprintf("Error getting %s: %s", MODIFY_HEADERS, err))
		} else {
			for _, v := range modifyHeaders {
				modifyHeader := v.(ModifyHeader)
				if !modifyHeader.Enabled {
					continue
				}
				if matches(modifyHeader.Request, resp.Request) {
					for _, v := range modifyHeader.Mods {
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
