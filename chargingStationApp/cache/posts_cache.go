package cache

import "charging-assignment1/models"

type ChargingStationCache interface {
	Set(key string, value []*models.ChargingStation)
	Get(key string) []*models.ChargingStation
}
type StartChargingCache interface {
	Set(key string, value []*models.StartCharging)
	Get(key string) []*models.StartCharging
}
