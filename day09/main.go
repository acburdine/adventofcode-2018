package main

import (
	"container/list"
	"fmt"
)

const (
	players    = 430
	lastMarble = 71588
)

func play(numPlayers, numMarbles int) int {
	marbles := list.New()
	scores := make(map[int]int, numPlayers)
	marbles.PushFront(0)

	currentPlayer, currentMarble := 0, marbles.Front()

	for i := 1; i <= numMarbles; i++ {
		if i%23 == 0 {
			for x := 0; x < 6; x++ {
				currentMarble = currentMarble.Prev()
				if currentMarble == nil {
					currentMarble = marbles.Back()
				}
			}
			marble := 0
			if currentMarble.Prev() == nil {
				marble = marbles.Remove(marbles.Back()).(int)
			} else {
				marble = marbles.Remove(currentMarble.Prev()).(int)
			}
			scores[currentPlayer] = scores[currentPlayer] + i + marble
		} else {
			currentMarble = currentMarble.Next()
			if currentMarble == nil {
				currentMarble = marbles.Front()
			}
			marbles.InsertAfter(i, currentMarble)
			currentMarble = currentMarble.Next()
		}

		currentPlayer = (currentPlayer + 1) % numPlayers
	}

	max := 0
	for _, s := range scores {
		if s > max {
			max = s
		}
	}
	return max
}

func main() {
	fmt.Printf("Day 9 Part 1: %d\n", play(players, lastMarble))
	fmt.Printf("Day 9 Part 2: %d\n", play(players, lastMarble*100))
}
