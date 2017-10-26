// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

import "fmt"

// ShareWith returns a quote with the given string included,
// if the given string is empty a placeholder will be used
func ShareWith(in string) string {
	if len(in) == 0 {
		in = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", in)
}
