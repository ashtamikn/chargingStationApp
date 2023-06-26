package main

import (
	"log"
	"mockServe/initializers"
	"mockServe/mockHandlers"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	// initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {

	r := mux.NewRouter()

	print("hi")
	rout := mockHandlers.Database{}
	r.HandleFunc("/{path:.*}", rout.AddMethod).Methods("POST")
	log.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal(err)
	}

}
