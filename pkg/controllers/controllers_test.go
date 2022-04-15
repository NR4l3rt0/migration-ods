package controllers

import "testing"

func TestHello(t *testing.T) {
	if Hello("World") != "Hello World" {
		t.Error("Method not returns appropriate value for input: {}. Expected {}; received: {}", "World", "Hello World")
	}
}

func TestTableHello(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{

		{"World", "Hello World"},
		{"Paul", "Hello Paul"},
		{"Cindy", "Hello Cindy"},
		{"Tim", "Hello Tim"},
	}

	for _, test := range tests {
		if output := Hello(test.input); output != test.expected {
			t.Error("Method not returns appropriate value for input: {}. Expected {}; received: {}", test.input, test.expected, output)
		}
	}
}
