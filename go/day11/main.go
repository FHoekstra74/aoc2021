package main

import (
	"fmt"
	"strconv"
)

type point struct {
	x int
	y int
}

func main() {
	_, _ = input, testinput
	grid, vectors, a, run := make([][]int, 10), []point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}, {1, 1}, {-1, 1}, {1, -1}, {-1, -1}}, 0, 0
	for i := range grid {
		grid[i] = make([]int, 10)
	}
	for y, row := range input {
		for x, val := range row {
			val, _ := strconv.Atoi(string(val))
			grid[x][y] = val
		}
	}
	for {
		run++
		flashes, flashcount := []point{}, 0
		for x := 0; x < 10; x++ {
			for y := 0; y < 10; y++ {
				if grid[x][y] == 9 {
					grid[x][y] = 0
					flashes = append(flashes, point{x, y})
					flashcount++
				} else {
					grid[x][y] = grid[x][y] + 1
				}
			}
		}
		for len(flashes) > 0 {
			p, newflashes := flashes[0], []point{}
			flashes = flashes[1:]
			for _, v := range vectors {
				if p.x+v.x >= 0 && p.x+v.x < 10 && p.y+v.y >= 0 && p.y+v.y < 10 {
					newflashes = append(newflashes, flashsurrounding(grid, point{p.x + v.x, p.y + v.y})...)
				}
			}
			for _, p := range newflashes {
				flashes = append(flashes, p)
				flashcount++
			}
		}
		a += flashcount
		if run == 100 {
			fmt.Println(a)
		}
		if flashcount == 100 {
			fmt.Println(run)
			break
		}
	}
}

func flashsurrounding(grid [][]int, p point) []point {
	newflashes := []point{}
	if grid[p.x][p.y] == 9 {
		grid[p.x][p.y] = 0
		newflashes = append(newflashes, point{p.x, p.y})
	} else {
		if grid[p.x][p.y] > 0 {
			grid[p.x][p.y] = grid[p.x][p.y] + 1
		}
	}
	return newflashes
}
