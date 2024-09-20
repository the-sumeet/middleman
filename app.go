package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	// "database/sql"

	"github.com/elazarl/goproxy"
	// _ "modernc.org/sqlite"
)

type State struct {
	IsCancelled  bool
	IsRedirected bool
}

type App struct {
	ctx             context.Context
	proxy           *goproxy.ProxyHttpServer
	proxyStartStoop chan bool
	database        Database
	requests        []http.Request
	config          Config
}

type ReturnValue struct {
	Redirects []Redirect `json:"redirects"`
	Cancels   []Cancel   `json:"cancels"`
	// Requests  []http.Request `json:"requests"`
	Error string `json:"error"`
}

func NewApp() *App {

	database := FileDatabase{
		filePath: "database.json",
	}
	database.load()

	proxy := goproxy.NewProxyHttpServer()
	config := getConfig()
	app := &App{
		proxy:           proxy,
		database:        &database,
		proxyStartStoop: make(chan bool),
		config:          config,
	}

	return app
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) getOnRequest() func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	return func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {

		ctx.UserData = &State{}

		// Check for cancels
		cancels := a.database.GetCancels()
		for _, cancel := range cancels {
			if cancel.matches(r) {
				ctx.UserData.(*State).IsCancelled = true
				res := &http.Response{
					Request:    r,
					StatusCode: 418,
					Body:       io.NopCloser(strings.NewReader("Request cancelled by Middleman")),
					Header:     make(http.Header),
				}
				return nil, res
			}
		}

		requestCopy := *r
		a.requests = append(a.requests, requestCopy)
		r.Header.Set("X-GoProxy", "yxorPoG-X")
		return r, nil
	}
}

func (a *App) getOnResponse() func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	return func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {

		if !ctx.UserData.(*State).IsCancelled {
			redirects := a.database.GetRedirects()
			for _, redirect := range redirects {
				if redirect.matches(resp) {
					ctx.UserData.(*State).IsRedirected = true
					resp.Header.Set("Location", redirect.ToValue)
					resp.StatusCode = 307
				}
			}
		}

		return resp
	}
}

func (a *App) StartProxy(port int) ReturnValue {
	certPath, certKey := getCertKeyPath()

	certytes, err := os.ReadFile(certPath)
	if err != nil {
		panic(err)
	}

	keybytes, err := os.ReadFile(certKey)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(certytes))
	fmt.Println(string(keybytes))
	goproxyCa, err := tls.X509KeyPair(certytes, keybytes)
	if err != nil {
		panic(err)
	}
	if goproxyCa.Leaf, err = x509.ParseCertificate(goproxyCa.Certificate[0]); err != nil {
		panic(err)
	}
	goproxy.GoproxyCa = goproxyCa
	goproxy.OkConnect = &goproxy.ConnectAction{Action: goproxy.ConnectAccept, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	goproxy.MitmConnect = &goproxy.ConnectAction{Action: goproxy.ConnectMitm, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	goproxy.HTTPMitmConnect = &goproxy.ConnectAction{Action: goproxy.ConnectHTTPMitm, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	goproxy.RejectConnect = &goproxy.ConnectAction{Action: goproxy.ConnectReject, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	//

	if err := PortAvailable("localhost", fmt.Sprintf("%d", port)); err == nil {
		log.Println("Port is not available")
		return ReturnValue{Error: fmt.Sprintf("Port %d is not available", port)}
	}

	portString := fmt.Sprintf(":%d", port)
	log.Println("Starting Proxy", portString)
	l, err := net.Listen("tcp", portString)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	log.Println("TCP listener started on ", portString)

	go func() {

		a.proxy.Verbose = true
		a.proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
		a.proxy.OnRequest().DoFunc(a.getOnRequest())
		a.proxy.OnResponse().DoFunc(a.getOnResponse())

		go func() {
			err := http.Serve(l, a.proxy)
			if err != nil {
				fmt.Println("Error starting server: ", err)
			}
			log.Println("Proxy serving started")
		}()

		log.Println("Proxy server goroutine started, waiting to stop")
		<-a.proxyStartStoop
		l.Close()
		log.Println("Proxy TCP listener closed")
	}()

	return ReturnValue{}
}

func (a *App) StopProxy() {
	log.Println("Stopping Proxy")
	a.proxyStartStoop <- true
}

func (a *App) GetConfig() Config {
	return a.config
}

// Config

func (a *App) AddConfigPort(port string) {
	a.config.ServerPort = port
	saveConfig(a.config)
}

// Redirects

func (a *App) GetRedirects() ReturnValue {
	return ReturnValue{
		Redirects: a.database.GetRedirects(),
	}
}

func (a *App) SaveRedirect(redirectId int, redirect Redirect) {
	a.database.SaveRedirect(redirectId, redirect)
}

func (a *App) AddRedirect(redirect Redirect) {
	a.database.AddRedirect(redirect)
}

func (a *App) RemoveRedirect(redirectId int) {
	a.database.RemoveRedirect(redirectId)
}

func (a *App) GenerateCert() {
	genCert()
}

// Cancels
func (a *App) GetCancels() ReturnValue {
	return ReturnValue{
		Cancels: a.database.GetCancels(),
	}
}

func (a *App) SaveCancel(cancelId int, cancel Cancel) {
	a.database.SaveCancel(cancelId, cancel)
}

func (a *App) AddCancel(cancel Cancel) {
	a.database.AddCancel(cancel)
}

func (a *App) RemoveCancel(cancelId int) {
	a.database.RemoveCancel(cancelId)
}
