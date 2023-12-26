package main

import (
	"fmt"
	"regexp"
	"strconv"

	"main/util"
)

type point struct {
	x, y int
}

type location int

const (
	UP   location = 2
	LEFT location = 1

	BR location = 0
	BL location = LEFT
	UR location = UP
	UL location = UP | LEFT
)

func main() {
	re := regexp.MustCompile(`([UDLR]) ([0-9]+) \(#([a-f0-9]+)\)`)

	lines := util.ReadInputLines("./input.txt")

	ps := make([]point, len(lines)+1)

	ps[0] = point{0, 0}
	inside := BR // assumed
	lastDx, lastDy := 0, 0

	for i, line := range lines {
		m := re.FindStringSubmatch(line)
		_, _, color := m[1], util.AssertInt(m[2]), m[3]

		hex, dir := color[:len(color)-1], color[len(color)-1:]
		cnt, err := strconv.ParseInt(hex, 16, 32)
		util.Check(err)

		dx, dy := 0, 0
		incr := 0

		if dir == "U" || dir == "3" {
			dy = -1

			if inside&UP == 0 {
				ps[i].x -= lastDx
				ps[i].y -= lastDy
				inside ^= LEFT
			} else {
				incr = 1
				inside ^= UP
			}
		} else if dir == "D" || dir == "1" {
			dy = 1

			if inside&UP != 0 {
				ps[i].x -= lastDx
				ps[i].y -= lastDy
				inside ^= LEFT
			} else {
				incr = 1
				inside ^= UP
			}
		} else if dir == "L" || dir == "2" {
			dx = -1

			if inside&LEFT == 0 {
				ps[i].x -= lastDx
				ps[i].y -= lastDy
				inside ^= UP
			} else {
				incr = 1
				inside ^= LEFT
			}
		} else if dir == "R" || dir == "0" {
			dx = 1

			if inside&LEFT != 0 {
				ps[i].x -= lastDx
				ps[i].y -= lastDy
				inside ^= UP
			} else {
				incr = 1
				inside ^= LEFT
			}
		}

		ps[i+1] = point{
			ps[i].x + dx*(int(cnt)+incr),
			ps[i].y + dy*(int(cnt)+incr),
		}

		lastDx, lastDy = dx, dy
	}

	//minX := util.Min(util.Map(ps, func(p point) int { return p.x })...)
	//maxX := util.Max(util.Map(ps, func(p point) int { return p.x })...)
	//minY := util.Min(util.Map(ps, func(p point) int { return p.y })...)
	//maxY := util.Max(util.Map(ps, func(p point) int { return p.y })...)

	//width := maxX - minX + 1
	//height := maxY - minY + 1

	// for i := range ps {
	// 	ps[i].y -= minY
	// }

	sum := int64(0)
	for i := range ps[:len(ps)-1] {
		p1, p2 := ps[i], ps[i+1]

		sum += int64(p1.y+p2.y) * int64(p1.x-p2.x)
	}

	fmt.Println(sum / 2)

	//fmt.Println(ps)
}
