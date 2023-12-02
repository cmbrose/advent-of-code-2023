package main

import (
	"fmt"
	"strings"

	"main/util"
)

type game struct {
	id    int
	pulls []pullResult
}

type pullResult struct {
	red, green, blue int
}

func main() {
	lines := util.ReadInputLines("./input.txt")

	games := util.Map(lines, func(line string) game {
		gameSplit := strings.Split(line, ": ")

		return game{
			id: util.AssertInt(strings.Split(gameSplit[0], " ")[1]),
			pulls: util.Map(strings.Split(gameSplit[1], "; "), func(pullString string) pullResult {
				colors := strings.Split(pullString, ", ")
				pull := pullResult{}

				for _, color := range colors {
					pair := strings.Split(color, " ")
					switch pair[1] {
					case "red":
						pull.red = util.AssertInt(pair[0])
					case "blue":
						pull.blue = util.AssertInt(pair[0])
					case "green":
						pull.green = util.AssertInt(pair[0])
					}
				}

				return pull
			}),
		}
	})

	var (
		maxRed   = 12
		maxBlue  = 14
		maxGreen = 13
	)

	validIdSum := 0

	for _, game := range games {
		for _, pull := range game.pulls {
			if pull.red > maxRed || pull.blue > maxBlue || pull.green > maxGreen {
				goto end
			}
		}

		validIdSum += game.id

	end:
	}

	fmt.Printf("%d\n", validIdSum)
}
