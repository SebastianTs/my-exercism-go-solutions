package accumulate

import "sync"

const testVersion = 1

// Accumulate accumulates a function to a list of strings
func Accumulate(list []string, fn func(string) string) []string {
	var wg sync.WaitGroup
	wg.Add(len(list))
	for i, v := range list {
		go func(i int, v string) {
			list[i] = fn(v)
			wg.Done()
		}(i, v)
	}
	wg.Wait()
	return list
}
