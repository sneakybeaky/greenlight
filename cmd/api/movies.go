package main

import (
	"fmt"
	"net/http"
)

func (s *server) createMovieHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "create a new movie")
	}
}

func (s *server) showMovieHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := s.readIDParam(r)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		fmt.Fprintf(w, "show the details of movie %d\n", id)

	}
}
