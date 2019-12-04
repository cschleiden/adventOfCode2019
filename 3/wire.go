package main

import (
	"strconv"
	"strings"
)

type Wire interface {
	PointInWire(p Point) bool
	Points() int
	Intersection(w2 Wire) []Point
}

type wire struct {
	path   string
	points map[Point]bool
}

func FromPath(path string) *wire {
	w := wire{
		path:   path,
		points: make(map[Point]bool),
	}

	segments := strings.Split(path, ",")

	var x, y int64

	for _, segment := range segments {

		steps, _ := strconv.Atoi(segment[1:len(segment)])
		for i := 0; i < steps; i++ {
			dx, dy := direction(segment[0])

			x += dx
			y += dy

			w.points[Point{x: x, y: y}] = true
		}
	}

	return &w
}

func (w *wire) PointInWire(p Point) bool {
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
