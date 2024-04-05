package api

import (
	"database/sql"
	"log"
	"net/http"
	"github.com/TimurNiki/go_api_tutorial/v4/services/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() error {
	// create router
	router := mux.NewRouter()
	// subrouter for users
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	// create user handler
	userHandler := user.NewHandler()
	// register routes
	userHandler.RegisterRoutes(subrouter)
	// start server on port
	log.Printf("Listening on port %s", s.addr)
	return http.ListenAndServe(s.addr, router)
}
