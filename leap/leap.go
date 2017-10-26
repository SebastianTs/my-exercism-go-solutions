package leap

const testVersion = 3

// IsLeapYear tests if a given int value represents a Leap Year
func IsLeapYear(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false
}
