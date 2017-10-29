package isogram

import (
	"strings"
)

const testVersion = 1

// IsIsogram is true when the given word is an isogram
func IsIsogram(word string) bool {
	m := make(map[string]int)
	word = strings.ToLower(word)
	word = strings.Replace(word, "-", "", -1)
	word = strings.Replace(word, " ", "", -1)
	for _, c := range word {
		m[string(c)]++
	}
	for _, v := range m {
		if v > 1 {
			return false
		}
	}
	return true
}
