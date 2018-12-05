package main

import (
	"fmt"
	"strings"

	"github.com/acburdine/adventofcode-2018"
)

func main() {
	polymer := strings.TrimSpace(adventofcode.Input()[0])

	fmt.Printf("Day 5 Part 1: %d\n", collapse([]rune(polymer)))
	fmt.Printf("Day 5 Part 2: %d\n", part2([]rune(polymer)))
}

func collapse(runes []rune) int {
	for {
		didSomething := false
		for i := 0; i < len(runes)-1; i++ {
			diff := int(runes[i]) - int(runes[i+1])
			if diff == 32 || diff == -32 {
				runes = append(runes[:i], runes[i+2:]...)
				didSomething = true
			}
		}
		if !didSomething {
			break
		}
	}

	return len(runes)
}

func part2(runes []rune) int {
	min := len(runes)

	for i := int('A'); i <= int('Z'); i++ {
		tmp := make([]rune, 0, len(runes))

		for j := 0; j < len(runes); j++ {
			v := int(runes[j])
			if v != i && v != i+32 {
				tmp = append(tmp, runes[j])
			}
		}

		test := collapse(tmp)
		if test < min {
			min = test
		}
	}

	return min
}
