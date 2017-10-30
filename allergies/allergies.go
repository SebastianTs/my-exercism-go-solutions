// Package allergies implements simple conversion
package allergies

const testVersion = 1

type allergenID struct {
	n uint
	s string
}

var list = []allergenID{
	{1, "eggs"},
	{2, "peanuts"},
	{4, "shellfish"},
	{8, "strawberries"},
	{16, "tomatoes"},
	{32, "chocolate"},
	{64, "pollen"},
	{128, "cats"},
}

// Allergies outputs the list of allergens to an id n
func Allergies(n uint) (out []string) {
	for _, entry := range list {
		if n&entry.n == entry.n {
			out = append(out, entry.s)
		}
	}
	return out

}

// AllergicTo tests if an id n is allgeric to an allergen in
func AllergicTo(n uint, in string) bool {
	for _, item := range Allergies(n) {
		if item == in {
			return true
		}
	}
	return false
}
