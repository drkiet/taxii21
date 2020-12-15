package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestProcessEndpoint(t *testing.T) {
	log.Println("test process endpoint")
	ApiRoots = TestApiRoots
	var expectedBody = "{\n  \"title\": \"TAXII 2.1 Server\",\n  \"description\": \"TAXII 2.1 Server in GoLang\",\n  \"contact\": \"Kiet T. Tran, Ph.D.\",\n  \"default\": \"https://student-VirtualBox/test_api1/\",\n  \"api_roots\": [\n    \"https://student-VirtualBox/test_api1/\",\n    \"https://student-VirtualBox/test_trustgroup1/\",\n    \"https://student-VirtualBox/test_trustgroup2/\"\n  ]\n}\n"
	w := httptest.NewRecorder()
	r := http.Request{}
	r.Header = http.Header{}
	r.URL, _ = url.Parse("/taxii2/")
	ProcessEndpoint(w, &r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var actualBody = string(body)

	log.Println("'" + actualBody+ "'")
	log.Println("'" + expectedBody+ "'")
	if actualBody != expectedBody {
		t.Error("test process endpoint fails")
	}
	log.Println("test process endpoint ends")
}

func TestProcessEndpointInvalidPath(t *testing.T) {
	log.Println("test process endpoint invalid path")
	ApiRoots = TestApiRoots
	w := httptest.NewRecorder()
	r := http.Request{}
	r.Header = http.Header{}
	r.URL, _ = url.Parse("/invalidpath/")
	ProcessEndpoint(w, &r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	if len(body) > 0 {
		t.Error("test process endpoint invalid path not empty fails")
	}
	log.Println("test process endpoint invalid path ends")
}