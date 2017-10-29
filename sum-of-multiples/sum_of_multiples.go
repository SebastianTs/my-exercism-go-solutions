package summultiples

const testVersion = 2

func SumMultiples(ds ...int) func(max int) int {
	return func(max int) int {
		sum := 0
		for i := 1; i < max; i++ {
			if isMultiple(i, ds) {
				sum++
			}
		}
		return sumcd 
	}
}

func isMultiple(n int, ds []int) bool {
	for _, d := range ds {
		if n%d == 0 {
			return true
		}
	}
	return false
}
