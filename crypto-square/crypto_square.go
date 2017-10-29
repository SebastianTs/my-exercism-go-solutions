package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

const testVersion = 2

// cleanText removes all characters that are not Numbers or Letters
func cleanText(in string) (out string) {
	for _, v := range in {
		if unicode.IsNumber(v) || unicode.IsLetter(v) {
			out += string(v)
		}
	}
	return out
}

// Encode encodes given plaintext to cipher Text
func Encode(in string) string {
	plainText := cleanText(in)
	plainText = strings.ToLower(plainText)
	size := int(math.Ceil(math.Sqrt(float64(len(plainText)))))
	cipherText := make([]string, size)

	for i := 0; i < size; i++ {
		for j := i; j < len(plainText); j += size {
			cipherText[i] += string(plainText[j])
		}
	}
	return strings.Join(cipherText, " ")
}
