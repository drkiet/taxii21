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
		var encoder = json.NewEncoder(w)
		encoder.SetIndent("", "  ")
		encoder.Encode(ProcessDiscovery())
	} else {
		log.Println("Unsupported request ... ", r.URL)
	}
}
