// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = 0 // not a triangle
	Equ      = 1
	Iso      = 2
	Sca      = 3
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}
	if a+b < c || b+c < a || a+c < b {
		return NaT
	}

	if a == b && b == c && a == c {
		return Equ
	}

	if a == b || b == c || a == c {
		return Iso
	}
	if a != b && b != c && a != c {
		return Sca
	}

	var k Kind
	return k
}
