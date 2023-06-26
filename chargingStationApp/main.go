package main

import (
	"log"
	"net/http"

	stations "charging-assignment1/handlers"
	"charging-assignment1/initializers"

	"charging-assignment1/middleware"

	"github.com/gorilla/mux"
)

func init() {
	// initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

// var retriever stations.ChargingInterface

// func header(w http.ResponseWriter, r *http.Request) stations.ChargingInterface {
// 	if r.Header.Get("Q-Testing") == "true" {
// 		fmt.Println("if")
// 		retriever = &stations.Mock{
// 			S1: "sfsfsf",
// 		}

// 		fmt.Println(retriever)
// 	} else {
// 		fmt.Println("else")
// 		retriever = &stations.Database{
// 			S1: "sfsfsf",
// 		}

// 		fmt.Println(retriever)
// 	}
// 	return retriever
// 	// You can also store the retriever instance in the request context if you want to access it in subsequent middleware or handlers:
// 	// ctx := context.WithValue(r.Context(), "retriever", retriever)
// 	// r = r.WithContext(ctx)
// }

// Http middleware
// var retriever stations.ChargingInterface = &stations.Database{}

// func retrieverMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		var retriever stations.ChargingInterface

// 		if r.Header.Get("Q-Testing") == "true" {
// 			retriever = &stations.Mock{}
// 		} else {
// 			retriever = &stations.Database{}
// 		}

// 		// Store the retriever object in the request context
// 		ctx := context.WithValue(r.Context(), "retriever", retriever)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

func main() {
	r := mux.NewRouter()
	r.Use(middleware.RetrieverMiddleware)

	// retriever = &stations.Database{
	// 	S1: "sfsfsf",
	// }

	// r.HandleFunc("/chargingStations", func(w http.ResponseWriter, r *http.Request) { header(w, r).GetLists(w, r) }).Methods("GET")
	// r.HandleFunc("/chargingStations/{id}", func(w http.ResponseWriter, r *http.Request) { header(w, r).GetListByID(w, r) }).Methods("GET")
	// r.HandleFunc("/chargingStations", func(w http.ResponseWriter, r *http.Request) { header(w, r).AddChargingStation(w, r) }).Methods("POST")
	// r.HandleFunc("/startCharging", func(w http.ResponseWriter, r *http.Request) { header(w, r).StartChargings(w, r) }).Methods("POST")
	// r.HandleFunc("/startCharging", func(w http.ResponseWriter, r *http.Request) { header(w, r).GetChargingStations(w, r) }).Methods("GET")
	// r.HandleFunc("/availableStations", func(w http.ResponseWriter, r *http.Request) { header(w, r).GetAvailableStation(w, r) }).Methods("GET")
	// r.HandleFunc("/occupiedStations", func(w http.ResponseWriter, r *http.Request) { header(w, r).OccupiedStation(w, r) }).Methods("GET")

	// Http middleware

	r.HandleFunc("/chargingStations", func(w http.ResponseWriter, r *http.Request) {
		retriever := r.Context().Value("retriever").(stations.ChargingInterface)
		retriever.GetLists(w, r)
	}).Methods("GET")
	r.HandleFunc("/chargingStations/{id}", func(w http.ResponseWriter, r *http.Request) {
		retriever := r.Context().Value("retriever").(stations.ChargingInterface)
		retriever.GetListByID(w, r)
	}).Methods("GET")
	r.HandleFunc("/chargingStations", func(w http.ResponseWriter, r *http.Request) {
		retriever := r.Context().Value("retriever").(stations.ChargingInterface)
		retriever.AddChargingStation(w, r)
	}).Methods("POST")
	r.HandleFunc("/startCharging", func(w http.ResponseWriter, r *http.Request) {
		retriever := r.Context().Value("retriever").(stations.ChargingInterface)
		retriever.StartChargings(w, r)
	}).Methods("POST")
	r.HandleFunc("/startCharging", func(w http.ResponseWriter, r *http.Request) {
		retriever := r.Context().Value("retriever").(stations.ChargingInterface)
		retriever.GetChargingStations(w, r)
	}).Methods("GET")
	r.HandleFunc("/availableStations", func(w http.ResponseWriter, r *http.Request) {
		retriever := r.Context().Value("retriever").(stations.ChargingInterface)
		retriever.GetAvailableStation(w, r)
	}).Methods("GET")
	r.HandleFunc("/occupiedStations", func(w http.ResponseWriter, r *http.Request) {
		retriever := r.Context().Value("retriever").(stations.ChargingInterface)
		retriever.OccupiedStation(w, r)
	}).Methods("GET")

	log.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal(err)
	}
}
