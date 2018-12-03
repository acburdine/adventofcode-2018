package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/acburdine/adventofcode-2018"
)

var lineRegex = regexp.MustCompile("([0-9]+)")

type claim struct {
	num    int
	left   int
	top    int
	width  int
	height int
}

type coor struct {
	x int
	y int
}

func main() {
	claims := mapClaims(adventofcode.Input())
	fabric := fabric(claims)

	fmt.Printf("Day 3 Part 1: %d\n", part1(fabric))
	fmt.Printf("Day 3 Part 2: %d\n", part2(fabric, claims))
}

func mapClaims(lines []string) []claim {
	claims := make([]claim, len(lines))
	for i, l := range lines {
		sp := lineRegex.FindAllString(l, -1)
		claim := claim{}
		claim.num, _ = strconv.Atoi(sp[0])
		claim.left, _ = strconv.Atoi(sp[1])
		claim.top, _ = strconv.Atoi(sp[2])
		claim.width, _ = strconv.Atoi(sp[3])
		claim.height, _ = strconv.Atoi(sp[4])
		claims[i] = claim
	}
	return claims
}

func fabric(claims []claim) map[coor]int {
	fabric := make(map[coor]int, 1000*1000)

	for _, c := range claims {
		for x := c.left; x < c.left+c.width; x++ {
			for y := c.top; y < c.top+c.height; y++ {
				xy := coor{x, y}
				fabric[xy] = fabric[xy] + 1
			}
		}
	}

	return fabric
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

func part2(fabric map[coor]int, claims []claim) int {
	for _, c := range claims {
		check := checkClaim(fabric, c)
		if check {
			return c.num
		}
	}
	return 0
}

func checkClaim(fabric map[coor]int, claim claim) bool {
	for x := claim.left; x < claim.left+claim.width; x++ {
		for y := claim.top; y < claim.top+claim.height; y++ {
			xy := coor{x, y}
			if fabric[xy] > 1 {
				return false
			}
		}
	}
	return true
}
