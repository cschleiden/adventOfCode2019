package main

import (
	"fmt"
	intcode "github.com/cschleiden/adventofcode/Intcode"
	"math"
)

type point struct {
	x, y int64
}

const (
	BLACK = iota
	WHITE
)

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

func main() {
	p := intcode.ReadFromFile("./input.txt")

	r := &intcode.Run{
		P:       p,
		Inputs:  make(chan int64),
		Outputs: make(chan int64),
	}

	w := make(map[point]int64)
	min := point{math.MaxInt64, math.MaxInt64}
	max := point{math.MinInt64, math.MinInt64}

	c := 0

	go (func() {

		direction := UP
		x := int64(100)
		y := int64(100)

		for {
			// Paint
			color := <-r.Outputs
			if _, ok := w[point{x, y}]; !ok {
				c++
			}
			w[point{x, y}] = color

			// Move
			turn := <-r.Outputs

			// Turn
			if turn == 0 {
				direction--
			} else {
				direction++
			}

			direction = (4 + direction) % 4

			// Move
			switch direction {
			case UP:
				{
					y--
				}
			case RIGHT:
				{
					x++
				}
			case DOWN:
				{
					y++
				}
			case LEFT:
				{
					x--
				}
			}

			// Read color at new location
			r.Inputs <- w[point{x, y}]

			fmt.Println(turn, " ", direction, " ", x, " ", y)

			if x < min.x {
				min.x = x
			}
			if y < min.y {
				min.y = y
			}
			if x > max.x {
				max.x = x
			}
			if y > max.y {
				max.y = y
			}
		}
	})()

	// First panel is black
	go func() { r.Inputs <- WHITE }()
	r.Execute()

	// Print board
	fmt.Println(min.x, min.y, max.x, max.y)
	for y := min.y; y <= max.y; y++ {
		for x := min.x; x <= max.x; x++ {
			c, ok := w[point{x, y}]
			if !ok {
				fmt.Print(" ")
			} else if c == WHITE {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}

	fmt.Println("Done.", c)
}
