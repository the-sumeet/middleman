package main

import (
	"net"
	"os"
	"strings"
	"time"
)

func getRequestEntity(entityName string, url, method, host string) string {
	switch entityName {
	case "url":
		return url
	case "method":
		return method
	case "host":
		return host
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
