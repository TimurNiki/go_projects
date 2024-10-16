package api

import (
	"database/sql"
	"log"
	"net/http"
	"v4/services/user"

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

	// create subrouter
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// create user store
	userStore := user.NewStore(s.db)
	
	// create user handler
	userHandler := user.NewHandler(userStore)
	// register routes
	userHandler.RegisterRoutes(subrouter)

	// productStore:=product.NewStore(s.db)
	// productHandler:=product.NewHandler(productStore)
	// productHandler.RegisterRoutes(subrouter)
	
	log.Printf("Listening on port %s", s.addr)
	return http.ListenAndServe(s.addr, router)
}
