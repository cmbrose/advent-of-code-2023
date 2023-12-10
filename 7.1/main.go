package main

import (
	"fmt"
	"sort"
	"strings"

	"main/util"
)

var cardPriorityMap = map[byte]int{
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'J': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

type handType int

const (
	handHighCard handType = iota
	handOnePair
	handTwoPair
	handThreeOfAKind
	handFullHouse
	handFourOfAKind
	handFiveOfAKind
)

type hand struct {
	raw      string
	handType handType
}

type handAndBid struct {
	hand
	bid int
}

func parseLine(line string) handAndBid {
	pair := strings.Fields(line)
	handStr, bidStr := pair[0], pair[1]

	h := hand{
		handStr,
		getHandType(handStr),
	}

	b := util.AssertInt(bidStr)

	return handAndBid{h, b}
}

func getHandType(s string) handType {
	countsOfCards := make(map[rune]int)
	for _, r := range s {
		countsOfCards[r] += 1
	}

	countsOfCounts := make(map[int]int)
	for _, c := range countsOfCards {
		countsOfCounts[c] += 1
	}

	if countsOfCounts[5] == 1 {
		return handFiveOfAKind
	}
	if countsOfCounts[4] == 1 {
		return handFourOfAKind
	}
	if countsOfCounts[2] == 1 && countsOfCounts[3] == 1 {
		return handFullHouse
	}
	if countsOfCounts[3] == 1 {
		return handThreeOfAKind
	}
	if countsOfCounts[2] == 2 {
		return handTwoPair
	}
	if countsOfCounts[2] == 1 {
		return handOnePair
	}
	return handHighCard
}

func (h1 hand) isLessThan(h2 hand) bool {
	if h1.handType != h2.handType {
		return h1.handType < h2.handType
	}

	for i := 0; i < len(h1.raw); i += 1 {
		p1 := cardPriorityMap[h1.raw[i]]
		p2 := cardPriorityMap[h2.raw[i]]

		if p1 != p2 {
			return p1 < p2
		}
	}

	return false
}

func main() {
	hands := util.Map(util.ReadInputLines("./input.txt"), parseLine)

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].hand.isLessThan(hands[j].hand)
	})

	total := 0
	for i, h := range hands {
		//fmt.Printf("%s %d\n", h.hand.raw, h.bid)
		total += (i + 1) * h.bid
	}

	fmt.Printf("%d\n", total)
}
