package anagram

import (
	"strings"
)

// Detect selects the correct sublist of anagrams to a given word
func Detect(subject string, candidates []string) (result []string) {
	for _, candidate := range candidates {
		if isAnagram(subject, candidate) {
			result = append(result, candidate)
		}
	}
	return result
}

func isAnagram(subject string, candidate string) bool {
	subject = strings.ToLower(subject)
	candidate = strings.ToLower(candidate)
	switch {
	case len(subject) != len(candidate):
		return false
	case subject == candidate:
		return false
	case SortString(subject) == SortString(candidate):
		return true
	default:
		return false
	}
}
