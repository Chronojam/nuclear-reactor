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

func (c *ControlRod) Operate() int {
	return c.m.AbsorbNeutrons()
}
