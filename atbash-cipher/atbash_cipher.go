package atbash

import "strings"

const groupSize = 5

func Atbash(str string) (result string) {
	str = strings.ToLower(str)
	offset := 'z' - 'a'
	i := 0
	for _, c := range str {
		switch {
		case c >= 'a' && c <= 'z':
			result += string('a' + offset - (c - 'a'))
			i++
		case c >= '0' && c <= '9':
			result += string(c)
			i++
		}
		if i == groupSize {
			result += " "
			i = 0
		}
	}
	result = strings.TrimSpace(result)
	return
}
