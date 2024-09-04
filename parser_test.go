package main

import (
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func NewRequest(method, url string, body string, headers map[string][]string) http.Request {
	reader := strings.NewReader(body)
	request, err := http.NewRequest(method, url, reader)
	if err != nil {
		panic(err)
	}
	for key, value := range headers {
		for _, v := range value {
			request.Header.Set(key, v)
		}
	}
	return *request
}

func TestEvaluateExpression(t *testing.T) {
	tests := []struct {
		input    string
		request  http.Request
		expected bool
	}{
		{"url contains \"example\"", http.Request{URL: &url.URL{Host: "example.com"}}, true},
		{"url contains \"example\" and url contains \"exam\"", http.Request{URL: &url.URL{Host: "example.com"}}, true},
		{"url contains \"example\" or url contains \"foo\"", http.Request{URL: &url.URL{Host: "example.com"}}, true},
		{"url contains \"foo\" or url contains \"bar\"", http.Request{URL: &url.URL{Host: "example.com"}}, false},
		{"url = \"//example.com\"", http.Request{URL: &url.URL{Host: "example.com"}}, true},
		{"url != \"//example.com\"", http.Request{URL: &url.URL{Host: "example.com"}}, false},
		{"body = \"foo\"", NewRequest("", "", "foo", map[string][]string{}), true},
		{"body = \"bar\"", NewRequest("", "", "foo", map[string][]string{}), false},
		{"body != \"bar\"", NewRequest("", "", "foo", map[string][]string{}), true},
		{"body != \"foo\"", NewRequest("", "", "foo", map[string][]string{}), false},
		{"body contains \"oo\"", NewRequest("", "", "foo", map[string][]string{}), true},
		{"body contains \"bar\"", NewRequest("", "", "foo", map[string][]string{}), false},
		{"header = \"name:value\"", NewRequest("", "", "foo", map[string][]string{"name": {"value"}}), true},
		{"header = \"namevalue\"", NewRequest("", "", "foo", map[string][]string{"name": {"value"}}), false},   // If header input is invalid, it should return false
		{"header = \"foo:bar:baz\"", NewRequest("", "", "foo", map[string][]string{"name": {"value"}}), false}, // If header input is invalid, it should return false
		{"header = \"foo:bar:baz\"", NewRequest("", "", "foo", map[string][]string{"name": {"value"}}), false}, // If header input is invalid, it should return false
		{"header = \"name:foo\"", NewRequest("", "", "foo", map[string][]string{"name": {"value"}}), false},
		{"header = \"name:foo\"", NewRequest("", "", "foo", map[string][]string{"name": {"value", "foo"}}), true},
		{"header = \"name:foo\"", NewRequest("", "", "foo", map[string][]string{"name": {"value", "value2"}}), false},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := parse(test.input, test.request)

			if result != test.expected {
				t.Error("Expected:", test.expected)
			}

		})
	}
}
