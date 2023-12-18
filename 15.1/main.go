package main

import (
	"fmt"
	"strings"

	"main/util"
)

func main() {
	line := util.ReadInputLines("./input.txt")[0]

	strs := strings.Split(line, ",")

	fmt.Printf("%d\n", util.Sum(util.Map(strs, hash)))
}

func hash(str string) int {
	var h byte

	for _, r := range str {
		h = (h + byte(r)) * 17
	}

	return int(h)
}
