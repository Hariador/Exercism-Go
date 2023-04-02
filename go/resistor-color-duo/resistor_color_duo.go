package resistorcolorduo

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	resitorValue := 0
	if len(colors) == 0 {
		return -1
	}
	resitorValue = colorCode(colors[0]) * 10
	if len(colors) > 1 {
		resitorValue += colorCode(colors[1])
	}

	return resitorValue

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
