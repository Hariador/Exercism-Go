package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	var temp = map[string]int{}
	for score, letters := range in {
		for _, letter := range letters {
			temp[strings.ToLower(letter)] = score
		}
	}
	return temp
}
