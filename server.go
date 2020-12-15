package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming request from ... ", r.Host)
	var tokens = strings.Split(r.Host, ":")
	log.Println("Host: ", tokens[0])

	if ProcessHeaders(w, r) == nil {
		if r.Method == http.MethodGet {
			ProcessEndpoint(w, r)
		} else {
			JSONError(w, makeHeaderError("Invalid method"), 405)
		}
	}
	log.Println("exiting rootHandler ...")
}

var HostName, _ = os.Hostname()

func StartServer(args[] string, rootHandler func(w http.ResponseWriter, r *http.Request)) {
	if len(args) == 1 {
		fmt.Println("Using default port number: ", DefaultServerPort)
	} else {
		ServerPort = ":" + args[1]
	}
	log.Println("Server is listening on", ServerPort)

	http.HandleFunc("/", rootHandler)
	//err := http.ListenAndServe(ServerPort, nil)

	err := http.ListenAndServeTLS(ServerPort, "server.crt", "server.key", nil)
	if err != nil {
		log.Println("ListenAndServeTLS: ", err)
		//os.Exit (10)
	}
	log.Println("TAXII Server exits!")
}
