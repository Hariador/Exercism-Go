package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	out := initial
	for _, v := range s {
		out = fn(out, v)
	}
	return out
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	out := initial
	for _, v := range s.Reverse() {
		out = fn(v, out)
	}
	return out
}

func (s IntList) Filter(fn func(int) bool) IntList {
	out := make(IntList, 0)
	for _, v := range s {
		if fn(v) {
			out = out.Append(IntList{v})
		}
	}
	return out
}

func (s IntList) Length() int {
	out := 0
	for range s {
		out++
	}
	return out
}

func (s IntList) Map(fn func(int) int) IntList {
	out := make(IntList, s.Length())
	for k, v := range s {
		out[k] = fn(v)
	}
	return out
}

func (s IntList) Reverse() IntList {
	l := s.Length()
	out := make(IntList, l)
	for k, v := range s {
		out[l-k-1] = v
	}
	return out
}

func (s IntList) Append(lst IntList) IntList {
	sl := s.Length()
	ll := lst.Length()
	out := make(IntList, sl+ll)
	for k, v := range s {
		out[k] = v
	}
	for k, v := range lst {
		out[sl+k] = v
	}
	return out
}

func (s IntList) Concat(lists []IntList) IntList {
	out := s
	for _, v := range lists {
		out = out.Append(v)
	}
	return out
}
