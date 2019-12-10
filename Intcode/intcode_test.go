package intcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeChannel(input []int64) chan int64 {
	c := make(chan int64, len(input))

	for _, v := range input {
		c <- v
	}

	return c
}

func TestInstructionDecoding(t *testing.T) {
	i := instruction(1002)
	assert.Equal(t, opcode(2), i.decodeOpcode())

	i = instruction(3)
	assert.Equal(t, opcode(3), i.decodeOpcode())
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

func TestEquals(t *testing.T) {
	r := &Run{
		P:       []int64{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
		Inputs:  makeChannel([]int64{8}),
		Outputs: make(chan int64, 5),
	}
	r.Execute()

	assert.Equal(t, int64(1), <-r.Outputs)
}

func TestEqualsImmediate(t *testing.T) {
	r := &Run{
		P:       []int64{3, 3, 1108, -1, 8, 3, 4, 3, 99},
		Inputs:  makeChannel([]int64{8}),
		Outputs: make(chan int64, 5),
	}
	r.Execute()
	assert.Equal(t, int64(1), <-r.Outputs)
}

func TestLessThan(t *testing.T) {
	r := &Run{
		P:       []int64{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
		Inputs:  makeChannel([]int64{5}),
		Outputs: make(chan int64, 5),
	}
	r.Execute()
	assert.Equal(t, int64(1), <-r.Outputs)
}

func TestLessThanImmediate(t *testing.T) {
	r := &Run{
		P:       []int64{3, 9, 1107, 9, 10, 9, 4, 9, 99, -1, 8},
		Inputs:  makeChannel([]int64{5}),
		Outputs: make(chan int64, 5),
	}
	r.Execute()
	assert.Equal(t, int64(1), <-r.Outputs)
}

type programTest struct {
	input           []int64
	Inputs          []int64
	expectedOutputs []int64
}

func TestProgram(t *testing.T) {
	testData := []programTest{
		{
			[]int64{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			[]int64{5},
			[]int64{999},
		},
		{
			[]int64{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			[]int64{8},
			[]int64{1000},
		},
		{
			[]int64{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			[]int64{10},
			[]int64{1001},
		},
		{
			[]int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
			[]int64{},
			[]int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		},
		{
			[]int64{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
			[]int64{},
			[]int64{1219070632396864},
		},
	}

	for _, test := range testData {
		r := &Run{
			P:       test.input,
			Inputs:  makeChannel(test.Inputs),
			Outputs: make(chan int64, len(test.expectedOutputs)),
		}

		go func() {
			r.Execute()
		}()
		<-r.Done

		assert.Equal(t, len(test.expectedOutputs), len(r.Outputs))

		for _, o := range test.expectedOutputs {
			assert.Equal(t, o, <-r.Outputs)
		}
	}
}
