package fuelrod

type FuelRod struct {
	// Type Indicates what this fuel rod is made of.
	Fuel FuelType
}

// New creates a new fuel rod for usage.
func New(fuel FuelType) *FuelRod {
	rod := &FuelRod{
		Fuel: fuel,
	}
	return rod
}
