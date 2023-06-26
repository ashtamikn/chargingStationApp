package main

import (
	"log"
	"mockApp/handlers"
	"mockApp/initializers"
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
	rout := handlers.Database{}
	r.HandleFunc("/{path:.*}", rout.GetMethod)
	log.Println("Starting server on port 8082...")
	err := http.ListenAndServe(":8082", r)

	if err != nil {
		log.Fatal(err)
	}

}
