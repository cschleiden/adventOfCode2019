package main

import (
	"fmt"
	"math"
	"strconv"
)

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

		for x := 0; x < w; x++ {
			l.pixels = append(l.pixels, make([]int, w))

			for y := 0; y < h; y++ {
				l.pixels[x][y], _ = strconv.Atoi(string(data[i]))

				i++
			}
		}

		img.layers = append(img.layers, *l)
	}

	return img
}

func main() {
	image := parseImage(25, 6, input1)

	// Find layers
	var result int
	minZeroes := math.MaxInt32
	for _, layer := range image.layers {
		zeroes := 0
		ones := 0
		twos := 0
		for x := 0; x < image.w; x++ {
			for y := 0; y < image.w; y++ {
				d := layer.pixels[x][y]

				if d == 0 {
					zeroes++
				} else if d == 1 {
					ones++
				} else if d == 2 {
					twos++
				}
			}
		}

		if zeroes < minZeroes {
			minZeroes = zeroes
			result = ones * twos
		}
	}

	fmt.Println(result)
}
