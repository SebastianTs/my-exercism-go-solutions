package grains

import "errors"

const testVersion = 1

// Square returns the number of grains on a chessboard n
func Square(n int) (uint64, error) {
	if n <= 0 || n > 64 {
		return 0, errors.New("Input out of range")
	}
	return (1 << uint64(n-1)), nil
}

// Total returns the sum of grains on a chessboard
func Total() uint64 {

	return (1 << uint64(64)) - 1
}
