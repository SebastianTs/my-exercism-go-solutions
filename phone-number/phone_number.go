package phonenumber

import (
	"errors"
	"fmt"
)

func Number(str string) (string, error) {
	number := ""
	if str[0] == '+' {
		str = str[2:]
	}
	if str[0] == '1' {
		str = str[1:]
	}
	for _, c := range str {
		if c >= '0' && c <= '9' {
			number += string(c)
		}
	}
	if number[0] <= '1' || number[3] <= '1' {
		return "", errors.New("Invalid area code")
	}
	if len(number) != 10 {
		return "", errors.New("Invalid length")
	}

	return number, nil
}

func AreaCode(str string) (string, error) {
	areaCode, err := Number(str)
	if err != nil {
		return "", err
	}
	return areaCode[:3], nil
}

func Format(str string) (string, error) {
	str, err := Number(str)
	if err != nil {
		return "", err
	}
	areaCode, err := AreaCode(str)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", areaCode, str[3:6], str[6:]), err
}
