package dna

import (
	"errors"
)

type Histogram map[byte]int
type Dna map[byte]int

func (d Dna) Counts() (Histogram, error) {
	if len(d) > 4 {
		return nil, errors.New("exist nucleotide invalid")
	}
	return Histogram(d), nil
}

func DNA(value string) Dna {
	dna := Dna{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, letter := range value {
		dna[byte(letter)]++
	}
	return dna
}
