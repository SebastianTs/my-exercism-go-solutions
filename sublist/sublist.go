package sublist

// Relation represents the Relation of two lists
type Relation string

// Sublist returns the Relation of two lists
func Sublist(one, two []int) (rel Relation) {
	switch {
	case len(one) == len(two):
		rel = "equal"
	case len(one) < len(two):
		rel = "sublist"
	case len(one) > len(two):
		one, two = two, one
		rel = "superlist"
	}

	for i := 0; i <= len(two)-len(one); i++ {
		if equals(one, two[i:i+len(one)]) {
			return rel
		}
	}
	return "unequal"

}

func equals(one, two []int) bool {
	for i := range one {
		if one[i] != two[i] {
			return false
		}
	}
	return true
}
