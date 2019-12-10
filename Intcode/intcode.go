package intcode

import (
	"fmt"
	"math"
)

type parameterMode int

const (
	positionMode  parameterMode = 0
	immediateMode               = 1
	relativeMode                = 2
)

type opcode int

const (
	add          opcode = 1
	multiply            = 2
	save                = 3
	output              = 4
	jumpIfTrue          = 5
	jumpIfFalse         = 6
	lessThan            = 7
	equals              = 8
	relativeBase        = 9
	halt                = 99
)

type instruction int

// Program state
type Program []int64 // can be instruction or data

type Execution struct {
	InstructionPointer int
	RelativeBase       int
	RAM                map[int64]int64
}

// Run of program
type Run struct {
	// Optional identifier for debugging
	Identifier *int
	P          Program
	Inputs     chan int64
	Outputs    chan int64
	Done       chan bool

	Execution Execution
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

func (r *Run) getParameterValue(ip int, mode parameterMode) (int64, int64) {
	pv := r.getValue(int64(ip))

	switch mode {
	case positionMode:
		{
			return r.getValue(pv), pv
		}
	case immediateMode:
		{
			return pv, pv
		}
	case relativeMode:
		{
			return r.getValue(int64(r.Execution.RelativeBase) + pv), int64(r.Execution.RelativeBase) + pv
		}
	}

	panic("invalid parameter mode")
}

func (r *Run) getParameters(ins instruction, num int, numOutputs int) []int64 {
	ip := r.Execution.InstructionPointer
	params := make([]int64, num)

	for i := 0; i < num; i++ {
		pm := ins.getParameterMode(i)
		if i >= num-numOutputs {
			if pm == relativeMode {
				_, rawpv := r.getParameterValue(ip+i+1, pm)
				params[i] = rawpv
			} else {
				// map to immediate mode
				params[i], _ = r.getParameterValue(ip+i+1, immediateMode)
			}
		} else {
			params[i], _ = r.getParameterValue(ip+i+1, pm)
		}
	}

	return params
}

func (r *Run) getValue(address int64) int64 {
	if v, ok := r.Execution.RAM[address]; ok {
		return v
	}

	if address < int64(len(r.P)) {
		return r.P[address]
	}

	return 0
}

func (r *Run) writeValue(address int64, value int64) {
	if address > int64(len(r.P)) {
		r.Execution.RAM[address] = value
	} else {
		r.P[address] = value
	}
}

var debug = true

// Execute program
func (r *Run) Execute() {
	myID := -1
	if r.Identifier != nil {
		myID = *r.Identifier
	}

	r.Execution = Execution{
		InstructionPointer: 0,
		RAM:                make(map[int64]int64),
		RelativeBase:       0,
	}

main:
	for r.Execution.InstructionPointer < len(r.P) {
		previousIP := r.Execution.InstructionPointer

		instruction := instruction(r.getValue(int64(r.Execution.InstructionPointer)))
		opcode := instruction.decodeOpcode()
		switch opcode {
		case add:
			{
				params := r.getParameters(instruction, 3, 1)
				r.writeValue(params[2], params[0]+params[1])
				r.Execution.InstructionPointer += 4
			}
		case multiply:
			{
				params := r.getParameters(instruction, 3, 1)
				r.writeValue(params[2], params[0]*params[1])
				r.Execution.InstructionPointer += 4
			}
		case save:
			{
				params := r.getParameters(instruction, 1, 1)
				if debug {
					fmt.Println(myID, " Waiting for input")
				}
				input := <-r.Inputs
				if debug {
					fmt.Println(myID, " Received ", input)
				}
				r.writeValue(params[0], input)
				r.Execution.InstructionPointer += 2
			}
		case jumpIfTrue:
			{
				params := r.getParameters(instruction, 2, 0)
				if params[0] != 0 {
					r.Execution.InstructionPointer = int(params[1])
				} else {
					r.Execution.InstructionPointer += 3
				}
			}
		case jumpIfFalse:
			{
				params := r.getParameters(instruction, 2, 0)
				if params[0] == 0 {
					r.Execution.InstructionPointer = int(params[1])
				} else {
					r.Execution.InstructionPointer += 3
				}
			}
		case lessThan:
			{
				params := r.getParameters(instruction, 3, 1)
				if params[0] < params[1] {
					r.writeValue(params[2], 1)
				} else {
					r.writeValue(params[2], 0)
				}
				r.Execution.InstructionPointer += 4
			}
		case equals:
			{
				params := r.getParameters(instruction, 3, 1)
				if params[0] == params[1] {
					r.writeValue(params[2], 1)
				} else {
					r.writeValue(params[2], 0)
				}
				r.Execution.InstructionPointer += 4
			}
		case output:
			{
				params := r.getParameters(instruction, 1, 0)
				if debug {
					fmt.Println(myID, " Outputting ", params[0])
				}
				r.Outputs <- params[0]
				r.Execution.InstructionPointer += 2
			}
		case relativeBase:
			{
				params := r.getParameters(instruction, 1, 0)
				r.Execution.RelativeBase += int(params[0])
				r.Execution.InstructionPointer += 2
			}
		case halt:
			if debug {
				fmt.Println(myID, " halt.")
			}
			break main
		default:
			panic(fmt.Sprint("Unknown opcode", r.Execution.InstructionPointer, opcode))
		}

		if previousIP == r.Execution.InstructionPointer {
			panic("InstructionPointer is stuck")
		}
	}

	r.Done <- true
}
