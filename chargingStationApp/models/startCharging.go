package models

import "time"

type StartCharging struct {
	StationID              int
	VehicleBatteryCapacity string
	CurrentVehicleCharge   string
	ChargingStartTime      time.Time
}
