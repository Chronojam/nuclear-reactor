package boron

import (
	"github.com/chronojam/nuclear-reactor/pkg/material"
)

const (
	NeutronsAbsorbedPerNanosecond = 1
)

type Boron struct {
	mass    int64
	quality int64
}

func New(mass, quality int64) *Boron {
	return &Boron{
		mass:    mass,
		quality: quality,
	}
}

func (g *Boron) Name() string {
	return "Boron"
}

func (g *Boron) Mass() int64 {
	return g.mass
}

func (g *Boron) Quality() int64 {
	return g.quality
}

// AbsorbNeutrons takes in the number of remaining neutrons,
// deducts the amount it absorbs, and returns the remainder
func (g *Boron) AbsorbNeutrons(remainingNeuts int64) int64 {
	atoms := material.GetAtomicCount(g)
	// Crudely model the chance of a neutron hitting this rod
	maxNumberOfNeutronsAbsorbed := (atoms / 8)
	g.quality = material.QualityAfterReaction(g, maxNumberOfNeutronsAbsorbed)
	return maxNumberOfNeutronsAbsorbed
}
