package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var dbClient IDatabaseClient
var ids chan int
func startServer(port int,dbName string) {
	router := mux.NewRouter()
	setupRoutes(router)
	dbClient = &DatabaseClient{}
	dbClient.setup(dbName)
	dbClient.initialize()
	ids = make(chan int)
	go incrementId(ids)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	if err != nil {
		fmt.Println("Error Occured")
		panic(err)
	}
	defer dbClient.teardown()
}

func main() {
	fmt.Println("Server starting...")
	startServer(8080,"todo.db")
	fmt.Println("Server stopping ...")
}
