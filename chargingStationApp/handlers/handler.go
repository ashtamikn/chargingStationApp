package handlers

import (
	"bytes"
	"charging-assignment1/cache"
	"charging-assignment1/initializers"
	"charging-assignment1/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	// "github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gorilla/mux"
)

var ChargingStationCache cache.ChargingStationCache
var StartChargingCache cache.StartChargingCache

type ChargingInterface interface {
	AddChargingStation(w http.ResponseWriter, r *http.Request)
	GetLists(w http.ResponseWriter, r *http.Request)
	GetListByID(w http.ResponseWriter, r *http.Request)
	DeleteListById(w http.ResponseWriter, r *http.Request)
	StartChargings(w http.ResponseWriter, r *http.Request)
	GetChargingStations(w http.ResponseWriter, r *http.Request)
	GetAvailableStation(w http.ResponseWriter, r *http.Request)
	OccupiedStation(w http.ResponseWriter, r *http.Request)
}
type Database struct {
	// S1 string
}

// func Getdatabase() *Database {
// 	return &Database{
// 		S1: "asht",
// 	}
// }

// func Getmock() *Mock {
// 	return &Mock{
// 		S1: "ashtami",
// 	}
// }

type Mock struct {
	// S1 string
}
type RequestBody struct {
	Method   string `json:"method"`
	Data     string `json:"data"`
	Endpoint string `json:"endpoint"`
}

func init() {
	ChargingStationCache = cache.NewRedisCache("localhost:6379", 0, 5*time.Minute)
	// StartChargingCache = cache.NewRedisCache("localhost:6380", 0, 5*time.Minute)
}

func (d Database) AddChargingStation(w http.ResponseWriter, r *http.Request) {
	fmt.Print("333333333333")
	// defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	fmt.Print(body, "  6789")

	if err != nil {
		log.Fatalln(err)
	}
	var cs models.ChargingStation
	json.Unmarshal(body, &cs)
	fmt.Print(cs, "444444444")

	fmt.Print(cs)
	cs.Availability = true

	// Append to the Books table
	if result := initializers.DB.Create(&cs); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode("Created")
	w.Header().Set("content-Type", "application/json")

	json.NewEncoder(w).Encode(cs)

}
func (m Mock) AddChargingStation(w http.ResponseWriter, r *http.Request) {
	mockServerURL := "http://localhost:3000/api/mock" // Replace with the actual mock server URL

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Extract the request method from the incoming request
	method := r.Method

	// Create the request body structure for the mock server
	requestBody := RequestBody{
		Method:   method,
		Data:     string(payload),
		Endpoint: "/chargingStations",
	}

	// Convert the request body to JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the POST request to the mock server
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, mockServerURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Charset", "UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body from the mock server
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the appropriate headers and write the response back to the caller
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(responseBody)

}

// func GetLists(w http.ResponseWriter, r *http.Request) {
// 	var lists []models.ChargingStation

// 	if result := initializers.DB.Find(&lists); result.Error != nil {
// 		fmt.Println(result.Error)
// 	}

// 	w.Header().Add("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(lists)
// }

func (d Database) GetLists(w http.ResponseWriter, r *http.Request) {
	var lists []models.ChargingStation
	if result := initializers.DB.Find(&lists); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lists)
}
func (m Mock) GetLists(w http.ResponseWriter, r *http.Request) {
	mockServerURL := "http://localhost:3000/api/mock"
	// if r.Method != http.MethodPost {
	// 	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	// 	return
	// }

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Extract the request method from the incoming request
	method := r.Method

	// Create the request body structure for the mock server
	requestBody := RequestBody{
		Method:   method,
		Data:     string(payload),
		Endpoint: "/chargingStations",
	}

	// Convert the request body to JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the POST request to the mock server
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, mockServerURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Charset", "UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body from the mock server
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the appropriate headers and write the response back to the caller
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(responseBody)
}

// func (m Mock) GetLists(w http.ResponseWriter, r *http.Request) {
// 	// isTesting := r.Header.Get("Q-Testing") == "true"

// 	mockURL := "https://run.mocky.io/v3/16bd7a60-652c-4910-835c-6d0c4649bea0"
// 	resp, err := http.Get(mockURL)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	mockData, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Add("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(mockData)

