package darts

import "math"

func Score(x, y float64) int {
	hyp := x*x + y*y
	dist := math.Sqrt(hyp)
	if dist <= 1.0 {
		return 10
	}
	if dist > 10 {
		return 0
	}
	if dist <= 10 && dist > 5 {
		return 1
	}
	return 5

}
