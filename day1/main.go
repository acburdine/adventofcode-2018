package main

import (
	"fmt"
	"strconv"

	"github.com/acburdine/adventofcode-2018"
)

func main() {
	ints := mapInt(adventofcode.Input())

	fmt.Printf("Day 1 Part 1: %d\n", part1(ints))
	fmt.Printf("Day 1 Part 2: %d\n", part2(ints))
}

func mapInt(arr []string) []int {
	intarr := make([]int, len(arr))
	for i, s := range arr {
		num, _ := strconv.Atoi(s)
		intarr[i] = num
	}
	return intarr
}

func part1(ints []int) int {
	sum := 0
	for _, i := range ints {
		sum += i
	}
	return sum
}

func part2(ints []int) int {
	mark := 0
	found := make(map[int]bool)

	for i := 0; true; i = (i + 1) % len(ints) {
		mark += ints[i]
		if _, ok := found[mark]; ok {
			return mark
		}
		found[mark] = true
	}
	return 0
}
