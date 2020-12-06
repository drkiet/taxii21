package main

import (
	"errors"
	"log"
	"net/http"
)

var (
	MediaTypes = map[string]int {
		"application/taxii+json": 1,
		"application/taxii+json;version=2.1":1,
	}
)

func VerifyAccept(w http.ResponseWriter,r *http.Request) (bool, error) {
	_, ok := MediaTypes[r.Header.Get("Accept")]
	if ok {
		return true, nil
	}
	return false, errors.New("invalid accept header: type/subtype")
}

func ProcessHeaders(w http.ResponseWriter,r *http.Request) (error) {
	result, err := VerifyAccept(w, r)
	log.Println("verify accepts: ", result, err)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	return err
}
