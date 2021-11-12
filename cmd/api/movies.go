package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (s *server) createMovieHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "create a new movie")
	}
}

func (s *server) showMovieHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := httprouter.ParamsFromContext(r.Context())

		// We can then use the ByName() method to get the value of the "id" parameter from
		// the slice. In our project all movies will have a unique positive integer ID, but
		// the value returned by ByName() is always a string. So we try to convert it to a
		// base 10 integer (with a bit size of 64). If the parameter couldn't be converted,
		// or is less than 1, we know the ID is invalid so we use the http.NotFound()
		// function to return a 404 Not Found response.

		id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
		if err != nil || id < 1 {
			http.NotFound(w, r)
			return
		}

		// Otherwise, interpolate the movie ID in a placeholder response.
		fmt.Fprintf(w, "show the details of movie %d\n", id)

	}
}
