package main

const (
	REDIRECT             = "redirect"
	CANCEL               = "cancel"
	DELAY                = "delay"
	MODIFY_HEADERS       = "modifyHeader"
	MODIFY_REQUEST_BODY  = "modifyRequestBody"
	MODIFY_RESPONSE_BODY = "modifyResponseBody"
)

type Header struct {
	Action    string `json:"action"`
	IsRequest bool   `json:"isRequest"`
	Name      string `json:"name"`
	Value     string `json:"value"`
}

type Rule struct {
	Id                 int      `json:"id"`
	Type               string   `json:"type"`
	Enabled            bool     `json:"enabled"`
	Entity             string   `json:"entity"`             // Like url, path, header etc
	Op                 string   `json:"op"`                 // Like contains, equals, etc
	Value              string   `json:"value"`              // The value to match
	RedirectTo         string   `json:"redirectTo"`         // For redirect rule
	DelaySec           int      `json:"delaySec"`           // For delay rule
	RequestBody        string   `json:"requestBody"`        // For modify request body rule
	ResponseBody       string   `json:"responseBody"`       // For modify response body rule
	RequestHeaderMods  []Header `json:"requestHeaderMods"`  // For modify header rule
	ResponseHeaderMods []Header `json:"responseHeaderMods"` // For modify header rule

}

type Database interface {
	GetManyRules(recordType string) ([]Rule, error)
	AddRule(rule Rule) (any, error)
}
