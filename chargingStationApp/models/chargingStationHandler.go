package models

type ChargingStation struct {
	StationID        int
	EnergyOutput     string
	Type             string
	Availability     bool
	AvailabilityTime string
}

// "stationId":0,
// "vehicleBatteryCapacity":"500kWh",
// "currentVehicleCharge":"100kWh",
// "chargingStartTime":"2022-07-14T13:30:00Z"
