package resistorcolortrio

import (
	"math"
	"strconv"
)

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	value := (10*colorCode(colors[0]) + colorCode(colors[1])) * int(math.Pow10(colorCode(colors[2])))

	postfix := ""
	switch value >= 0 {
	case value < 1000:
		postfix = " ohms"

	case value >= 1000 && value < 1000000:
		{
			value = value / 1000
			postfix = " kiloohms"
		}

	case value >= 1000000 && value < 1000000000:
		{
			value = value / 1000000
			postfix = " megaohms"
		}
	case value >= 1000000000 && value < 1000000000000:
		{
			value = value / 1000000000
			postfix = " gigaohms"
		}
	}
	return strconv.Itoa(value) + postfix
}

func colors() []string {
	return []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}
}

// ColorCode returns the resistance value of the given color.
func colorCode(color string) int {
	for k, v := range colors() {

		if v == color {
			return k
		}
	}

	return -1
}
