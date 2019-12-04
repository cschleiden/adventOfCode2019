package main

import (
	"strconv"
	"strings"
)

type Wire interface {
	PointInWire(p Point) bool
	StepsToPoint(p Point) int64
	Points() int
	Intersection(w2 Wire) []Point
}

type wire struct {
	path string
	// Point to number of steps it took to reach the point
	points map[Point]int64
}

func FromPath(path string) *wire {
	w := wire{
		path:   path,
		points: make(map[Point]int64),
	}

	segments := strings.Split(path, ",")

	var x, y, s int64

	for _, segment := range segments {

		steps, _ := strconv.Atoi(segment[1:len(segment)])
		for i := 0; i < steps; i++ {
			s++

			dx, dy := direction(segment[0])

			x += dx
			y += dy

			w.points[Point{x: x, y: y}] = s
		}
	}

	return &w
}

func (w *wire) PointInWire(p Point) bool {
	return w.points[p] != 0
}

func (w *wire) StepsToPoint(p Point) int64 {
	return w.points[p]
}

func (w *wire) Points() int {
	return len(w.points)
}

func (w *wire) Intersection(w2 Wire) []Point {
	intersection := make([]Point, 0)

	for p := range w.points {
		if w2.PointInWire(p) {
			intersection = append(intersection, p)
		}
	}

	return intersection
}

func direction(c byte) (int64, int64) {
	switch c {
	case 'R':
		return 1, 0
	case 'L':
		return -1, 0
	case 'U':
		return 0, -1
	case 'D':
		return 0, 1
	default:
		return 0, 0
	}
}
