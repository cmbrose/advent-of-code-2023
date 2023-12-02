package main

import (
	"fmt"

	"main/util"
)

func main() {
	curr := 0

	for _, line := range util.ReadInputLines("./input.txt") {
		var (
			first rune
			last  rune
		)

		for _, first = range line {
			if util.IsNumber(first) {
				break
			}
		}
		for i := range line {
			last = (rune)(line[len(line)-i-1])
			if util.IsNumber(last) {
				break
			}
		}

		value := (int)(10*(first-'0') + (last - '0'))

		curr += value
	}

	fmt.Printf("%d\n", curr)
}
