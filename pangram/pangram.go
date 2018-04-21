package pangram

import (
	"strings"
	"unicode"
)

var quantityLettersAlphabet = 26
var emptySpace int32 = 32

func isEmpty(word string) bool {
	return len(strings.TrimSpace(word)) == 0
}

func IsPangram(sentence string) bool {
	if isEmpty(sentence) {
		return false
	}

	letters := make(map[int32]bool)
	sentence = strings.ToLower(sentence)

	for _, letter := range sentence {
		if letter != emptySpace && unicode.IsLetter(letter) {
			if _, exist := letters[letter]; !exist {
				letters[letter] = true
			}
		}
	}
	return len(letters) == quantityLettersAlphabet
}
