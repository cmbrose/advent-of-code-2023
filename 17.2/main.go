package main

import (
	"container/heap"
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

type move struct {
	x, y                     int
	consecutiveStraightMoves int
	direction                direction
}

type weightedMove struct {
	move
	weight int
}

func main() {
	grid := util.ParseIntGrid()

	mins := make(map[move]int)
	proc := make(map[move]bool)
	from := make(map[move]move)

	q := util.NewPriorityQueue(func(m1, m2 weightedMove) bool {
		return m1.weight < m2.weight
	})

	heap.Push(q, weightedMove{move{0, 0, 0, RIGHT}, 0})

	moveIfValid := func(m weightedMove, dx, dy int, consecutiveStraightMoves int, d direction) {
		x, y := m.x+dx, m.y+dy
		if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) {
			return
		}

		w := m.weight + grid[y][x]
		nextMove := weightedMove{move{x, y, consecutiveStraightMoves + 1, d}, w}

		if w, ok := mins[nextMove.move]; ok && w <= nextMove.weight {
			return
		}
		mins[nextMove.move] = w

		from[nextMove.move] = m.move
		heap.Push(q, nextMove)
	}

	var end weightedMove

	for q.Len() > 0 {
		m := heap.Pop(q).(weightedMove)
		if m.x == len(grid[0])-1 && m.y == len(grid)-1 && m.consecutiveStraightMoves >= 4 {
			end = m
			break
		}

		if proc[m.move] {
			continue
		}
		proc[m.move] = true

		if m.consecutiveStraightMoves < 10 {
			dx, dy := getMovement(m.direction)
			moveIfValid(m, dx, dy, m.consecutiveStraightMoves, m.direction)
		}

		if m.consecutiveStraightMoves >= 4 {
			x1, y1, d1, x2, y2, d2 := getTurn(m.direction)
			moveIfValid(m, x1, y1, 0, d1)
			moveIfValid(m, x2, y2, 0, d2)
		}
	}

	printFrom(from, end.move)

	fmt.Println(end.weight)
}

func printFrom(from map[move]move, end move) {
	grid := util.ReadInputRuneGrid("./input.txt")

	m := end
	for m.x != 0 || m.y != 0 {
		f := from[m]
		grid[m.y][m.x] = getMovementChar(m.direction)
		m = f
	}

	util.PrintGrid(grid, "%c")
}

func getMovementChar(d direction) rune {
	if d == UP {
		return '^'
	}
	if d == DOWN {
		return 'v'
	}
	if d == LEFT {
		return '<'
	}
	if d == RIGHT {
		return '>'
	}

	panic("Bad direction")
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

func getTurn(d direction) (int, int, direction, int, int, direction) {
	if d == UP {
		return -1, 0, LEFT, 1, 0, RIGHT
	}
	if d == DOWN {
		return -1, 0, LEFT, 1, 0, RIGHT
	}
	if d == LEFT {
		return 0, -1, UP, 0, 1, DOWN
	}
	if d == RIGHT {
		return 0, -1, UP, 0, 1, DOWN
	}

	panic("Bad direction")
}
