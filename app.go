package main

import (
	"bytes"
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
	"time"

	"github.com/elazarl/goproxy"
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
	Redirects          []Redirect           `json:"redirects"`
	Cancels            []Cancel             `json:"cancels"`
	Delays             []Delay              `json:"delays"`
	ModifyHeaders      []ModifyHeader       `json:"modifyHeaders"`
	ModifyRequestBody  []ModifyRequestBody  `json:"modifyRequestBody"`
	ModifyResponseBody []ModifyResponseBody `json:"modifyResponseBody"`
	Error              string               `json:"error"`
}

type InValue struct {
	Redirect           Redirect           `json:"redirect"`
	Cancel             Cancel             `json:"cancel"`
	Delay              Delay              `json:"delay"`
	ModifyHeader       ModifyHeader       `json:"modifyHeader"`
	ModifyRequestBody  ModifyRequestBody  `json:"modifyRequestBody"`
	ModifyResponseBody ModifyResponseBody `json:"modifyResponseBody"`
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

	if !certKeyPresent() {
		genCert()
	}
}

func (a *App) getOnRequest() func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	return func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {

		ctx.UserData = &State{}

		// Check for cancels
		cancels, err := a.database.GetMany("cancel")
		if err != nil {
			fmt.Println("Error getting cancels: ", err)
		} else {
			for _, v := range cancels {
				cancel := v.(Cancel)
				if !cancel.Enabled {
					continue
				}
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
			redirects, err := a.database.GetMany("redirect")
			if err != nil {
				fmt.Println("Error getting redirects: ", err)
			} else {
				for _, v := range redirects {
					redirect := v.(Redirect)
					if !redirect.Enabled {
						continue
					}
					if redirect.matches(resp) {
						ctx.UserData.(*State).IsRedirected = true
						resp.Header.Set("Location", redirect.ToValue)
						resp.StatusCode = 307
					}
				}
			}

			// Modify response body
			modResBodies, err := a.database.GetMany(MODIFY_RESPONSE_BODY)
			if err != nil {
				fmt.Println("Error getting modify resposne bodies: ", err)
			} else {
				for _, v := range modResBodies {
					modResBody := v.(ModifyResponseBody)
					if !modResBody.Enabled {
						continue
					}
					if modResBody.matches(resp) {
						resp.Body = io.NopCloser(bytes.NewReader([]byte(modResBody.Body)))
					}
				}
			}
		}

		// Add delay
		delays, err := a.database.GetMany("delay")
		if err != nil {
			fmt.Println("Error getting delays: ", err)
		} else {
			for _, v := range delays {
				delay := v.(Delay)
				if !delay.Enabled {
					continue
				}
				if delay.matches(resp) {
					time.Sleep(time.Duration(delay.DelaySec) * time.Second)
				}
			}
		}

		// Change headers
		modifyHeaders, err := a.database.GetMany(MODIFY_HEADERS)
		if err != nil {
			fmt.Println("Error getting modify headers: ", err)
		} else {
			for _, v := range modifyHeaders {
				modifyHeader := v.(ModifyHeader)
				if !modifyHeader.Enabled {
					continue
				}
				if (&modifyHeader).matches(resp.Request) {
					for _, v := range modifyHeader.Mods {
						if v.IsRequest {
							continue
						}
						if v.Action == "add" {
							fmt.Println("Adding header: ", v.Name, v.Value)
							resp.Header.Add(v.Name, v.Value)
						}
					}
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

func (a *App) AddConfigPort(port string) {
	a.config.ServerPort = port
	saveConfig(a.config)
}

func (a *App) GetMany(recordType string) ReturnValue {

	res, err := a.database.GetMany(recordType)
	if err != nil {
		return ReturnValue{Error: err.Error()}
	}

	if recordType == REDIRECT {
		redirects := make([]Redirect, len(res))
		for i, v := range res {
			redirects[i] = v.(Redirect)
		}
		return ReturnValue{
			Redirects: redirects,
		}
	}
	if recordType == CANCEL {
		cancels := make([]Cancel, len(res))
		for i, v := range res {
			cancels[i] = v.(Cancel)
		}
		return ReturnValue{
			Cancels: cancels,
		}
	}
	if recordType == DELAY {
		delays := make([]Delay, len(res))
		for i, v := range res {
			delays[i] = v.(Delay)
		}
		return ReturnValue{
			Delays: delays,
		}
	}

	if recordType == MODIFY_HEADERS {
		modifyHeader := make([]ModifyHeader, len(res))
		for i, v := range res {
			modifyHeader[i] = v.(ModifyHeader)
		}
		return ReturnValue{
			ModifyHeaders: modifyHeader,
		}
	}

	if recordType == MODIFY_REQUEST_BODY {
		modifyRequestBody := make([]ModifyRequestBody, len(res))
		for i, v := range res {
			modifyRequestBody[i] = v.(ModifyRequestBody)
		}
		return ReturnValue{
			ModifyRequestBody: modifyRequestBody,
		}
	}

	if recordType == MODIFY_RESPONSE_BODY {
		modifyResponseBody := make([]ModifyResponseBody, len(res))
		for i, v := range res {
			modifyResponseBody[i] = v.(ModifyResponseBody)
		}
		return ReturnValue{
			ModifyResponseBody: modifyResponseBody,
		}
	}
	return ReturnValue{}
}

func (a *App) Save(recordType string, recordId int, input InValue) ReturnValue {
	if recordType == REDIRECT {
		err := a.database.Save(recordType, recordId, input.Redirect)
		if err != nil {
			return ReturnValue{Error: err.Error()}
		}
	}
	if recordType == CANCEL {
		err := a.database.Save(recordType, recordId, input.Cancel)
		if err != nil {
			return ReturnValue{Error: err.Error()}
		}
	}
	if recordType == DELAY {
		err := a.database.Save(recordType, recordId, input.Delay)
		if err != nil {
			return ReturnValue{Error: err.Error()}
		}
	}
	if recordType == MODIFY_HEADERS {
		fmt.Println(input.ModifyHeader.Mods)
		err := a.database.Save(recordType, recordId, input.ModifyHeader)
		if err != nil {
			return ReturnValue{Error: err.Error()}
		}
	}

	if recordType == MODIFY_REQUEST_BODY {
		err := a.database.Save(recordType, recordId, input.ModifyRequestBody)
		if err != nil {
			return ReturnValue{Error: err.Error()}
		}
	}

	if recordType == MODIFY_RESPONSE_BODY {
		err := a.database.Save(recordType, recordId, input.ModifyResponseBody)
		if err != nil {
			return ReturnValue{Error: err.Error()}
		}
	}

	return ReturnValue{}
}

func (a *App) Remove(recordType string, recordId int) ReturnValue {
	err := a.database.Remove(recordType, recordId)
	if err != nil {
		return ReturnValue{Error: err.Error()}
	}
	return ReturnValue{}
}

func (a *App) Add(recordType string, records InValue) ReturnValue {

	if recordType == REDIRECT {
		err := a.database.Add(recordType, records.Redirect)
		if err != nil {
			return ReturnValue{Error: err.Error()}
		}
	}
	if recordType == CANCEL {
		err := a.database.Add(recordType, records.Cancel)
		if err != nil {
			return ReturnValue{Error: err.Error()}
		}
	}
	if recordType == DELAY {
		err := a.database.Add(recordType, records.Delay)
		if err != nil {
			return ReturnValue{Error: err.Error()}
		}
	}
	if recordType == MODIFY_HEADERS {
		err := a.database.Add(recordType, records.ModifyHeader)
		if err != nil {
			return ReturnValue{Error: err.Error()}
		}
	}

	if recordType == MODIFY_REQUEST_BODY {
		err := a.database.Add(recordType, records.ModifyRequestBody)
		if err != nil {
			return ReturnValue{Error: err.Error()}
		}
	}

	if recordType == MODIFY_RESPONSE_BODY {
		err := a.database.Add(recordType, records.ModifyResponseBody)
		if err != nil {
			return ReturnValue{Error: err.Error()}
		}
	}

	return ReturnValue{}
}

func (a *App) GenerateCert() {
	genCert()
}
