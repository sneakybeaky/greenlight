package main

import (
	"net/http"
)

func (s *server) routes() {
	s.router.HandlerFunc(http.MethodGet, "/v1/healthcheck", s.healthcheckHandler())
	s.router.HandlerFunc(http.MethodPost, "/v1/movies", s.createMovieHandler())
	s.router.HandlerFunc(http.MethodGet, "/v1/movies/:id", s.showMovieHandler())
}
