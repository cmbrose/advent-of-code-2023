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

	t := util.AssertInt(strings.Split(strings.Replace(lines[0], " ", "", -1), ":")[1])
	d := util.AssertInt(strings.Split(strings.Replace(lines[1], " ", "", -1), ":")[1])

	cnt := 0

	for h := 1; h < t-1; h += 1 {
		if calcAchievedDistance(t, h) > d {
			cnt += 1
		}
	}

	fmt.Printf("%d\n", cnt)
}
