package main

import (
	"log"
	"strings"
	"testing"
)

var TestApiRoots = []string{
"/test_api1/",
"/test_trustgroup1/",
"/test_trustgroup2/",
}

func TestProcessDiscovery(t *testing.T) {
	log.Println("test process discovery")
	ApiRoots = TestApiRoots
	var dr = ProcessDiscovery()
	log.Println(dr)
	if dr.Title != "TAXII 2.1 Server" {
		t.Error("test process discovery fails")
	}
	log.Println("test process discovery ends")
	for _, url := range dr.ApiRoots {
		if !StringInSlice(url, TestApiRoots) {
			t.Error("test process discovery unexpected url fails", url)
		}
	}
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		log.Println(a, b)
		if strings.Contains(a, b) {
			return true
		}
	}
	return false
}