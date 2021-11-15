package main

import (
	"net/http"
)

func (s *server) healthcheckHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Create a map which holds the information that we want to send in the response.
		data := map[string]string{
			"status":      "available",
			"environment": s.config.env,
			"version":     version,
		}

		err := s.writeJSON(w, http.StatusOK, data, nil)
		if err != nil {
			s.logger.Println(err)
			http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		}
	}
}
