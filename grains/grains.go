package grains

import "errors"

const testVersion = 1

// Square returns the number of grains on a chessboard n
func Square(n int) (uint64, error) {
	if n <= 0 || n > 64 {
		return 0, errors.New("Input out of range")
	}
	return exp_rec(2, uint(n)-1), nil
}

// Total returns the sum of grains on a chessboard
func Total() uint64 {
	var result uint64
	for i := 1; i <= 64; i++ {
		nb, err := Square(i)
		result += nb
		if err != nil {
			return 0
		}
	}
	return result
}

// exp_rec calculates a^b
func exp_rec(a, b uint) uint64 {
	if b == 0 {
		return 1
	}
	if b%2 == 1 {
		return uint64(a) * exp_rec(a, b-1)
	}
	p := exp_rec(a, b/2)
	return p * p
}
