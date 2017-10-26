package pangram

import "strings"

const testVersion = 2
const alphaLen = 26

// IsPangram checks if a given string is a Pangram
func IsPangram(in string) bool {
	var chars [alphaLen]bool
	pos := int('A')

	if len(in) >= alphaLen {
		in = strings.ToUpper(in)
		for _, v := range in {
			cur := (int(v) - pos)
			if cur >= 0 && cur < alphaLen {
				chars[cur] = true
			}
		}
		for _, containsLetter := range chars {
			if !containsLetter {
				return false
			}
		}
		return true
	}
	return false

}
