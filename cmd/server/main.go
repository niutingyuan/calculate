package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/niutingyuan/calculate/internal/handlers"
)

func main() {
	router := httprouter.New()
	router.POST("/calculate", handlers.CalculateHandler)

	log.Fatal(http.ListenAndServe(":8989", router))
}
