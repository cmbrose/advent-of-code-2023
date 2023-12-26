package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"main/util"
)

type check func(p part) string

type part map[string]int

func main() {
	blocks := util.ReadInputBlocks("./input.txt")

	checks := make(map[string]check)

	for _, line := range blocks[0] {
		pair := strings.Split(line, "{")

		name := pair[0]
		checkStrs := strings.TrimSuffix(pair[1], "}")

		checks[name] = parseCheck(checkStrs)
	}

	sum := 0

	checks["A"] = func(p part) string {
		sum += util.Sum(util.Values(p))
		return ""
	}
	checks["R"] = func(p part) string {
		return ""
	}

	for _, p := range util.Map(blocks[1], func(line string) (p part) {
		line = strings.ReplaceAll(line, "=", "\":")
		line = strings.ReplaceAll(line, ",", ",\"")
		line = strings.ReplaceAll(line, "{", "{\"")

		err := json.Unmarshal([]byte(line), &p)
		util.Check(err)
		return
	}) {
		n := "in"
		for n != "" {
			c := checks[n]
			n = c(p)
		}
	}

	fmt.Printf("%d\n", sum)
}

func parseCheck(s string) check {
	ss := strings.Split(s, ",")

	def := ss[len(ss)-1]

	cs := util.Map(ss[:len(ss)-1], parseSingleCheck)

	return func(p part) string {
		for _, c := range cs {
			if t := c(p); t != "" {
				return t
			}
		}
		return def
	}
}

func parseSingleCheck(s string) check {
	r := s[:1]
	o := rune(s[1])

	pair := strings.Split(s[2:], ":")
	v := util.AssertInt(pair[0])
	t := pair[1]

	if o == '>' {
		return func(p part) string {
			a := p[r]
			if a > v {
				return t
			}
			return ""
		}
	}

	return func(p part) string {
		a := p[r]
		if a < v {
			return t
		}
		return ""
	}
}
