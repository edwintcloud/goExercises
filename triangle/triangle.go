// Package triangle is used for determining what kind of triangle from a, b, c sides
package triangle

import "math"

// Kind is a type for what type of triangle
type Kind int

const (
	NaT Kind = 0
	Equ Kind = 1
	Iso Kind = 2
	Sca Kind = 3
)

// KindFromSides determines what kind of triangle from sides a, b, c
func KindFromSides(a, b, c float64) Kind {

	// first check to make sure input is actually a triangle
	if (a <= 0 || b <= 0 || c <= 0) || (a+b < c || a+c < b || c+b < a) ||
		math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) || math.IsInf(a, 0) ||
		math.IsInf(b, 0) || math.IsInf(c, 0) {
		// return NaT
		return NaT
	}

	// then check what type of triangle
	if (a == b) && (a == c) && (b == c) {
		return Equ
	} else if (a != b) && (a != c) && (b != c) {
		return Sca
	} else {
		return Iso
	}
}
