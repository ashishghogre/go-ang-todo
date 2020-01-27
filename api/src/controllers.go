package main

import (
	"fmt"
	"encoding/json"
	"strconv"
	"github.com/sirupsen/logrus"
	"net/http"
	"github.com/gorilla/mux"
)


func createItem(w http.ResponseWriter, r *http.Request) {
	var item todoItem
	if handleError(json.NewDecoder(r.Body).Decode(&item), w,
		400, "Invalid json format") {
		return
	}
	logrus.Infof("Creating item %s", item.GoString())
	id := createItemInDb(item)
	var res = createResponse{}
	res.ID = id
	setResponseHeaders(w, 201)
	logrus.Infof("Item with ID: %d created successfully", id)
	defer json.NewEncoder(w).Encode(res)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	var item todoItem
	id := mux.Vars(r)["id"]
	if handleError(json.NewDecoder(r.Body).Decode(&item), w,
		400, "Invalid json format") {
		return
	}
	logrus.Infof("Updating item with ID: %s", id)
	item.CreatedTime = getItemFromDb(id).CreatedTime
	item.ID, _ = strconv.Atoi(id)
	updateItemInDb(id, item)
	setResponseHeaders(w, 204)
	logrus.Infof("Item with ID: %s successfully updated", id)
}

func getItems(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Getting all items")
	setResponseHeaders(w, 200)
	err := json.NewEncoder(w).Encode(getItemsFromDb())
	if handleError(err, w, 500, "Internal server error") {
		return
	}
	logrus.Info("Returning all items")
}

func getItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	logrus.Infof("Get item for ID %s", id)
	item := getItemFromDb(id)
	if item.ID == 0 {
		setResponseHeaders(w, 404)
		json.NewEncoder(w).Encode(fmt.Sprintf("Item not found for id: %s", id))
		logrus.Warnf("Item not found for id: %s", id)
		return
	}
	setResponseHeaders(w, 200)
	if handleError(json.NewEncoder(w).Encode(item), w,
		500, "Internal server error") {
		return
	}
	logrus.Infof("Returning item for ID %s {%s}", id,item.GoString())
}
