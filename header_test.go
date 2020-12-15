package main

import (
	"encoding/base64"
	"github.com/drkiet/accounts"
	"log"
	"net/http"
	"net/http/httptest"
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

func TestVerifyBasicAuth(t *testing.T) {
	log.Println("VerifyBasicAuth test")
	r := http.Request{}
	r.Header = http.Header{}
	r.Header.Set("Authorization", "Basic dXNlcjE6UGFzc3dvcmQxMjM0IUAjJA==")
	status, err := VerifyBasicAuth(&r)
	if err != nil {
		t.Error(err)
	}
	if status != http.StatusOK {
		t.Error(status)
	}
	log.Println("VerifyBasicAuth successful")
}

func TestVerifyBasicAuthStatusUnauthorized(t *testing.T) {
	log.Println("VerifyBasicAuth StatusUnauthorized test")
	r := http.Request{}
	r.Header = http.Header{}
	status, err := VerifyBasicAuth(&r)
	if err == nil {
		t.Error(err)
	}
	if status != http.StatusUnauthorized {
		t.Error(status)
	}
	log.Println("VerifyBasicAuth StatusUnauthorized successful")
}

func TestVerifyBasicAuthStatusForbidden(t *testing.T) {
	log.Println("VerifyBasicAuth StatusForbidden test")
	r := http.Request{}
	r.Header = http.Header{}
	r.Header.Set("Authorization", "Basic dXNlcjE6UGFzc3dvcmQxMjM0IUAjNA==")
	status, err := VerifyBasicAuth(&r)
	if err == nil {
		t.Error(err)
	}
	if status != http.StatusForbidden {
		t.Error(status)
	}
	log.Println("VerifyBasicAuth StatusForbidden successful")
}

func makeHttpBasicAuth(user, password string) string {
	CreateUser(user, password)
	var base64UserPassword = "Basic " + base64.StdEncoding.EncodeToString([]byte(user+":"+password))
	log.Println("base64UserPassword:", base64UserPassword)
	return base64UserPassword
}
func makeHttpBasicAuthNonExistentUser(user, password string) string {
	var base64UserPassword = "Basic " + base64.StdEncoding.EncodeToString([]byte(user+":"+password))
	log.Println("base64UserPassword:", base64UserPassword)
	return base64UserPassword
}

func TestProcessHeaders(t *testing.T) {
	log.Println("test process headers")
	accounts.Config([]string{"localhost:2379", "localhost:2380"}, 5)

	w := httptest.NewRecorder()
	r := http.Request{}
	r.Header = http.Header{}
	r.Header.Set("Authorization", makeHttpBasicAuth("test_user", "test_password"))
	r.Header.Set("Accept","application/taxii+json")

	err := ProcessHeaders(w, &r)
	log.Println(err)
	if err != nil {
		t.Error("test process headers fails")
	}
	log.Println(w.Header())
	if DefaultMediaTypeVersion != w.Header().Get("Accept") ||
		DefaultMediaTypeVersion != w.Header().Get("Content-Type") {
		t.Error("test process headers fails bad headers")
	}
	log.Println("test process headers ends")
}

func TestProcessHeadersInvalidMediaType(t *testing.T) {
	log.Println("test process headers invalid media type")
	w := httptest.NewRecorder()
	r := http.Request{}
	r.Header = http.Header{}
	r.Header.Set("Accept","application/json")
	err := ProcessHeaders(w, &r)
	log.Println(err)
	if err == nil {
		t.Error("test process headers invalid media type fails")
	}

	log.Println("test process headers invalid media type ends")
}

func TestProcessHeadersUnauthorized(t *testing.T) {
	log.Println("test process headers unauthorized")
	w := httptest.NewRecorder()
	r := http.Request{}
	r.Header = http.Header{}
	r.Header.Set("Accept","application/taxii+json")
	err := ProcessHeaders(w, &r)
	log.Println(err)
	if err == nil {
		t.Error("test process headers unauthorized fails")
	}

	log.Println("test process headers unauthorized ends")
}

func TestProcessHeadersForbidden(t *testing.T) {
	log.Println("test process headers forbidden")
	w := httptest.NewRecorder()
	r := http.Request{}
	r.Header = http.Header{}
	r.Header.Set("Accept","application/taxii+json")
	r.Header.Set("Authorization", makeHttpBasicAuthNonExistentUser("test_userxyz", "test_password"))
	err := ProcessHeaders(w, &r)
	log.Println(err)
	if err == nil {
		t.Error("test process headers forbidden fails")
	}

	log.Println("test process headers forbidden ends")
}