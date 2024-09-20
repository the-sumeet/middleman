package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Request struct {
	Entity string `json:"entity"`
	Op     string `json:"op"`
	Value  string `json:"value"`
	ToType string `json:"toType"`
}

type Redirect struct {
	Request
	ToValue string `json:"toValue"`
	Enabled bool   `json:"enabled"`
}

func (r *Redirect) matches(req *http.Response) bool {
	entityValue := getRequestEntity(r.Entity, req.Request.URL.Path, req.Request.Method, req.Request.URL.Host)
	return evalOp(r.Op, entityValue, r.Value)
}

type Cancel struct {
	Request
}

func (r *Cancel) matches(req *http.Request) bool {
	entityValue := getRequestEntity(r.Entity, req.URL.Path, req.Method, req.URL.Host)
	return evalOp(r.Op, entityValue, r.Value)
}

type Delay struct {
	Request
	DelaySec int `json:"delaySec"`
}

type Database interface {
	// InsertCollection(record Record) error
	// GetCollection() (Collection, error)

	// Redirects
	GetRedirects() []Redirect
	SaveRedirect(redirectId int, redirect Redirect) error
	RemoveRedirect(redirectId int) error
	AddRedirect(redirect Redirect) error

	// Cancels
	GetCancels() []Cancel
	SaveCancel(cancelId int, cancel Cancel) error
	RemoveCancel(cancelId int) error
	AddCancel(cancel Cancel) error
}

type FileDatabase struct {
	filePath  string
	redirects []Redirect
	cancels   []Cancel
	delays    []Delay
}

func (f *FileDatabase) load() {
	data, err := os.ReadFile(f.filePath)
	if err != nil {
		panic(err)
	}

	var database struct {
		Redirects []Redirect `json:"redirects"`
		Cancels   []Cancel   `json:"cancels"`
	}
	err = json.Unmarshal(data, &database)

	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	f.redirects = database.Redirects
	f.cancels = database.Cancels
}

func (f *FileDatabase) store() {
	data, err := json.Marshal(map[string]any{
		"redirects": f.redirects,
		"cancels":   f.cancels,
	})
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(f.filePath, data, 0644)
	if err != nil {
		panic(err)
	}
}

// Redirects
func (f *FileDatabase) GetRedirects() []Redirect {
	return f.redirects
}

func (f *FileDatabase) SaveRedirect(redirectId int, redirect Redirect) error {

	if redirectId >= len(f.redirects) {
		return fmt.Errorf("redirect with id %s not found", redirectId)
	}

	f.redirects[redirectId] = redirect
	f.store()
	return nil
}

func (f *FileDatabase) RemoveRedirect(redirectId int) error {

	if redirectId >= len(f.redirects) {
		return fmt.Errorf("redirect with id %s not found", redirectId)
	}

	f.redirects = append(f.redirects[:redirectId], f.redirects[redirectId+1:]...)

	f.store()
	return nil
}

func (f *FileDatabase) AddRedirect(redirect Redirect) error {
	f.redirects = append(f.redirects, redirect)
	f.store()
	return nil
}

// Cancels
func (f *FileDatabase) GetCancels() []Cancel {
	return f.cancels
}

func (f *FileDatabase) SaveCancel(cancelId int, cancel Cancel) error {

	if cancelId >= len(f.cancels) {
		return fmt.Errorf("redirect with id %s not found", cancelId)
	}

	f.cancels[cancelId] = cancel
	f.store()
	return nil
}

func (f *FileDatabase) RemoveCancel(cancelId int) error {

	if cancelId >= len(f.cancels) {
		return fmt.Errorf("redirect with id %s not found", cancelId)
	}

	f.cancels = append(f.cancels[:cancelId], f.cancels[cancelId+1:]...)
	f.store()
	return nil
}

func (f *FileDatabase) AddCancel(cancel Cancel) error {
	f.cancels = append(f.cancels, cancel)
	f.store()
	return nil
}

// Delays
func (f *FileDatabase) GetDelays() []Delay {
	return f.delays
}

func (f *FileDatabase) SaveDelay(delayId int, delay Delay) error {

	if delayId >= len(f.delays) {
		return fmt.Errorf("delay with id %d not found", delayId)
	}

	f.delays[delayId] = delay
	f.store()
	return nil
}

func (f *FileDatabase) RemoveDelay(delayId int) error {

	if delayId >= len(f.delays) {
		return fmt.Errorf("delay with id %d not found", delayId)
	}

	f.delays = append(f.delays[:delayId], f.delays[delayId+1:]...)
	f.store()
	return nil
}

func (f *FileDatabase) AddDelay(delay Delay) error {
	f.delays = append(f.delays, delay)
	f.store()
	return nil
}
