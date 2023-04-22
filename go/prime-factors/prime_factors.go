package prime

func Factors(n int64) []int64 {
	return split(n)
}

func split(n int64) []int64 {
	for i := int64(2); i <= n/2; i++ {
		if n%i == 0 {
			return append(split(n/i), i)
		}
	}
	if n == 1 {
		return []int64{}
	}
	return []int64{n}
}
