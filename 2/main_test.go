package main

import (
	"testing"
)

func TestExecuteProgram_Addition(t *testing.T) {
	p := []int{1, 0, 0, 0, 99}
	ExecuteProgram(p)

	if p[0] != 2 {
		t.Fail()
	}
}

func TestExecuteProgram_Multiplication(t *testing.T) {
	p := []int{2, 3, 0, 3, 99}
	ExecuteProgram(p)

	if p[3] != 6 {
		t.Fail()
	}
}
