package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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
type program []int // can be instruction or data

type run struct {
	p       program
	inputs  []int
	outputs []int
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

func (p *program) getParameterValue(ip int, mode parameterMode) int {
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

func (p *program) getParameters(ip int, ins instruction, num int, numOutputs int) []int {
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

func (p *program) getValue(address int) int {
	return (*p)[address]
}

func (p *program) writeValue(address int, value int) {
	(*p)[address] = value
}

func (r *run) execute() {
	p := r.p
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
				p.writeValue(params[0], r.inputs[inputPtr])
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
				r.outputs = append(r.outputs, params[0])
				i += 2
			}
		case halt:
			return
		default:
			panic(fmt.Sprint("Unknown opcode", i, opcode))
		}
	}
}

func main() {
	file, err := os.Open("./input5.txt")
	if err != nil {
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	inputLine := scanner.Text()
	inputStrings := strings.Split(inputLine, ",")

	p := make(program, len(inputStrings))
	for i, v := range inputStrings {
		p[i], err = strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	r := run{
		p:      p,
		inputs: []int{5},
	}

	r.execute()

	for _, o := range r.outputs {
		fmt.Println(o)
	}
}
