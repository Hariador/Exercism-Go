package prime

import (
	"errors"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("invalid input")
	}

	found := 0
	num := 2
	var isPrime bool
	for true {
		mid := num / 2
		isPrime = true
		for i := 2; i <= mid; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			found++
			if found == n {
				return num, nil
			}
		}
		num++
	}

	return 0, nil
}
