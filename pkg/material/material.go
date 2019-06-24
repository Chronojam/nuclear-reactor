package material

const (
	AvogadrosNumber int64 = 6.022e10 * 23
	JoulesPerMeV          = 1.60218e-13
)

type Material interface {
	Name() string
	Mass() int64
	Quality() int64
}

type FissileMaterial interface {
	Material
	DoFission() (int64, int64)
}

type ControlMaterial interface {
	Material
	AbsorbNeutrons(current int64) int64
}

func GetAtomicCount(m Material) int64 {
	return (m.Mass() * (m.Quality() / 100)) * AvogadrosNumber
}

// After a reaction, what should we set our new quality too?
func QualityAfterReaction(m Material, atomsLostPerReaction int64) int64 {
	return (GetAtomicCount(m) - atomsLostPerReaction) /
		AvogadrosNumber /
		m.Mass() * 100.0
}
