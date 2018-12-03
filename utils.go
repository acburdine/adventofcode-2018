package adventofcode

import (
	"io/ioutil"
	"log"
	"strings"
)

const inputFile = "./input.txt"

func Input() []string {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(strings.TrimSpace(string(data)), "\n")
}
