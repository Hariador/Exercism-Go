package luhn

import (
	"strconv"
)

func Valid(id string) bool {

	var digits []int
	for _, c := range id {
		if c != ' ' {
			d, err := strconv.Atoi(string(c))
			if err != nil {
				return false
			}
			digits = append(digits, d)
		}
	}
	double := false
	sum := 0
	if len(digits) <= 1 {
		return false
	}
	for i := len(digits) - 1; i >= 0; i-- {
		if double {
			t := digits[i] * 2
			if t > 9 {
				t = t - 9
			}
			sum += t
		} else {
			sum += digits[i]
		}
		double = !double
	}

	return sum%10 == 0

}
