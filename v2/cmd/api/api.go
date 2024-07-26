package api

import (
	"log"
	"net/http"
	"v2/store"
"v2/tasks"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr  string
	store store.Store
}

func NewAPIServer(addr string, store store.Store) *APIServer {
	return &APIServer{addr, store}
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	tasksService := tasks.NewTasksService(s.store)
	tasksService.RegisterRoutes(subrouter)

	log.Println("Starting server on port", s.addr)

	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
