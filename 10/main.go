package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

type polar struct {
	angle float64
	dist  float64
}

type point struct {
	x, y int
}

type starfield struct {
	w, h      int
	asteroids map[point]bool
}

func (s *starfield) printStarfield() {
	for y := 0; y < s.h; y++ {
		for x := 0; x < s.w; x++ {
			if s.asteroids[point{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}
}

func getPolar(start *point, target *point) polar {
	dx := float64(start.x - target.x)
	dy := float64(start.y - target.y)
	return polar{
		math.Atan2(dy, dx),
		math.Sqrt(dx*dx + dy*dy),
	}
}

func (s *starfield) findVisibleNeighbors(asteroid *point) int {
	c := 0

	seenAsteroidAtAngle := make(map[int]bool)

	for a := range s.asteroids {
		if a == *asteroid {
			continue
		}

		p := getPolar(asteroid, &a)
		x := int(math.Round(p.angle * 10000.0))
		if !seenAsteroidAtAngle[x] {
			seenAsteroidAtAngle[x] = true
			c++
		}

	}

	return c
}

func (s *starfield) findBestPosition() (point, int) {
	var best point
	max := math.MinInt32

	for a := range s.asteroids {
		c := s.findVisibleNeighbors(&a)

		if c > max {
			max = c
			best = a
		}
	}

	return best, max
}

type target struct {
	asteroid point
	coord    polar
}

type targetList []target

func (t targetList) Len() int {
	return len(t)
}

func (t targetList) Less(i, j int) bool {
	return t[i].coord.dist < t[j].coord.dist
}

func (t targetList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (s *starfield) determineZapOrder(asteroid *point, goal int) *point {
	targets := make(map[int]targetList)

	// build map
	for a := range s.asteroids {
		if a == *asteroid {
			// Skip start
			continue
		}

		p := getPolar(asteroid, &a)
		x := int(math.Round(p.angle * 1000.0))
		targets[x] = append(targets[x], target{a, p})
		sort.Sort(targets[x])
		// fmt.Println(x, targets[x])
	}

	// Roootate
	zapped := 1
	angle := math.Pi / 2.0 * 1000.0
	for {
		x := int(math.Round(angle))

		if _, ok := targets[x]; ok {
			target := targets[x][0]
			fmt.Println(zapped, target.asteroid)
			zapped++
			if zapped == goal {
				return &target.asteroid
			}

			if len(targets[x]) > 1 {
				targets[x] = targets[x][1:]
			} else {
				delete(targets, x)
			}
		}

		angle += 1.0
		if angle > math.Pi*1000 {
			angle *= -1.0
		}
	}
}

func readStarfield(filename string) *starfield {
	input, _ := ioutil.ReadFile(filename)

	return parseStarfield(string(input))
}

func parseStarfield(input string) *starfield {
	s := &starfield{
		asteroids: make(map[point]bool),
	}

	lines := strings.Split(input, "\n")

	s.h = len(lines)

	for y, line := range lines {
		s.w = int(math.Max(float64(len(line)), float64(s.w)))

		for x, v := range line {
			if string(v) == "#" {
				s.asteroids[point{x, y}] = true
			}
		}
	}

	return s
}

func main() {
	s := readStarfield("./input10.txt")
	a, c := s.findBestPosition()
	fmt.Println(a, c)
	fmt.Println(s.determineZapOrder(&a, 200))
}
