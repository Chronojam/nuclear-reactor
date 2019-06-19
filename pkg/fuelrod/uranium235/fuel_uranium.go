package uranium235

// fissions per second = (thermal energy produced per second)/(energy per fission)
const (
	// EnergyProducedPerFission in Mev per Fission
	EnergyProducedPerFission = 200
	// How many joules of energy per MeV
	JoulesPerMeV = 1.60218e-13
	// the .25 is to simulate conversion to U238, which we arnt modelling,
	// (yet!), the Real value here should be 1.25 - but we set that to 10x
	// as we dont want to wait literal years.
	ApproximateNoAtomsLostAfterEachFission = 12.5
	AvogadrosNumber                        = 6.022e10 * 23
	// CriticalMass of U235
	CriticalMass          = 52000
	NumNeutronsPerFission = 2.5
)

type Uranium235Oxide struct {
	Mass    float64
	Quality float64
}

// DoFission, performs a fission, then returns the thermal energy it produced
// and the number of neutrons its released into the system
func (u *Uranium235Oxide) DoFission() (float64, float64) {
	// How many u235 nucleui are there
	u235AtomicQuantity := (u.Mass * (u.Quality / 100.0)) * AvogadrosNumber
	energyPerFissionJ := EnergyProducedPerFission * JoulesPerMeV

	// Degrade our quality by the appropriate amount
	k := (u235AtomicQuantity - ApproximateNoAtomsLostAfterEachFission) / AvogadrosNumber
	p := k / u.Mass
	u.Quality = p * 100.0

	return u235AtomicQuantity * energyPerFissionJ, 2.5 * u235AtomicQuantity
}
