package wordy

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const num string = `-?\d+`

type operator func(a, b int) int

func Answer(question string) (int, bool) {
	ops := strings.Join(getOps(), "|")
	parseRegex := fmt.Sprintf(`^What is (%s)((?: (?:%s) %s)+)\?$`, num, ops, num)
	match := regexp.MustCompile(parseRegex).FindStringSubmatch(question)
	if len(match) != 3 {
		return 0, false
	}

	result, _ := strconv.Atoi(match[1])
	tokenizer := fmt.Sprintf(" (%s) (%s)", ops, num)
	tokens := regexp.MustCompile(tokenizer).FindAllStringSubmatch(match[2], -1)
	for _, t := range tokens {
		op, _ := operators[t[1]]
		n, _ := strconv.Atoi(t[2])
		result = op(result, n)
	}
	return result, true
}

func getOps() []string {
	var opsList = make([]string, 0, len(operators))
	for k := range operators {
		opsList = append(opsList, k)
	}
	return opsList
}

var operators = map[string]operator{
	"plus":          plus,
	"minus":         minus,
	"multiplied by": mult,
	"divided by":    div,
}

func plus(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func mult(a, b int) int {
	return a * b
}

func div(a, b int) int {
	if b != 0 {
		return a / b
	}
	return 0
}
