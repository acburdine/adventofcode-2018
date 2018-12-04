package adventofcode

import (
	"io/ioutil"
	"log"
	"strings"
)

const inputFile = "./input.txt"

func Input() []string {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(strings.TrimSpace(string(data)), "\n")
}
