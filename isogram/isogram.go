package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {
	if isEmpty(word) {
		return true
	}

	word = strings.Replace(word, "-", "", -1)
	word = strings.Replace(word, " ", "", -1)
	word = strings.ToLower(word)

	values := map[int32]bool{}

	for _, char := range word {
		if _, exists := values[char]; !exists {
			values[char] = true
		} else {
			return false
		}
	}
	return true
}

func isEmpty(word string) bool {
	return len(strings.TrimSpace(word)) == 0
}
