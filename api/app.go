package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var dbClient IDatabaseClient

func startServer(port int) {
	router := mux.NewRouter()
	setupRoutes(router)
	dbClient = &DatabaseClient{}
	dbClient.setup()
	dbClient.initialize()
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	if err != nil {
		fmt.Println("Error Occured")
		panic(err)
	}
	defer dbClient.teardown()
}

func main() {
	fmt.Println("Server starting...")
	startServer(8080)
	fmt.Println("Server stopping ...")
}
