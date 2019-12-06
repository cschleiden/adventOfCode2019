package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	//
	orbits := make(map[string][]string)

	// Parse
	file, err := os.Open("./input6.txt")
	if err != nil {
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()

		input := strings.Split(line, ")")
		// can move in both directions
		orbits[input[1]] = append(orbits[input[1]], input[0])
		orbits[input[0]] = append(orbits[input[0]], input[1])
	}

	min := math.MaxInt32
	var candidates stack
	for _, o := range orbits["YOU"] {
		candidates = candidates.Push(e{o, 0})
	}
	visited := make(map[string]bool)
	for len(candidates) > 0 {
		var x e
		candidates, x = candidates.Pop()
		visited[x.object] = true

		if x.object == "SAN" {
			min = int(math.Min(float64(min), float64(x.path)))
		} else {
			// Enqueue children
			for _, next := range orbits[x.object] {
				if !visited[next] {
					candidates = candidates.Push(e{next, x.path + 1})
				}
			}
		}
	}

	// We don't have to reach SAN, just get into the same orbit
	fmt.Println(min - 1)
}
