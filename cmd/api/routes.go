package main

import (
	"net/http"
)

func (s *server) routes() {
	s.router.HandlerFunc(http.MethodGet, "/v1/healthcheck", s.healthcheckHandler())
}
