package reactor

import (
	"fmt"
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

	// RELATIVE_INFINITY
	if r.EffectiveNeutronMultiplicationFactor >= 1000000.0 {
		fmt.Printf("3.5 Roentgens, not great but not terrible")
		return
	}

	for _, rod := range r.reactorMatrix {
		count := 0
		enr := r.EffectiveNeutronMultiplicationFactor
		for enr > 0.0 {
			enr -= rand.Float64()

			if enr > 0 {
				fissionHits++
				a, u := rod.Fuel.DoFission()
				additionalReactorEnergy += a
				updateNeutrons += u
			}
			count++
		}
		fmt.Printf("Hit count: %v\n", count)
	}

	fmt.Printf("FISSION HITS %v\n", fissionHits)

	fmt.Printf("FISSION HITS %v\n", updateNeutrons)
	fmt.Printf("FISSION HITS %v\n", StepCoefficent)
	fmt.Printf("FISSION HITS %v\n", fissionHits)

	r.EffectiveNeutronMultiplicationFactor = (updateNeutrons * StepCoefficent / r.lastUpdateNeutrons * StepCoefficent)
	r.lastUpdateNeutrons = updateNeutrons
	r.currentReactorEnergy = r.currentReactorEnergy + additionalReactorEnergy
}
