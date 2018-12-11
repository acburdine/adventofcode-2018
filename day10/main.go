package main

import (
	"fmt"

	"github.com/acburdine/adventofcode-2018"
)

const yThreshold = 20

type point struct {
	x, y, velx, vely int
}

func main() {
	points := parse(adventofcode.Input())
	var xmin, xmax, ymin, ymax, ydiff, t int

	for {
		xmin, ymin = 100000, 100000
		xmax, ymax = -100000, -100000

		for _, p := range points {
			p.x += p.velx
			p.y += p.vely

			if p.x < xmin {
				xmin = p.x
			} else if p.x > xmax {
				xmax = p.x
			}

			if p.y < ymin {
				ymin = p.y
			} else if p.y > ymax {
				ymax = p.y
			}
		}

		diff := abs(ymax - ymin)

		if ydiff > 0 && diff > ydiff {
			// TODO: this is a bad idea and needs cleaning up
			for _, p := range points {
				p.x -= p.velx
				p.y -= p.vely
			}

			break
		} else {
			ydiff = diff
		}

		t++
	}

	canvas := make([][]bool, abs(ymax-ymin)+1)
	for i := range canvas {
		canvas[i] = make([]bool, abs(xmax-xmin)+1)
	}

	for _, p := range points {
		canvas[abs(p.y-ymin)][abs(p.x-xmin)] = true
	}

	for y := 0; y < len(canvas); y++ {
		for x := 0; x < len(canvas[y]); x++ {
			if canvas[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Print("\n")
	}

	fmt.Println(t)
}

func abs(val int) int {
	if val < 0 {
		return -1 * val
	}
	return val
}

func parse(lines []string) []*point {
	result := make([]*point, len(lines))
	for i, l := range lines {
		point := &point{}
		fmt.Sscanf(l, "position=<%d, %d> velocity=<%d, %d>", &point.x, &point.y, &point.velx, &point.vely)
		result[i] = point
	}
	return result
}
