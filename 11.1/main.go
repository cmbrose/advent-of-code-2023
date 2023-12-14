package main

import (
	"fmt"

	"main/util"
)

type point struct {
	x, y int
}

func main() {
	var (
		lines = util.ReadInputLines("./input.txt")

		w = len(lines[0])
		h = len(lines)

		galaxies        []point
		columnHasGalaxy = make([]bool, w)
		rowHasGalaxy    = make([]bool, h)
	)

	for y, line := range lines {
		for x, r := range line {
			if r == '#' {
				galaxies = append(galaxies, point{x, y})

				columnHasGalaxy[x] = true
				rowHasGalaxy[y] = true
			}
		}
	}

	xOffsets := make([]int, w)
	for i, hasGalaxy := range columnHasGalaxy {
		incr := 0
		if !hasGalaxy {
			incr = 1
		}

		if i == 0 {
			xOffsets[i] = incr
		} else {
			xOffsets[i] = xOffsets[i-1] + incr
		}
	}

	yOffsets := make([]int, h)
	for i, hasGalaxy := range rowHasGalaxy {
		incr := 0
		if !hasGalaxy {
			incr = 1
		}

		if i == 0 {
			yOffsets[i] = incr
		} else {
			yOffsets[i] = yOffsets[i-1] + incr
		}
	}

	for i := range galaxies {
		galaxies[i].x += xOffsets[galaxies[i].x]
		galaxies[i].y += yOffsets[galaxies[i].y]
	}

	sum := 0
	for i := 0; i < len(galaxies)-1; i += 1 {
		for j := i + 1; j < len(galaxies); j += 1 {
			dx, dy := galaxies[i].x-galaxies[j].x, galaxies[i].y-galaxies[j].y

			sum += util.Abs(dx) + util.Abs(dy)
		}
	}

	fmt.Printf("%d\n", sum)
}
