package series

func All(n int, s string) []string {
	var output []string
	for i := 0; i+n <= len(s); i++ {
		output = append(output, s[i:i+n])
	}
	return output
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}
