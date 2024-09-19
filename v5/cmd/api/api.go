package main

import (
	// "errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type application struct {
	config config
}

type config struct {
	addr string
}

type authConfig struct {
	basic basicConfig
	token tokenConfig
}

type tokenConfig struct {
	secret string
	exp    time.Duration
	iss    string
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}
type basicConfig struct {
	user string
	pass string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)

	})

	return r
}

func (app *application) run(mux http.Handler) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	jnjnj
	shutdown := make(chan error)

	return srv.ListenAndServe()
	// err := srv.ListenAndServe()
	// if !errors.Is(err, http.ErrServerClosed) {
	// 	return err
	// }

	// return nil
}
