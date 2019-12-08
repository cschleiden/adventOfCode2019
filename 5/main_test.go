package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstructionDecoding(t *testing.T) {
	i := instruction(1002)
	assert.Equal(t, 2, i.decodeOpcode())

	i = instruction(3)
	assert.Equal(t, 3, i.decodeOpcode())
}

type testinput struct {
	instrution instruction
	pIdx       int
	pm         parameterMode
}

func TestParameterMode(t *testing.T) {
	var tests = []testinput{
		{1002, 0, positionMode},
		{1002, 1, immediateMode},
		{1002, 2, positionMode},
		{4, 2, positionMode},
		{1144, 0, immediateMode},
		{1144, 1, immediateMode},
	}

	for _, x := range tests {
		assert.Equal(t, x.pm, x.instrution.getParameterMode(x.pIdx))
	}
}
