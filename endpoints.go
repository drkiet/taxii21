package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func ProcessEndpoint(w http.ResponseWriter,r *http.Request) {
	path := r.URL.Path
	if path == "/taxii2/" {
		log.Println("returning APIRoots")
		body := ProcessDiscovery()
		w.Header().Set("Content-Type", "application/taxii+json")
		json.NewEncoder(w).Encode(body)
	} else {
		log.Println("Unsupported request ... ", r.URL)
	}
}
