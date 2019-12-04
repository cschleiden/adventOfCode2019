package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func OptimizeDistance() func(p *Point) int64 {
	return func(p *Point) int64 {
		return p.DistanceFromOrigin()
	}
}

func OptimizeSteps(w1 *wire, w2 *wire) func(p *Point) int64 {
	return func(p *Point) int64 {
		return w1.StepsToPoint(*p) + w2.StepsToPoint(*p)
	}
}

func IntersectWires(w1 *wire, w2 *wire, valueFunc func(p *Point) int64) int64 {
	intersection := w1.Intersection(w2)

	var min int64 = math.MaxInt64
	for _, p := range intersection {
		min = int64(math.Min(float64(min), float64(valueFunc(&p))))
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

	fmt.Println(IntersectWires(w1, w2, OptimizeSteps(w1, w2)))
}