// }
func (d Database) GetListByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Print(vars)

	id, err := strconv.Atoi(vars["id"])
	fmt.Print(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var cs models.ChargingStation

	if result := initializers.DB.First(&cs, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cs)
}
func (d Database) DeleteListById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var cs models.ChargingStation

	if result := initializers.DB.First(&cs, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Delete that book
	initializers.DB.Delete(&cs)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")

}
func (m Mock) GetListByID(w http.ResponseWriter, r *http.Request) {
	mockServerURL := "http://localhost:3000/api/mock" // Replace with the actual mock server URL

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	method := r.Method
	id := mux.Vars(r)["id"]
	endpoint := fmt.Sprintf("/chargingStations/%s", id)

	// Create the request body structure for the mock server
	requestBody := RequestBody{
		Method:   method,
		Data:     string(payload),
		Endpoint: endpoint,
	}
	// Extract the ID from the request URL
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create the request URL with the ID
	// requestURL := fmt.Sprintf("%s/%s", mockServerURL, id)

	// Send the GET request to the mock server
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, mockServerURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Charset", "UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body from the mock server
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Remove the backslashes from the JSON data string
	var responseData map[string]interface{}
	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the response data back to JSON without backslashes
	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the appropriate headers and write the response back to the caller
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(responseJSON)
}
func (m Mock) DeleteListById(w http.ResponseWriter, r *http.Request) {
}
func (d Database) StartChargings(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var sc models.StartCharging
	var cs models.ChargingStation
	json.Unmarshal(body, &sc)
	id := sc.StationID
	if resu := initializers.DB.First(&cs, id); resu.Error != nil {
		fmt.Println(resu.Error)
	}
	cs.Availability = false
	// initializers.DB.Save(&cs)

	initializers.DB.Model(&cs).Where("station_id = ?", id).Update("availability", false)
	// Append to the Books table
	if result := initializers.DB.Create(&sc); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode("Created")
	json.NewEncoder(w).Encode(sc)

}
func (m Mock) StartChargings(w http.ResponseWriter, r *http.Request) {
	mockServerURL := "http://localhost:3000/api/mock" // Replace with the actual mock server URL

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Extract the request method from the incoming request
	method := r.Method

	// Create the request body structure for the mock server
	requestBody := RequestBody{
		Method:   method,
		Data:     string(payload),
		Endpoint: "/startCharging",
	}

	// Convert the request body to JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the POST request to the mock server
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, mockServerURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Charset", "UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body from the mock server
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the appropriate headers and write the response back to the caller
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(responseBody)

}
func (d Database) GetChargingStations(w http.ResponseWriter, r *http.Request) {
	var stations []models.StartCharging

	if result := initializers.DB.Find(&stations); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stations)
}
func (m Mock) GetChargingStations(w http.ResponseWriter, r *http.Request) {
	// isTesting := r.Header.Get("Q-Testing") == "true"

	mockServerURL := "http://localhost:3000/api/mock"
	// if r.Method != http.MethodPost {
	// 	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	// 	return
	// }

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Extract the request method from the incoming request
	method := r.Method

	// Create the request body structure for the mock server
	requestBody := RequestBody{
		Method:   method,
		Data:     string(payload),
		Endpoint: "/startCharging",
	}

	// Convert the request body to JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the POST request to the mock server
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, mockServerURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Charset", "UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body from the mock server
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the appropriate headers and write the response back to the caller
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(responseBody)

	// mockURL := "https://run.mocky.io/v3/9c9c0a02-b926-4be2-8075-2f14d46410b3"
	// resp, err := http.Get(mockURL)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// defer resp.Body.Close()
	// mockData, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// w.Header().Add("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(mockData)

}

// func GetChargingStations(w http.ResponseWriter, r *http.Request) {
// 	var stations []models.StartCharging
// 	var cachedSessions []*models.StartCharging = StartChargingCache.Get("all")
// 	if cachedSessions == nil {
// 		if result := initializers.DB.Find(&stations); result.Error != nil {
// 			fmt.Println(result.Error)
// 		}
// 		cachedSessions = make([]*models.StartCharging, len(stations))
// 		for i, v := range stations {
// 			sc := v
// 			cachedSessions[i] = &sc
// 			fmt.Print(cachedSessions[i])
// 		}
// 		fmt.Print("db sc .")
// 		StartChargingCache.Set("all", cachedSessions)
// 	}
// 	list := make([]models.StartCharging, len(cachedSessions))
// 	for i, v := range cachedSessions {
// 		list[i] = *v
// 	}
// 	fmt.Print("cache sc ")

