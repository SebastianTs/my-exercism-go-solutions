package allyourbase

import (
	"errors"
)

// ConvertToBase converts a number, represented as a sequence of digits in one base, to any other base.
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}
	// Assumption 4
	if len(inputDigits) == 0 {
		return []int{0}, nil
	}

	remainder, err := convertFromBase(inputBase, inputDigits)
	if err != nil {
		return []int{}, err
	}
	// Assumption 1
	if remainder == 0 {
		return []int{0}, nil
	}

	multiplier := 1
	digits := []int{}
	for remainder > 0 {
		multiplier *= outputBase
		value := remainder % multiplier
		digit := value / (multiplier / outputBase)

		digits = append(digits, digit)
		remainder -= value
	}

	digits = reverse(digits)

	return digits, nil
}

func reverse(digits []int) []int {
	l := len(digits) - 1
	for i := 0; i < len(digits)/2; i++ {
		digits[i], digits[l-i] = digits[l-i], digits[i]
	}
	return digits
}

func convertFromBase(inputBase int, digits []int) (value int, err error) {
	for _, v := range digits {
		if v < 0 || v >= inputBase {
			return 0, errors.New("all digits must satisfy 0 <= d < input base")
		}
		value = value*inputBase + v
	}
	return value, nil
}
