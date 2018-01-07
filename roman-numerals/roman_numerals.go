package romannumerals

import "errors"

func ToRomanNumeral(n int) (out string, err error) {
	if n < 1 || n > 3000 {
		return "", errors.New("Input out of Range")
	}

	rn := []string{"I", "V", "X", "L", "C", "D", "M"}
	i := 0
	for n > 0 {
		switch n % 10 {
		case 1:
			out = rn[i] + out
			break
		case 2:
			out = rn[i] + rn[i] + out
			break
		case 3:
			out = rn[i] + rn[i] + rn[i] + out
			break
		case 4:
			out = rn[i] + rn[i+1] + out
			break
		case 5:
			out = rn[i+1] + out
			break
		case 6:
			out = rn[i+1] + rn[i] + out
			break
		case 7:
			out = rn[i+1] + rn[i] + rn[i] + out
			break
		case 8:
			out = rn[i+1] + rn[i] + rn[i] + rn[i] + out
			break
		case 9:
			out = rn[i] + rn[i+2] + out
			break
		}
		n = n / 10
		i += 2
	}
	return out, nil

}
