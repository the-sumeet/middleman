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
	Logs         []string         `json:"logs"`
	HttpRequests []HttpRequestLog `json:"httpRequests"`
	Error        string           `json:"error"`
	InsertedId   any              `json:"insertedId"`
	Rules        []Rule           `json:"rules"`
}

type InValue struct {
	Id   any  `json:"id"`
	Rule Rule `json:"rule"`
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
					} else {
						a.logger.Error(fmt.Sprintf("Error starting web server: %s", err))
						log.Fatal("Error starting web server: ", err)
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
	a.logger.Info("Stopping proxy")
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
					} else {
						a.logger.Error(fmt.Sprintf("Error starting web server: %s", err))
						log.Fatal("Error starting web server: ", err)
					}
				}
				a.logger.Error(fmt.Sprintf("Error starting web server: %s", err))
				log.Fatal("Error starting web server: ", err)
			}
		}()

		a.logger.Info("Web server goroutine started, waiting to stop")
		<-a.webStartStop
		l.Close()
	}()

	return ReturnValue{}
}

func (a *App) StopWebServer() {
	a.logger.Info("Stopping web server")
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

func (a *App) GetWebServerPath() string {
	return a.webServerPath
}

func (a *App) GenerateCert() {
	genCert()
}

func (a *App) GetLogs() ReturnValue {
	return ReturnValue{
		HttpRequests: a.httpRequests,
	}
}

// Rule CRUD

func (a *App) UpdateRule(ruleIn InValue) ReturnValue {
	updatedRule, err := a.database.UpdateRule(ruleIn.Id, ruleIn.Rule)
	if err != nil {
		return ReturnValue{Error: err.Error()}
	}
	return ReturnValue{Rules: []Rule{updatedRule}}
}

func (a *App) GetOneRule(id int) ReturnValue {
	rule, err := a.database.GetOneRule(id)
	if err != nil {
		return ReturnValue{Error: err.Error()}
	}
	return ReturnValue{Rules: []Rule{rule}}
}

func (a *App) GetManyRules(recordType string) ReturnValue {

	res, err := a.database.GetManyRules(recordType)
	if err != nil {
		return ReturnValue{Error: err.Error()}
	}

	return ReturnValue{Rules: res}
}

func (a *App) RemoveRule(recordId int) ReturnValue {
	err := a.database.RemoveRule(recordId)
	if err != nil {
		return ReturnValue{Error: err.Error()}
	}
	return ReturnValue{}
}

func (a *App) AddRule(records InValue) ReturnValue {

	fmt.Println("Adding rule", records.Rule)
	id, err := a.database.AddRule(records.Rule)
	if err != nil {
		return ReturnValue{Error: err.Error()}
	}
	return ReturnValue{InsertedId: id, Rules: []Rule{records.Rule}}
}
