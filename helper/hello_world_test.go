package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Hilmi")
	if result != "Hello Hilmi" {
		// error
		panic("Result is not 'Hello Hilmi'")
	}
}

func TestHelloWorldYahya(t *testing.T) {
	result := HelloWorld("Yahya")
	if result != "Hello Yahya" {
		// error
		panic("Result is not 'Hello Yahya'")
	}
}