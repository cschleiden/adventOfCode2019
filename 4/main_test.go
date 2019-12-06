package main

import "testing"

type testpair struct {
	input   int
	isValid bool
}

var tests = []testpair{
	{112233, true},
	{123444, false},
	{124444, false},
	{113444, true},
	{111122, true},
	{112233, true},
	{123789, false},
	{223450, false},
	{111445, true},
	{689999, false},
}

func Test_Validate(t *testing.T) {
	for _, pair := range tests {
		if isValid(pair.input) != pair.isValid {
			t.Error(pair.input, " should be ", pair.isValid)
		}
	}
}
