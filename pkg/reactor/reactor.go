package reactor

import (
	"math/rand"

	"github.com/chronojam/nuclear-reactor/pkg/fuelrod"
)

const (
	// Exists because StepUpdate() isnt fast enough to run lots per second
	// So we boost the results here by a %
	StepCoefficent = 1.3
)

type Reactor struct {
	reactorMatrix                        []*fuelrod.FuelRod
	currentReactorEnergy                 float64
	EffectiveNeutronMultiplicationFactor float64

	lastUpdateNeutrons float64
}

func (r *Reactor) StepUpdate() {
	additionalReactorEnergy := 0.0
	updateNeutrons := 0.0
	fissionHits := 0
	for _, rod := range r.reactorMatrix {
		// Subcritical, fission reactions are dropping
		if r.EffectiveNeutronMultiplicationFactor < 1.0 {
			if rand.Float64() > r.EffectiveNeutronMultiplicationFactor {
				// No fission reaction
				continue
			}
		}
		// Critical, fission reactions are not gains or declining
		fissionHits += 1
		a, u := rod.Fuel.DoFission()
		additionalReactorEnergy += a
		updateNeutrons += u

		// SuperCritical, fission reactions are increasing
		if r.EffectiveNeutronMultiplicationFactor > 1 {
			// maths is hard, so ive simplified this bit
			if rand.Float64() < 2-r.EffectiveNeutronMultiplicationFactor {
				// oh dear
				a, u := rod.Fuel.DoFission()
				additionalReactorEnergy += a
				updateNeutrons += u
			}
		}
	}

	r.EffectiveNeutronMultiplicationFactor = (updateNeutrons * StepCoefficent / r.lastUpdateNeutrons * StepCoefficent)
	r.lastUpdateNeutrons = updateNeutrons
	r.currentReactorEnergy = r.currentReactorEnergy + additionalReactorEnergy
}
