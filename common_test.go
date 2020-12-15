package main

import (
	"log"
	"testing"
)

func TestMarshalError(t *testing.T) {
	log.Println("test Marshal error")
	var error = Error{Title: "Test error message"}
	var expectedJsonError = "{\n  \"title\": \"Test error message\"\n}"
	var actualJsonError = MarshalError(error)
	log.Println(actualJsonError)

	if expectedJsonError != actualJsonError {
		t.Error("MarshalError fails")
	}
	log.Println("test Marshal error ends")
}
