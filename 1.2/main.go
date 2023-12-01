package main

import (
	"fmt"
	"strings"

	"main/util"
)

func main() {
	curr := 0

	for _, line := range util.ReadInputLines("./input.txt") {
		var (
			first = 0
			last  = 0
		)

		for i := 0; i < len(line); i += 1 {
			if util.IsNumber(line[i]) {
				n := (int)(line[i] - '0')
				if first == 0 {
					first = n
				}
				last = n
			}

			if n := GetNumberWord(line[i:]); n != -1 {
				if first == 0 {
					first = n
				}
				last = n
			}
		}

		value := 10*first + last

		fmt.Printf("%s -> %d, %d\n", line, first, last)

		curr += value
	}

	fmt.Printf("%d\n", curr)
}

func GetNumberWord(s string) int {
	if strings.HasPrefix(s, "one") {
		return 1
	}
	if strings.HasPrefix(s, "two") {
		return 2
	}
	if strings.HasPrefix(s, "three") {
		return 3
	}
	if strings.HasPrefix(s, "four") {
		return 4
	}
	if strings.HasPrefix(s, "five") {
		return 5
	}
	if strings.HasPrefix(s, "six") {
		return 6
	}
	if strings.HasPrefix(s, "seven") {
		return 7
	}
	if strings.HasPrefix(s, "eight") {
		return 8
	}
	if strings.HasPrefix(s, "nine") {
		return 9
	}

	return -1
}
