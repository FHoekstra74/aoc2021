package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	_, _ = input, testinput
	heightmap, a, b := [100][100]int{}, 0, []int{}
	for y, row := range input {
		for x, val := range row {
			val, _ := strconv.Atoi(string(val))
			heightmap[x][y] = val
		}
	}
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			ok, val := true, heightmap[x][y]
			if y > 0 && val >= heightmap[x][y-1] {
				ok = false
			}
			if y < 99 && val >= heightmap[x][y+1] {
				ok = false
			}
			if x > 0 && val >= heightmap[x-1][y] {
				ok = false
			}
			if x < 99 && val >= heightmap[x+1][y] {
				ok = false
			}
			if ok {
				a += val
				b = append(b, basinsize(x, y, heightmap))
			}
		}
	}
	fmt.Println(a + len(b))
	sort.Ints(b)
	fmt.Println(b[len(b)-1] * b[len(b)-2] * b[len(b)-3])
}

func basinsize(x, y int, heightmap [100][100]int) int {
	res := 1
	heightmap[x][y] = 10
	for {
		onefound := false
		for x := 0; x < 100; x++ {
			for y := 0; y < 100; y++ {
				ok, val := false, heightmap[x][y]
				if val < 9 && y > 0 && heightmap[x][y-1] == 10 {
					ok = true
				}
				if val < 9 && y < 99 && heightmap[x][y+1] == 10 {
					ok = true
				}
				if val < 9 && x > 0 && heightmap[x-1][y] == 10 {
					ok = true
				}
				if val < 9 && x < 99 && heightmap[x+1][y] == 10 {
					ok = true
				}
				if ok {
					heightmap[x][y] = 10
					onefound = true
					res += 1
				}
			}
		}
		if !onefound {
			break
		}
	}
	return res
}
