package main

import (
	"fmt"
	"strings"
)

func main() {
	_, _ = input, testinput
	a, b, onea, oneb := 0, 0, 0, 0
	for _, i := range input {
		onea, oneb = decode(strings.TrimSpace(strings.Split(i, "|")[0]), strings.TrimSpace(strings.Split(i, "|")[1]))
		a += onea
		b += oneb
	}
	fmt.Println(a)
	fmt.Println(b)
}

func decode(patterns string, output string) (unique int, outputval int) {
	org, decoded, decodemap, out := [10]string{}, [10]string{}, make(map[string]int), strings.Split(output, " ")
	for i, j := range strings.Split(patterns, " ") {
		sorted := day8sort(j)
		switch len(sorted) {
		case 2:
			decoded[1] = sorted
		case 4:
			decoded[4] = sorted
		case 3:
			decoded[7] = sorted
		case 7:
			decoded[8] = sorted
		default:
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
		z := 0
		for _, c := range decoded[6] {
			if strings.Contains(j, string(c)) {
				z += 1
			}
		}
		if z == 5 {
			decoded[5] = j
		} else if z == 4 {
			decoded[2] = j
		}
	}
	for i, j := range decoded {
		decodemap[j] = i
	}
	for i, j := range []int{1000, 100, 10, 1} {
		if len(out[i]) == 2 || len(out[i]) == 4 || len(out[i]) == 3 || len(out[i]) == 7 {
			unique++
		}
		outputval += decodemap[day8sort(out[i])] * j
	}
	return unique, outputval
}

func day8contains(input string, teststring string) bool {
	for _, c := range teststring {
		if !strings.Contains(input, string(c)) {
			return false
		}
	}
	return true
}

func day8sort(input string) (sorted string) {
	var chars = [7]string{"a", "b", "c", "d", "e", "f", "g"}
	for _, c := range chars {
		if strings.Contains(input, c) {
			sorted += c
		}
	}
	return sorted
}
