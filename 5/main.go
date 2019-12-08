package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	intcode "github.com/cschleiden/adventofcode/Intcode"
)

func main() {
	file, err := os.Open("./input5.txt")
	if err != nil {
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	inputLine := scanner.Text()
	inputStrings := strings.Split(inputLine, ",")

	p := make(intcode.Program, len(inputStrings))
	for i, v := range inputStrings {
		p[i], err = strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	r := intcode.Run{
		P:      p,
		Inputs: []int{5},
	}

	r.Execute()

	for _, o := range r.Outputs {
		fmt.Println(o)
	}
}
