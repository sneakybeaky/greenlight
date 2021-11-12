package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGETHealthcheck(t *testing.T) {

	app := Application{Config: Config{
		Env: "test",
	}}

	t.Run("Get status", func(t *testing.T) {

		request, _ := http.NewRequest(http.MethodGet, "/v1/healthcheck", nil)
		response := httptest.NewRecorder()

		app.healthcheckHandler(response, request)

		got := response.Body.String()
		want := "status: available"

		if !strings.Contains(got, want) {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
