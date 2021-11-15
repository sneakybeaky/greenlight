package main

import (
	"github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETHealthcheck(t *testing.T) {

	is := is.New(t)

	srv, err := newServer()
	is.NoErr(err)

	req := httptest.NewRequest(http.MethodGet, "/v1/healthcheck", nil)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	is.Equal(w.Code, http.StatusOK)
	is.Equal(w.Body.String(), "{\"status\": \"available\", \"environment\": \"development\", \"version\": \"1.0.0\"}")

}
