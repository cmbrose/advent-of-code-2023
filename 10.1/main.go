package main

import (
	"fmt"

	"main/util"
)

func getAllowedMoves(r rune) (int, int, int, int) {
	switch r {
	case '.':
		return 0, 0, 0, 0
	case '|':
		return 0, 1, 0, -1
	case '-':
		return 1, 0, -1, 0
	case 'L':
		return 1, 0, 0, -1
	case 'J':
		return -1, 0, 0, -1
	case '7':
		return -1, 0, 0, 1
	case 'F':
		return 1, 0, 0, 1
	}

	panic(fmt.Sprintf("Unexpected char: %c", r))
}

func main() {
	grid := util.Map(util.ReadInputLines("./input.txt"), func(line string) []rune {
		return []rune(line)
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

	fmt.Printf("Found S at (%d, %d)\n", sx, sy)

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

	fmt.Printf("Found S adjacent to (%d, %d) and (%d, %d)\n", tx1, ty1, tx2, ty2)

	// pt => "previous travel"
	ptx1, pty1, ptx2, pty2 := sx, sy, sx, sy

	c := 1 // Already made one move to get off of S

	for tx1 != tx2 || ty1 != ty2 {
		c += 1
		r1, r2 := grid[ty1][tx1], grid[ty2][tx2]

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

	fmt.Printf("%d\n", c)
}
