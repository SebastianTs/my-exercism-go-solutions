// Package protein implements the protein exercise
package protein

const testVersion = 1

// FromCondon translates one RNA Codons to Proteins
func FromCodon(in string) string {
	var m map[string]string = map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP",
		"UAG": "STOP",
		"UGA": "STOP",
	}
	return m[in]
}

// FromRNA translates polypetides to a sequence of proteins
// stops if a STOP codon is read
func FromRNA(in string) (out []string) {
	for i := 0; i < len(in); i += 3 {
		cur := FromCodon(in[i : i+3])
		if cur == "STOP" {
			break
		}
		out = append(out, cur)
	}
	return out
}
