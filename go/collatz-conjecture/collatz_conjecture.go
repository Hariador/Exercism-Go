package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("n must be larger than 0")
	}
	if n == 1 {
		return 0, nil
	}
	if n%2 == 0 {
		count, err := CollatzConjecture(n / 2)
		return count + 1, err
	} else {
		count, err := CollatzConjecture((3 * n) + 1)
		return count + 1, err
	}

}
