package dna

import (
	"errors"
)

const testVersion = 2

type DNA string
type nucleotide byte

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

func (n nucleotide) isValid() bool {
	switch n {
	case 'A', 'C', 'G', 'T':
		return true
	default:
		return false
	}

}

func (d DNA) isValid() bool {
	for _, v := range d {
		if !nucleotide(v).isValid() {
			return false
		}
	}
	return true
}

type Histogram map[rune]int

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
