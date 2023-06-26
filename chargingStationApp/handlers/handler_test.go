package handlers

import (
	"net/http"
	"testing"
)

func TestDatabase_AddChargingStation(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		d    *Database
		args args
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.AddChargingStation(tt.args.w, tt.args.r)
		})
	}
}

func TestMock_AddChargingStation(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		m    *Mock
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.AddChargingStation(tt.args.w, tt.args.r)
		})
	}
}
