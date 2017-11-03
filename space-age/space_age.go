package space

import (
	"math"
)

const testVersion = 1

// Planet describes the planet of our solar system
type Planet string

const earthYearInSec = 31557600.00

var table = map[Planet]float64{
	"Earth":   1.00,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age returns the age given in seconds as orbital year of a given Planet
func Age(sec float64, p Planet) float64 {
	if factor, ok := table[p]; ok {
		exactTimeInYears := sec / (earthYearInSec * factor)
		return roundToPlace(exactTimeInYears, 2)
	}
	return 0.0

}

func round(val float64) float64 {
	return math.Floor(val + .5)
}

func roundToPlace(val float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return round(val*shift) / shift
}
