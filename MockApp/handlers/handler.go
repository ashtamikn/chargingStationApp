package handlers

import (
	"encoding/json"
	"mockApp/initializers"
	"mockApp/models"

	"net/http"
)

type MockInterface interface {
	GetMethod(w http.ResponseWriter, r *http.Request)
}

type Database struct {
}

func (d Database) GetMethod(w http.ResponseWriter, r *http.Request) {
	endpoint := r.URL.Path
	method := r.Method

	// Query the database to find a matching MockTable entry
	var mockTable models.MockTable
	result := initializers.DB.Where("end_point = ? AND method = ?", endpoint, method).First(&mockTable)
	if result.Error != nil {
		// If no matching entry is found, return an appropriate response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Not Found")
		return
	}

	// If a matching entry is found, return the response from the database
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(mockTable.Response))

}
