package main

import (
	"github.com/drkiet/accounts"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestRootHandler(t *testing.T) {
	log.Println("test root handler")
	accounts.Config([]string{"localhost:2379", "localhost:2380"}, 5)

	w := httptest.NewRecorder()
	r := http.Request{}
	r.Header = http.Header{}
	r.Header = http.Header{}
	r.Header.Set("Authorization", makeHttpBasicAuth("test_user", "test_password"))
	r.Header.Set("Accept","application/taxii+json")
	r.Method = http.MethodGet
	r.URL, _= url.Parse("/taxii2/")
	r.Host = HostName
	RootHandler(w, &r)
	log.Println(w)
	if w.Result().StatusCode != http.StatusOK {
		t.Error("root handler status fails", w.Result().Status)
	}
	log.Println("test root handler ends")
}

func TestRootHandlerNotGetMethod(t *testing.T) {
	log.Println("test process headers not get")
	accounts.Config([]string{"localhost:2379", "localhost:2380"}, 5)

	w := httptest.NewRecorder()
	r := http.Request{}
	r.Header = http.Header{}
	r.Header = http.Header{}
	r.Header.Set("Authorization", makeHttpBasicAuth("test_user", "test_password"))
	r.Header.Set("Accept","application/taxii+json")
	r.Method = http.MethodPost
	r.URL, _= url.Parse("/taxii2/")
	r.Host = HostName
	RootHandler(w, &r)
	log.Println(w)
	if w.Result().StatusCode != http.StatusMethodNotAllowed {
		t.Error("test root handler status fails", w.Result().Status)
	}
	log.Println("test root headers not get ends")
}

func TestStartServer(t *testing.T) {
	var args = []string {"testprogram", "1"}
	StartServer(args, RootHandlerTest)
}

func RootHandlerTest(w http.ResponseWriter, r *http.Request) {

}