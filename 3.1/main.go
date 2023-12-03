package main

import (
	"fmt"
	"main/util"
)

type partNumber struct {
	value     int
	x1, x2, y int
}

type symbol struct {
	value byte
	x, y  int
}

func main() {
	var (
		partNumbers []partNumber
		symbols     []symbol
	)

	for y, line := range util.ReadInputLines("./input.txt") {
		for x := 0; x < len(line); x += 1 {
			c := line[x]

			if line[x] == '.' {
				continue
			}

			if util.IsNumber(c) {
				v := 0
				x1 := x
				for ; x < len(line) && util.IsNumber(line[x]); x += 1 {
					v = v*10 + (int)(line[x]-'0')
				}
				x -= 1
				x2 := x

				partNumbers = append(partNumbers, partNumber{v, x1, x2, y})
			} else {
				symbols = append(symbols, symbol{c, x, y})
			}
		}
	}

	sum := 0

	for _, p := range partNumbers {
		for _, s := range symbols {
			if isAdjacent(p, s) {
				fmt.Printf("%d at %d-%d, %d is adjacent to %c at %d, %d\n", p.value, p.x1, p.x2, p.y, s.value, s.x, s.y)
				sum += p.value
				goto end
			}
		}

		fmt.Printf("%d at %d-%d, %d has no adjacencies\n", p.value, p.x1, p.x2, p.y)

	end:
	}

	println(sum)
}

func isAdjacent(p partNumber, s symbol) bool {
	return s.x >= p.x1-1 && s.x <= p.x2+1 && s.y >= p.y-1 && s.y <= p.y+1
}
