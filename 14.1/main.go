package main

import (
	"fmt"
	"main/util"
)

func main() {
	grid := util.Map(util.ReadInputLines("./input.txt"), func(line string) []rune { return []rune(line) })

	upperBlockers := make([]int, len(grid[0]))

	totalScore := 0

	for y, row := range grid {
		for x, r := range row {
			if r == 'O' {
				target := upperBlockers[x]
				upperBlockers[x] += 1

				totalScore += len(grid) - target

				grid[y][x] = '.'
				grid[target][x] = 'O'
			} else if r == '#' {
				upperBlockers[x] = y + 1
			}
		}
	}

	util.PrintGrid(grid, "%c")

	fmt.Printf("%d\n", totalScore)
}
