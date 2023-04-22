package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	if equals(l1, l2) {
		return RelationEqual
	}
	if isSublist(l1, l2) {
		return RelationSublist
	}
	if isSublist(l2, l1) {
		return RelationSuperlist
	}

	return RelationUnequal
}

func isSublist(l1, l2 []int) bool {
	if len(l1) > len(l2) {
		return false
	}
	for i := 0; i+len(l1) <= len(l2); i++ {
		if equals(l1, l2[i:i+len(l1)]) {
			return true
		}
	}
	return false
}

func equals(l1, l2 []int) bool {
	if len(l1) != len(l2) {
		return false
	}
	equal := true
	for k, v := range l1 {
		if l2[k] != v {
			equal = false
			break
		}
	}
	return equal
}
