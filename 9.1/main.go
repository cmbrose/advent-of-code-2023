package main

import (
	"fmt"
	"strings"

	"main/util"
)

func main() {
	valueLists := util.Map(util.ReadInputLines("./input.txt"), func(line string) []int {
		return util.Map(strings.Fields(line), util.AssertInt)
	})

	nexts := util.Map(valueLists, predictNext)

	sum := util.Sum(nexts)

	fmt.Printf("%d\n", sum)
}

func predictNext(values []int) int {
	if util.All(values, func(i int) bool { return i == 0 }) {
		return 0
	}

	diffs := make([]int, len(values)-1)

	for i := 0; i < len(values)-1; i += 1 {
		diffs[i] = values[i+1] - values[i]
	}

	next := predictNext(diffs)

	return values[len(values)-1] + next
}
