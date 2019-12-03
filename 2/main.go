package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ExecuteProgram(program []int) {
	for i := 0; i < len(program); i += 4 {
		switch program[i] {
		case 1:
			{
				// Addition
				lhs := program[program[i+1]]
				rhs := program[program[i+2]]

				program[program[i+3]] = lhs + rhs
			}
		case 2:
			{
				// Multiplication
				lhs := program[program[i+1]]
				rhs := program[program[i+2]]

				program[program[i+3]] = lhs * rhs
			}
		case 99:
			return
		}
	}
}

func main() {
	file, err := os.Open("./input2.txt")
	if err != nil {
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	inputLine := scanner.Text()
	inputStrings := strings.Split(inputLine, ",")

	inputs := make([]int, len(inputStrings))
	for i, v := range inputStrings {
		inputs[i], err = strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// Modify
	inputs[1] = 12
	inputs[2] = 2

	ExecuteProgram(inputs)

	// Output
	fmt.Println(inputs[0])
}
