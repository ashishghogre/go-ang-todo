package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var items []todoItem

func createItem(w http.ResponseWriter, r *http.Request) {
	var item todoItem
	if handleError(json.NewDecoder(r.Body).Decode(&item), w,
		400, "Invalid json format") {
		return
	}
	items = append(items, item)

	id := createItemInDb(item)
	var res = createResponse{}
	res.Id = id
	json.NewEncoder(w).Encode(res)
	setResponseHeaders(w, 201)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	var item todoItem
	id := mux.Vars(r)["id"]
	if handleError(json.NewDecoder(r.Body).Decode(&item), w,
		400, "Invalid json format") {
		return
	}
	items = append(items, item)
	updateItemInDb(id, item)
	setResponseHeaders(w, 204)
}

func getItems(w http.ResponseWriter, r *http.Request) {
	setResponseHeaders(w, 200)
	err := json.NewEncoder(w).Encode(getItemsFromDb())
	fmt.Println("here")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("xyz", "application/json")
	fmt.Println(w.Header())
	if handleError(err, w, 500, "Internal server error") {
		return
	}
}

func getItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	item := getItemFromDb(id)
	/*	if nil == item {
		setResponseHeaders(w, 404)
		return
	}*/
	setResponseHeaders(w, 200)
	if handleError(json.NewEncoder(w).Encode(item), w,
		500, "Internal server error") {
		return
	}
}
