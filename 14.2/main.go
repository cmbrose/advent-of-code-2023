package main

import (
	"fmt"
	"main/util"
)

func main() {
	grid := util.Map(util.ReadInputLines("./input.txt"), func(line string) []rune { return []rune(line) })

	cache := make(map[string]int)
	repeatIdx, loopSize := -1, -1

	target := 1000000000
	for i := 0; i < target; i += 1 {
		k := fmt.Sprintf("%v", grid)
		if c, ok := cache[k]; ok {
			repeatIdx = c
			loopSize = i - c
			break
		}

		cache[k] = i

		rotate(grid)
	}

	target -= repeatIdx
	target %= loopSize

	for i := 0; i < target; i += 1 {
		rotate(grid)
	}

	totalScore := 0

	for y, row := range grid {
		for _, r := range row {
			if r == 'O' {
				totalScore += len(grid) - y
			}
		}
	}

	fmt.Printf("%d\n", totalScore)
}

func rotate(grid [][]rune) {
	upperBlockers := make([]int, len(grid[0]))
	for y, row := range grid {
		for x, r := range row {
			if r == 'O' {
				target := upperBlockers[x]
				upperBlockers[x] += 1

				grid[y][x] = '.'
				grid[target][x] = 'O'
			} else if r == '#' {
				upperBlockers[x] = y + 1
			}
		}
	}

	leftBlockers := make([]int, len(grid))
	for y, row := range grid {
		for x, r := range row {
			if r == 'O' {
				target := leftBlockers[y]
				leftBlockers[y] += 1

				grid[y][x] = '.'
				grid[y][target] = 'O'
			} else if r == '#' {
				leftBlockers[y] = x + 1
			}
		}
	}

	downBlockers := util.Repeat(len(grid)-1, len(grid[0]))
	for y := len(grid) - 1; y >= 0; y -= 1 {
		row := grid[y]
		for x, r := range row {
			if r == 'O' {
				target := downBlockers[x]
				downBlockers[x] -= 1

				grid[y][x] = '.'
				grid[target][x] = 'O'
			} else if r == '#' {
				downBlockers[x] = y - 1
			}
		}
	}

	rightBlockers := util.Repeat(len(grid[0])-1, len(grid))
	for y, row := range grid {
		for x := len(row) - 1; x >= 0; x -= 1 {
			r := row[x]
			if r == 'O' {
				target := rightBlockers[y]
				rightBlockers[y] -= 1

				grid[y][x] = '.'
				grid[y][target] = 'O'
			} else if r == '#' {
				rightBlockers[y] = x - 1
			}
		}
	}
}
