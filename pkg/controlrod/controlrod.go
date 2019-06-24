package controlrod

import (
	"github.com/chronojam/nuclear-reactor/pkg/material"
)

type ControlRod struct {
	m material.ControlMaterial
}

// New creates a new fuel rod for usage.
func New(m material.ControlMaterial) *ControlRod {
	rod := &ControlRod{
		m: m,
	}
	return rod
}

func (c *ControlRod) Operate(f int64) int64 {
	return c.m.AbsorbNeutrons(f)
}
