package space

type Planet string

const (
	Earth        Planet  = "Earth"
	Mercury      Planet  = "Mercury"
	Venus        Planet  = "Venus"
	Mars         Planet  = "Mars"
	Jupiter      Planet  = "Jupiter"
	Saturn       Planet  = "Saturn"
	Uranus       Planet  = "Uranus"
	Neptune      Planet  = "Neptune"
	EarthSeconds float64 = 31557600
	MercuryDiv   float64 = 0.2408467
	VenusDiv     float64 = 0.61519726
	MarsDiv      float64 = 1.8808158
	JupiterDiv   float64 = 11.862615
	SaturnDiv    float64 = 29.447498
	UranusDiv    float64 = 84.016846
	NeptuneDiv   float64 = 164.79132
)

func Age(seconds float64, planet Planet) float64 {
	switch planet {
	case Mercury:
		return seconds / EarthSeconds / MercuryDiv
	case Venus:
		return seconds / EarthSeconds / VenusDiv
	case Mars:
		return seconds / EarthSeconds / MarsDiv
	case Jupiter:
		return seconds / EarthSeconds / JupiterDiv
	case Saturn:
		return seconds / EarthSeconds / SaturnDiv
	case Uranus:
		return seconds / EarthSeconds / UranusDiv
	case Neptune:
		return seconds / EarthSeconds / NeptuneDiv
	default:
		return seconds / EarthSeconds
	}
}
