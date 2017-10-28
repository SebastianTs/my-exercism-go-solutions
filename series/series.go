package series

const testVersion = 2

// All returns the substrings of length n of s
func All(n int, s string) (out []string) {

	for i := 0; i <= len(s)-n; i++ {
		out = append(out, s[i:i+n])
	}
	return out
}

// UnsafeFirst returns the first n characters of a string
func UnsafeFirst(n int, s string) string {
	return s[0:n]
}

// First returns the first n characters of a string and
// a bool to signal that the operation is possible
func First(n int, s string) (first string, ok bool) {
	if len(s) >= n {
		return s[0:n], true
	}
	return "", false
}
