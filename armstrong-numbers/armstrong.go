package armstrong

import (
	"fmt"
	"math"
)

//IsNumber checks if a given int is an armstrong number
func IsNumber(number int) bool {
	str := fmt.Sprintf("%d", number)
	sum := 0.0
	exp := float64(len(str))
	for _, c := range str {
		base := float64((int(c) - '0'))
		sum += math.Pow(base, exp)
	}
	return int(sum) == number
}
