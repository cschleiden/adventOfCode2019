package main

import (
	"fmt"
	"strconv"
)

func isValid(input int) bool {
	adjacent := false
	previousDigit := -1

	s := strconv.Itoa(input)
	for _, x := range s {
		digit, _ := strconv.Atoi(string(x))

		if previousDigit == digit && !adjacent {
			adjacent = true
		}

		if previousDigit > digit {
			return false
		}

		previousDigit = digit
	}

	return adjacent
}

func main() {
	rangeStart := 147981
	rangeEnd := 691423

	isValid(1223)

	valid := 0
	for i := rangeStart; i <= rangeEnd; i++ {
		if isValid(i) {
			valid++
		}
	}

	fmt.Println(valid)
}
