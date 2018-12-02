package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	answer1 := part1(lines)
	fmt.Printf("Answer to part 1: %d\n", answer1)

	answer2 := part2(lines)
	fmt.Printf("Answer to part 2: %s\n", answer2)
}

func part1(lines []string) int {
	twoCount := 0
	threeCount := 0

	for _, l := range lines {
		found := make(map[rune]int, len(l))

		for _, char := range l {
			if char == ' ' {
				continue
			}

			v, ok := found[char]
			if ok {
				found[char] = v + 1
			} else {
				found[char] = 1
			}
		}

		hasTwo := false
		hasThree := false

		for _, count := range found {
			if count == 2 {
				hasTwo = true
			} else if count == 3 {
				hasThree = true
			}
		}

		if hasTwo {
			twoCount++
		}

		if hasThree {
			threeCount ++
		}
	}

	return twoCount * threeCount
}

func part2(lines []string) string {
	for i := 0; i < len(lines); i++ {
		for j := i; j < len(lines); j++ {
			line1 := []rune(lines[i])
			line2 := []rune(lines[j])

			diffCount := 0
			diffPos := 0

			for x, r1 := range line1 {
				r2 := line2[x]

				if r1 == r2 {
					continue
				}

				diffPos = x
				diffCount++
				if diffCount > 1 {
					break
				}
			}

			if diffCount == 1 {
				result := make([]rune, 0, len(line1))
				for x, char := range line1 {
					if x != diffPos {
						result = append(result, char)
					}
				}
				return string(result)
			}
		}
	}

	return ""
}

