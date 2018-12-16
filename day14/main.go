package main

import (
	"fmt"
	"strconv"
	"strings"
)

const input = 330121

var strinput = fmt.Sprintf("%d", input)

func main() {
	scores := []byte{'3', '7'}
	elfOne, elfTwo := 0, 1

	for len(scores) < 50000000 {
		score := []byte(strconv.Itoa(int(scores[elfOne] - '0' + scores[elfTwo] - '0')))
		scores = append(scores, score...)

		elfOne = (elfOne + 1 + int(scores[elfOne]-'0')) % len(scores)
		elfTwo = (elfTwo + 1 + int(scores[elfTwo]-'0')) % len(scores)
	}

	fmt.Printf("Day 14 Part 1: %s\n", string(scores[input:input+10]))
	fmt.Printf("Day 14 Part 2: %d\n", strings.Index(string(scores), strconv.Itoa(input)))
}
