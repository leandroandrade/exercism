package hamming

import (
	"strings"
	"errors"
)

func processHamming(s1, s2 string) int {
	if len(s1) == 1 && len(s2) == 1 && s1 != s2 {
		return 1
	}

	var total int

	s1Split := strings.Split(s1, "")
	s2Split := strings.Split(s2, "")

	for i := 0; i < len(s1); i++ {
		if s1Split[i] != s2Split[i] {
			total++
		}
	}
	return total
}

func Distance(s1, s2 string) (int, error) {
	if s1 == "" && s2 == "" || s1 == s2 {
		return 0, nil
	}

	if len(s1) != len(s2) {
		return -1, errors.New("nucleotides are different")
	}
	return processHamming(s1, s2), nil
}
