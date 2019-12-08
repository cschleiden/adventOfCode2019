package main

import (
	"fmt"
	"strconv"

	colorfmt "github.com/fatih/color"
)

type color int

const (
	black       color = 0
	white             = 1
	transparent       = 2
)

func (c *color) print() {
	switch *c {
	case black:
		colorfmt.New(colorfmt.BgBlack, colorfmt.FgBlack).Print("░")
	case white:
		colorfmt.New(colorfmt.BgWhite, colorfmt.FgWhite).Print("█")
	case transparent:
		fmt.Print(" ")
	}
}

type layer struct {
	pixels [][]int
}

type image struct {
	layers []layer
	w, h   int
}

func parseImage(w, h int, data string) *image {
	img := &image{
		w: w,
		h: h,
	}

	for i := 0; i < len(data); {
		l := &layer{}

		for y := 0; y < h; y++ {
			l.pixels = append(l.pixels, make([]int, w))

			for x := 0; x < w; x++ {
				l.pixels[y][x], _ = strconv.Atoi(string(data[i]))

				i++
			}
		}

		img.layers = append(img.layers, *l)
	}

	return img
}

func (i *image) print() {
	for y := 0; y < i.h; y++ {
		for x := 0; x < i.w; x++ {
			var pixel color

			for _, layer := range i.layers {
				p := layer.pixels[y][x]
				if p != transparent {
					pixel = color(p)
					break
				}
			}

			pixel.print()
		}

		fmt.Println()
	}
}

func main() {
	image := parseImage(25, 6, input1)
	image.print()

	// image := parseImage(2, 2, "0222112222120000")
	// image.print()
}