// 	w.Header().Add("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(list)
// }

// Available Charging Stations

func (d Database) GetAvailableStation(w http.ResponseWriter, r *http.Request) {
	var stations []models.ChargingStation

	initializers.DB.Find(&stations)
	var availablestations []models.ChargingStation

	for _, cs := range stations {
		// fmt.Print("him........", stations)
		if cs.Availability == true {
			availablestations = append(availablestations, cs)

		}
	}
	w.Header().Set("content-type", "application/json")

	json.NewEncoder(w).Encode(availablestations)
	// fmt.Println(availablestations)

}
func (m Mock) GetAvailableStation(w http.ResponseWriter, r *http.Request) {
	// isTesting := r.Header.Get("Q-Testing") == "true"

	mockURL := "https://run.mocky.io/v3/85c66040-3d96-4cdd-ae95-917726ddc118"
	resp, err := http.Get(mockURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	mockData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(mockData)

}

// // Occupied Charging Stations:

func (d Database) OccupiedStation(w http.ResponseWriter, r *http.Request) {
	var chargingStations []models.ChargingStation
	var occupiedStations []models.ChargingStation

	initializers.DB.Find(&chargingStations)
	var charging []models.StartCharging
	initializers.DB.Find(&charging)
	for _, session := range charging {
		for _, station := range chargingStations {
			if station.StationID == session.StationID {
				availabilityTime := calculateAvailabilityTime(session, station)
				w.Header().Set("content-type", "application/json")

				// json.NewEncoder(w).Encode(models.ChargingStation{
				// 	StationID:        station.StationID,
				// 	EnergyOutput:     station.EnergyOutput,
				// 	Type:             station.Type,
				// 	AvailabilityTime: availabilityTime,
				// })

				occupiedStations = append(occupiedStations, models.ChargingStation{
					StationID:        station.StationID,
					EnergyOutput:     station.EnergyOutput,
					Type:             station.Type,
					AvailabilityTime: availabilityTime,
				})
				break
			}
		}
	}
	json.NewEncoder(w).Encode(occupiedStations)
}
func (m Mock) OccupiedStation(w http.ResponseWriter, r *http.Request) {
	// isTesting := r.Header.Get("Q-Testing") == "true"

	mockURL := "https://run.mocky.io/v3/18c2978c-6485-4fef-b3e3-d18537dbfb82"
	resp, err := http.Get(mockURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	mockData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(mockData)

}

func calculateAvailabilityTime(session models.StartCharging, station models.ChargingStation) string {
	// batteryCapacity, _ := strconv.Atoi(session.VehicleBatteryCapacity)
	batteryCapacity, _ := strconv.Atoi(strings.TrimSuffix(session.VehicleBatteryCapacity, "kWh"))
	fmt.Println(batteryCapacity)
	// currentCharge, _ := strconv.Atoi(session.CurrentVehicleCharge)
	currentCharge, _ := strconv.Atoi(strings.TrimSuffix(session.CurrentVehicleCharge, "kWh"))

	fmt.Println(currentCharge)

	energyOutput, _ := strconv.Atoi(strings.TrimSuffix(station.EnergyOutput, "kWh"))
	fmt.Println(energyOutput)

	totalChargeTime := float64(batteryCapacity-currentCharge) / float64(energyOutput)
	fmt.Println(totalChargeTime)
	fmt.Println(session.ChargingStartTime.Format(time.RFC3339))
	availabilityTime := session.ChargingStartTime.Add(time.Hour * time.Duration(totalChargeTime))
	fmt.Println(availabilityTime.Format(time.RFC3339))
	return availabilityTime.Format(time.RFC3339)
}

// var ocs models.OccupiedStations

// for _, cs := range chargingStations {
// 	// if cs.StationID == sc.StationID {
// 	if cs.Availability == false {
// 		ocs.StationID = cs.StationID
// 		ocs.EnergyOutput = cs.EnergyOutput
// 		ocs.Type = cs.Type
// 		for _, sc := range charging {
// 			if cs.StationID == sc.StationID {

// 				break
// 			}

// 			ocs.AvailabilityTime = t + sc.chargingStartTime
// 			occupiedStations = append(occupiedStations, ocs)

// 		}
// 	}
// }
