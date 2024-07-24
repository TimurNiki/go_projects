package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/TimurNiki/go_api_tutorial/v2/store"

)

type APIServer struct {
	addr  string
	store store.Store
}

func newAPIServer(addr string, store store.Store) *APIServer {
	return &APIServer{addr, store}
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	tasksService := NewTasksService(s.store)
	tasksService.RegisterRoutes(subrouter)

	log.Println("Starting server on port", s.addr)

	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
