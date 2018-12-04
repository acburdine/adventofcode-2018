package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/acburdine/adventofcode-2018"
)

const timeConst = "2006-01-02 15:04"
const asleep = "falls asleep"
const awake = "wakes up"

var regex = regexp.MustCompile("^\\[([0-9-: ]+)\\] ([A-Za-z0-9# ]+)$")
var idRegex = regexp.MustCompile("([0-9]+)")

type entry struct {
	time time.Time
	text string
}

func (e entry) String() string {
	return fmt.Sprintf("Time: %s, text: %s", e.time.Format(time.RFC3339), e.text)
}

type entrylist []entry

func (l entrylist) Len() int {
	return len(l)
}

func (l entrylist) Less(i, j int) bool {
	return l[i].time.Before(l[j].time)
}

func (l entrylist) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

type guard struct {
	id      int
	total   int
	minutes map[int]int
}

func main() {
	lines := adventofcode.Input()
	tlist := entries(lines)
	sort.Sort(tlist)
	guards := guards(tlist)

	fmt.Printf("Day 4 Part 1: %d\n", part1(guards))
	fmt.Printf("Day 4 Part 2: %d\n", part2(guards))
}

func entries(lines []string) entrylist {
	list := make([]entry, len(lines))
	for i, l := range lines {
		sp := regex.FindStringSubmatch(strings.TrimSpace(l))
		t, _ := time.Parse(timeConst, sp[1])
		list[i] = entry{t, sp[2]}
	}
	return entrylist(list)
}

func guards(list entrylist) map[int]*guard {
	guards := make(map[int]*guard, len(list))
	var curGuard *guard
	var sleep int

	for _, ent := range list {
		id := idRegex.FindString(ent.text)
		if id != "" {
			idnum, _ := strconv.Atoi(id)
			g, ok := guards[idnum]
			if ok {
				curGuard = g
			} else {
				curGuard = &guard{idnum, 0, make(map[int]int)}
				guards[idnum] = curGuard
			}
		} else if ent.text == asleep {
			sleep = ent.time.Minute()
		} else if ent.text == awake {
			for i := sleep; i < ent.time.Minute(); i++ {
				curGuard.total++
				curGuard.minutes[i] = curGuard.minutes[i] + 1
			}
		}
	}

	return guards
}

func part1(guards map[int]*guard) int {
	var maxGuard *guard
	max := 0

	for _, guard := range guards {
		if guard.total > max {
			maxGuard = guard
			max = guard.total
		}
	}

	max, minute := 0, 0
	for min, count := range maxGuard.minutes {
		if count > max {
			minute = min
			max = count
		}
	}

	return minute * maxGuard.id
}

func part2(guards map[int]*guard) int {
	minute, max, id := 0, 0, 0
	for _, guard := range guards {
		for min, count := range guard.minutes {
			if count > max {
				minute, max, id = min, count, guard.id
			}
		}
	}

	return minute * id
}
