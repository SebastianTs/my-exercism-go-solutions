package isbn

// IsValidISBN checks if an given string represents a valid ISBN
func IsValidISBN(isbn string) bool {

	sum := 0
	pos := 10
	for _, c := range isbn {
		switch {
		case c >= '0' && c <= '9':
			sum += (int(c) - '0') * pos
			pos--
		case c == 'X':
			sum += 10
			pos--
		}
	}
	return sum%11 == 0 && pos == 0
}
