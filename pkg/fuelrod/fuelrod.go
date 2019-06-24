package fuelrod

import (
	"github.com/chronojam/nuclear-reactor/pkg/material"
)

type FuelRod struct {
	// Type Indicates what this fuel rod is made of.
	m material.FissileMaterial
}

// New creates a new fuel rod for usage.
func New(m material.FissileMaterial) *FuelRod {
	rod := &FuelRod{
		m: m,
	}
	return rod
}

func (f *FuelRod) Operate() (int64, int64) {
	return f.m.DoFission()
}
