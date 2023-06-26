package main

import (
	"fmt"
	"log"
	"mockk/mockHandler"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router using Gorilla Mux
	router := mux.NewRouter()

	// Define the route and its corresponding handler
	router.HandleFunc("/api/mock", mockHandler.HandlePostRequest).Methods("POST")

	// Start the server
	fmt.Println("Mock server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
