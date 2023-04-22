package bottlesong

import (
	"fmt"
	"strings"
)

func Recite(startBottles, takeDown int) []string {
	numbers := []string{"no", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}
	var output []string
	for i := startBottles; i > startBottles-takeDown; i-- {
		output = append(output, hang(i, numbers))
		output = append(output, hang(i, numbers))
		output = append(output, "And if one green bottle should accidentally fall,")
		output = append(output, final(i-1, numbers))
		output = append(output, "")

	}

	return output[:len(output)-1]
}

func hang(i int, numbers []string) string {
	if i == 1 {
		return fmt.Sprintf("%v green bottle hanging on the wall,", numbers[i])
	} else {
		return fmt.Sprintf("%v green bottles hanging on the wall,", numbers[i])
	}

}

func final(i int, numbers []string) string {
	out := strings.ToLower(numbers[i])
	if i == 1 {
		return fmt.Sprintf("There'll be %v green bottle hanging on the wall.", out)
	} else {
		return fmt.Sprintf("There'll be %v green bottles hanging on the wall.", out)
	}
}
