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
}

func main() {

	lines := util.ReadInputLines("./input.txt")
	cards := util.Map(lines, func(line string) card {
		idAndNums := strings.Split(line, ": ")
		id := strings.Fields(idAndNums[0])[1]

		winningAndChosen := strings.Split(idAndNums[1], " | ")
		winning := strings.Fields(winningAndChosen[0])
		chosen := strings.Fields(winningAndChosen[1])

		return card{
			id:          util.AssertInt(id),
			winningNums: util.Map(winning, util.AssertInt),
			chosenNums:  util.Map(chosen, util.AssertInt),
		}
	})

	sum := 0

	for _, c := range cards {
		sort.Ints(c.winningNums)
		sort.Ints(c.chosenNums)

		intersection := util.Intersect(c.winningNums, c.chosenNums)
		cnt := len(intersection)

		if cnt > 0 {
			sum += (1 << (cnt - 1))
		}
	}

	fmt.Printf("%d\n", sum)
}
