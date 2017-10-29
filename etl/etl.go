package etl

import "strings"

const testVersion = 1

// Transform makes scrabble great again
func Transform(in map[int][]string) map[string]int {
	out := make(map[string]int)
	for k, list := range in {
		for _, letter := range list {
			out[strings.ToLower(letter)] = k
		}
	}
	return out
}
