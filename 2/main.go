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

	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {
			fmt.Println("Trying", i, j)

			// Duplicate program to reset inputs
			p := make([]int, len(inputs))
			copy(p, inputs)

			p[1] = i
			p[2] = j

			ExecuteProgram(p)

			if p[0] == 19690720 {
				fmt.Println(i, j, 100*i+j)
				return
			}
		}
	}
}
