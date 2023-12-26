package main

import (
	"fmt"
	"strings"

	"main/util"
)

type branch struct {
	n string
	p part
}

type check func(p part) []branch

type minMax struct {
	min, max int
}

type part map[string]minMax

func main() {
	blocks := util.ReadInputBlocks("./input.txt")

	checks := make(map[string]check)

	for _, line := range blocks[0] {
		pair := strings.Split(line, "{")

		name := pair[0]
		checkStrs := strings.TrimSuffix(pair[1], "}")

		checks[name] = parseCheck(checkStrs)
	}

	var combos int64

	checks["A"] = func(p part) []branch {
		prod := int64(1)
		for _, mm := range p {
			prod *= int64(mm.max - mm.min + 1)
		}

		combos += prod
		return nil
	}
	checks["R"] = func(p part) []branch {
		return nil
	}

	q := []branch{
		{
			"in",
			part{
				"x": minMax{1, 4000},
				"m": minMax{1, 4000},
				"a": minMax{1, 4000},
				"s": minMax{1, 4000},
			},
		},
	}

	for len(q) > 0 {
		n, p := q[0].n, q[0].p
		q = q[1:]

		bs := checks[n](p)
		q = append(q, bs...)
	}

	fmt.Printf("%d\n", combos)
}

func parseCheck(s string) check {
	ss := strings.Split(s, ",")

	def := ss[len(ss)-1]

	cs := util.Map(ss[:len(ss)-1], parseSingleCheck)

	return func(p part) []branch {
		var bs []branch

		for _, c := range cs {
			subs := c(p)
			bs = append(bs, subs...)
		}

		bs = append(bs, branch{def, p})

		return bs
	}
}

func parseSingleCheck(s string) check {
	r := s[:1]
	o := rune(s[1])

	pair := strings.Split(s[2:], ":")
	v := util.AssertInt(pair[0])
	t := pair[1]

	if o == '>' {
		return func(p part) []branch {
			mm := p[r]

			if mm.max <= v {
				return nil
			}

			bmm := minMax{util.Max(mm.min, v+1), mm.max}
			bs := []branch{
				{t, p.cloneWith(r, bmm)},
			}

			p[r] = minMax{mm.min, v}
			return bs
		}
	}

	return func(p part) []branch {
		mm := p[r]
		if mm.min >= v {
			return nil
		}

		bmm := minMax{mm.min, util.Min(mm.max, v-1)}
		bs := []branch{
			{t, p.cloneWith(r, bmm)},
		}

		p[r] = minMax{v, mm.max}
		return bs
	}
}

func (p part) cloneWith(nk string, nv minMax) part {
	c := make(part)

	for k, v := range p {
		c[k] = v
	}

	c[nk] = nv
	return c
}
