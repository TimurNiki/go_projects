package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type APIServer struct {
	addr  string
	store Store
}

func newAPIServer(addr string, store Store) *APIServer {
	return &APIServer{addr, store}
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	log.Println("Starting server on port", s.addr)

	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
