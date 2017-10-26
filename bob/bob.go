package bob

import (
	"strings"
)

const testVersion = 3

// Hey implements the respnses of an teenager names Bob
func Hey(in string) string {
	in = strings.TrimSpace(in)
	if len(in) > 0 {
		if strings.ToUpper(in) == in && strings.ContainsAny(in, "ABCDEFGHIJKLMOPQRSTVWXYZ") {
			return "Whoa, chill out!"
		}
		if in[len(in)-1] == '?' {
			return "Sure."
		}
		return "Whatever."
	}
	return "Fine. Be that way!"

}
