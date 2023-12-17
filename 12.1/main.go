package main

import (
	"fmt"
	"strings"

	"main/util"
)

type input struct {
	springs string
	counts  []int
}

func main() {
	counts := util.Map(util.ReadInputLines("./input.txt"), func(line string) int {
		pair := strings.Fields(line)

		return countVariations(
			[]rune(pair[0]),
			util.Map(strings.Split(pair[1], ","), util.AssertInt),
			0,
		)
	})

	total := util.Sum(counts)

	fmt.Printf("%d\n", total)
}

func countVariations(springs []rune, counts []int, require rune) int {
	if len(springs) == 0 {
		if len(counts) != 0 {
			return 0
		}
		return 1
	}
	if len(counts) == 0 {
		if len(util.Filter(springs, func(r rune) bool { return r == '#' })) != 0 {
			return 0
		}
		return 1
	}

	if springs[0] == '#' {
		if require == '.' {
			return 0
		}

		counts[0] -= 1
		recurseCounts, require := counts, '#'
		if recurseCounts[0] == 0 {
			recurseCounts = recurseCounts[1:]
			require = '.'
		}

		cnt := countVariations(springs[1:], recurseCounts, require)

		counts[0] += 1
		return cnt
	}

	if springs[0] == '.' {
		if require == '#' {
			return 0
		}

		return countVariations(springs[1:], counts, 0)
	}

	cnt := 0

	springs[0] = '#'
	cnt += countVariations(springs, counts, require)

	springs[0] = '.'
	cnt += countVariations(springs, counts, require)

	springs[0] = '?'
	return cnt
}
