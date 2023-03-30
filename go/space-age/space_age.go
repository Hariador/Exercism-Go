package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	switch planet {
	case "Mercury":
		return getYears(seconds) / 0.2408467
	case "Venus":
		return getYears(seconds) / 0.61519726
	case "Earth":
		return getYears(seconds)
	case "Mars":
		return getYears(seconds) / 1.8808158
	case "Jupiter":
		return getYears(seconds) / 11.862615
	case "Saturn":
		return getYears(seconds) / 29.447498
	case "Uranus":
		return getYears(seconds) / 84.016846
	case "Neptune":
		return getYears(seconds) / 164.79132
	}

	return -1.0
}

func getYears(seconds float64) float64 {
	return seconds / 31557600
}
