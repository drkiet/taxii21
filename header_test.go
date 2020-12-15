package main

import (
	"log"
	"net/http"
	"testing"
)

func TestVerifyAcceptGood(t *testing.T) {
	log.Println("VerifyAccept test")
	r := http.Request{}
	r.Header = http.Header{}
	r.Header.Set("Accept","application/taxii+json")
	acceptable, err := VerifyAccept(&r)
	if err != nil {
		t.Error(err)
	}
	if !acceptable {
		t.Error("Bad Accept")
	}
	log.Println("VerifyAccept successful")
}

func TestVerifyAcceptBad(t *testing.T) {
	log.Println("Negative VerifyAccept test")
	r := http.Request{}
	r.Header = http.Header{}
	r.Header.Set("Accept","application/json")
	acceptable, err := VerifyAccept(&r)
	if err == nil {
		t.Error(err)
	}
	if acceptable {
		t.Error("Unexpected Accept")
	}
	log.Println("Negative VerifyAccept successful")
}
