package main

import (
	"github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"strings"
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
	is.True(strings.Contains(w.Body.String(), "status: available"))

}
