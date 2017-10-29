package luhn

import (
	"unicode"
)

const testVersion = 2

// cleanText removes all characters that are not Numbers
func cleanText(in string) (out []int, ok bool) {
	for _, v := range in {
		if unicode.IsNumber(v) {
			out = append(out, int(v)-int('0'))
		} else if unicode.IsSpace(v) {
			// Do nothing
		} else {
			return nil, false
		}
	}
	return out, true
}

func Valid(in string) bool {
	nb, ok := cleanText(in)
	if !ok {
		return false
	}
	if len(nb) < 2 {
		return false
	}
	sum := 0
	parity := len(nb) % 2
	for i := 0; i < len(nb); i++ {
		digit := nb[i]
		if i%2 == parity {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}

	return sum%10 == 0
}
