package main

import (
	"fmt"
	"math"
	"strings"

	"main/util"
)

type mapping struct {
	label   string
	entries []mapEntry
}

type mapEntry struct {
	destStart   int
	sourceStart int
	length      int
}

func main() {
	blocks := util.ReadInputBlocks("./input.txt")

	seeds := util.Map(strings.Fields(blocks[0][0])[1:], util.AssertInt)

	mappings := util.Map(blocks[1:], func(lines []string) mapping {
		return mapping{
			lines[0],
			util.Map(lines[1:], func(line string) mapEntry {
				fields := strings.Fields(line)
				return mapEntry{
					util.AssertInt(fields[0]),
					util.AssertInt(fields[1]),
					util.AssertInt(fields[2]),
				}
			}),
		}
	})

	minLocation := math.MaxInt
	for _, v := range seeds {
		for _, m := range mappings {
			v = m.doMapping(v)
		}

		if v < minLocation {
			minLocation = v
		}
	}

	fmt.Println(minLocation)
}

func (m mapping) doMapping(src int) int {
	for _, e := range m.entries {
		srcOffset := src - e.sourceStart
		if srcOffset >= 0 && srcOffset < e.length {
			return e.destStart + srcOffset
		}
	}

	return src
}
