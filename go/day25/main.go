package main

import (
	"fmt"
)

type point struct {
	x, y int
}

func main() {
	_, _ = input, testinput
	grid := make(map[point]string)
	for y, line := range input {
		for x, c := range line {
			grid[point{x, y}] = string(c)
		}
	}
	maxx, maxy, step := len(input[0])-1, len(input)-1, 1
	for {
		moved, new := false, make(map[point]string)
		for y := 0; y <= maxy; y++ {
			for x := 0; x <= maxx; x++ {
				new[point{x, y}] = "."
			}
		}
		for y := 0; y <= maxy; y++ {
			for x := 0; x <= maxx; x++ {
				if grid[point{x, y}] == ">" {
					xx := x + 1
					if xx > maxx {
						xx = 0
					}
					if grid[point{xx, y}] == "." {
						new[point{xx, y}] = ">"
						moved = true
					} else {
						new[point{x, y}] = ">"
					}
				}
			}
		}
		for y := 0; y <= maxy; y++ {
			for x := 0; x <= maxx; x++ {
				if grid[point{x, y}] == "v" {
					yy := y + 1
					if yy > maxy {
						yy = 0
					}
					if new[point{x, yy}] == "." && grid[point{x, yy}] != "v" {
						new[point{x, yy}] = "v"
						moved = true
					} else {
						new[point{x, y}] = "v"
					}
				}
			}
		}
		grid = new
		if !moved {
			fmt.Println(step)
			break
		}
		step++
	}
}
