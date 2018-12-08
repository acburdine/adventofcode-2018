package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/acburdine/adventofcode-2018"
)

var regex = regexp.MustCompile("[A-Z]")

func main() {
	input := parse(adventofcode.Input())

	part1, _ := doWork(input, 1, 0)
	_, part2 := doWork(input, 5, 60)

	fmt.Printf("Day 7 Part 1: %s\n", part1)
	fmt.Printf("Day 7 Part 2: %d\n", part2)
}

func doWork(tasks [][]string, worker_count int, work_factor int) (string, int) {
	remaining := make(map[string]bool, len(tasks))
	done := ""

	for _, t := range tasks {
		remaining[t[0]] = true
		remaining[t[1]] = true
	}

	goal := len(remaining)
	getNextTask := func() string {
		doable := make([]string, 0, len(remaining))
		for task := range remaining {
			if strings.Contains(done, task) {
				delete(remaining, task)
				continue
			}

			canDo := true
			for _, t := range tasks {
				if t[1] == task && !strings.Contains(done, t[0]) {
					canDo = false
					break
				}
			}

			if canDo {
				doable = append(doable, task)
			}
		}
		if len(doable) > 0 {
			sort.Strings(doable)
			return doable[0]
		}
		return ""
	}

	workers := make([]*worker, worker_count)
	t := 0

	for {
		for i, w := range workers {
			if w == nil {
				continue
			}

			if t == w.done {
				done = done + w.task
				workers[i] = nil
			}
		}

		if len(done) == goal {
			break
		}

		for i, w := range workers {
			if w == nil {
				task := getNextTask()
				if task == "" {
					continue
				}
				workers[i] = &worker{
					done: t + (work_factor + 1 + (int(task[0]) - int('A'))),
					task: task,
				}
				delete(remaining, task)
			}
		}

		t++
	}

	return done, t
}

type worker struct {
	done int
	task string
}

func parse(lines []string) [][]string {
	result := make([][]string, len(lines))
	for i, l := range lines {
		sp := regex.FindAllString(strings.TrimPrefix(l, "Step"), -1)
		result[i] = []string{sp[0], sp[1]}
	}
	return result
}
