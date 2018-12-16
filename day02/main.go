package main

import (
	"fmt"
	"strings"

	"github.com/acburdine/adventofcode-2018"
)

func main() {
	lines := mapRunes(adventofcode.Input())

	fmt.Printf("Day 2 Part 1: %d\n", part1(lines))
	fmt.Printf("Day 2 Part 2: %s\n", part2(lines))
}

func mapRunes(lines []string) [][]rune {
	result := make([][]rune, len(lines))
	for i, l := range lines {
		result[i] = []rune(strings.TrimSpace(l))
	}
	return result
}

func part1(lines [][]rune) int {
	two := 0
	three := 0

	for _, r := range lines {
		found := make(map[rune]int, len(r))

		for _, char := range r {
			found[char] = found[char] + 1
		}

		hasTwo := false
		hasThree := false

		for _, count := range found {
			hasTwo = (count == 2) || hasTwo
			hasThree = (count == 3) || hasThree
		}
		if hasTwo {
			two++
		}
		if hasThree {
			three++
		}
	}
	return two * three
}

func part2(lines [][]rune) string {
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			line1 := lines[i]
			line2 := lines[j]

			same := make([]rune, 0, len(line1))
			for x, r1 := range line1 {
				if r1 == line2[x] {
					same = append(same, r1)
				}
			}

			diff := len(line1) - len(same)
			if diff > 1 {
				continue
			}

			return string(same)
		}
	}

	return ""
}
