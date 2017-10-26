package twelve

import "fmt"

const testVersion = 1

var (
	days = []string{"a Partridge", "two Turtle Doves",
		"three French Hens", "four Calling Birds", "five Gold Rings",
		"six Geese-a-Laying", "seven Swans-a-Swimming",
		"eight Maids-a-Milking", "nine Ladies Dancing",
		"ten Lords-a-Leaping", "eleven Pipers Piping",
		"twelve Drummers Drumming"}
	numbers = []string{"first", "second", "third", "fourth", "fifth", "sixth",
		"seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}

	text = "On the %s day of Christmas my true love gave to me, %s in a Pear Tree."
)

// Verse outputs the n-th verse of the song
func Verse(n int) string {
	present := ""
	for i := n; i > 0; i-- {
		present += days[i-1]
		if n > 1 && i == 2 {
			present += ", and "
		}
		if i != 2 && i != 1 {
			present += ", "
		}
	}
	return fmt.Sprintf(text, numbers[n-1], present)
}

// Song outputs all 12 Verses of the song.
func Song() string {
	str := ""
	for i := 1; i <= 12; i++ {
		str += Verse(i) + "\n"
	}
	return str
}
