package dna

import (
	"errors"
)

const testVersion = 2

// DNA is the model for an DNA string
type DNA string
type nucleotide byte

// Count returns the number ob nucleotides of an DNA string or an Error
// if the string contains invalid characters
func (d DNA) Count(b byte) (int, error) {
	n := nucleotide(b)
	if !n.isValid() {
		return 0, errors.New("Value is not a valid alphabet of DNA")
	}
	sum := 0
	for _, v := range d {
		if byte(v) == b {
			sum++
		}
	}
	return sum, nil
}

// isValid checks if a nucleotide is a valid DNA character
func (n nucleotide) isValid() bool {
	switch n {
	case 'A', 'C', 'G', 'T':
		return true
	default:
		return false
	}

}

// isValid checks if a string is a valid DNA string
func (d DNA) isValid() bool {
	for _, v := range d {
		if !nucleotide(v).isValid() {
			return false
		}
	}
	return true
}

// Histogram is a map that counts the number of character in an DNA string
type Histogram map[rune]int

// Counts parts of an dna string and returns an error if an dna is not valid
func (d DNA) Counts() (h Histogram, err error) {
	if !d.isValid() {
		return nil, errors.New("DNA is not valid")
	}
	alpha := []byte{'A', 'C', 'G', 'T'}
	h = make(map[rune]int)
	for _, v := range alpha {
		h[rune(v)], err = d.Count(v)
		if err != nil {
			return nil, err
		}
	}
	return h, nil
}
