package main

import (
	"fmt"
	"sort"
	"strings"

	"main/util"
)

type card struct {
	id          int
	winningNums []int
	chosenNums  []int
	copies      int
}

func main() {

	lines := util.ReadInputLines("./input.txt")
	cards := util.Map(lines, func(line string) *card {
		idAndNums := strings.Split(line, ": ")
		id := strings.Fields(idAndNums[0])[1]

		winningAndChosen := strings.Split(idAndNums[1], " | ")
		winning := strings.Fields(winningAndChosen[0])
		chosen := strings.Fields(winningAndChosen[1])

		return &card{
			id:          util.AssertInt(id),
			winningNums: util.Map(winning, util.AssertInt),
			chosenNums:  util.Map(chosen, util.AssertInt),
			copies:      1,
		}
	})

	sum := 0

	for i, c := range cards {
		sum += c.copies

		sort.Ints(c.winningNums)
		sort.Ints(c.chosenNums)

		intersection := util.Intersect(c.winningNums, c.chosenNums)
		cnt := len(intersection)

		for j := 0; j < cnt; j += 1 {
			cards[i+j+1].copies += c.copies
		}
	}

	fmt.Printf("%d\n", sum)
}
