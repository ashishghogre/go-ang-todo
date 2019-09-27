package main

import "github.com/gorilla/mux"

func setupRoutes(router *mux.Router) {
	router.HandleFunc("/item", createItem).Methods("POST")
	router.HandleFunc("/items", getItems).Methods("GET")
}