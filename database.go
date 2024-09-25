package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	REDIRECT = "redirect"
	CANCEL   = "cancel"
	DELAY    = "delay"
)

type Request struct {
	Entity string `json:"entity"`
	Op     string `json:"op"`
	Value  string `json:"value"`
}

type Redirect struct {
	Request
	ToType  string `json:"toType"`
	ToValue string `json:"toValue"`
	Enabled bool   `json:"enabled"`
}

func (r *Redirect) matches(req *http.Response) bool {
	entityValue := getRequestEntity(r.Entity, req.Request.URL.Path, req.Request.Method, req.Request.URL.Host)
	return evalOp(r.Op, entityValue, r.Value)
}

type Cancel Request

func (r *Cancel) matches(req *http.Request) bool {
	entityValue := getRequestEntity(r.Entity, req.URL.Path, req.Method, req.URL.Host)
	return evalOp(r.Op, entityValue, r.Value)
}

type Delay struct {
	Request
	DelaySec int `json:"delaySec"`
}

func (r *Delay) matches(res *http.Response) bool {
	entityValue := getRequestEntity(r.Entity, res.Request.URL.Path, res.Request.Method, res.Request.URL.Host)
	return evalOp(r.Op, entityValue, r.Value)
}

type Database interface {
	GetMany(recordType string) ([]any, error)
	Save(recordType string, id int, value any) error
	Remove(recordType string, id int) error
	Add(recordType string, value any) error
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
		Delays    []Delay    `json:"delays"`
	}
	err = json.Unmarshal(data, &database)

	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	f.redirects = database.Redirects
	f.cancels = database.Cancels
	f.delays = database.Delays
}

func (f *FileDatabase) store() {
	data, err := json.Marshal(map[string]any{
		"redirects": f.redirects,
		"cancels":   f.cancels,
		"delays":    f.delays,
	})
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(f.filePath, data, 0644)
	if err != nil {
		panic(err)
	}
}

// Generic methods
func (f *FileDatabase) GetMany(recordType string) ([]any, error) {
	if recordType == REDIRECT {
		interfaceSlice := make([]interface{}, len(f.redirects))
		for i, v := range f.redirects {
			interfaceSlice[i] = v
		}
		return interfaceSlice, nil
	}
	if recordType == CANCEL {
		interfaceSlice := make([]interface{}, len(f.cancels))
		for i, v := range f.cancels {
			interfaceSlice[i] = v
		}
		return interfaceSlice, nil
	}
	if recordType == DELAY {
		interfaceSlice := make([]interface{}, len(f.delays))
		for i, v := range f.delays {
			interfaceSlice[i] = v
		}
		return interfaceSlice, nil
	}
	return []any{}, errors.New("invalid record type")
}

func (f *FileDatabase) Save(recordType string, id int, value any) error {
	if recordType == REDIRECT {
		if id >= len(f.redirects) {
			return fmt.Errorf("redirect with id %d not found", id)
		}
		f.redirects[id] = value.(Redirect)
	}
	if recordType == CANCEL {
		if id >= len(f.cancels) {
			return fmt.Errorf("cancel with id %d not found", id)
		}
		f.cancels[id] = value.(Cancel)
	}
	if recordType == DELAY {
		if id >= len(f.delays) {
			return fmt.Errorf("delay with id %d not found", id)
		}
		f.delays[id] = value.(Delay)
	}
	f.store()
	return nil
}

func (f *FileDatabase) Remove(recordType string, id int) error {
	if recordType == REDIRECT {
		if id >= len(f.redirects) {
			return fmt.Errorf("redirect with id %d not found", id)
		}

		f.redirects = append(f.redirects[:id], f.redirects[id+1:]...)
		f.store()
	}

	if recordType == CANCEL {
		if id >= len(f.cancels) {
			return fmt.Errorf("cancel with id %d not found", id)
		}
		f.cancels = append(f.cancels[:id], f.cancels[id+1:]...)
		f.store()
	}

	if recordType == DELAY {
		if id >= len(f.delays) {
			return fmt.Errorf("delay with id %d not found", id)
		}
		f.delays = append(f.delays[:id], f.delays[id+1:]...)
		f.store()
	}

	return nil
}

func (f *FileDatabase) Add(recordType string, value any) error {
	if recordType == REDIRECT {
		f.redirects = append(f.redirects, value.(Redirect))
		f.store()
	}

	if recordType == CANCEL {
		f.cancels = append(f.cancels, value.(Cancel))
		f.store()
	}

	if recordType == DELAY {
		f.delays = append(f.delays, value.(Delay))
		f.store()
	}

	return nil
}
