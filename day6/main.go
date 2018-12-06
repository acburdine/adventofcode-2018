package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/acburdine/adventofcode-2018"
)

const maxDist = 10000

type pos struct {
	x int
	y int
}

func main() {
	input := mapLines(adventofcode.Input())
	xmax, xmin, ymax, ymin := minmax(input)

	infinite := make(map[pos]bool, len(input))
	bound := make(map[pos]int, len(input))
	within := 0

	for y := ymin; y <= ymax; y++ {
		is_edge_y := y == ymin || y == ymax

		for x := xmin; x <= xmax; x++ {
			is_edge_x := x == xmin || x == xmax
			best := -1
			shortest := xmax + ymax
			sum := 0

			for i, p := range input {
				dist := abs(p.x-x) + abs(p.y-y)
				sum += dist
				if dist < shortest {
					best = i
					shortest = dist
				} else if dist == shortest {
					best = -1
				}
			}

			if sum < maxDist {
				within += 1
			}

			if best < 0 {
				continue
			}

			bestPos := input[best]

			if is_edge_y && is_edge_x {
				infinite[bestPos] = true
			} else {
				bound[bestPos] += 1
			}
		}
	}

	max := 0

	for _, p := range input {
		if infinite[p] {
			continue
		}
		s := bound[p]
		if s > max {
			max = s
		}
	}

	fmt.Printf("Day 6 Part 1: %d\n", max)
	fmt.Printf("Day 6 Part 2: %d\n", within)
}

func mapLines(lines []string) []pos {
	result := make([]pos, len(lines))
	for i, l := range lines {
		sp := strings.Split(strings.TrimSpace(l), ", ")
		x, _ := strconv.Atoi(sp[0])
		y, _ := strconv.Atoi(sp[1])
		result[i] = pos{x, y}
	}
	return result
}

func minmax(list []pos) (xmax int, xmin int, ymax int, ymin int) {
	xmin, ymin = list[0].x, list[0].y

	for _, pos := range list {
		if pos.x > xmax {
			xmax = pos.x
		} else if pos.x < xmin {
			xmin = pos.x
		}

		if pos.y > ymax {
			ymax = pos.y
		} else if pos.y < ymin {
			ymin = pos.y
		}
	}

	return
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
