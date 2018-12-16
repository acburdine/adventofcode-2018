package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/novalagung/gubrak"
)

type inst func([]int, int, int, int)

func (in inst) Name() string {
	return runtime.FuncForPC(reflect.ValueOf(in).Pointer()).Name()
}

var instructions = []inst{
	addr, addi, mulr, muli,
	banr, bani, borr, bori,
	setr, seti, gtir, gtri,
	gtrr, eqir, eqri, eqrr,
}

var dr = regexp.MustCompile("\\d+")

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sp := strings.Split(string(data), "\n\n\n\n")
	parts := strings.Split(sp[0], "\n\n")
	cases := make([]testcase, len(parts))

	for i, p := range parts {
		sp := strings.Split(p, "\n")
		cases[i] = testcase{
			before: mapint(dr.FindAllString(sp[0], -1)),
			instr:  mapint(dr.FindAllString(sp[1], -1)),
			after:  mapint(dr.FindAllString(sp[2], -1)),
		}
	}

	sum := 0
	opmap := make(map[int][]int, 16)

	for _, c := range cases {
		csum := 0

		for i, in := range instructions {
			tmp := make([]int, len(c.before))
			copy(tmp, c.before)
			in(tmp, c.instr[1], c.instr[2], c.instr[3])
			if reflect.DeepEqual(tmp, c.after) {
				csum++
				res, _ := gubrak.Uniq(append(opmap[c.instr[0]], i))
				opmap[c.instr[0]] = res.([]int)
			}
		}

		if csum >= 3 {
			sum++
		}
	}

	fmt.Printf("Day 16 Part 1: %d\n", sum)

	foundcodes := make(map[int]int, 16)
	for len(opmap) > 0 {
		for c, ps := range opmap {
			if len(ps) == 1 {
				foundcodes[ps[0]] = c
				delete(opmap, c)
				continue
			}

			for i, s := range ps {
				if _, ok := foundcodes[s]; ok {
					if i == len(ps)-1 {
						opmap[c] = ps[0:i]
					} else {
						opmap[c] = append(ps[0:i], ps[i+1:]...)
					}
				}
			}
		}
	}

	operations := make(map[int]inst, 16)
	for in, op := range foundcodes {
		operations[op] = instructions[in]
	}

	registers := make([]int, 4)
	actions := strings.Split(strings.TrimSpace(sp[1]), "\n")
	for _, l := range actions {
		in := mapint(dr.FindAllString(l, -1))
		operations[in[0]](registers, in[1], in[2], in[3])
	}

	fmt.Printf("Day 16 Part 2: %d\n", registers[0])
}

type testcase struct {
	before []int
	instr  []int
	after  []int
}

func mapint(in []string) []int {
	result := make([]int, len(in))
	for i, s := range in {
		result[i], _ = strconv.Atoi(s)
	}
	return result
}

func addr(registers []int, a, b, c int) {
	registers[c] = registers[a] + registers[b]
}

func addi(registers []int, a, b, c int) {
	registers[c] = registers[a] + b
}

func mulr(registers []int, a, b, c int) {
	registers[c] = registers[a] * registers[b]
}

func muli(registers []int, a, b, c int) {
	registers[c] = registers[a] * b
}

func banr(registers []int, a, b, c int) {
	registers[c] = registers[a] & registers[b]
}

func bani(registers []int, a, b, c int) {
	registers[c] = registers[a] & b
}

func borr(registers []int, a, b, c int) {
	registers[c] = registers[a] | registers[b]
}

func bori(registers []int, a, b, c int) {
	registers[c] = registers[a] | b
}

func setr(registers []int, a, b, c int) {
	registers[c] = registers[a]
}

func seti(registers []int, a, b, c int) {
	registers[c] = a
}

func gtir(registers []int, a, b, c int) {
	if a > registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func gtri(registers []int, a, b, c int) {
	if registers[a] > b {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func gtrr(registers []int, a, b, c int) {
	if registers[a] > registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func eqir(registers []int, a, b, c int) {
	if a == registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func eqri(registers []int, a, b, c int) {
	if registers[a] == b {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func eqrr(registers []int, a, b, c int) {
	if registers[a] == registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}
