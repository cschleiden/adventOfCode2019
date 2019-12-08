package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	intcode "github.com/cschleiden/adventofcode/Intcode"
)

const numAmplifiers = 5
const numPhaseSettings = 4

func run(p intcode.Program, phases []int) int {
	previousResult := 0

	for a := 0; a < numAmplifiers; a++ {
		phaseSetting := phases[a]

		pa := make([]int, len(p))
		copy(pa, p)

		r := &intcode.Run{
			P:      pa,
			Inputs: []int{phaseSetting, previousResult},
		}
		r.Execute()

		previousResult = r.Outputs[0]
	}

	return previousResult
}

func generatePermutations(f func(input []int)) {
	perm([]int{0, 1, 2, 3, 4}, f, 0)
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
