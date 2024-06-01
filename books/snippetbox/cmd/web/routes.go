package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	// create a new serve mux
	// mux := http.NewServeMux()

	//Initialize the router
	router := httprouter.New()

	// Create a handler function which wraps our notFound() helper, and then
	// assign it as the custom handler for 404 Not Found responses. You can also
	// set a custom handler for 405 Method Not Allowed responses by setting
	// router.MethodNotAllowed in the same way too.
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Update the pattern for the route for the static files.
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	// Create a new middleware chain containing the middleware specific to our
	// dynamic application routes. For now, this chain will only contain the
	// LoadAndSave session middleware but we'll add more to it later.
	dynamic := alice.New(app.sessionManager.LoadAndSave)

	// And then create the routes using the appropriate methods, patterns and
	// handlers.

	 // Update these routes to use the new dynamic middleware chain followed by
 // the appropriate handler function. Note that because the alice ThenFunc()
 // method returns a http.Handler (rather than a http.HandlerFunc) we also
 // need to switch to registering the route using the router.Handler() method
	router.HandlerFunc(http.MethodGet, "/", dynamic.ThenFunc(app.home))
	router.HandlerFunc(http.MethodGet, "/snippet/view/:id", dynamic.ThenFunc(app.snippetView))
	router.HandlerFunc(http.MethodGet, "/snippet/create", dynamic.ThenFunc(app.snippetCreate))
	router.HandlerFunc(http.MethodPost, "/snippet/create", dynamic.ThenFunc(app.snippetCreatePost))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	// mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Register the other application routes as normal.
	// mux.HandleFunc("/", app.home)
	// mux.HandleFunc("/snippet/view", app.snippetView)
	// mux.HandleFunc("/snippet/create", app.snippetCreate)
	// return mux

	// Pass the servemux as the 'next' parameter to the secureHeaders middleware.
	// Because secureHeaders is just a function, and the function returns a
	// http.Handler we don't need to do anything else.
	// return app.recoverPanic(app.logRequest(secureHeaders(mux)))

	// Create a middleware chain containing our 'standard' middleware
	// which will be used for every request our application receives.
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	// Return the 'standard' middleware chain followed by the servemux.
	return standard.Then(router)
}
