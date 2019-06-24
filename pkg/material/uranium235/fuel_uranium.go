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
	mass    int64
	quality int64
}

func New(mass, quality int64) *Uranium235Oxide {
	return &Uranium235Oxide{
		mass:    mass,
		quality: quality,
	}
}

func (u *Uranium235Oxide) Name() string {
	return "Uranium235Oxide"
}

func (u *Uranium235Oxide) Mass() int64 {
	return u.mass
}

func (u *Uranium235Oxide) Quality() int64 {
	return u.quality
}

// DoFission, performs a fission, then returns the thermal energy it produced
// and the number of neutrons its released into the system
func (u *Uranium235Oxide) DoFission() (int64, int64) {
	u235AtomicQuantity := material.GetAtomicCount(u)
	energyPerFissionJ := EnergyProducedPerFission * material.JoulesPerMeV
	u.quality = material.QualityAfterReaction(u, ApproximateNoAtomsLostAfterEachFission)

	return int64(float64(u235AtomicQuantity) * energyPerFissionJ), NumNeutronsPerFission * u235AtomicQuantity
}
