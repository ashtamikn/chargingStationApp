package mockHandlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mockServe/initializers"
	"mockServe/models"
	"net/http"
)

type MockInterface interface {
	AddMethod(w http.ResponseWriter, r *http.Request)
}

type Database struct {
}

func (d Database) AddMethod(w http.ResponseWriter, r *http.Request) {
	method := r.Header.Get("Method")
	endpoints := r.URL.Path

	body, _ := ioutil.ReadAll(r.Body)

	var mk models.MockTable
	mk.EndPoint = endpoints
	mk.Method = method
	mk.Response = string(body)
	fmt.Println(mk)
	if result := initializers.DB.Create(&mk); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode("Created")
	w.Header().Set("content-Type", "application/json")

	json.NewEncoder(w).Encode(mk)
}
