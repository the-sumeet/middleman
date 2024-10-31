package main

import (
	"net/http"
	"time"
)

type HttpRequestLog struct {
	Timestamp              time.Time   `json:"timestamp"`
	Scheme                 string      `json:"scheme"`
	Method                 string      `json:"method"`
	Host                   string      `json:"host"`
	Path                   string      `json:"path"`
	RequestHeaders         http.Header `json:"requestHeaders"`
	ResponseHeaders        http.Header `json:"responseHeaders"`
	RequestBody            string      `json:"requestBody"`
	ResponseBody           string      `json:"responseBody"`
	Cancelled              bool        `json:"cancelled"`
	Redirected             bool        `json:"redirected"`
	RequestHeaderModified  bool        `json:"requestHeaderModified"`
	ResponseHeaderModified bool        `json:"responseHeaderModified"`
	RequestBodyModified    bool        `json:"requestBodyModified"`
	ResponseBodyModified   bool        `json:"responseBodyModified"`
	Delayed                int         `json:"delayed"`
}
