package sieve

import (
	"fmt"
	"sort"
)

func Sieve(limit int) []int {
	sieve := make(map[int]bool)
	for i := 2; i <= limit; i++ {
		sieve[i] = true
	}
	fmt.Println(fmt.Sprintf("%v", sieve))
	for k, v := range sieve {
		if v {
			for i := 2; i*k <= limit; i++ {
				sieve[i*k] = false
			}
		}
	}
	var output []int
	for k, v := range sieve {
		if v {
			output = append(output, k)
		}
	}
	sort.Ints(output)
	return output
}
