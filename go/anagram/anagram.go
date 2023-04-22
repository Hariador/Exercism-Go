package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	sortedSubject := string(sortRunes([]rune(strings.ToLower(subject))))
	var output []string
	for _, candiate := range candidates {
		if strings.ToLower(subject) != strings.ToLower(candiate) {
			if sortedSubject == string(sortRunes([]rune(strings.ToLower(candiate)))) {
				output = append(output, candiate)
			}
		}
	}
	return output
}

func sortRunes(word []rune) []rune {
	sort.Slice(word, func(i, j int) bool {
		return word[i] < word[j]
	})
	return word
}
