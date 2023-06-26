package handlers

import (
	as "connectAssign1/go-client-generated"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type StationInterface interface {
	GetAggregatedStations(w http.ResponseWriter, r *http.Request)
}
type ApiClient struct {
	Ac *as.APIClient
}

func (c *ApiClient) GetAggregatedStations(w http.ResponseWriter, r *http.Request) {
	// var client = as.NewAPIClient(as.NewConfiguration())
	fmt.Print("6")
	ctx := context.Background()
	// apiResponse, _, err := client.ChargingStationApi.GetAvailableStations(ctx)
	apiResponse, _, err := c.Ac.ChargingStationApi.GetAvailableStations(ctx)

	// occupiedApiResponse, _, err := client.ChargingStationApi.OccupiedStation(ctx)
	occupiedApiResponse, _, err := c.Ac.ChargingStationApi.OccupiedStation(ctx)
	fmt.Print("7")

	jsonData, err := json.Marshal(apiResponse)
	if err != nil {
		panic(err)
	}
	var stations []as.ChargingStation

	err = json.Unmarshal(jsonData, &stations)
	if err != nil {
		panic(err)
	}
	jsonDataO, err := json.Marshal(occupiedApiResponse)
	if err != nil {
		panic(err)
	}
	var stationsOcuupied []as.ChargingStation
	err = json.Unmarshal(jsonDataO, &stationsOcuupied)
	if err != nil {
		panic(err)
	}
	jsonResponse, err := json.Marshal(map[string][]as.ChargingStation{
		"available stations are ": stations,
		"occupied stations are ":  stationsOcuupied,
	})
	if err != nil {
		panic(err)
	}

	// Write the JSON response to the HTTP response writer
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

	fmt.Print(string(jsonResponse))
	// GET /occupiedStations

}
func NewConnect() StationInterface {
	fmt.Print("2")
	apiClient := as.NewAPIClient(as.NewConfiguration())

	nc := &ApiClient{Ac: apiClient}
	return nc
}
