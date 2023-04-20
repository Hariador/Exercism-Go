package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span == 0 {
		return 1, nil
	}

	if span < 0 {
		return 0, errors.New("invalid span length")
	}

	end := span
	if end > len(digits) {
		return 0, errors.New("span length great than input")
	}
	ints, err := convertToIntArray(digits)
	if err != nil {
		return 0, err
	}
	start := 0
	var largest int64
	var product int64
	for end <= len(digits) {
		product = 1
		for i := start; i < end; i++ {
			product = product * int64(ints[i])
		}
		if product > largest {
			largest = product
		}
		start++
		end++
	}
	return largest, nil
}

func convertToIntArray(digits string) ([]int, error) {
	var ints []int
	for _, v := range digits {
		digit, err := strconv.Atoi(string(v))
		if err != nil {
			return nil, err
		}
		ints = append(ints, digit)
	}
	return ints, nil
}
