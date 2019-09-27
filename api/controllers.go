package main

import (
	"encoding/json"
	"net/http"
)

var items []todoItem


func createItem(w http.ResponseWriter, r *http.Request) {
	var item todoItem
	if handleError(json.NewDecoder(r.Body).Decode(&item), w,
		400, "Invalid json format") {
		return
	}
	items = append(items, item)

	dbClient.upsertItem("1",item)
	w.WriteHeader(201)
}

func getItems(w http.ResponseWriter, r *http.Request) {
	/*if handleError(json.NewEncoder(w).Encode(items), w,
		500, "Internal server error") {
		return
	}*/if handleError(json.NewEncoder(w).Encode(dbClient.getItem("1")), w,
		500, "Internal server error") {
		return
	}
}
