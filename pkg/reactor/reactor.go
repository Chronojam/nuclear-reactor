package reactor

import (
	"fmt"
	"math/rand"

	"github.com/chronojam/nuclear-reactor/pkg/controlrod"
	"github.com/chronojam/nuclear-reactor/pkg/fuelrod"
)

const (
	// Exists because StepUpdate() isnt fast enough to run lots per second
	// So we boost the results here by a %
	StepCoefficent = 1.3
)

type Reactor struct {
	fuelRods                             []*fuelrod.FuelRod
	controlRods                          []*controlrod.ControlRod
	currentReactorEnergy                 int64
	EffectiveNeutronMultiplicationFactor int64
}

func New(fuelRods []*fuelrod.FuelRod, controlRods []*controlrod.ControlRod) *Reactor {
	return &Reactor{
		EffectiveNeutronMultiplicationFactor: 1.0,
		fuelRods:                             fuelRods,
		controlRods:                          controlRods,
	}
}

func (r *Reactor) StepUpdate() {
	additionalReactorEnergy := 0.0
	updateNeutrons := 0.0

	for _, rod := range r.fuelRods {
		count := 0
		// 1.146
		enr := r.EffectiveNeutronMultiplicationFactor
		for enr > 0.0 {
			a, u := rod.Operate()
			additionalReactorEnergy += a
			updateNeutrons += u
			count++
			enr -= rand.Float64()
		}
	}

	numNeutronsAbsorbed := 0.0
	for _, rod := range r.controlRods {
		numNeutronsAbsorbed += rod.Operate(updateNeutrons)
	}
	updateNeutrons -= numNeutronsAbsorbed
	if updateNeutrons > numNeutronsAbsorbed {
		fmt.Printf("Reaction is increasing next iteration \n")
	} else {
		fmt.Printf("Reaction is decreasing next iteration \n")
	}

	//fmt.Printf("%v: %v\n", updateNeutrons, r.lastUpdateNeutrons)
	r.EffectiveNeutronMultiplicationFactor = (updateNeutrons / numNeutronsAbsorbed)
	r.currentReactorEnergy = r.currentReactorEnergy + additionalReactorEnergy
	//fmt.Printf("Reactor Energy: %vJ\n", r.currentReactorEnergy)
}
