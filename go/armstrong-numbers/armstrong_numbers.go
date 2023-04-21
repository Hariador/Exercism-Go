package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	sN := strconv.Itoa(n)
	sum := 0
	for _, v := range sN {
		sum = sum + int(math.Pow(float64(v-48), float64(len(sN))))
	}

	return sum == n
}
