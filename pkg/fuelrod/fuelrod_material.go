package fuelrod

type FuelType interface {
	// ThermalEnergyPerFission
	DoFission() (float64, float64)
}
