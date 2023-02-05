package logs

import (
	"fmt"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {

	for _, r := range log {
		switch r {
		case '\U00002757':
			return "recommendation"
		case '\U0001F50D':
			return "search"
		case '\u2600':
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	fmt.Println(fmt.Sprintf("%v", len(log)))
	var temp []int32
	for _, r := range log {
		if oldRune == r {
			temp = append(temp, newRune)
		} else {
			temp = append(temp, r)
		}
	}
	fmt.Println(fmt.Sprintf("%v", len(temp)))
	return string(temp)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
