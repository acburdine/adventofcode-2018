package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const linePattern = "^#(?P<Num>[0-9]+) @ (?P<Left>[0-9]+),(?P<Right>[0-9]+): (?P<Width>[0-9]+)x(?P<Height>[0-9]+)$"
var lineRegex = regexp.MustCompile(linePattern)

type claim struct {
	num int
	left int
	top int
	width int
	height int
}

type coor struct {
	x int
	y int
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	claims, err := parseLines(lines)
	if err != nil {
		log.Fatal(err)
	}
	fabric := fabric(claims)

	fmt.Printf("Answer 1: %d\n", part1(fabric))
	fmt.Printf("Answer 2: %d\n", part2(fabric, claims))
}

func part1(fabric map[coor]int) int {
	sum := 0
 	for _, val := range fabric {
 		if val >= 2 {
 			sum++
		}
	}

 	return sum
}

func part2(fabric map[coor]int, claims []*claim) int {
	for _, c := range claims {
		check := checkClaim(fabric, c)
		if check {
			return c.num
		}
	}
	return 0
}

func checkClaim(fabric map[coor]int, claim *claim) bool {
	for x := claim.left; x < claim.left + claim.width; x++ {
		for y := claim.top; y < claim.top + claim.height; y++ {
			coord := coor{x, y}
			val := fabric[coord]
			if val > 1 {
				return false
			}
		}
	}
	return true
}

func fabric(claims []*claim) map[coor]int {
	fabric := make(map[coor]int, 1000 * 1000)

	for _, c := range claims {
		for x := c.left; x < c.left + c.width; x++ {
			for y := c.top; y < c.top + c.height; y++ {
				coord := coor{x, y}
				val := fabric[coord]
				fabric[coord] = val + 1
			}
		}
	}

	return fabric
}


func parseLines(lines []string) ([]*claim, error) {
	claims := make([]*claim, len(lines))

	for i, l := range lines {
		line := strings.TrimSpace(l)
		if l == "" {
			continue
		}

		if !lineRegex.MatchString(line) {
			return nil, fmt.Errorf("Line did not match regex: %s", line)
		}

		split := lineRegex.FindStringSubmatch(line)
		num, _ := strconv.Atoi(split[1])
		left, _ := strconv.Atoi(split[2])
		top, _ := strconv.Atoi(split[3])
		width, _ := strconv.Atoi(split[4])
		height, _ := strconv.Atoi(split[5])

		claims[i] = &claim{num, left, top, width, height}
	}

	return claims, nil
}