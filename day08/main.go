package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/acburdine/adventofcode-2018"
)

type node struct {
	children []*node
	metadata []int
}

func main() {
	var test bool
	var input string

	flag.BoolVar(&test, "test", false, "")
	flag.Parse()

	if test {
		input = "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"
	} else {
		input = adventofcode.Input()[0]
	}

	data := parse(input)
	root, _ := buildTree(data)

	fmt.Printf("Day 8 Part 1: %d\n", sumTree(root))
	fmt.Printf("Day 8 Part 2: %d\n", nodeValue(root))
}

func nodeValue(nd *node) int {
	value := 0
	for _, m := range nd.metadata {
		if len(nd.children) == 0 {
			value += m
		}

		if m > len(nd.children) {
			continue
		}

		value += nodeValue(nd.children[m-1])
	}
	return value
}

func sumTree(nd *node) int {
	sum := 0
	for _, n := range nd.children {
		sum += sumTree(n)
	}
	for _, i := range nd.metadata {
		sum += i
	}
	return sum
}

func buildTree(data []int) (*node, int) {
	children := data[0]
	metadata := data[1]

	node := &node{
		children: make([]*node, children),
	}

	dataIndex := 2
	for i := 0; i < children; i++ {
		childNode, used := buildTree(data[dataIndex:])
		dataIndex += used
		node.children[i] = childNode
	}

	end := dataIndex + metadata
	node.metadata = data[dataIndex:end]

	return node, end
}

func parse(line string) []int {
	sp := strings.Split(line, " ")
	result := make([]int, len(sp))
	for i, s := range sp {
		result[i], _ = strconv.Atoi(s)
	}
	return result
}
