package main

import (
	"fmt"
	"strconv"
)

type point struct {
	x, y int
}

func main() {
	_, _ = input, testinput
	algorithm := input[0]
	startimage := make(map[point]string)
	for y, row := range input {
		if y > 1 {
			for x, v := range row {
				startimage[point{x, y - 2}] = string(v)
			}
		}
	}
	maxx, maxy, minx, miny := 0, 0, 99999, 99999
	for i := 0; i < 50; i++ {
		new := make(map[point]string)
		for k := range startimage {
			if k.x < minx {
				minx = k.x
			} else if k.x > maxx {
				maxx = k.x
			}
			if k.y < miny {
				miny = k.y
			} else if k.y > maxy {
				maxy = k.y
			}
		}
		for x := minx - 1; x <= maxx+1; x++ {
			for y := miny - 1; y <= maxy+1; y++ {
				arg := ""
				for _, yy := range []int{y - 1, y, y + 1} {
					for _, xx := range []int{x - 1, x, x + 1} {
						v, ok := startimage[point{xx, yy}]
						if ok {
							arg += string(v)
						} else {
							if i%2 == 0 {
								arg += "."
							} else {
								arg += "#"
							}
						}
					}
				}
				binnr := ""
				for _, v := range arg {
					if v == '#' {
						binnr += "1"
					} else {
						binnr += "0"
					}
				}
				num, _ := strconv.ParseInt(binnr, 2, 64)
				new[point{x, y}] = string(algorithm[num])
			}
		}
		startimage = new
		if i == 1 || i == 49 {
			res := 0
			for _, v := range startimage {
				if v == "#" {
					res += 1
				}
			}
			fmt.Println(res)
		}
	}
}

func printgrid(image map[point]string, minx int, maxx int, miny int, maxy int) {
	for y := miny; y <= maxy; y++ {
		line := ""
		for x := minx; x <= maxx; x++ {
			line += string(image[point{x, y}])
		}
		fmt.Println(line)
	}
}
