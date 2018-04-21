package raindrops

import (
	"strconv"
)

var plingPlangPlong = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

func Convert(number int) string {
	var primeUsed int
	var message string

	defaultNumber := number
	prime := 2

	for number > 1 {
		if number%prime == 0 {
			number = number / prime

			if value, exist := plingPlangPlong[prime]; exist && primeUsed != prime {
				message += value
				primeUsed = prime
			}
		} else {
			prime++
		}
	}

	if len(message) == 0 {
		return strconv.Itoa(defaultNumber)
	}
	return message
}
