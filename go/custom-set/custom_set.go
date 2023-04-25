package stringset

import (
	"reflect"
	"strings"
)

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

type Set struct {
	elems map[string]bool
}

func New() Set {
	return Set{map[string]bool{}}
}

func NewFromSlice(l []string) Set {
	var t = make(map[string]bool, len(l))
	for _, v := range l {
		_, exists := t[v]
		if !exists {
			t[v] = true
		}
	}
	return Set{t}
}

func (s Set) String() string {
	sb := strings.Builder{}
	sb.WriteRune('{')
	first := true
	for v := range s.elems {
		if first {
			first = false
		} else {
			sb.WriteString(", ")
		}
		sb.WriteRune('"')
		sb.WriteString(v)
		sb.WriteRune('"')
	}
	sb.WriteRune('}')
	return sb.String()
}

func (s Set) IsEmpty() bool {
	return len(s.elems) == 0
}

func (s Set) Has(elem string) bool {
	_, exists := s.elems[elem]

	return exists
}

func (s Set) Add(elem string) {
	_, exists := s.elems[elem]
	if !exists {
		s.elems[elem] = true
	}
}

func Subset(s1, s2 Set) bool {
	if s1.IsEmpty() {
		return true
	}
	for v := range s1.elems {
		_, exists := s2.elems[v]
		if !exists {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	if s1.IsEmpty() {
		return true
	}
	for v := range s1.elems {
		_, exists := s2.elems[v]
		if exists {
			return false
		}
	}
	return true
}
func Equal(s1, s2 Set) bool {
	return reflect.DeepEqual(s1, s2)
}

func Intersection(s1, s2 Set) Set {
	for v := range s1.elems {
		_, exists := s2.elems[v]
		if !exists {
			delete(s1.elems, v)
		}
	}
	return s1
}

func Difference(s1, s2 Set) Set {
	for v := range s2.elems {
		delete(s1.elems, v)
	}
	return s1
}

func Union(s1, s2 Set) Set {
	for v := range s2.elems {
		s1.Add(v)
	}
	return s1
}
