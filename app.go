package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/elazarl/goproxy"
)

type App struct {
	ctx             context.Context
	proxy           *goproxy.ProxyHttpServer
	proxyStartStoop chan bool
	database        Database
}

type ReturnValue struct {
	Redirects []Redirect `json:"redirects"`
	Error     string     `json:"error"`
}

func NewApp() *App {

	database := FileDatabase{
		filePath: "database.json",
	}
	database.load()
	genCert()

	proxy := goproxy.NewProxyHttpServer()

	app := &App{
		proxy:           proxy,
		database:        &database,
		proxyStartStoop: make(chan bool),
	}

	return app
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) StartProxy() {
	setCA()
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	go func() {

		log.Println("Proxy Starting")
		a.proxy.Verbose = true
		a.proxy.OnRequest().DoFunc(
			func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
				fmt.Println(r.URL)
				r.Header.Set("X-GoProxy", "yxorPoG-X")
				return r, nil
			})

		go func() {
			err := http.Serve(l, a.proxy)
			if err != nil {
				fmt.Println("Proxy Stopped: ", err)
			}
			log.Println("Proxy Started")
		}()

		fmt.Println("Waiting for stop")
		<-a.proxyStartStoop
		fmt.Println("Stopping")
		l.Close()

		log.Println("Proxy Stopped")
	}()
	log.Println("Exit")

}

func (a *App) StopProxy() {
	log.Println("Proxy Stopping")
	a.proxyStartStoop <- true
	log.Println("Proxy Stopped")
}

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
