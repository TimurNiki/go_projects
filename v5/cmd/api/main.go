package main

import (
	"log"
	"v5/internal/env"
)

func main() {

	cfg := config{
		addr: env.GetString("ADDR", ":9595"),
		
	}
	app := &application{
		config: cfg,
	}
	mux := app.mount()

	log.Fatal(app.run(mux))
}
