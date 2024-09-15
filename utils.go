package main

import (
	"net"
	"net/http"
	"strings"
	"time"
)

func getRequestEntity(entityName string, r *http.Response) string {
	switch entityName {
	case "url":
		return r.Request.URL.Path
	case "method":
		return r.Request.Method
	case "host":
		return r.Request.Host
	default:
		return ""
	}
}

func evalOp(op string, input string, value string) bool {
	switch op {
	case "equals":
		return input == value
	case "contains":
		return strings.Contains(input, value)
	case "startsWith":
		return strings.HasPrefix(input, value)
	case "endsWith":
		return strings.HasSuffix(input, value)
	default:
		return false
	}
}

func PortAvailable(host string, port string) error {
	timeout := time.Second
	_, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	return err
}
