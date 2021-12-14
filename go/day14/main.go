package main

import (
	"fmt"
)

func main() {
	_, _ = input, testinput
	template := input[0]
	rules := make(map[string]string)
	for i, j := range input {
		if i > 1 {
			rules[string(j[:2])] = string(j[6])
		}
	}
	polymer := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		polymer[string(template[i])+string(template[i+1])] += 1
	}
	for i := 0; i < 40; i++ {
		newpolymer := make(map[string]int)
		for k, l := range polymer {
			newval := rules[k]
			newpolymer[string(k[0])+newval] += l
			newpolymer[newval+string(k[1])] += l
		}
		polymer = newpolymer
		if i == 9 || i == 39 {
			counts := make(map[string]int)
			for k, l := range polymer {
				counts[string(k[0])] += l
			}
			counts[string(template[len(template)-1])] += 1
			min, max := 0, 0
			for _, j := range counts {
				if min == 0 {
					min = j
					max = j
				}
				if j > max {
					max = j
				} else if j < min {
					min = j
				}
			}
			fmt.Println(max - min)
		}
	}
}
