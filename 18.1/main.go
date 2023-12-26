package main

import (
	"fmt"
	"regexp"

	"main/util"
)

func main() {
	re := regexp.MustCompile(`([UDLR]) ([0-9]+) \(#([a-f0-9]+)\)`)

	grid := [][]rune{
		{'#'},
	}

	x, y := 0, 0

	for _, line := range util.ReadInputLines("./input.txt") {
		m := re.FindStringSubmatch(line)
		dir, cnt, _ := m[1], util.AssertInt(m[2]), m[3]

		for i := 0; i < cnt; i += 1 {
			if dir == "U" {
				y -= 1
				if y < 0 {
					grid = extendGridUp(grid)
					y = 0
				}
			} else if dir == "D" {
				y += 1
				if y >= len(grid) {
					grid = extendGridDown(grid)
				}
			} else if dir == "L" {
				x -= 1
				if x < 0 {
					grid = extendGridLeft(grid)
					x = 0
				}
			} else if dir == "R" {
				x += 1
				if x >= len(grid[y]) {
					grid = extendGridRight(grid)
				}
			}

			grid[y][x] = '#'
		}
	}

	// Pad
	grid = extendGridDown(grid)
	grid = extendGridUp(grid)
	grid = extendGridRight(grid)
	grid = extendGridLeft(grid)

	type point struct {
		x, y int
	}
	q := []point{
		{0, 0},
	}

	for len(q) > 0 {
		x, y := q[0].x, q[0].y
		q = q[1:]

		if grid[y][x] != '.' {
			continue
		}

		grid[y][x] = ' '

		if x > 0 && grid[y][x-1] == '.' {
			q = append(q, point{x - 1, y})
		}
		if x < len(grid[y])-1 && grid[y][x+1] == '.' {
			q = append(q, point{x + 1, y})
		}
		if y > 0 && grid[y-1][x] == '.' {
			q = append(q, point{x, y - 1})
		}
		if y < len(grid)-1 && grid[y+1][x] == '.' {
			q = append(q, point{x, y + 1})
		}
	}

	cnt := 0

	for _, row := range grid {
		for _, r := range row {
			if r != ' ' {
				cnt += 1
			}
		}
	}

	util.PrintGrid(grid, "%c")

	fmt.Printf("%d\n", cnt)
}

func extendGridUp(grid [][]rune) [][]rune {
	newRow := util.Repeat('.', len(grid[0]))
	return append([][]rune{newRow}, grid...)
}
func extendGridDown(grid [][]rune) [][]rune {
	newRow := util.Repeat('.', len(grid[0]))
	return append(grid, newRow)
}
func extendGridRight(grid [][]rune) [][]rune {
	for i := range grid {
		grid[i] = append(grid[i], '.')
	}
	return grid
}
func extendGridLeft(grid [][]rune) [][]rune {
	for i := range grid {
		grid[i] = append([]rune{'.'}, grid[i]...)
	}
	return grid
}
