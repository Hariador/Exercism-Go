// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"fmt"
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimRight(remark, " \n\t\r")

	if isSilence(remark) {
		return "Fine. Be that way!"
	}
	if isQuestion(remark) && isShouting(remark) {
		return "Calm down, I know what I'm doing!"
	}

	if isQuestion(remark) {
		return "Sure."
	}

	if isShouting(remark) {
		return "Whoa, chill out!"
	}
	return "Whatever."
}

func isQuestion(remark string) bool {
	last := remark[len(remark)-1:]
	return last == "?"
}

func isShouting(remark string) bool {
	remarkRunes := []rune(remark)
	allUpper := false
	for _, v := range remarkRunes {
		fmt.Println(fmt.Sprintf("Rune: %c", v))
		if unicode.IsLetter(v) {
			if unicode.IsLower(v) {
				return false
			} else {
				allUpper = true
			}
		}
	}

	return allUpper
}

func isSilence(remark string) bool {
	if len(remark) < 1 {
		return true
	}

	return false
}
