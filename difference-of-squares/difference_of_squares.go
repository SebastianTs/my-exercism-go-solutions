package diffsquares

const testVersion = 1

// SquareOfSums returns the square of the sum of a given number n
// using the "little Gauß"-Formula
func SquareOfSums(n int) int {
	sum := (n*n + n) / 2
	return sum * sum
}

// SumOfSquares returns the sum of the squares of a given number n
func SumOfSquares(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

// Difference calculates the Difference of the SquareOfSums and the SumOfSquares
// for a given number n
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
