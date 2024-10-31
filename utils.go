package main

import (
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func getRequestEntity(entityName string, r *http.Request) string {
	switch entityName {
	case "path":
		return r.URL.Path
	case "method":
		return r.Method
	case "host":
		return r.Host
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

func certKeyPresent() bool {
	certPath, keyPath := getCertKeyPath()
	// Check if files exists
	if _, err := os.Stat(certPath); os.IsNotExist(err) {
		return false
	}
	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		return false
	}
	return true
}

func getRequestLogValues(r *http.Request, args ...any) []any {
	res := []any{
		"host", r.Host,
		"path", r.URL.Path,
		"method", r.Method,
	}
	res = append(res, args...)
	return res
}

func getResponseLogValues(r *http.Response, args ...any) []any {
	return getRequestLogValues(r.Request, args...)
}

func matches(r Rule, hr *http.Request) bool {
	entityValue := getRequestEntity(r.Entity, hr)
	if entityValue == "" || r.Value == "" {
		return false
	}
	return evalOp(r.Op, entityValue, r.Value)
}
