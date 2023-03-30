package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var temp Ints
	for _, v := range i {
		if filter(v) {
			temp = append(temp, v)
		}
	}

	return temp
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var temp Ints
	for _, v := range i {
		if !filter(v) {
			temp = append(temp, v)
		}
	}

	return temp
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var temp Lists
	for _, v := range l {
		if filter(v) {
			temp = append(temp, v)
		}
	}

	return temp
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var temp Strings
	for _, v := range s {
		if filter(v) {
			temp = append(temp, v)
		}
	}

	return temp
}
