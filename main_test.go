package main

import "testing"

func TestHelloFunc(t *testing.T) {
	t.Log("Chekcing trhat function prints what's expected...")
	result := GetHello()
	if result != "Hello World" {
		t.Error("Expected 'Hello World' but got", result)
	}
}