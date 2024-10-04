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
	REDIRECT             = "redirect"
	CANCEL               = "cancel"
	DELAY                = "delay"
	MODIFY_HEADERS       = "modifyHeader"
	MODIFY_REQUEST_BODY  = "modifyRequestBody"
	MODIFY_RESPONSE_BODY = "modifyResponseBody"
)

type Request struct {
	Enabled bool   `json:"enabled"`
	Entity  string `json:"entity"`
	Op      string `json:"op"`
	Value   string `json:"value"`
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

type ModifyRequestBody struct {
	Request
	Body string `json:"body"`
}

func (m *ModifyRequestBody) matches(req *http.Request) bool {
	entityValue := getRequestEntity(m.Entity, req.URL.Path, req.Method, req.URL.Host)
	return evalOp(m.Op, entityValue, m.Value)
}

type ModifyResponseBody struct {
	Request
	Body string `json:"body"`
}

func (m *ModifyResponseBody) matches(res *http.Response) bool {
	entityValue := getRequestEntity(m.Entity, res.Request.URL.Path, res.Request.Method, res.Request.URL.Host)
	return evalOp(m.Op, entityValue, m.Value)
}

type Header struct {
	Action    string `json:"action"`
	IsRequest bool   `json:"isRequest"`
	Name      string `json:"name"`
	Value     string `json:"value"`
}
type ModifyHeader struct {
	Request
	Mods []Header `json:"mods"`
}

func (r *ModifyHeader) matches(req *http.Request) bool {
	entityValue := getRequestEntity(r.Entity, req.URL.Path, req.Method, req.URL.Host)
	return evalOp(r.Op, entityValue, r.Value)
}

type Database interface {
	GetMany(recordType string) ([]any, error)
	Save(recordType string, id int, value any) error
	Remove(recordType string, id int) error
	Add(recordType string, value any) error
}

type FileDatabase struct {
	filePath           string
	redirects          []Redirect
	cancels            []Cancel
	delays             []Delay
	modifyHeaders      []ModifyHeader
	modifyRequestBody  []ModifyRequestBody
	modifyResponseBody []ModifyResponseBody
}

func (f *FileDatabase) load() {
	data, err := os.ReadFile(f.filePath)
	if err != nil {
		panic(err)
	}

	var database struct {
		Redirects          []Redirect           `json:"redirects"`
		Cancels            []Cancel             `json:"cancels"`
		Delays             []Delay              `json:"delays"`
		ModifyHeaders      []ModifyHeader       `json:"modifyHeaders"`
		ModifyRequestBody  []ModifyRequestBody  `json:"modifyRequestBody"`
		ModifyResponseBody []ModifyResponseBody `json:"modifyResponseBody"`
	}
	err = json.Unmarshal(data, &database)

	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	f.redirects = database.Redirects
	f.cancels = database.Cancels
	f.delays = database.Delays
	f.modifyHeaders = database.ModifyHeaders
	f.modifyRequestBody = database.ModifyRequestBody
	f.modifyResponseBody = database.ModifyResponseBody
}

func (f *FileDatabase) store() {
	data, err := json.Marshal(map[string]any{
		"redirects":          f.redirects,
		"cancels":            f.cancels,
		"delays":             f.delays,
		"modifyHeaders":      f.modifyHeaders,
		"modifyRequestBody":  f.modifyRequestBody,
		"modifyResponseBody": f.modifyResponseBody,
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
	if recordType == MODIFY_HEADERS {
		interfaceSlice := make([]interface{}, len(f.modifyHeaders))
		for i, v := range f.modifyHeaders {
			interfaceSlice[i] = v
		}
		return interfaceSlice, nil
	}
	if recordType == MODIFY_REQUEST_BODY {
		interfaceSlice := make([]interface{}, len(f.modifyRequestBody))
		for i, v := range f.modifyRequestBody {
			interfaceSlice[i] = v
		}
		return interfaceSlice, nil
	}
	if recordType == MODIFY_RESPONSE_BODY {
		interfaceSlice := make([]interface{}, len(f.modifyResponseBody))
		for i, v := range f.modifyResponseBody {
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
	if recordType == MODIFY_HEADERS {
		if id >= len(f.modifyHeaders) {
			return fmt.Errorf("modify header with id %d not found", id)
		}
		f.modifyHeaders[id] = value.(ModifyHeader)
	}
	if recordType == MODIFY_REQUEST_BODY {
		if id >= len(f.modifyRequestBody) {
			return fmt.Errorf("modify request body with id %d not found", id)
		}
		f.modifyRequestBody[id] = value.(ModifyRequestBody)
	}
	if recordType == MODIFY_RESPONSE_BODY {
		if id >= len(f.modifyResponseBody) {
			return fmt.Errorf("modify response body with id %d not found", id)
		}
		f.modifyResponseBody[id] = value.(ModifyResponseBody)
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

	if recordType == MODIFY_HEADERS {
		fmt.Println("Removing record")
		if id >= len(f.modifyHeaders) {
			return fmt.Errorf("modify header with id %d not found", id)
		}
		f.modifyHeaders = append(f.modifyHeaders[:id], f.modifyHeaders[id+1:]...)
		f.store()
	}
	if recordType == MODIFY_REQUEST_BODY {
		if id >= len(f.modifyRequestBody) {
			return fmt.Errorf("modify request body with id %d not found", id)
		}
		f.modifyRequestBody = append(f.modifyRequestBody[:id], f.modifyRequestBody[id+1:]...)
		f.store()
	}
	if recordType == MODIFY_RESPONSE_BODY {
		if id >= len(f.modifyResponseBody) {
			return fmt.Errorf("modify response body with id %d not found", id)
		}
		f.modifyResponseBody = append(f.modifyResponseBody[:id], f.modifyResponseBody[id+1:]...)
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

	if recordType == MODIFY_HEADERS {
		f.modifyHeaders = append(f.modifyHeaders, value.(ModifyHeader))
		f.store()
	}

	if recordType == MODIFY_REQUEST_BODY {
		f.modifyRequestBody = append(f.modifyRequestBody, value.(ModifyRequestBody))
		f.store()
	}

	if recordType == MODIFY_RESPONSE_BODY {
		f.modifyResponseBody = append(f.modifyResponseBody, value.(ModifyResponseBody))
		f.store()
	}

	return nil
}
