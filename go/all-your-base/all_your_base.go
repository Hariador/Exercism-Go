package allyourbase

import (
	"errors"
	"math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {

	if inputBase < 2 {
		return []int{}, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return []int{}, errors.New("output base must be >= 2")
	}
	if len(inputDigits) == 0 {
		return []int{0}, nil
	}
	if len(inputDigits) == 1 && inputDigits[0] == 0 {
		return []int{0}, nil
	}
	numDigits := len(inputDigits)
	base10 := 0
	for i := 0; i < numDigits; i++ {
		if inputDigits[i] < 0 || inputDigits[i] >= inputBase {
			return []int{}, errors.New("all digits must satisfy 0 <= d < input base")
		}
		base10 = base10 + inputDigits[i]*int(math.Pow(float64(inputBase), float64(numDigits-i-1)))
	}
	var output []int
	Done := false
	for !Done {
		n := base10 / outputBase
		r := base10 % outputBase
		output = append(output, r)
		if n == 0 {
			Done = true
		} else {
			base10 = n
		}
	}
	for i, j := 0, len(output)-1; i < j; i, j = i+1, j-1 {
		output[i], output[j] = output[j], output[i]
	}

	return output, nil
}
