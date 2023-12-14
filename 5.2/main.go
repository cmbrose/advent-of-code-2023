package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"main/util"
)

type mapping struct {
	label   string
	entries []mapEntry
}

type mapEntry struct {
	destStart int
	srcStart  int
	length    int
}

type mapRange struct {
	start  int
	length int
}

func main() {
	blocks := util.ReadInputBlocks("./input.txt")

	seedInts := util.Map(strings.Fields(blocks[0][0])[1:], util.AssertInt)
	var seedRanges []mapRange
	for i := 0; i < len(seedInts); i += 2 {
		seedRanges = append(seedRanges, mapRange{
			seedInts[i],
			seedInts[i+1],
		})
	}

	mappings := util.Map(blocks[1:], func(lines []string) mapping {
		entries := util.Map(lines[1:], func(line string) mapEntry {
			fields := strings.Fields(line)
			return mapEntry{
				util.AssertInt(fields[0]),
				util.AssertInt(fields[1]),
				util.AssertInt(fields[2]),
			}
		})

		sort.Slice(entries, func(i, j int) bool {
			return entries[i].srcStart < entries[j].srcStart
		})

		return mapping{
			lines[0],
			entries,
		}
	})

	minLocation := math.MaxInt
	for _, r := range seedRanges {
		ranges := []mapRange{r}

		for _, m := range mappings {
			var next []mapRange
			for _, r := range ranges {
				next = append(next, m.mapRange(r)...)
			}

			sort.Slice(next, func(i, j int) bool {
				return next[i].start < next[j].start
			})

			ranges = next
		}

		if ranges[0].start < minLocation {
			minLocation = ranges[0].start
		}
	}

	fmt.Println(minLocation)
}

func (m mapping) mapRange(r mapRange) []mapRange {
	var mapped []mapRange

	for _, e := range m.entries {
		if r.start < e.srcStart {
			diff := e.srcStart - r.start
			length := util.Min(diff, r.length)
			r2 := mapRange{
				r.start,
				length,
			}

			mapped = append(mapped, r2)

			r.start += length
			r.length -= length
		}

		if r.length == 0 {
			break
		}

		offset := r.start - e.srcStart
		if offset < e.length {
			length := r.length
			if offset+length > e.length {
				length = e.length - offset
			}
			r2 := mapRange{
				e.destStart + offset,
				length,
			}

			mapped = append(mapped, r2)

			r.start += length
			r.length -= length
		}

		if r.length == 0 {
			break
		}
	}

	if r.length > 0 {
		mapped = append(mapped, r)
	}

	return mapped
}
