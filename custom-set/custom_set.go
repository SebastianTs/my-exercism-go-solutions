package stringset

// Set represents a set of strings
type Set struct {
	data map[string]struct{}
}

// New creates an empty set
func New() Set {
	return Set{data: make(map[string]struct{})}
}

// NewFromSlice creates an Set from a given list of strings
func NewFromSlice(list []string) Set {
	set := New()
	for _, element := range list {
		set.data[element] = struct{}{}
	}
	return set
}

// String creates a string from a set
func (s Set) String() string {
	out := "{"
	for v := range s.data {
		out += "\"" + v + "\", "
	}
	// remove last ", "
	if len(out) != 1 {
		out = out[:len(out)-2]
	}
	return out + "}"
}

// Length returns the number of elements a set contains
func (s Set) Length() int {
	return len(s.data)
}

// IsEmpty is true if an Set is Empty
func (s Set) IsEmpty() bool {
	return s.Length() == 0
}

// Has checks if the given element is in the set
func (s Set) Has(elem string) bool {
	_, ok := s.data[elem]
	return ok
}

// Subset is true when s1 is a subset of s2
func Subset(s1, s2 Set) bool {
	if s1.IsEmpty() {
		return true
	}
	if s2.IsEmpty() {
		return false
	}
	for k := range s1.data {
		if !s2.Has(k) {
			return false
		}
	}
	return true
}

// Disjoint is true when the two given sets have no common element
func Disjoint(s1, s2 Set) bool {
	if (s1.IsEmpty() && !s2.IsEmpty()) ||
		(!s1.IsEmpty() && s2.IsEmpty()) {
		return true
	}
	for k := range s2.data {
		if s1.Has(k) {
			return false
		}
	}
	return true
}

// Equal is true when the given sets have (only) the same elements
func Equal(s1, s2 Set) bool {
	if s1.Length() != s2.Length() {
		return false
	}
	for k := range s1.data {
		if !s2.Has(k) {
			return false
		}
	}
	return true
}

// Add adds a string to the set
func (s *Set) Add(elem string) {
	s.data[elem] = struct{}{}
}

// Intersection returns the intersection of two sets
func Intersection(s1, s2 Set) Set {
	intersect := New()
	for k := range s1.data {
		if s2.Has(k) {
			intersect.Add(k)
		}
	}
	return intersect
}

// Difference returns the difference of s1 and s2
func Difference(s1, s2 Set) Set {
	difference := New()
	if s1.IsEmpty() {
		return difference
	}
	for k := range s1.data {
		if !s2.Has(k) {
			difference.Add(k)
		}
	}
	return difference
}

// Clone clone a set and create a new one
func Clone(s Set) Set {
	s1 := New()
	for key := range s.data {
		s1.Add(key)
	}
	return s1
}

// Union returns the union of two sets
func Union(s1, s2 Set) Set {
	union := Clone(s1)
	for k := range s2.data {
		union.Add(k)
	}
	return union
}
