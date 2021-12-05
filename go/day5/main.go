package main

import (
	"fmt"
)

type point struct {
	x int
	y int
}

type overlap struct {
	straight int
	diagonal int
}

func main() {
	_, _ = input, testinput
	points := make(map[point]overlap)
	for _, l := range input {
		x1, y1, x2, y2 := 0, 0, 0, 0
		n, _ := fmt.Sscanf(l, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if n != 4 {
			break
		}
		dx, dy := x2-x1, y2-y1
		if dx < 1 {
			dx *= -1
		}
		if dy < 1 {
			dy *= -1
		}
		if dx == 0 || dy == 0 {
			for i := 0; i <= dx+dy; i++ {
				px, py := x1, y1
				if x1 < x2 {
					px += i
				} else if x1 > x2 {
					px -= i
				} else if y1 < y2 {
					py += i
				} else if y1 > y2 {
					py -= i
				}
				p := points[point{px, py}]
				p.straight += 1
				points[point{px, py}] = p
			}
		}
		if dx == dy {
			for i := 0; i <= dx; i++ {
				px, py := x1, y1
				if x1 > x2 && y1 < y2 {
					px -= i
					py += i
				} else if x1 > x2 && y1 > y2 {
					px -= i
					py -= i
				} else if x1 < x2 && y1 < y2 {
					px += i
					py += i
				} else if x1 < x2 && y1 > y2 {
					px += i
					py -= i
				}
				p := points[point{px, py}]
				p.diagonal += 1
				points[point{px, py}] = p
			}
		}
	}
	a, b := 0, 0
	for _, v := range points {
		if v.straight > 1 {
			a += 1
		}
		if v.straight+v.diagonal > 1 {
			b += 1
		}
	}
	fmt.Println(a)
	fmt.Println(b)
}
