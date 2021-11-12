package main

import (
	"fmt"
	"net/http"
)

func (s *server) healthcheckHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "status: available")
	}
}
