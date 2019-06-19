package reactor

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/chronojam/nuclear-reactor/pkg/fuelrod"
	"github.com/chronojam/nuclear-reactor/pkg/fuelrod/uranium235"
)

func TestReactor(t *testing.T) {
	// 57000 fuel rods, each with ~ 1.5kg @ 3% u235
	// total reactor mass ~87000KG
	rand.Seed(time.Now().UnixNano())
	rods := []*fuelrod.FuelRod{}
	for i := 0; i < 1000; i++ {
		rods = append(rods, fuelrod.New(&uranium235.Uranium235Oxide{
			Mass:    1500,
			Quality: 2,
		}))
	}
	r := &Reactor{
		reactorMatrix:                        rods,
		EffectiveNeutronMultiplicationFactor: 1.0,
	}

	for i := 0; i < 100; i++ {
		fmt.Printf("%v\n", r.EffectiveNeutronMultiplicationFactor)
		r.StepUpdate()
	}
}

func BenchmarkStepUpdate(b *testing.B) {
	rods := []*fuelrod.FuelRod{}
	for i := 0; i < 1; i++ {
		rods = append(rods, fuelrod.New(&uranium235.Uranium235Oxide{
			Mass:    1500,
			Quality: 3,
		}))
	}
	r := &Reactor{
		reactorMatrix:                        rods,
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
