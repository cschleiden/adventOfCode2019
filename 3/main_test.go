package main

import "testing"

func TestIntersectWires_Distane(t *testing.T) {
	w1 := FromPath("R75,D30,R83,U83,L12,D49,R71,U7,L72")
	w2 := FromPath("U62,R66,U55,R34,D71,R55,D58,R83")

	m := IntersectWires(w1, w2, OptimizeDistance())
	if m != 159 {
		t.Error(m)
	}
}

func TestIntersectWires_Steps(t *testing.T) {
	w1 := FromPath("R75,D30,R83,U83,L12,D49,R71,U7,L72")
	w2 := FromPath("U62,R66,U55,R34,D71,R55,D58,R83")

	m := IntersectWires(w1, w2, OptimizeSteps(w1, w2))
	if m != 610 {
		t.Error(m)
	}
}
