package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming request from ... ", r.Host)
	var tokens = strings.Split(r.Host, ":")
	log.Println("Host: ", tokens[0])

	if ProcessHeaders(w, r) == nil {
		ProcessEndpoint(w, r)
	}
}

var HostName, _ = os.Hostname()

func StartServer(args[] string) {
	if len(args) == 1 {
		fmt.Println("Using default port number: ", DefaultServerPort)
	} else {
		ServerPort = ":" + args[1]
	}

	http.HandleFunc("/", rootHandler)
	//err := http.ListenAndServe(ServerPort, nil)
	err := http.ListenAndServeTLS(":9443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
		os.Exit (10)
	}
}