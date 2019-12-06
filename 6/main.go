package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//
	// orbiting := make(map[string][]string)
	orbits := make(map[string]string)
	objects := make([]string, 0)

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
		// orbiting[input[0]] = append(orbiting[input[0]], input[1])
		orbits[input[1]] = input[0]
		objects = append(objects, input[1])
	}

	total := 0

	for _, object := range objects {
		o := orbits[object]
		for o != "" {
			total++
			o = orbits[o]
		}
	}

	fmt.Println(total)
}
