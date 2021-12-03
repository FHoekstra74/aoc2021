package main

import (
	"fmt"
	"strconv"
)

func calc_rating(mostcommon bool) int64 {
	inputmap := make(map[string]bool)
	for i := 0; i < len(input); i++ {
		inputmap[input[i]] = true
	}
	for i := 0; i < len(input[0]); i++ {
		c := []int{0, 0}
		for k, v := range inputmap {
			if k[i] == 48 && v {
				c[0]++
			}
			if k[i] == 49 && v {
				c[1]++
			}
		}
		if c[0]+c[1] == 1 {
			break
		}
		if c[0] == c[1] {
			c[0] = 0
		}
		for k, v := range inputmap {
			inputmap[k] = v && !(k[i] == 49 && c[1] < c[0] && mostcommon || k[i] == 48 && c[0] < c[1] && mostcommon || k[i] == 49 && c[0] < c[1] && !mostcommon || k[i] == 48 && c[1] < c[0] && !mostcommon)
		}
	}
	for k, v := range inputmap {
		if v {
			l, _ := strconv.ParseInt(k, 2, 32)
			return l
		}
	}
	return 0
}

func main() {
	_, _ = input, testinput
	gamma, epsilon := "", ""
	for i := 0; i < len(input[0]); i++ {
		c := []int{0, 0}
		for _, v := range input {
			if v[i] == 48 {
				c[0]++
			} else {
				c[1]++
			}
		}
		if c[0] < c[1] {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}
	g, _ := strconv.ParseInt(gamma, 2, 32)
	e, _ := strconv.ParseInt(epsilon, 2, 32)
	fmt.Println(g * e)
	fmt.Println(calc_rating(true) * calc_rating(false))
}
