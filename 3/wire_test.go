package main

import "testing"

type testpair struct {
	path           string
	numberOfPoints int
}

var tests = []testpair{
	{"R2", 2},
	{"R4", 4},
	{"R2,U2", 4},
}

func TestFromPath(t *testing.T) {
	for _, pair := range tests {
		if FromPath(pair.path).Points() != pair.numberOfPoints {
			t.Fail()
		}
	}
}

func TestPointInWire(t *testing.T) {
	w := FromPath("R2")

	if w.PointInWire(Point{x: 1, y: 0}) == false {
		t.Fail()
	}
}
