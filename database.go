package main

const (
	REDIRECT                = "redirect"
	CANCEL                  = "cancel"
	DELAY                   = "delay"
	MODIFY_REQUEST_HEADERS  = "modifyRequestHeader"
	MODIFY_RESPONSE_HEADERS = "modifyResponseHeader"
	MODIFY_REQUEST_BODY     = "modifyRequestBody"
	MODIFY_RESPONSE_BODY    = "modifyResponseBody"
)

type Database interface {
	GetManyRules(recordType string) ([]Rule, error)
	AddRule(rule Rule) (any, error)
	RemoveRule(id any) error
	UpdateRule(id any, rule Rule) (Rule, error)
	GetOneRule(id any) (Rule, error)
	// Requests CRUD
	AddRequest(requestId any, request *HttpRequestLog) (any, error)
	GetManyRequests(string, []string, []string, []string, []string, int) ([]HttpRequestLog, error)
	GetOneRequest(id string) (HttpRequestLog, error)
	RemoveRequest(id any) error
}
