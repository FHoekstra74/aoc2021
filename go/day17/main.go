package main

import (
	"fmt"
)

func main() {
	a, b := 0, 0
	for i := -500; i < 500; i++ {
		for j := -500; j < 500; j++ {
			top, hit := check(i, j)
			if hit {
				b += 1
				if top > a {
					a = top
				}
			}
		}
	}
	fmt.Println(a)
	fmt.Println(b)
}

func check(xlouch, ylounch int) (int, bool) {
	xmintaget := 201
	xmaxtarget := 230
	ymintarget := -99
	ymaxtarget := -65
	x, y, ymax := 0, 0, 0
	for {
		x += xlouch
		y += ylounch
		if xlouch < 0 {
			xlouch += 1
		} else if xlouch > 0 {
			xlouch -= 1
		}
		ylounch -= 1
		if y > ymax {
			ymax = y
		}
		if x >= xmintaget && x <= xmaxtarget && y >= ymintarget && y <= ymaxtarget {
			return ymax, true
		}
		if x > xmaxtarget || y < ymintarget {
			return 0, false
		}
	}
}
