package diffsquares

import (
	"math"
)

func SquareOfSums(number int) int {
	var sum int
	for i := number; i >= 1; i-- {
		sum += i
	}
	return int(math.Pow(float64(sum), 2))
}

func SumOfSquares(number int) int {
	var sum int
	for i := 1; i <= number; i++ {
		sum += int(math.Pow(float64(i), 2))
	}
	return sum
}

func Difference(number int) int {
	squareOfSums := SquareOfSums(number)
	sumOfSquares := SumOfSquares(number)

	return squareOfSums - sumOfSquares
}
