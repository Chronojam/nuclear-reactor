package material

const (
	AvogadrosNumber = 6.022e10 * 23
	JoulesPerMeV    = 1.60218e-13
	MevPerJoule     = 6241506479963
)

type Material interface {
	Name() string
	Mass() int
	Quality() float64
}

type FissileMaterial interface {
	Material
	DoFission() (int, int)
}

type ControlMaterial interface {
	Material
	AbsorbNeutrons() int
}

func GetAtomicCount(m Material) int {
	percentQuality := m.Quality() / 100.0
	floatingCount := float64(m.Mass()) * percentQuality * AvogadrosNumber
	return int(floatingCount)
}

// After a reaction, what should we set our new quality too?
func QualityAfterReaction(m Material, atomsLostPerReaction int) float64 {
	r := float64((GetAtomicCount(m) - atomsLostPerReaction)) /
		AvogadrosNumber /
		float64(m.Mass()) * 100.0

	//fmt.Printf("Quality Remaining: %v\n ", r)

	return r
}
