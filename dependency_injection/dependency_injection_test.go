package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Samuel")

	result := buffer.String()
	expected := "Hello, Samuel"

	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}
