package main

import (
	"fmt"
	"strings"

	"github.com/acburdine/adventofcode-2018"
)

func main() {
	input := adventofcode.Input()
	state := parseInitialState(input[0])
	notes := parseNotes(input[2:])

	fmt.Printf("Day 12 Part 1: %d\n", run(state, notes, 20))
	fmt.Printf("Day 12 Part 2: %d\n", run(state, notes, 50000000000))
}

func run(state map[int]bool, notes map[note]bool, iterations int) int {
	stateCopy := make(map[int]bool, len(state))
	for i, b := range state {
		stateCopy[i] = b
	}

	lastMin, lastMax := minmax(stateCopy)

	for i := 0; i < iterations; i++ {
		newState := make(map[int]bool, len(stateCopy))

		for x := lastMin - 2; x <= lastMax+2; x++ {
			for note, next := range notes {
				if note.minusTwo != stateCopy[x-2] ||
					note.minusOne != stateCopy[x-1] ||
					note.zero != stateCopy[x] ||
					note.one != stateCopy[x+1] ||
					note.two != stateCopy[x+2] {
					continue
				}

				if !next && stateCopy[x] {
					newState[x] = next
				} else if next {
					newState[x] = next
				}
			}
		}

		// Eventually things will settle into a pattern
		// so we check for that and if it has settled into a pattern we just extrapolate out the rest
		min, max := minmax(newState)
		if (max-min) == (lastMax-lastMin) && str(newState, min, max) == str(stateCopy, lastMin, lastMax) {
			diff := sum(newState) - sum(stateCopy)
			left := iterations - i - 1

			return sum(newState) + (diff * left)
		}

		stateCopy = newState
		lastMin, lastMax = min, max
	}

	return sum(stateCopy)
}

func sum(state map[int]bool) int {
	res := 0
	for i, b := range state {
		if b {
			res += i
		}
	}
	return res
}

func str(state map[int]bool, min, max int) string {
	builder := &strings.Builder{}
	for i := min; i <= max; i++ {
		b := state[i]
		if b {
			builder.WriteString("#")
		} else {
			builder.WriteString(".")
		}
	}
	return builder.String()
}

func minmax(state map[int]bool) (int, int) {
	min, max := 10000, 0
	for i := range state {
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}
	return min, max
}

func parseInitialState(line string) map[int]bool {
	line = strings.TrimPrefix(line, "initial state: ")
	parts := []rune(line)
	state := make(map[int]bool, len(parts))
	for i, p := range parts {
		state[i] = p == '#'
	}
	return state
}

type note struct {
	minusTwo, minusOne, zero, one, two bool
}

func parseNotes(lines []string) map[note]bool {
	notes := make(map[note]bool, len(lines))
	for _, l := range lines {
		sp := strings.Split(l, " => ")
		pre := sp[0]
		note := note{
			minusTwo: pre[0] == '#',
			minusOne: pre[1] == '#',
			zero:     pre[2] == '#',
			one:      pre[3] == '#',
			two:      pre[4] == '#',
		}
		notes[note] = sp[1][0] == '#'
	}
	return notes
}
