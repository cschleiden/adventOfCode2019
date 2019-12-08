package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"

	intcode "github.com/cschleiden/adventofcode/Intcode"
)

const numAmplifiers = 5
const numPhaseSettings = 4

func run(p intcode.Program, phases []int) int {
	var wg sync.WaitGroup

	var inOut [numAmplifiers]chan int

	for a := 0; a < numAmplifiers; a++ {
		inOut[a] = make(chan int, 1)
		inOut[a] <- phases[a]
	}

	for a := 0; a < numAmplifiers; a++ {
		pa := make([]int, len(p))
		copy(pa, p)

		wg.Add(1)

		go func(i int, p2 intcode.Program) {
			defer wg.Done()

			r := &intcode.Run{
				Identifier: &i,
				P:          p2,
				Inputs:     inOut[i],
				Outputs:    inOut[(i+1)%numAmplifiers],
			}

			r.Execute()
		}(a, pa)
	}

	// Provide A with its initial input
	inOut[0] <- 0

	// Wait for all programs to halt
	wg.Wait()

	// Capture E's output
	result := <-inOut[0]
	return result
}

func generatePermutations(f func(input []int)) {
	perm([]int{5, 6, 7, 8, 9}, f, 0)
}

// being lazy: https://yourbasic.org/golang/generate-permutation-slice-string/
func perm(a []int, f func([]int), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func main() {
	file, err := os.Open("./input7.txt")
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

	// Program is parsed, generate permutations and run computers
	result := math.MinInt32

	generatePermutations(func(input []int) {
		result = int(math.Max(float64(result), float64(run(p, input))))
	})

	fmt.Println(result)
}
