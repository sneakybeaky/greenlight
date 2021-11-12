package main

import (
	"github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShowMovies(t *testing.T) {

	is := is.New(t)

	srv, err := newServer()
	is.NoErr(err)

	t.Run("Happy path", func(t *testing.T) {
		is := is.New(t)

		req := httptest.NewRequest(http.MethodGet, "/v1/movies/1", nil)
		w := httptest.NewRecorder()
		srv.ServeHTTP(w, req)
		is.Equal(w.Code, http.StatusOK)
		is.True(strings.Contains(w.Body.String(), "movie 1"))
	})

	t.Run("Id isn't numeric", func(t *testing.T) {
		is := is.New(t)

		req := httptest.NewRequest(http.MethodGet, "/v1/movies/blah", nil)
		w := httptest.NewRecorder()
		srv.ServeHTTP(w, req)
		is.Equal(w.Code, http.StatusNotFound)
	})

	t.Run("Id less than 1", func(t *testing.T) {
		is := is.New(t)

		req := httptest.NewRequest(http.MethodGet, "/v1/movies/0", nil)
		w := httptest.NewRecorder()
		srv.ServeHTTP(w, req)
		is.Equal(w.Code, http.StatusNotFound)
	})

}
