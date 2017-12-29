package brackets

func Bracket(in string) (bool, error) {
	opener := map[rune]rune{
		']': '[',
		')': '(',
		'}': '{',
	}
	stack := make([]rune, 0)
	for _, c := range in {
		switch c {
		case '[', '(', '{':
			stack = append(stack, c)
		case ']', ')', '}':
			if len(stack) == 0 || stack[len(stack)-1] != opener[c] {
				return false, nil
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0, nil
}
