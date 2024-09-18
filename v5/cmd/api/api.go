package main

import (
	// "errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

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

	return srv.ListenAndServe()
	// err := srv.ListenAndServe()
	// if !errors.Is(err, http.ErrServerClosed) {
	// 	return err
	// }

	// return nil
}