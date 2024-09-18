package main

import "log"

func main() {

	cfg := config{
		addr: ":9595",
		// addr: env.GetString("ADDR", ":9595"),
	}
	app := &application{
		config: cfg,
	}
	mux := app.mount()

	log.Fatal(app.run(mux))
}
