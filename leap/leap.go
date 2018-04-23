package leap

func divisibleBy(year, divisor int) bool {
	return year%divisor == 0
}

func IsLeapYear(year int) bool {
	return (divisibleBy(year, 4) && !divisibleBy(year, 100)) || divisibleBy(year, 400)
}
