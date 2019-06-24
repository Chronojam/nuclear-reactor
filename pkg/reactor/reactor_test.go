package reactor

import (
	"math/rand"
	"testing"
	"time"

	"github.com/chronojam/nuclear-reactor/pkg/controlrod"
	"github.com/chronojam/nuclear-reactor/pkg/fuelrod"
	"github.com/chronojam/nuclear-reactor/pkg/material/boron"
	"github.com/chronojam/nuclear-reactor/pkg/material/uranium235"
)

func TestReactor(t *testing.T) {
	// 57000 fuel rods, each with ~ 1.5kg @ 3% u235
	// total reactor mass ~87000KG
	rand.Seed(time.Now().UnixNano())
	fuelRods := []*fuelrod.FuelRod{}
	for i := 0; i < 10; i++ {
		fuelRods = append(fuelRods, fuelrod.New(
			uranium235.New(1500, 3)),
		)
	}
	controlRods := []*controlrod.ControlRod{}
	for i := 0; i < 5; i++ {
		controlRods = append(controlRods, controlrod.New(
			boron.New(1500, 100),
		))
	}
	r := New(
		fuelRods,
		controlRods,
	)

	for i := 0; i < 30; i++ {
		r.StepUpdate()
	}

}

func BenchmarkStepUpdate(b *testing.B) {
	fuelRods := []*fuelrod.FuelRod{}
	for i := 0; i < 1; i++ {
		fuelRods = append(fuelRods, fuelrod.New(uranium235.New(1500, 3)))
	}
	controlRods := []*controlrod.ControlRod{}
	for i := 0; i < 1; i++ {
		controlRods = append(controlRods, controlrod.New(boron.New(1500, 3)))
	}

	r := &Reactor{
		fuelRods:                             fuelRods,
		controlRods:                          controlRods,
		EffectiveNeutronMultiplicationFactor: 1.0,
	}
	for i := 0; i < b.N; i++ {
		r.StepUpdate()
	}
}

func BenchmarkRandom64(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		rand.Float64()
	}
}
func BenchmarkRandom32(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		rand.Float32()
	}
}
