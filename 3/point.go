package main

import "math"

type Point struct {
	x int64
	y int64
}

func (p *Point) DistanceFromOrigin() int {
	return int(math.Abs(float64(p.x)) + math.Abs(float64(p.y)))
}
