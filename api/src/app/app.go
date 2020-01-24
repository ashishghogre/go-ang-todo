package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"github.com/gorilla/mux"
)

var dbClient IDatabaseClient
var ids chan int

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func startServer(port int, dbName string) {
	logrus.Infof("Server starting on port %d ...", port)
	router := mux.NewRouter()
	setupRoutes(router)
	dbClient = &DatabaseClient{}
	dbClient.setup(dbName)
	dbClient.initialize()
	ids = make(chan int)
	go incrementID(ids)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	if err != nil {
		fmt.Println("Error Occured")
		panic(err)
	}
	defer dbClient.teardown()
}

func main() {

	startServer(8080, "todo.db")
	logrus.Info("Server stopping ...")
}
