package reactor

import (
	"math/rand"
	"testing"
	"time"

	"github.com/chronojam/nuclear-reactor-discord/pkg/fuelrod"
	"github.com/chronojam/nuclear-reactor-discord/pkg/fuelrod/uranium235"
)

func TestReactor(t *testing.T) {
	// 57000 fuel rods, each with ~ 1.5kg @ 3% u235
	// total reactor mass ~87000KG
	rand.Seed(time.Now().UnixNano())
	rods := []*fuelrod.FuelRod{}
	for i := 0; i < 57000; i++ {
		rods = append(rods, fuelrod.New(&uranium235.Uranium235Oxide{
			Mass:    1500,
			Quality: 100,
		}))
	}
	r := &Reactor{
		reactorMatrix:                        rods,
		EffectiveNeutronMultiplicationFactor: 1.0,
	}

	for i := 0; i < 10; i++ {
		t.Logf("%v", r.EffectiveNeutronMultiplicationFactor)
		r.StepUpdate()
	}
}
