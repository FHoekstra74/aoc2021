package main

import (
	"fmt"
	"strings"
	"unicode"
)

type cave struct {
	name    string
	tocaves []string
	lower   bool
}

func (c *cave) nextcaves(m map[string]cave, prev string, smalltwice bool) []string {
	var res = []string{}
	if c.name == "end" {
		res = append(res, "end")
		return res
	}
	for _, n := range c.tocaves {
		valid, to, lowcount := true, m[n], 0
		if to.lower {
			for _, k := range strings.Split(prev, "-") {
				if k == to.name {
					lowcount += 1
				}
			}
			if lowcount > 1 || (to.name == "end" && lowcount == 1) || (lowcount == 1 && !smalltwice) {
				valid = false
			}
		}
		if valid {
			for _, l := range to.nextcaves(m, prev+"-"+c.name, smalltwice && lowcount == 0) {
				res = append(res, c.name+"-"+l)
			}
		}
	}
	return res
}

func main() {
	_, _ = input, testinput
	m := make(map[string]cave)
	for _, l := range input {
		from, to := strings.Split(l, "-")[0], strings.Split(l, "-")[1]
		if _, exists := m[from]; !exists {
			c := cave{name: from, lower: unicode.IsLower(rune(from[0]))}
			m[from] = c
		}
		if _, exists := m[to]; !exists {
			c := cave{name: to, lower: unicode.IsLower(rune(to[0]))}
			m[to] = c
		}
		if to != "start" && from != "end" {
			c := m[from]
			c.tocaves = append(c.tocaves, to)
			m[from] = c
		}
		if from != "start" && to != "end" {
			c := m[to]
			c.tocaves = append(c.tocaves, from)
			m[to] = c
		}
	}
	start := m["start"]
	a := start.nextcaves(m, "", false)
	fmt.Println(len(a))
	b := start.nextcaves(m, "", true)
	fmt.Println(len(b))
}
