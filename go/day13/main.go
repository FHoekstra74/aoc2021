package main

import (
	"fmt"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func main() {
	_, _ = input, testinput
	grid, xsize, ysize, a := make(map[point]bool), 1000, 1000, 0
	for _, j := range input {
		if len(j) > 0 && !strings.HasPrefix(j, "fold") {
			x, _ := strconv.Atoi(strings.Split(j, ",")[0])
			y, _ := strconv.Atoi(strings.Split(j, ",")[1])
			grid[point{x, y}] = true
		} else if len(j) > 0 {
			direction := strings.Split(strings.Split(j, " ")[2], "=")[0]
			pos, _ := strconv.Atoi(strings.Split(strings.Split(j, " ")[2], "=")[1])
			newgrid := make(map[point]bool)
			for j := range grid {
				newx, newy := j.x, j.y
				if direction == "y" && j.y > pos {
					newy = pos - (j.y - pos)
				} else if direction == "x" && j.x > pos {
					newx = pos - (j.x - pos)
				}
				newgrid[point{newx, newy}] = true
			}
			grid = newgrid
			if a == 0 {
				a = len(grid)
			}
			if direction == "x" {
				xsize = pos
			} else {
				ysize = pos
			}
		}
	}
	fmt.Println("A: ", a)
	for y := 0; y < ysize; y++ {
		line := ""
		for x := 0; x < xsize; x++ {
			_, exists := grid[point{x, y}]
			if exists {
				line += "#"
			} else {
				line += " "
			}
		}
		fmt.Println(line)
	}
}
