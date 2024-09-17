package main

import (
	// "errors"
	"net/http"

	"github.com/gorilla/mux"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) run() error {

	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    app.config.addr,
		Handler: mux,
	}

	return srv.ListenAndServe()
	// err := srv.ListenAndServe()
	// if !errors.Is(err, http.ErrServerClosed) {
	// 	return err
	// }

	// return nil
}
