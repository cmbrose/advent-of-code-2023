package main

import (
	"fmt"
	"strings"

	"main/util"
)

func calcAchievedDistance(raceTime, holdTime int) int {
	return holdTime * (raceTime - holdTime)
}

func main() {
	lines := util.ReadInputLines("./input.txt")
	times := util.Map(strings.Fields(lines[0])[1:], util.AssertInt)
	distances := util.Map(strings.Fields(lines[1])[1:], util.AssertInt)

	possibilities := 1
	for i := 0; i < len(times); i += 1 {
		t, d := times[i], distances[i]

		cnt := 0

		for h := 1; h < t-1; h += 1 {
			if calcAchievedDistance(t, h) > d {
				cnt += 1
			}
		}

		possibilities *= cnt
	}

	fmt.Printf("%d\n", possibilities)
}
