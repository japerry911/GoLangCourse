package main

import "testing"

func TestReturnValue(t *testing.T) {
	hello := helloWorld()

	if hello != "Hello World" {
		t.Errorf("Expected 'Hello World' but received %s.", hello)
	}
}
