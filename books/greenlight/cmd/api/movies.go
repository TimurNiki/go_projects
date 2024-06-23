package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/TimurNiki/go_api_tutorial/books/greenlight/internal/data"
	"github.com/julienschmidt/httprouter"
)

// Add a createMovieHandler for the "POST /v1/movies" endpoint. For now we simply
// return a plain-text placeholder response.
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

// Add a showMovieHandler for the "GET /v1/movies/:id" endpoint. For now, we retrieve
// the interpolated "id" parameter from the current URL and include it in a placeholder
// response.
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	
	id, err :=app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	// Otherwise, interpolate the movie ID in a placeholder response.
	fmt.Fprintf(w, "show the details of movie %d\n", id)

	movie:=data.Movie{
		ID:id,
		CreatedAt:time.Now(),
		Title:"Casablanca",
		Runtime:102,
		Genres:[]string{"drama","romance", "war"},
		Version:1,
	}

	err:=app.writeJSON(w,http.StatusOK,movie,nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered an error", http.StatusInternalServerError)
	}
}
