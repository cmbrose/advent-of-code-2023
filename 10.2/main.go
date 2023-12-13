package main

import (
	"fmt"
	"strings"

	"main/util"
)

func getAllowedMoves(r rune) (int, int, int, int) {
	switch r {
	case '.':
		return 0, 0, 0, 0
	case '│':
		return 0, 1, 0, -1
	case '─':
		return 1, 0, -1, 0
	case '└':
		return 1, 0, 0, -1
	case '┘':
		return -1, 0, 0, -1
	case '┐':
		return -1, 0, 0, 1
	case '┌':
		return 1, 0, 0, 1
	}

	panic(fmt.Sprintf("Unexpected char: %c", r))
}

func convertToBoxChars(s string) string {
	s = strings.ReplaceAll(s, "J", "┘")
	s = strings.ReplaceAll(s, "L", "└")
	s = strings.ReplaceAll(s, "|", "│")
	s = strings.ReplaceAll(s, "-", "─")
	s = strings.ReplaceAll(s, "F", "┌")
	s = strings.ReplaceAll(s, "7", "┐")

	return s
}

func getStartShape(sx, sy, ax1, ay1, ax2, ay2 int) rune {
	dx1, dy1, dx2, dy2 := ax1-sx, ay1-sy, ax2-sx, ay2-sy

	matches := func(r rune) bool {
		mx1, my1, mx2, my2 := getAllowedMoves(r)

		return (mx1 == dx1 && my1 == dy1 && mx2 == dx2 && my2 == dy2) ||
			(mx2 == dx1 && my2 == dy1 && mx1 == dx2 && my1 == dy2)
	}

	shapes := []rune{'┘', '└', '│', '─', '┌', '┐'}

	matched := util.Filter(shapes, matches)

	if len(matched) != 1 {
		panic(fmt.Sprintf("Cound not find start shaped - options were %v", matched))
	}

	return matched[0]
}

func main() {
	grid := util.Map(util.ReadInputLines("./input.txt"), func(line string) []rune {
		return []rune(convertToBoxChars(line))
	})

	var sx, sy int

	w, h := len(grid[0]), len(grid)
	for x := 0; x < w; x += 1 {
		for y := 0; y < h; y += 1 {
			if grid[y][x] == 'S' {
				sx, sy = x, y
				goto foundS
			}
		}
	}
foundS:

	// t => "travel"
	var tx1, ty1, tx2, ty2 int

	for dx := -1; dx <= 1; dx += 1 {
		for dy := -1; dy <= 1; dy += 1 {
			x, y := sx+dx, sy+dy

			if x < 0 || x >= w || y < 0 || y >= h {
				continue
			}

			if dx == 0 && dy == 0 {
				continue
			}

			mx1, my1, mx2, my2 := getAllowedMoves(grid[y][x])
			if (mx1*-1 == dx && my1*-1 == dy) || (mx2*-1 == dx && my2*-1 == dy) {
				tx1, ty1 = tx2, ty2
				tx2, ty2 = x, y
			}
		}
	}

	onlyPipe := util.FillGrid[rune](w, h, ' ')
	onlyPipe[sy][sx] = getStartShape(sx, sy, tx1, ty1, tx2, ty2)

	// pt => "previous travel"
	ptx1, pty1, ptx2, pty2 := sx, sy, sx, sy

	for tx1 != tx2 || ty1 != ty2 {
		r1, r2 := grid[ty1][tx1], grid[ty2][tx2]

		onlyPipe[ty1][tx1] = r1
		onlyPipe[ty2][tx2] = r2

		mx11, my11, mx21, my21 := getAllowedMoves(r1)
		if tx1+mx11 != ptx1 || ty1+my11 != pty1 {
			ptx1, pty1, tx1, ty1 = tx1, ty1, tx1+mx11, ty1+my11
		} else {
			ptx1, pty1, tx1, ty1 = tx1, ty1, tx1+mx21, ty1+my21
		}

		mx12, my12, mx22, my22 := getAllowedMoves(r2)
		if tx2+mx12 != ptx2 || ty2+my12 != pty2 {
			ptx2, pty2, tx2, ty2 = tx2, ty2, tx2+mx12, ty2+my12
		} else {
			ptx2, pty2, tx2, ty2 = tx2, ty2, tx2+mx22, ty2+my22
		}
	}

	onlyPipe[ty1][tx1] = grid[ty1][tx1]

	// Pad empty cells around the entire thing - note: after this, w and h are off by 2
	// Could I just update their values? Yes. Will I? No.
	onlyPipe = append([][]rune{util.Repeat(' ', w)}, onlyPipe...)
	onlyPipe = append(onlyPipe, util.Repeat(' ', w))

	for i := 0; i <= h+1; i += 1 {
		onlyPipe[i] = append([]rune{' '}, onlyPipe[i]...)
		onlyPipe[i] = append(onlyPipe[i], ' ')
	}

	util.PrintGrid(onlyPipe, "%c")
	fmt.Println()

	type point struct{ x, y int }
	q := []point{
		{0, 0},
	}
	seen := util.FillGrid(w+2, h+2, false)

	for len(q) > 0 {
		x, y := q[0].x, q[0].y
		q = q[1:]

		if seen[y][x] {
			continue
		}

		seen[y][x] = true

		// The trick here is: envision you are at the top left corner of the current cell.
		// Try to move from there, to the top left of an adjecent cell.
		// Under which conditions are you blocked from making that move?
		// Note: if you move up or left, then it is the next cell that blocks you, but
		// if you move right or down, then it is actually the current cell.

		// Try go up
		if y > 0 && !seen[y-1][x] {
			r := onlyPipe[y-1][x]
			if r != '─' && r != '┐' && r != '┘' {
				q = append(q, point{x, y - 1})
			}
		}
		// Try go down
		if y <= h && !seen[y+1][x] {
			r := onlyPipe[y][x]
			if r != '─' && r != '┐' && r != '┘' {
				q = append(q, point{x, y + 1})
			}
		}
		// Try go right
		if x <= w && !seen[y][x+1] {
			r := onlyPipe[y][x]
			if r != '│' && r != '└' && r != '┘' {
				q = append(q, point{x + 1, y})
			}
		}
		// Try go left
		if x > 0 && !seen[y][x-1] {
			r := onlyPipe[y][x-1]
			if r != '│' && r != '└' && r != '┘' {
				q = append(q, point{x - 1, y})
			}
		}
	}

	cnt := 0
	for x := 0; x <= w+1; x += 1 {
		for y := 0; y <= h+1; y += 1 {
			if onlyPipe[y][x] != ' ' {
				continue
			}

			if seen[y][x] {
				onlyPipe[y][x] = 'O'
			} else {
				onlyPipe[y][x] = 'I'
				cnt += 1
			}
		}
	}

	util.PrintGrid(onlyPipe, "%c")
	fmt.Println()

	fmt.Println(cnt)
}
