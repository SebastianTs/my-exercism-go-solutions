package hamming

import "errors"

const testVersion = 6

// Distance returns the Hamming distance of two strings or
// an error the strings have not the same length
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Distance: Length of String a and b not equal")
	}
	d := 0
	for i, v := range a {
		if v != rune(b[i]) {
			d++
		}
	}
	return d, nil
}
