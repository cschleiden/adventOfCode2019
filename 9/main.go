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
	file, err := os.Open("./input9.txt")
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
		v, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			p[i] = int64(v)
		}
	}

	r := &intcode.Run{
		P:       p,
		Inputs:  make(chan int64, 1),
		Outputs: make(chan int64),
		Done:    make(chan bool),
	}

	go func() {
		for {
			r := <-r.Outputs
			fmt.Println(r)
		}
	}()

	go func() { r.Inputs <- 2 }()

	r.Execute()
}
