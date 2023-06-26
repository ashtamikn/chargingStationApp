package mockHandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type RequestBody struct {
	Method   string `json:"method"`
	Data     string `json:"data"`
	Endpoint string `json:"endpoint"`
}

type Response struct {
	Status int             `json:"status"`
	Data   json.RawMessage `json:"data"`
}

func HandlePostRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Parse the request body
	var reqBody RequestBody
	err = json.Unmarshal(requestBody, &reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := generateResponse(reqBody)

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Status)
	w.Write(responseJSON)
}

func generateResponse(requestBody RequestBody) Response {
	endpoint := requestBody.Endpoint
	method := requestBody.Method
	details := requestBody.Data
	// print(requestBody.Endpoint)
	print(method, endpoint)
	switch method {
	case http.MethodGet:
		switch endpoint {
		case "/stationCharging":
			// Generate the response for GET /chargingStations
			return Response{
				Status: http.StatusOK,
				Data:   json.RawMessage(`{"message": "Response for GET /stationsCharging"}`),
			}
		case "/chargingStations":
			responseData := []map[string]interface{}{
				{
					"StationID":        1,
					"EnergyOutput":     "50kWh",
					"Type":             "AC",
					"Availability":     true,
					"AvailabilityTime": "",
				},
				{
					"StationID":        2,
					"EnergyOutput":     "10kWh",
					"Type":             "AC",
					"Availability":     false,
					"AvailabilityTime": "",
				},
				{
					"StationID":        5,
					"EnergyOutput":     "120kWh",
					"Type":             "AC",
					"Availability":     true,
					"AvailabilityTime": "",
				},
			}
			responseDataBytes, _ := json.Marshal(responseData)

			return Response{
				Status: http.StatusOK,
				Data:   json.RawMessage(responseDataBytes),
			}

		default:
			// Handle unknown GET endpoint
			if strings.HasPrefix(endpoint, "/chargingStations/") {
				// Extract the ID from the endpoint
				id := strings.TrimPrefix(endpoint, "/chargingStations/")

				// Generate the response for GET /chargingStations/{id}
				responseData := fetchDetailsByID(id)
				return Response{
					Status: http.StatusOK,
					Data:   responseData,
				}
			}

			// Handle unknown GET endpoint
			return Response{
				Status: http.StatusBadRequest,
				Data:   json.RawMessage(`{"message": "Unknown endpoint for GET method"}`),
			}
		}
	case http.MethodPost:
		switch endpoint {
		case "/startCharging":
			// Generate the response for POST /startCharging
			return Response{
				Status: http.StatusOK,
				Data:   json.RawMessage(details),
			}
		case "/chargingStations":
			// Generate the response for POST /startCharging
			return Response{
				Status: http.StatusOK,
				Data:   json.RawMessage(details),
			}

		default:
			// Handle unknown POST endpoint
			return Response{
				Status: http.StatusBadRequest,
				Data:   json.RawMessage(`{"message": "Unknown endpoint for POST method"}`),
			}
		}
	// case http.MethodPut:
	// 	switch endpoint {
	// 	case "/updateChargingStation":
	// 		// Generate the response for PUT /updateChargingStation
	// 		return Response{
	// 			Status: http.StatusOK,
	// 			Data:   json.RawMessage(`{"message": "Response for PUT /updateChargingStation"}`),
	// 		}
	// 	default:
	// 		// Handle unknown PUT endpoint
	// 		return Response{
	// 			Status: http.StatusBadRequest,
	// 			Data:   json.RawMessage(`{"message": "Unknown endpoint for PUT method"}`),
	// 		}
	// 	}
	// case http.MethodDelete:
	// 	switch endpoint {
	// 	case "/deleteChargingStation":
	// 		// Generate the response for DELETE /deleteChargingStation
	// 		return Response{
	// 			Status: http.StatusOK,
	// 			Data:   json.RawMessage(`{"message": "Response for DELETE /deleteChargingStation"}`),
	// 		}
	// 	default:
	// 		// Handle unknown DELETE endpoint
	// 		return Response{
	// 			Status: http.StatusBadRequest,
	// 			Data:   json.RawMessage(`{"message": "Unknown endpoint for DELETE method"}`),
	// 		}
	// 	}
	default:
		// Handle unknown method
		return Response{
			Status: http.StatusBadRequest,
			Data:   json.RawMessage(`{"message": "Unknown method"}`),
		}
	}
}

func fetchDetailsByID(id string) json.RawMessage {
	// Fetch details based on the provided ID
	// Implement the logic to fetch the details based on your requirements
	// For example, you can query a database or an API with the ID
	// Return the response data as a json.RawMessage
	responseData := fmt.Sprintf(`{"StationID": "%s", "EnergyOutput": "200kWh", "Type": "AC","Availability":"true","AvailabilityTime": ""}`, id)
	return json.RawMessage(responseData)
}
