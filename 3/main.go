package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func IntersectWires(w1 *wire, w2 *wire) int64 {
	intersection := w1.Intersection(w2)

	var min int64 = math.MaxInt64
	for _, p := range intersection {
		min = int64(math.Min(float64(min), float64(p.DistanceFromOrigin())))
	}

	return min
}

func main() {
	file, _ := os.Open("./input3.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	w1 := FromPath(scanner.Text())

	scanner.Scan()
	w2 := FromPath(scanner.Text())

	fmt.Println(IntersectWires(w1, w2))
}
