package triangle

import "math"

const (
	testVersion      = 3
	Equ         Kind = "Equ" // equilateral (all sides equal)
	Iso         Kind = "Iso" // isosceles (2-sides of equal length
	NaT         Kind = "Nat" // not a triangle (! z <= x + y)
	Sca         Kind = "Sca" // scalene (all sides unequal)
)

type Kind string

// KindFromSides determines the +Kind+ of triangle given the length
// of each side. When invalid, returns 'NaT'.
func KindFromSides(a, b, c float64) Kind {
	if isNaT(a, b, c) {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || b == c || a == c {
		return Iso
	}

	if a != b && b != c && a != c {
		return Sca
	}

	return NaT
}

// Private Methods
// ---------------

func isNaT(a, b, c float64) bool {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return true
	}

	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return true
	}

	if a <= 0 || b <= 0 || c <= 0 {
		return true
	}

	if a+b < c || a+c < b || b+c < a {
		return true
	}

	return false
}
