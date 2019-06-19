package reactor

import (
	"fmt"
	"math/rand"

	"github.com/chronojam/nuclear-reactor-discord/pkg/fuelrod"
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
		if r.EffectiveNeutronMultiplicationFactor < 1 {
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
	fmt.Printf("fission reactions: %v\n", fissionHits)
	fmt.Printf("k: %v\n", r.EffectiveNeutronMultiplicationFactor)
	fmt.Printf("%v: %v\n", updateNeutrons, r.lastUpdateNeutrons)
	fmt.Printf("======================================\n")
	r.EffectiveNeutronMultiplicationFactor = (updateNeutrons / r.lastUpdateNeutrons) / 18
	r.lastUpdateNeutrons = updateNeutrons
	r.currentReactorEnergy = r.currentReactorEnergy + additionalReactorEnergy
}
