package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("invalid square")
	}
	ratio := 2.0
	firstTerm := 1.0
	convN := float64(number)
	term := firstTerm * math.Pow(ratio, convN-1.0)
	return uint64(term), nil
}

func Total() uint64 {
	sum := uint64(0)
	for i := 1; i < 65; i++ {
		val, _ := Square(i)
		sum = sum + val
	}
	return sum
}
