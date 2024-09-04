package main

import (
	"errors"
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
		errors   []error
	}{
		{"url contains \"example\"", http.Request{URL: &url.URL{Host: "example.com"}}, true, nil},
		{"url contains \"example\" and url contains \"exam\"", http.Request{URL: &url.URL{Host: "example.com"}}, true, nil},
		{"url contains \"example\" or url contains \"foo\"", http.Request{URL: &url.URL{Host: "example.com"}}, true, nil},
		{"url contains \"foo\" or url contains \"bar\"", http.Request{URL: &url.URL{Host: "example.com"}}, false, nil},
		{"url = \"//example.com\"", http.Request{URL: &url.URL{Host: "example.com"}}, true, nil},
		{"url != \"//example.com\"", http.Request{URL: &url.URL{Host: "example.com"}}, false, nil},
		{"body = \"foo\"", NewRequest("", "", "foo", map[string][]string{}), true, nil},
		{"body = \"bar\"", NewRequest("", "", "foo", map[string][]string{}), false, nil},
		{"body != \"bar\"", NewRequest("", "", "foo", map[string][]string{}), true, nil},
		{"body != \"foo\"", NewRequest("", "", "foo", map[string][]string{}), false, nil},
		{"body contains \"oo\"", NewRequest("", "", "foo", map[string][]string{}), true, nil},
		{"body contains \"bar\"", NewRequest("", "", "foo", map[string][]string{}), false, nil},
		{"header = \"name:value\"", NewRequest("", "", "foo", map[string][]string{"name": {"value"}}), true, nil},
		{"header = \"namevalue\"", NewRequest("", "", "foo", map[string][]string{"name": {"value"}}), false, nil},   // If header input is invalid, it should return false
		{"header = \"foo:bar:baz\"", NewRequest("", "", "foo", map[string][]string{"name": {"value"}}), false, nil}, // If header input is invalid, it should return false
		{"header = \"foo:bar:baz\"", NewRequest("", "", "foo", map[string][]string{"name": {"value"}}), false, nil}, // If header input is invalid, it should return false
		{"header = \"name:foo\"", NewRequest("", "", "foo", map[string][]string{"name": {"value"}}), false, nil},
		{"header = \"name:foo\"", NewRequest("", "", "foo", map[string][]string{"name": {"value", "foo"}}), true, nil},
		{"header = \"name:foo\"", NewRequest("", "", "foo", map[string][]string{"name": {"value", "value2"}}), false, nil},
		{"ksfmoi sduf sdf 'sef'", NewRequest("", "", "foo", map[string][]string{"name": {"value", "value2"}}), false, []error{errors.New("line 1:7 no viable alternative at input 'ksfmoisduf'")}},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result, parserErrors := parse(test.input, test.request)

			if len(parserErrors) == len(test.errors) {
				if test.errors == nil {
					t.Error("Expected no errors")
				} else if len(test.errors) != len(parserErrors) {
					t.Error("Expected:", test.errors)
				} else {
					for i, err := range parserErrors {
						if err.Error() != test.errors[i].Error() {
							t.Error("Expected:", test.errors[i])
						}
					}
				}
			} else {
				if result != test.expected {
					t.Error("Expected:", test.expected)
				}
			}

		})
	}
}
