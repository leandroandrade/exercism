package cipher

import (
	"strings"
	"regexp"
)

type Cipher interface {
	Encode(text string) string
	Decode(text string) string
}

type Caesar struct {
}

func NewCaesar() *Caesar {
	return &Caesar{}
}

func (c Caesar) Encode(word string) string {
	return process(prepareMessage(word), 3)
}

func (c Caesar) Decode(word string) string {
	return process(prepareMessage(word), -3)
}

func process(word string, diff int) string {
	letters := make([]string, len(word))
	for _, char := range word {
		letters = append(letters, string(wordAlphabet(char, diff)))
	}
	return strings.Join(letters, "")
}

func wordAlphabet(r rune, diff int) rune {
	n := r + rune(diff)
	if n > 'z' {
		n -= 'z' - 'a' + 1
	}
	if n < 'a' {
		n += 'z' - 'a' + 1
	}
	return n
}

func prepareMessage(m string) string {
	m = strings.ToLower(m)
	re := regexp.MustCompile("[^a-z]")
	return re.ReplaceAllString(m, "")
}
