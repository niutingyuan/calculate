package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/niutingyuan/calculate/internal/handlers"
)

// This is the main entry point to the application. It sets up the HTTP server and routes using `httprouter` package
func main() {
	router := httprouter.New() // Setting up the router using 3rd party package
	router.POST("/calculate", handlers.CalculateHandler)

	log.Fatal(http.ListenAndServe(":8989", router)) // Setting up the server
}
