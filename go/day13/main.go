package main

import (
	"fmt"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type foldinstruction struct {
	line string
	pos  int
}

func main() {
	_, _ = input, testinput
	var grid = []point{}
	var folds = []foldinstruction{}
	for _, j := range input {
		if len(j) > 0 {
			if strings.HasPrefix(j, "fold") {
				instr := strings.Split(j, " ")[2]
				l := strings.Split(instr, "=")[0]
				pos, _ := strconv.Atoi(strings.Split(instr, "=")[1])
				folds = append(folds, foldinstruction{l, pos})
			} else {
				x, _ := strconv.Atoi(strings.Split(j, ",")[0])
				y, _ := strconv.Atoi(strings.Split(j, ",")[1])
				grid = append(grid, point{x, y})
			}
		}
	}
	for i, j := range folds {
		grid = dofold(grid, j)
		if i == 0 {
			fmt.Println("A: ", len(grid))
		}
	}
	print(grid)
}

func dofold(grid []point, instruction foldinstruction) []point {
	var res = []point{}
	for _, j := range grid {
		newx, newy := j.x, j.y
		if instruction.line == "y" && j.y > instruction.pos {
			newy = instruction.pos - (j.y - instruction.pos)
		} else if instruction.line == "x" && j.x > instruction.pos {
			newx = instruction.pos - (j.x - instruction.pos)
		}
		if !exists(res, newx, newy) {
			res = append(res, point{newx, newy})
		}
	}
	return res
}

func exists(grid []point, x int, y int) bool {
	for _, i := range grid {
		if i.x == x && i.y == y {
			return true
		}
	}
	return false
}

func print(grid []point) {
	maxx, maxy := 0, 0
	for _, i := range grid {
		if i.x > maxx {
			maxx = i.x
		}
		if i.y > maxy {
			maxy = i.y
		}
	}
	for y := 0; y <= maxy; y++ {
		line := ""
		for x := 0; x <= maxx; x++ {
			if exists(grid, x, y) {
				line += "#"
			} else {
				line += " "
			}
		}
		fmt.Println(line)
	}
}
