package intcode

import (
	"fmt"
	"math"
)

type parameterMode int

const (
	positionMode  parameterMode = 0
	immediateMode               = 1
)

type opcode int

const (
	add         opcode = 1
	multiply           = 2
	save               = 3
	output             = 4
	jumpIfTrue         = 5
	jumpIfFalse        = 6
	lessThan           = 7
	equals             = 8
	halt               = 99
)

type instruction int

// Program state
type Program []int // can be instruction or data

// Run of program
type Run struct {
	P       Program
	Inputs  []int
	Outputs []int
}

func (i *instruction) decodeOpcode() opcode {
	return opcode(int(*i) % 100)
}

func (i *instruction) getParameterMode(parameterIdx int) parameterMode {
	// Strip of opcode
	x := int(*i) / 100

	// Get to desired position
	x = x / int(math.Pow10(parameterIdx))

	return parameterMode(x % 10)
}

func (p *Program) getParameterValue(ip int, mode parameterMode) int {
	pv := p.getValue(ip)

	switch mode {
	case positionMode:
		{
			return p.getValue(pv)
		}
	case immediateMode:
		{
			return pv
		}
	}

	panic("invalid parameter mode")
}

func (p *Program) getParameters(ip int, ins instruction, num int, numOutputs int) []int {
	params := make([]int, num)

	for i := 0; i < num; i++ {
		pm := ins.getParameterMode(i)
		if i >= num-numOutputs {
			pm = immediateMode
		}
		params[i] = p.getParameterValue(ip+i+1, pm)
	}

	return params
}

func (p *Program) getValue(address int) int {
	return (*p)[address]
}

func (p *Program) writeValue(address int, value int) {
	(*p)[address] = value
}

// Execute program
func (r *Run) Execute() {
	p := r.P
	inputPtr := 0

	for i := 0; i < len(p); {
		instruction := instruction(p.getValue(i))
		opcode := instruction.decodeOpcode()
		switch opcode {
		case add:
			{
				params := p.getParameters(i, instruction, 3, 1)
				p.writeValue(params[2], params[0]+params[1])
				i += 4
			}
		case multiply:
			{
				params := p.getParameters(i, instruction, 3, 1)
				p.writeValue(params[2], params[0]*params[1])
				i += 4
			}
		case save:
			{
				params := p.getParameters(i, instruction, 1, 1)
				p.writeValue(params[0], r.Inputs[inputPtr])
				inputPtr++
				i += 2
			}
		case jumpIfTrue:
			{
				params := p.getParameters(i, instruction, 2, 0)
				if params[0] != 0 {
					i = params[1]
				} else {
					i += 3
				}
			}
		case jumpIfFalse:
			{
				params := p.getParameters(i, instruction, 2, 0)
				if params[0] == 0 {
					i = params[1]
				} else {
					i += 3
				}
			}
		case lessThan:
			{
				params := p.getParameters(i, instruction, 3, 1)
				if params[0] < params[1] {
					p.writeValue(params[2], 1)
				} else {
					p.writeValue(params[2], 0)
				}
				i += 4
			}
		case equals:
			{
				params := p.getParameters(i, instruction, 3, 1)
				if params[0] == params[1] {
					p.writeValue(params[2], 1)
				} else {
					p.writeValue(params[2], 0)
				}
				i += 4
			}
		case output:
			{
				params := p.getParameters(i, instruction, 1, 0)
				r.Outputs = append(r.Outputs, params[0])
				i += 2
			}
		case halt:
			return
		default:
			panic(fmt.Sprint("Unknown opcode", i, opcode))
		}
	}
}
