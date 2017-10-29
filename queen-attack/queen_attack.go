package queenattack

import (
	"errors"
	"math"
)

const testVersion = 2

func CanQueenAttack(w, b string) (attack bool, err error) {
	if w == b {
		return false, errors.New("Both Queens are on the same square")
	}
	if w[0] == b[0] || w[1] == b[1] {
		return true, nil
	}
	wx, wy := runeToRow(w[0]), runeToRow(w[1])
	bx, by := runeToRow(b[0]), runeToRow(b[1])
	if wx == 0 || wy == 0 || bx == 0 || by == 0 {
		return false, errors.New("At least one Queen is out of the Board")
	}
	if math.Abs(float64(wx-bx)) == math.Abs(float64(wy-by)) {
		return true, nil
	}
	return false, nil
}

// runeToRow is a stupid map from chars to board rows
// sure there is a more elegant solution
// still do not like the x - int('a') approch
func runeToRow(r byte) int {
	switch r {
	case 'a', '1':
		return 1
	case 'b', '2':
		return 2
	case 'c', '3':
		return 3
	case 'd', '4':
		return 4
	case 'e', '5':
		return 5
	case 'f', '6':
		return 6
	case 'g', '7':
		return 7
	case 'h', '8':
		return 8
	default:
		return 0
	}

}
