package main

import (
	"fmt"

	"main/util"
)

func main() {
	scores := util.Map(util.ReadInputBlocks("./input.txt"), func(block []string) int {
		for i := 1; i < len(block); i += 1 {
			if searchForMirrorAtRow(block, i) {
				return i * 100
			}
		}

		block = rotate(block)

		for i := 1; i < len(block); i += 1 {
			if searchForMirrorAtRow(block, i) {
				return i
			}
		}

		panic("Did not find a reflection line!")
	})

	fmt.Printf("%d\n", util.Sum(scores))
}

func rotate(block []string) []string {
	rotated := util.Grid[rune](len(block), len(block[0]))

	for y := 0; y < len(block); y += 1 {
		for x, r := range block[y] {
			rotated[x][y] = r
		}
	}

	return util.Map(rotated, func(line []rune) string { return string(line) })
}

func searchForMirrorAtRow(block []string, i int) bool {
	for j := 0; i+j < len(block) && i-j-1 >= 0; j += 1 {
		if block[i+j] != block[i-j-1] {
			return false
		}
	}

	return true
}
