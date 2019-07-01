package uranium235

import (
	"github.com/chronojam/nuclear-reactor/pkg/material"
)

// fissions per second = (thermal energy produced per second)/(energy per fission)
const (
	// EnergyProducedPerFission in Mev per Fission
	EnergyProducedPerFission = 200
	// the .25 is to simulate conversion to U238, which we arnt modelling,
	// (yet!), the Real value here should be 1.25 - but we set that to 10x
	// as we dont want to wait literal years.
	ApproximateNoAtomsLostAfterEachFission = 13
	// CriticalMass of U235
	CriticalMass          = 52000
	NumNeutronsPerFission = 3

	FissionReactionsPerNanosecond = 0.01
)

type Uranium235Oxide struct {
	mass    int
	quality float64
}

func New(mass int, quality float64) *Uranium235Oxide {
	return &Uranium235Oxide{
		mass:    mass,
		quality: quality,
	}
}

func (u *Uranium235Oxide) Name() string {
	return "Uranium235Oxide"
}

func (u *Uranium235Oxide) Mass() int {
	return u.mass
}

func (u *Uranium235Oxide) Quality() float64 {
	return u.quality
}

//1 joule is equal to 6241506479963.2 MeV.
// DoFission, performs a fission, then returns the thermal energy it produced
// and the number of neutrons its released into the system
func (u *Uranium235Oxide) DoFission() (int, int) {
	u235AtomicQuantity := material.GetAtomicCount(u)
	energyPerFissionJ := EnergyProducedPerFission * material.JoulesPerMeV
	u.quality = material.QualityAfterReaction(u, ApproximateNoAtomsLostAfterEachFission)
	if NumNeutronsPerFission*u235AtomicQuantity <= 0 {
		panic("u235 <= 0")
	}
	return u235AtomicQuantity * int(energyPerFissionJ), NumNeutronsPerFission * u235AtomicQuantity
}
