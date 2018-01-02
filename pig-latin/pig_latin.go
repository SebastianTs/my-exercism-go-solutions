package piglatin

import (
	"strings"
)

func Sentence(str string) string {

	list := strings.Split(str, " ")
	sentence := make([]string, len(list))
	for i, word := range list {
		sentence[i] = Latenize(word)
	}
	return strings.Join(sentence, " ")
}

func Latenize(in string) string {
	out := ""
	for i, s := range in {
		if isVowel(byte(s)) && (i < 1 || (s != 'u' || in[i-1] != 'q')) {
			out = in[i:] + in[:i]
			break
		} else if (s == 'y' || s == 'x') && (i < len(in)-1 && !isVowel(in[i+1])) {
			out = in[i:] + in[:i]
			break
		} else if s == 'y' && i == 1 {
			out = in[i:] + in[:i]
			break
		}
	}
	return out + "ay"
}

func isVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}
