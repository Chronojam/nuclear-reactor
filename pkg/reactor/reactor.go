package reactor

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/chronojam/nuclear-reactor/pkg/controlrod"
	"github.com/chronojam/nuclear-reactor/pkg/fuelrod"
)

const (
	// Exists because StepUpdate() isnt fast enough to run lots per second
	// So we boost the results here by a %
	StepCoefficent = 1.3

	// Maximum uint64 value: 1111111111111111111111111111111111111111111111111111111111111111
	RelativeInfinity = math.MaxUint64
)

type Reactor struct {
	fuelRods                             []*fuelrod.FuelRod
	controlRods                          []*controlrod.ControlRod
	currentReactorEnergy                 int
	EffectiveNeutronMultiplicationFactor float64
}

func New(fuelRods []*fuelrod.FuelRod, controlRods []*controlrod.ControlRod) *Reactor {
	return &Reactor{
		EffectiveNeutronMultiplicationFactor: 9.0,
		fuelRods:                             fuelRods,
		controlRods:                          controlRods,
	}
}

func (r *Reactor) StepUpdate() {
	additionalReactorEnergy := 0
	updateNeutrons := uint64(0)

	for _, rod := range r.fuelRods {
		// 1.146
		enr := r.EffectiveNeutronMultiplicationFactor
		for enr > 0.0 {
			a, u := rod.Operate()
			additionalReactorEnergy += a
			updateNeutrons += uint64(u)
			enr -= rand.Float64()
		}
	}
	fmt.Printf("Neutrons Produced: %v\n", updateNeutrons)

	numNeutronsAbsorbed := 0.0
	for _, rod := range r.controlRods {
		v := rod.Operate()
		if updateNeutrons-uint64(v) <= 0 {
			fmt.Printf("Absorbed All neutrons.\n")
			numNeutronsAbsorbed = float64(updateNeutrons)
			break
		}
		updateNeutrons -= uint64(v)
		numNeutronsAbsorbed += float64(v)
	}
	//updateNeutrons -= int(numNeutronsAbsorbed)
	if updateNeutrons > uint64(numNeutronsAbsorbed) {
		fmt.Printf("Reaction is increasing next iteration \n")
	} else {
		fmt.Printf("Reaction is decreasing next iteration \n")
	}

	fmt.Printf("%v: %v\n", updateNeutrons, numNeutronsAbsorbed)
	r.EffectiveNeutronMultiplicationFactor = (float64(updateNeutrons) / numNeutronsAbsorbed)
	fmt.Printf("%v: \n", r.EffectiveNeutronMultiplicationFactor)
	r.currentReactorEnergy = r.currentReactorEnergy + additionalReactorEnergy
	//fmt.Printf("Reactor Energy: %vJ\n", r.currentReactorEnergy)
	fmt.Printf("\n\n==== NEXT ITERATION ====\n\n")
}
