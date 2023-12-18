package main

import (
	"fmt"

	"main/util"
)

type direction int

const (
	UP direction = iota
	DOWN
	LEFT
	RIGHT
)

type beam struct {
	x, y      int
	direction direction
}

func main() {
	grid := util.ReadInputRuneGrid("./input.txt")

	max := 0
	for i := 0; i < len(grid); i += 1 {
		maxLeft := countEnergizedTiles(grid, beam{len(grid[0]) - 1, i, LEFT})
		maxRight := countEnergizedTiles(grid, beam{0, i, RIGHT})

		max = util.Max(max, maxLeft, maxRight)
	}

	for i := 0; i < len(grid[0]); i += 1 {
		maxUp := countEnergizedTiles(grid, beam{i, len(grid) - 1, UP})
		maxDown := countEnergizedTiles(grid, beam{i, 0, DOWN})

		max = util.Max(max, maxUp, maxDown)
	}

	fmt.Printf("%d\n", max)
}

func printEnergized(energized [][]bool, grid [][]rune) {
	for y, row := range energized {
		for x, b := range row {
			if grid[y][x] != '.' && b {
				fmt.Printf("%c", grid[y][x])
			} else if b {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func countEnergizedTiles(grid [][]rune, b beam) int {
	energized := util.Grid[bool](len(grid[0]), len(grid))
	cache := make(map[beam]bool)

	q := []beam{b}

	for len(q) > 0 {
		b := q[0]
		q = q[1:]
		b1, b2 := runToEnd(b, grid, energized, cache)

		if b1 != nil {
			q = append(q, *b1)
		}
		if b2 != nil {
			q = append(q, *b2)
		}
	}

	return util.Sum(util.Map(energized, func(row []bool) int {
		return len(util.Filter(row, func(b bool) bool { return b }))
	}))
}

func runToEnd(b beam, grid [][]rune, energized [][]bool, cache map[beam]bool) (*beam, *beam) {
	for b.x >= 0 && b.x < len(grid[0]) && b.y >= 0 && b.y < len(grid) && !cache[b] {
		energized[b.y][b.x] = true
		cache[b] = true

		dx, dy := getMovement(b.direction)
		r := grid[b.y][b.x]

		if r == '.' {
			b.x += dx
			b.y += dy
		} else if r == '|' {
			if b.direction == UP || b.direction == DOWN {
				b.x += dx
				b.y += dy
			} else {
				return &beam{
						b.x,
						b.y - 1,
						UP,
					},
					&beam{
						b.x,
						b.y + 1,
						DOWN,
					}
			}
		} else if r == '-' {
			if b.direction == LEFT || b.direction == RIGHT {
				b.x += dx
				b.y += dy
			} else {
				return &beam{
						b.x + 1,
						b.y,
						RIGHT,
					},
					&beam{
						b.x - 1,
						b.y,
						LEFT,
					}
			}
		} else {
			rx, ry, rd := getReflection(b.direction, r)
			return &beam{
				b.x + rx,
				b.y + ry,
				rd,
			}, nil
		}
	}

	return nil, nil
}

func getMovement(d direction) (int, int) {
	if d == UP {
		return 0, -1
	}
	if d == DOWN {
		return 0, 1
	}
	if d == LEFT {
		return -1, 0
	}
	if d == RIGHT {
		return 1, 0
	}

	panic("Bad direction")
}

func getReflection(d direction, mirror rune) (int, int, direction) {
	if mirror == '/' {
		if d == UP {
			return 1, 0, RIGHT
		}
		if d == DOWN {
			return -1, 0, LEFT
		}
		if d == LEFT {
			return 0, 1, DOWN
		}
		if d == RIGHT {
			return 0, -1, UP
		}
	}
	if mirror == '\\' {
		if d == UP {
			return -1, 0, LEFT
		}
		if d == DOWN {
			return 1, 0, RIGHT
		}
		if d == LEFT {
			return 0, -1, UP
		}
		if d == RIGHT {
			return 0, 1, DOWN
		}
	}

	panic("Bad mirror")
}
