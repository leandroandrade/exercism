package isbn

import (
	"strings"
	"unicode"
	"strconv"
)

var x int32 = 120

func isEmpty(word string) bool {
	return len(strings.TrimSpace(word)) == 0
}

func isLastValueIsDigitX(isbn string) bool {
	return int32(isbn[len(isbn)-1]) == x
}

func calculateIsbn(isbn string) bool {
	var total int
	count := 10

	for _, number := range isbn {
		if unicode.IsNumber(number) {
			number, _ := strconv.Atoi(string(number))
			total += number * count
			count -= 1
		}
	}

	if isLastValueIsDigitX(isbn) {
		total += 10
	}
	return total%11 == 0
}

func IsValidISBN(isbn string) bool {
	if isEmpty(isbn) {
		return false
	}

	isbn = strings.Replace(isbn, "-", "", -1)

	if (len(isbn)) > 10 {
		return false
	}
	return calculateIsbn(strings.ToLower(isbn))
}
