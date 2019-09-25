package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type todoItem struct {
	Title string `json:"title"`
}

var items []todoItem

func handleError(err error, w http.ResponseWriter, statusCode int, errorMessage string) bool {
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(errorMessage)
		return true
	}
	return false
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var item todoItem
	if handleError(json.NewDecoder(r.Body).Decode(&item), w,
		400, "Invalid json format") {
		return
	}
	items = append(items, item)
	w.WriteHeader(201)
}

func getItems(w http.ResponseWriter, r *http.Request) {
	if handleError(json.NewEncoder(w).Encode(items), w,
		500, "Internal server error") {
		return
	}
}

func startServer(port int) {
	go func() {
		router := mux.NewRouter()
		router.HandleFunc("/item", createItem).Methods("POST")
		router.HandleFunc("/items", getItems).Methods("GET")
		err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)
		if err != nil{
			fmt.Println("Error Occured")
			panic(err)
		}
	}()
}

func main() {
	fmt.Println("Server starting...")
	startServer(8080)
	fmt.Println("Server stopping ...")
}
