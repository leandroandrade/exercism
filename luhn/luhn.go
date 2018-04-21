package luhn

import (
	"strings"
	"unicode"
	"strconv"
)

func invalidLength(word string) bool {
	return len(strings.TrimSpace(word)) <= 1
}

func isNotBlank(r rune) bool {
	return len(strings.TrimSpace(string(r))) > 0
}

func notOnlyNumbers(letter string) bool {
	for _, value := range letter {
		if isNotBlank(value) {
			if !unicode.IsNumber(value) {
				return true
			}
		}
	}
	return false
}

func calculateNumber(number int, total *int) {
	number = number * 2
	if number > 9 {
		*total += number - 9
	} else {
		*total += number
	}
}

func invalidNumber(number string) bool {
	var total int
	number = strings.Replace(number, " ", "", -1)

	first, _ := strconv.Atoi(string(number[0]))
	if first <= 0 {
		number = number[1:]
	}

	for index, n := range number {
		i, _ := strconv.Atoi(string(n))
		if index%2 == 0 {
			calculateNumber(i, &total)
		} else {
			total += i
		}
	}
	return total%10 != 0
}

func Valid(value string) bool {
	if invalidLength(value) || notOnlyNumbers(value) || invalidNumber(value) {
		return false
	}
	return true
}
