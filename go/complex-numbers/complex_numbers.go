package complexnumbers

import (
	"math"
)

type Number struct {
	a, b float64
}

func (n Number) Real() float64 {
	return n.a
}

func (n Number) Imaginary() float64 {
	return n.b
}

func (n1 Number) Add(n2 Number) Number {
	return Number{a: n1.a + n2.a, b: n1.b + n2.b}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{a: n1.a - n2.a, b: n1.b - n2.b}
}

func (n1 Number) Multiply(n2 Number) Number {
	a := n1.a*n2.a - n1.b*n2.b
	b := n1.b*n2.a + n1.a*n2.b
	return Number{a: a, b: b}
}

func (n Number) Times(factor float64) Number {
	return Number{a: n.a * factor, b: n.b * factor}
}

func (n1 Number) Divide(n2 Number) Number {
	nom1 := n1.a*n2.a + n1.b*n2.b
	denom1 := n2.a*n2.a + n2.b*n2.b
	nom2 := n1.b*n2.a - n1.a*n2.b
	denom2 := n2.a*n2.a + n2.b*n2.b
	return Number{a: nom1 / denom1, b: nom2 / denom2}
}

func (n Number) Conjugate() Number {
	return Number{a: n.a, b: n.b * -1}
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.a*n.a + n.b*n.b)
}

func (n Number) Exp() Number {
	a := math.Exp(n.a)
	return Number{a: a * math.Cos(n.b), b: a * math.Sin(n.b)}
}
