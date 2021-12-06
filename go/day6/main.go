package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	_, _ = input, testinput
	var fishes [9]int
	for _, j := range strings.Split(input[0], ",") {
		j, _ := strconv.Atoi(j)
		fishes[j]++
	}
	for i := 0; i < 256; i++ {
		if i == 80 {
			a := 0
			for _, v := range fishes {
				a += v
			}
			fmt.Println(a)
		}
		prev := fishes
		for j := 0; j < 8; j++ {
			fishes[j] = prev[j+1]
		}
		fishes[8] = prev[0]
		fishes[6] += prev[0]
	}
	b := 0
	for _, v := range fishes {
		b += v
	}
	fmt.Println(b)
}
