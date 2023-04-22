// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import "unicode"

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	output := ""
	pick := true
	for _, v := range s {
		if pick && unicode.IsLetter(v) {
			output = output + string(unicode.ToUpper(v))
			pick = false
		} else if v == ' ' || v == '-' {
			pick = true
		}
	}
	return output
}
