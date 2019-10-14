package main

import (
	"encoding/json"
	"net/http"
)

func handleError(err error, w http.ResponseWriter, statusCode int, errorMessage string) bool {
	if err != nil {
		setResponseHeaders(w, statusCode)
		json.NewEncoder(w).Encode(errorMessage)
		return true
	}
	return false
}

func setResponseHeaders(w http.ResponseWriter, statusCode int){
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Header", "*")
	w.Header().Add("Access-Control-Allow-Method", "*")
	w.WriteHeader(statusCode)
}