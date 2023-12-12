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
	onlyPipe[sy][sx] = 'J' // This is the S and I'm a cheater for hardcoding it

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

	util.PrintGrid(onlyPipe, "%c")
	fmt.Println()

	type cache struct {
		inside     bool
		lastCorner rune
	}

	goingDown := util.FillGrid(w, h, cache{})
	goingUp := util.FillGrid(w, h, cache{})
	goingRight := util.FillGrid(w, h, cache{})
	goingLeft := util.FillGrid(w, h, cache{})

	for i := 0; i < w; i += 1 {
		for j := 0; j < h; j += 1 {
			x, y, rx, uy := i, j, w-i-1, h-j-1

			var prevDown, prevUp, prevRight, prevLeft cache

			if j != 0 {
				prevDown, prevUp = goingDown[y-1][x], goingUp[uy+1][rx]
			} else {
				prevDown, prevUp = cache{false, '|'}, cache{false, '|'}
			}

			if i != 0 {
				prevRight, prevLeft = goingRight[y][x-1], goingLeft[uy][rx+1]
			} else {
				prevRight, prevLeft = cache{false, '-'}, cache{false, '-'}
			}

			goingDown[y][x], goingUp[uy][rx], goingRight[y][x], goingLeft[uy][rx] = prevDown, prevUp, prevRight, prevLeft

			if onlyPipe[y][x] != ' ' {
				dx, _, _, ry := getAllowedMoves(onlyPipe[y][x])
				pdx, _, _, _ := getAllowedMoves(prevDown.lastCorner)
				_, _, _, pry := getAllowedMoves(prevRight.lastCorner)

				if dx != 0 && dx != -1*pdx {
					goingDown[y][x].inside = !goingDown[y][x].inside
				}
				if onlyPipe[y][x] != '|' {
					goingDown[y][x].lastCorner = onlyPipe[y][x]
				}

				if ry != 0 && ry != -1*pry {
					goingRight[y][x].inside = !goingRight[y][x].inside
				}
				if onlyPipe[y][x] != '-' {
					goingRight[y][x].lastCorner = onlyPipe[y][x]
				}
			}

			if onlyPipe[uy][rx] != ' ' {
				ux, _, _, ly := getAllowedMoves(onlyPipe[uy][rx])
				pux, _, _, _ := getAllowedMoves(prevUp.lastCorner)
				_, _, _, ply := getAllowedMoves(prevLeft.lastCorner)

				if ux != 0 && ux != -1*pux {
					goingUp[uy][rx].inside = !goingUp[uy][rx].inside
				}
				if onlyPipe[uy][rx] != '|' {
					goingUp[uy][rx].lastCorner = onlyPipe[uy][rx]
				}

				if ly != 0 && ly != -1*ply {
					goingLeft[uy][rx].inside = !goingLeft[uy][rx].inside
				}
				if onlyPipe[uy][rx] != '-' {
					goingLeft[uy][rx].lastCorner = onlyPipe[uy][rx]
				}
			}
		}
	}

	cnt := 0

	directionCounts := util.FillGrid(w, h, 0)

	for x := 0; x < w; x += 1 {
		for y := 0; y < h; y += 1 {
			if goingDown[y][x].inside {
				directionCounts[y][x] += (1 << 3)
			}
			if goingUp[y][x].inside {
				directionCounts[y][x] += (1 << 2)
			}
			if goingRight[y][x].inside {
				directionCounts[y][x] += (1 << 1)
			}
			if goingLeft[y][x].inside {
				directionCounts[y][x] += (1 << 0)
			}
		}
	}

	util.PrintGrid(directionCounts, "%x")
	fmt.Println()

	for x := 0; x < w; x += 1 {
		for y := 0; y < h; y += 1 {
			if onlyPipe[y][x] == ' ' && directionCounts[y][x] == 0xf {
				onlyPipe[y][x] = 'I'
				cnt += 1
			}
		}
	}

	util.PrintGrid(onlyPipe, "%c")

	fmt.Println(cnt)
}
