package main

import (
	"fmt"
	"sort"
)

func main() {
	_, _ = input, testinput
	a, b, closers, ascore, bscore := 0, []int{}, map[string]string{"{": "}", "(": ")", "<": ">", "[": "]"}, map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}, map[string]int{")": 1, "]": 2, "}": 3, ">": 4}
	for _, row := range input {
		corrupt, completion, bs := false, "", 0
		for _, c := range row {
			switch string(c) {
			case "(", "<", "{", "[":
				completion = closers[string(c)] + completion
			case ")", ">", "}", "]":
				if string(completion[0]) == string(c) {
					completion = completion[1:]
				} else {
					corrupt = true
					a += ascore[string(c)]
				}
			}
			if corrupt {
				break
			}
		}
		if !corrupt {
			for _, c := range completion {
				bs *= 5
				bs += bscore[string(c)]
			}
			b = append(b, bs)
		}
	}
	fmt.Println(a)
	sort.Ints(b)
	fmt.Println(b[(len(b) / 2)])
}
