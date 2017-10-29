package letter

import (
	"fmt"
)

type FreqMap map[rune]int

const testVersion = 1

// Frequency counts the letters of a text
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the letters of more than one text in parallel
func ConcurrentFrequency(s []string) FreqMap {
	c := make(chan FreqMap, len(s))

	for _, t := range s {
		t := t
		go func() {
			c <- Frequency(t)
		}()
	}
	result := <-c

	for i := 0; i < len(s)-1; i++ {
		for k, v := range <-c {
			result[k] += v
		}
	}
	return result
}

// NonConcurrentFrequency for benchmark purposes
func NonConcurrentFrequency(s []string) FreqMap {

	interimList := make([]FreqMap, len(s))
	for i, text := range s {
		interimList[i] = Frequency(text)
	}

	// merge results
	result := FreqMap{}
	for _, m := range interimList {
		for k, v := range m {
			fmt.Println(k, v)
			result[k] += v
		}
	}
	return result
}
