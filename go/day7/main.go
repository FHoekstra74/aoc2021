package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	_, _ = input, testinput
	high, low := 0, 0
	var positions = []int{}
	for i, j := range strings.Split(input[0], ",") {
		j, _ := strconv.Atoi(j)
		if j < low || i == 0 {
			low = j
		}
		if j > high || i == 0 {
			high = j
		}
		positions = append(positions, j)
	}
	fuelb, fuela := 0, 0
	for i := low; i <= high; i++ {
		a, b := 0, 0
		for _, crab := range positions {
			t := crab - i
			if t < 0 {
				t *= -1
			}
			for j := 0; j < t; j++ {
				b += (j + 1)
			}
			a += t
		}
		if i == low || b < fuelb {
			fuelb = b
		}
		if i == low || a < fuela {
			fuela = a
		}
	}
	fmt.Println(fuela)
	fmt.Println(fuelb)
}
