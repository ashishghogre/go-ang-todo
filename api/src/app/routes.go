package main

import "github.com/gorilla/mux"

func setupRoutes(router *mux.Router) {
	router.HandleFunc("/item", createItem).Methods("POST")
	router.HandleFunc("/item/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/item/{id}", getItem).Methods("GET")
	router.HandleFunc("/items", getItems).Methods("GET")
}