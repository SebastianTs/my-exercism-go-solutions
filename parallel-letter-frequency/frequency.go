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
	c := make(chan int, len(s))
	interimList := make([]FreqMap, len(s))

	for i := 0; i < len(s); i++ {
		go func(s string, i int, c chan int) {
			interimList[i] = Frequency(s)
			c <- 1
		}(s[i], i, c)
	}

	for i := 0; i < len(s); i++ {
		<-c
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
