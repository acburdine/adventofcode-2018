package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	answer1, err := part1(lines)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Answer to Part 1: %d\n", answer1)

	answer2, err := part2(lines)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Answer to Part 2: %d\n", answer2)
}

func part1(lines []string) (int, error) {
	freq := 0

	for _, l := range lines {
		if l == "" {
			continue
		}

		num, err := strconv.Atoi(l)
		if err != nil {
			return 0, err
		}
		freq += num
	}

	return freq, nil
}

func part2(lines []string) (int, error) {
	freq := 0
	i := 0
	found := make(map[int]bool)

	for {
		l := lines[i]
		if l == "" {
			i = (i + 1) % len(lines)
			continue
		}

		num, err := strconv.Atoi(l)
		if err != nil {
			return 0, err
		}
		freq += num

		if _, ok := found[freq]; ok {
			return freq, nil
		}

		found[freq] = true
		i = (i + 1) % len(lines)
	}
}

