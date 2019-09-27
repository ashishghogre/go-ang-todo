package main

import (
	"encoding/json"
	"net/http"
)

func handleError(err error, w http.ResponseWriter, statusCode int, errorMessage string) bool {
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(errorMessage)
		return true
	}
	return false
}
