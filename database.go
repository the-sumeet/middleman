package main

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
}
