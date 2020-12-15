package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

const DefaultMediaTypeVersion = "application/taxii+json;version=2.1"
const DefaultMediaTypePlain =   "application/taxii+json"

var (
	MediaTypes = map[string]int {
		DefaultMediaTypePlain:   1,
		DefaultMediaTypeVersion: 1,
	}
)

func VerifyAccept(r *http.Request) (bool, error) {
	_, ok := MediaTypes[r.Header.Get("Accept")]
	if ok {
		return true, nil
	}
	return false, errors.New("invalid accept header: type/subtype")
}

func ProcessHeaders(w http.ResponseWriter,r *http.Request) error {
	log.Println("Processing headers ...")
	w.Header().Set("content-type",DefaultMediaTypeVersion)
	w.Header().Set("accept", DefaultMediaTypeVersion)

	result, err := VerifyAccept(r)
	log.Println("verify accepts: ", result, err)
	if !(err == nil && result) {
		JSONError(w, makeHeaderError("The media type provided in the Accept header is invalid"), 406)
	} else {
		var status int
		status, err = VerifyBasicAuth(r)
		switch status {
		case http.StatusInternalServerError:
			log.Println("Internal error:", err)
			JSONError(w, makeHeaderError("Internal Server Error"), 500)
		case http.StatusUnauthorized:
			log.Println("Unauthorized:", err)
			w.Header().Set("WWW-Authenticate", "Basic realm=\"simple\"")
			JSONError(w, makeHeaderError("The client needs to authenticate"), 401)
		case http.StatusForbidden:
			log.Println("Forbidden:", err)
			JSONError(w, makeHeaderError("The client does not have access to this resource"), 403)
		case http.StatusOK:
			return nil
		}
	}

	log.Println("Process headers ends ...", err)
	return err
}

func makeHeaderError(errorMessage string) Error {
	return Error{Title: errorMessage}
}

func VerifyBasicAuth(r *http.Request) (int, error) {
	user, psw, ok := r.BasicAuth()
	log.Println("basic: ", user, psw)
	if !ok || len(user) == 0 {
		return http.StatusUnauthorized, errors.New("no authentication")
	}

	if !VerifyUser(user, psw) {
		return http.StatusForbidden, errors.New("unauthorized")
	}
	return http.StatusOK, nil
}

func JSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", DefaultMediaTypeVersion)
	w.WriteHeader(code)

	var encoder = json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	encoder.Encode(err)
}