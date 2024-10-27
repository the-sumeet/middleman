package main

// This file contains the methods that are exposed to the frontend

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/elazarl/goproxy"
)

type ReturnValue struct {
	Redirects          []Redirect           `json:"redirects"`
	Cancels            []Cancel             `json:"cancels"`
	Delays             []Delay              `json:"delays"`
	ModifyHeaders      []ModifyHeader       `json:"modifyHeaders"`
	ModifyRequestBody  []ModifyRequestBody  `json:"modifyRequestBody"`
	ModifyResponseBody []ModifyResponseBody `json:"modifyResponseBody"`
	Logs               []string             `json:"logs"`
	HttpRequests       []HttpRequestLog     `json:"httpRequests"`
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

func (a *App) StartProxy() ReturnValue {
	certPath, certKey := getCertKeyPath()

	certytes, err := os.ReadFile(certPath)
	if err != nil {
		panic(err)
	}

	keybytes, err := os.ReadFile(certKey)
	if err != nil {
		panic(err)
	}

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

	if err := PortAvailable("localhost", a.config.ProxyServerPort); err == nil {
		a.logger.Info(fmt.Sprintf("Port %s is not available", a.config.ProxyServerPort))
		return ReturnValue{Error: fmt.Sprintf("Port %s is not available", a.config.ProxyServerPort)}
	}

	portString := fmt.Sprintf(":%s", a.config.ProxyServerPort)
	log.Println("Starting Proxy", portString)
	l, err := net.Listen("tcp", portString)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Error listening: %s", err))
		return ReturnValue{Error: fmt.Sprintf("Error listening: %s", err)}
	}
	a.logger.Info(fmt.Sprintf("Started listening on: %s", portString))

	go func() {

		// a.proxy.Verbose = true
		a.proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
		a.proxy.OnRequest().DoFunc(a.getOnRequest())
		a.proxy.OnResponse().DoFunc(a.getOnResponse())

		go func() {

			err = http.Serve(l, a.proxy)
			if err != nil {
				opErr, ok := err.(*net.OpError)
				if ok {
					if errors.Is(opErr.Err, net.ErrClosed) {
						a.logger.Info("Proxy TCP listener closed")
						return
					}
				}
				a.logger.Error(fmt.Sprintf("Error starting server: %s", err))
				log.Fatal("Error starting server: ", err)
			}
		}()

		a.logger.Info("Proxy server goroutine started, waiting to stop")
		<-a.proxyStartStop
		l.Close()
	}()

	return ReturnValue{}
}

func (a *App) StopProxy() {
	log.Println("Stopping proxy")
	a.proxyStartStop <- true
}

func (a *App) StartWebServer() ReturnValue {
	if err := PortAvailable("localhost", a.config.WebServerPort); err == nil {
		a.logger.Info(fmt.Sprintf("Port %s is not available for web server", a.config.WebServerPort))
		return ReturnValue{Error: fmt.Sprintf("Port %s is not available", a.config.WebServerPort)}
	}

	portString := fmt.Sprintf(":%s", a.config.WebServerPort)
	a.logger.Info(fmt.Sprintf("Starting web server on: %s", portString))
	l, err := net.Listen("tcp", portString)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Error starting web server: %s", err))
		log.Fatal("Error start listening: ", err)
	}
	a.logger.Info(fmt.Sprintf("Started listening on: %s for web server", portString))

	go func() {
		go func() {

			mux := http.NewServeMux()
			mux.HandleFunc(a.webServerPath, a.middlemanWeb)
			mux.HandleFunc(fmt.Sprintf("%s/cert", a.webServerPath), a.downloadCert)

			err = http.Serve(l, mux)
			if err != nil {
				opErr, ok := err.(*net.OpError)
				if ok {
					if errors.Is(opErr.Err, net.ErrClosed) {
						a.logger.Info("Web TCP listener closed")
						return
					}
				}
				a.logger.Error(fmt.Sprintf("Error starting web server: %s", err))
				log.Fatal("Error starting web server: ", err)
			}
			a.logger.Info("Web server started")
		}()

		a.logger.Info("Web server goroutine started, waiting to stop")
		<-a.webStartStop
		l.Close()
		a.logger.Info("Web TCP listener closed")
	}()

	return ReturnValue{}
}

func (a *App) StopWebServer() {
	log.Println("Stopping web server")
	a.webStartStop <- true
}

func (a *App) GetConfig() Config {
	return a.config
}

func (a *App) AddProxyPort(proxyPort string) ReturnValue {
	if len(proxyPort) != 4 {
		return ReturnValue{Error: "Port must be 4 digits"}
	}
	if proxyPort == a.config.WebServerPort {
		return ReturnValue{Error: "Proxy and Web server port cannot be same"}
	}
	a.config.ProxyServerPort = proxyPort
	saveConfig(a.config)
	return ReturnValue{}
}

func (a *App) AddWebPort(webPort string) ReturnValue {
	if len(webPort) != 4 {
		return ReturnValue{Error: "Port must be 4 digits"}
	}
	if webPort == a.config.ProxyServerPort {
		return ReturnValue{Error: "Proxy and Web server port cannot be same"}
	}
	a.config.WebServerPort = webPort
	saveConfig(a.config)
	return ReturnValue{}
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

func (a *App) GetLogs() ReturnValue {
	return ReturnValue{
		HttpRequests: a.httpRequests,
	}
}

func (a *App) GetWebServerPath() string {
	return a.webServerPath
}
