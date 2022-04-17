package controllers

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	if output := Hello("World"); output != "Hello World" {
		t.Errorf("TestHello FAILED: Expected %q, received: %q", "Hello World", output)
	} else {
		t.Logf("TestHello PASSED: Expected %q, received: %q", "Hello World", output)
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
			t.Errorf("TestTableHello FAILED: input: %q, expected %q, received: %q", test.input, test.expected, output)
		} else {
			t.Logf("TestTableHello PASSED: input: %q, expected %q, received: %q", test.input, test.expected, output)
		}
	}
}

func TestGetData(t *testing.T) {

	if rs, err := GetData(); err != nil {
		t.Errorf("FAILED: %s", err)
	} else {
		t.Logf("SUCCESS")
		for _, row := range rs {
			fmt.Printf("Row: %s\n", row)
		}
	}
}
