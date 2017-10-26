package raindrops

import (
	"strconv"
)

const testVersion = 3

// Convert returns a string representation of the given number,
// calculated based on the factors of the input
func Convert(n int) string {
	out := ""
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			switch i {
			case 3:
				out += "Pling"
			case 5:
				out += "Plang"
			case 7:
				out += "Plong"
			}
		}
	}
	if len(out) == 0 {
		out = strconv.Itoa(n)
	}
	return out
}
