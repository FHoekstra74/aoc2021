package main

import (
	"fmt"
	"strings"
)

func main() {
	_, _ = input, testinput
	a, b := 0, 0
	for _, i := range input {
		t := strings.TrimSpace(strings.Split(i, "|")[1])
		for _, j := range strings.Split(t, " ") {
			if len(j) == 2 || len(j) == 4 || len(j) == 3 || len(j) == 7 {
				a++
			}
		}
	}
	fmt.Println(a)
	for _, i := range input {
		t := strings.Split(i, "|")
		b += decode(strings.TrimSpace(t[0]), strings.TrimSpace(t[1]))
	}
	fmt.Println(b)
}

func decode(patterns string, output string) int {
	var org = [10]string{}
	var decoded = [10]string{}
	for i, j := range strings.Split(patterns, " ") {
		sorted := day8sort(j)
		if len(sorted) == 2 {
			decoded[1] = sorted
		} else if len(sorted) == 4 {
			decoded[4] = sorted
		} else if len(sorted) == 3 {
			decoded[7] = sorted
		} else if len(sorted) == 7 {
			decoded[8] = sorted
		} else {
			org[i] = sorted
		}
	}
	for i, j := range org {
		if len(j) == 6 && day8contains(j, decoded[1]) && day8contains(j, decoded[4]) && day8contains(j, decoded[7]) {
			decoded[9] = j
			org[i] = ""
		}
	}
	for i, j := range org {
		if len(j) == 6 && day8contains(j, decoded[1]) {
			decoded[0] = j
			org[i] = ""
		}
	}
	for i, j := range org {
		if len(j) == 5 && day8contains(j, decoded[1]) {
			decoded[3] = j
			org[i] = ""
		} else if len(j) == 6 {
			decoded[6] = j
			org[i] = ""
		}
	}
	for _, j := range org {
		if len(j) > 0 {
			z := 0
			for _, c := range decoded[6] {
				if strings.Contains(j, string(c)) {
					z += 1
				}
			}
			if z == 5 {
				decoded[5] = j
			} else {
				decoded[2] = j
			}
		}
	}
	res := 0
	for k, j := range strings.Split(output, " ") {
		sorted := day8sort(j)
		for i, j := range decoded {
			if j == sorted {
				if k == 0 {
					res += i * 1000
				}
				if k == 1 {
					res += i * 100
				}
				if k == 2 {
					res += i * 10
				}
				if k == 3 {
					res += i * 1
				}
			}
		}
	}
	return res
}

func day8contains(input string, teststring string) bool {
	for _, c := range teststring {
		if !strings.Contains(input, string(c)) {
			return false
		}
	}
	return true
}

func day8sort(input string) string {
	sorted := ""
	var chars = [7]string{"a", "b", "c", "d", "e", "f", "g"}
	for _, c := range chars {
		if strings.Contains(input, c) {
			sorted += c
		}
	}
	return sorted
}
