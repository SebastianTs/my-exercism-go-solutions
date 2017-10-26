package acronym

import (
	"strings"
)

const testVersion = 3

// Abbreviate creates an acronym out of a given string
func Abbreviate(in string) string {
	pos := 0
	out := ""
	for i, v := range in {
		if v == ' ' || v == '-' || i == len(in)-1 {
			out += strings.ToUpper(in[pos : pos+1])
			pos = i + 1
		}
	}
	return out
}
