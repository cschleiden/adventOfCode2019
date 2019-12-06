package main

import (
	"fmt"
	"strconv"
)

func isValid(input int) bool {
	previousDigit := -1
	digitCount := make(map[int]int)

	s := strconv.Itoa(input)
	for _, x := range s {
		digit, _ := strconv.Atoi(string(x))

		digitCount[digit] = digitCount[digit] + 1

		if previousDigit > digit {
			return false
		}

		previousDigit = digit
	}

	for _, v := range digitCount {
		if v == 2 {
			return true
		}
	}

	return false
}

func main() {
	rangeStart := 147981
	rangeEnd := 691423

	valid := 0
	for i := rangeStart; i <= rangeEnd; i++ {
		if isValid(i) {
			valid++
		}
	}

	fmt.Println(valid)
}
