package summultiples

import "fmt"

func SumMultiples(limit int, divisors ...int) int {
	factors := make(map[int]bool, limit)
	for _, v := range divisors {
		mult := 1
		stop := false

		for !stop && v > 0 {
			fmt.Println(v * mult)
			if v*mult >= limit {
				stop = true
			} else {
				factors[v*mult] = true
				mult++
			}
		}
	}
	sum := 0
	fmt.Println(factors)
	for num, exists := range factors {
		if exists {
			sum = sum + num
		}
	}

	return sum
}
