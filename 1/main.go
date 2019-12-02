package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calculateFuel(input int) int {
	return input/3 - 2
}

func main() {
	fmt.Println("Hello World")

	file, err := os.Open("./input1.txt")
	if err != nil {
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var result int

	for scanner.Scan() {
		line := scanner.Text()
		input, err := strconv.Atoi(line)
		if err != nil {
			os.Exit(2)
		}
		result += calculateFuel(input)
	}

	fmt.Println(result)
}
