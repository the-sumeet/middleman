package main

import (
	"testing"
)

func TestEvaluateExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected Expression
	}{
		{"a = b", Expression{Op: "", Value: "true", Type: STRING}},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := parse(test.input)

			if result == test.expected {
				t.Error("Expected:", test.expected)
			}

		})
	}
}
