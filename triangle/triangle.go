package triangle

import (
	"math"
)

const testVersion = 3

// KindFromSides returns the kind of a given triangle
func KindFromSides(a, b, c float64) Kind {
	if math.IsNaN(a) || math.IsInf(a, 1) || math.IsNaN(b) || math.IsInf(b, 1) || math.IsNaN(c) || math.IsInf(c, 1) {
		return NaT
	}
	if a <= 0 || b <= 0 || c <= 0 || a+b < c || b+c < a || a+c < b {
		return NaT
	}
	if a == b && a == c {
		return Equ
	}
	if a == b || b == c || a == c {
		return Iso
	}

	return Sca
}

// Kind represents the Kind of an Triangle
type Kind byte

// Types of Triangles see README.md for further information
const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)
