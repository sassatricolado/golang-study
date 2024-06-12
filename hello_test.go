package main

import "testing"

func TestHello(t *testing.T) {
	verifyMessage := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("result: %s, expected: %s", result, expected)
		}
	}
	t.Run("Say hello to all people", func(t *testing.T) {
		result := Hello("Samuel")
		expected := "Hello, Samuel"
		verifyMessage(t, result, expected)
	})

	t.Run("Say hello to the world if the name is empty", func(t *testing.T) {
		result := Hello("")
		expected := "Hello, World"
		verifyMessage(t, result, expected)
	})
}
