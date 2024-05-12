package main

import (
	"log"
	"net/http"
	
)

func main() {
	// create a new serve mux
	mux := http.NewServeMux()
	// register the handler
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	log.Println("Starting server on port 5000")
	err := http.ListenAndServe(":5000", mux)
	log.Fatal(err)
}

