package main

import (
	"fmt"
	"strings"

	"main/util"
)

func main() {
	counts := util.Map(util.ReadInputLines("./input.txt"), func(line string) int {
		pair := strings.Fields(line)

		springs := strings.Join(util.Repeat(pair[0], 5), "?")
		counts := strings.Join(util.Repeat(pair[1], 5), ",")

		cache := make(map[string]int)

		cnt := countVariations(
			[]rune(springs),
			util.Map(strings.Split(counts, ","), util.AssertInt),
			0,
			cache,
		)

		return cnt
	})

	total := util.Sum(counts)

	fmt.Printf("%d\n", total)
}

func countVariations(springs []rune, counts []int, require rune, cache map[string]int) (cnt int) {
	k := fmt.Sprintf("%v %v %c", springs, counts, require)
	if c, ok := cache[k]; ok {
		return c
	}

	defer func() {
		cache[k] = cnt
	}()

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

		cnt := countVariations(springs[1:], recurseCounts, require, cache)

		counts[0] += 1
		return cnt
	}

	if springs[0] == '.' {
		if require == '#' {
			return 0
		}

		return countVariations(springs[1:], counts, 0, cache)
	}

	springs[0] = '#'
	cnt += countVariations(springs, counts, require, cache)

	springs[0] = '.'
	cnt += countVariations(springs, counts, require, cache)

	springs[0] = '?'
	return cnt
}
