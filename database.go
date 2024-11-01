package main

import "net/http"

const (
	REDIRECT             = "redirect"
	CANCEL               = "cancel"
	DELAY                = "delay"
	MODIFY_HEADERS       = "modifyHeader"
	MODIFY_REQUEST_BODY  = "modifyRequestBody"
	MODIFY_RESPONSE_BODY = "modifyResponseBody"
)

type Database interface {
	GetManyRules(recordType string) ([]Rule, error)
	AddRule(rule Rule) (any, error)
	RemoveRule(id any) error
	UpdateRule(id any, rule Rule) (Rule, error)
	GetOneRule(id any) (Rule, error)
	// Requests CRUD
	AddRequest(request HttpRequestLog) (any, error)
	UpdateRequest(id string, request HttpRequestLog) (HttpRequestLog, error)
	GetOneRequest(id string) (HttpRequestLog, error)
	GetManyRequests() ([]HttpRequestLog, error)
	RemoveRequest(id any) error
	SetRequestHeaders(id any, headers http.Header, isRequest bool) error
	SetRequestBody(id any, body string, isRequest bool) error
	SetRequestCancelled(id any, value bool) error
	SetRequestDelayed(id any, value int) error
	SetRequestRedirected(id any, value bool) error
	SetRequestHeaderModified(id any, value bool) error
	SetResponseHeaderModified(id any, value bool) error
	SetRequestBodyModified(id any, value bool) error
	SetResponseBodyModified(id any, value bool) error
}
