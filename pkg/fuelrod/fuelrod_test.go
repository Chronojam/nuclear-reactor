package fuelrod

import (
	"testing"
	"time"
)

func TestRod(t *testing.T) {
	f := New(time.Second * 1)

	for !f.IsDepleted() {
		time.Sleep(time.Second)
		t.Logf("%v", f.FuelRemaining())
	}
}
