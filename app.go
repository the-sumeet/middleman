package main

import (
	"context"
)

type App struct {
	ctx      context.Context
	proxy    *Proxy
	database Database
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

	certPath, keyPath := getCertKeyPath()

	return &App{
		proxy:    NewProxy(certPath, keyPath),
		database: &database,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
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
