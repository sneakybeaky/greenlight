package main

import (
	"encoding/json"
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

	type statusResponse struct {
		Environment string `json:"environment"`
		Status      string `json:"status"`
		Version     string `json:"version"`
	}

	want := statusResponse{
		Environment: "development",
		Status:      "available",
		Version:     "1.0.0",
	}

	got := statusResponse{}
	err = json.Unmarshal(w.Body.Bytes(), &got)
	is.NoErr(err)

	is.Equal(want, got)

}
