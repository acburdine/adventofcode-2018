package main

import "fmt"

const input = 6303

type xy struct {
	x, y int
}

func main() {
	cells := make(map[xy]int, 300*300)
	for x := 1; x <= 300; x++ {
		for y := 1; y <= 300; y++ {
			rackid := x + 10
			value := digit(((rackid*y)+input)*rackid) - 5
			cells[xy{x, y}] = value
		}
	}

	maxVal := 0
	var maxCoord xy

	for x := 1; x <= 300-2; x++ {
		for y := 1; y <= 300-2; y++ {
			sum := cells[xy{x, y}] +
				cells[xy{x, y + 1}] +
				cells[xy{x, y + 2}] +
				cells[xy{x + 1, y}] +
				cells[xy{x + 1, y + 1}] +
				cells[xy{x + 1, y + 2}] +
				cells[xy{x + 2, y}] +
				cells[xy{x + 2, y + 1}] +
				cells[xy{x + 2, y + 2}]

			if sum > maxVal {
				maxVal = sum
				maxCoord = xy{x, y}
			}
		}
	}

	fmt.Printf("Day 11 Part 1: %d,%d\n", maxCoord.x, maxCoord.y)

	var maxSize int
	maxVal = 0

	// These were just guesses as to when the size would start to get too large for max sums
	// Gotta love O(n^5) solutions
	for s := 3; s <= 20; s++ {
		for x := 1; x <= 301-s; x++ {
			for y := 1; y <= 301-s; y++ {
				sum := 0

				for i := 0; i < s; i++ {
					for j := 0; j < s; j++ {
						sum += cells[xy{x + i, y + j}]
					}
				}

				if sum > maxVal {
					maxVal = sum
					maxCoord = xy{x, y}
					maxSize = s
				}
			}
		}
	}

	fmt.Printf("Day 11 Part 2: %d,%d,%d\n", maxCoord.x, maxCoord.y, maxSize)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func digit(num int) int {
	if num < 100 {
		return 0
	}

	return (num / 100) % 10
}
