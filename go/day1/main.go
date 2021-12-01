package main

import (
	"fmt"
)

func main() {
	a, b := 0, 0
	for i, v := range input {
		if i > 0 && v > input[i-1] {
			a += 1
		}
		if i > 2 && input[i] > input[i-3] {
			b += 1
		}
	}
	fmt.Println(a)
	fmt.Println(b)
}
