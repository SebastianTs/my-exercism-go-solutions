package house

import "fmt"

const testVersion = 1

var terms = [][]string{
	{"house that Jack built.", ""},
	{"malt", "lay in"},
	{"rat", "ate"},
	{"cat", "killed"},
	{"dog", "worried"},
	{"cow with the crumpled horn", "tossed"},
	{"maiden all forlorn", "milked"},
	{"man all tattered and torn", "kissed"},
	{"priest all shaven and shorn", "married"},
	{"rooster that crowed in the morn", "woke"},
	{"farmer sowing his corn", "kept"},
	{"horse and the hound and the horn", "belonged to"},
}

// Song builds the song with all verses
func Song() (out string) {
	for i := 1; i <= len(terms); i++ {
		out += Verse(i)
		if i != len(terms) {
			out += "\n\n"
		}
	}
	fmt.Println(out)
	return out
}

// Verse returns the i-th Verse
func Verse(i int) string {
	return buildVerse(i, true)
}

// buildVerse builds the Verse recursive
func buildVerse(i int, withBeginning bool) (out string) {
	if i > 0 && i <= len(terms) {
		line := terms[i-1]

		if withBeginning {
			out += "This is the "
		}
		out += line[0]

		if line[1] != "" {
			out += "\nthat " + line[1] + " the " + buildVerse(i-1, false)
		}
	}
	return out
}
