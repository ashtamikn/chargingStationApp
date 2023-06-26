package main

import (
	"connectAssign1/handlers"
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Print("hello")
	// var ac *as.APIClient

	// apiClient := as.NewAPIClient(as.NewConfiguration())
	// api := &handlers.ApiClient{Ac: apiClient}
	api := handlers.NewConnect()
	fmt.Print("5")
	r := mux.NewRouter()
	// r.HandleFunc("/aggregatedStations", handlers.GetAggregatedStations).Methods("GET")
	r.HandleFunc("/aggregatedStations", api.GetAggregatedStations).Methods("GET")

	err := http.ListenAndServe(":8082", r)
	log.Println("Starting server on port 8082...")
	if err != nil {
		log.Fatal("failed to connect to 8082")
	}
}
