package main

import (
	"container/heap"
	"fmt"
	"strconv"
)

func main() {
	_, _ = input, testinput
	//	input = testinput
	grid := make([][]int, len(input[0]))
	for i := range grid {
		grid[i] = make([]int, len(input[0]))
	}
	for y, row := range input {
		for x, val := range row {
			val, _ := strconv.Atoi(string(val))
			grid[x][y] = val
		}
	}
	fmt.Println(calcshortest(grid))

	grid = make([][]int, len(input[0])*5)
	for i := range grid {
		grid[i] = make([]int, len(input[0])*5)
	}
	for y, row := range input {
		for x, val := range row {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					newx, newy := i*len(input[0])+y, j*len(input[0])+x
					val, _ := strconv.Atoi(string(val))
					grid[newx][newy] = val + (i + j)
					if grid[newx][newy] > 9 {
						grid[newx][newy] -= 9
					}
				}
			}
		}
	}
	fmt.Println(calcshortest(grid))
}

func calcshortest(grid [][]int) int {
	maxx, maxy := len(grid), len(grid[0])
	quickest := make([][]int, maxx)
	for i := range grid {
		quickest[i] = make([]int, maxy)
		for j := 0; j < maxy; j++ {
			quickest[i][j] = 9999999
		}
	}
	vectors := []point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	h := &pointheap{{coordinate: point{0, 0}, totalrisk: 0, route: make([]point, 1)}}
	for {
		p := heap.Pop(h).(pointplus)
		if p.coordinate.x == maxx-1 && p.coordinate.y == maxy-1 {
			//fmt.Println(p.route)
			return p.totalrisk
		}
		for _, v := range vectors {
			newx := p.coordinate.x + v.x
			newy := p.coordinate.y + v.y
			if newx < 0 || newx >= maxx || newy < 0 || newy >= maxy {
				continue
			}
			risk := p.totalrisk + grid[newx][newy]
			if risk >= quickest[newx][newy] {
				continue
			}
			quickest[newx][newy] = risk
			var newroute []point
			newroute = append(newroute, p.route...)
			heap.Push(h, pointplus{coordinate: point{newx, newy}, totalrisk: risk, route: append(newroute, point{newx, newy})})
		}
	}
}

type point struct {
	x, y int
}
type pointplus struct {
	coordinate point
	totalrisk  int
	route      []point //for debugging
}
type pointheap []pointplus

//https://pkg.go.dev/container/heap@go1.17.5
func (h pointheap) Len() int            { return len(h) }
func (h pointheap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h pointheap) Less(i, j int) bool  { return h[i].totalrisk < h[j].totalrisk }
func (h *pointheap) Push(x interface{}) { *h = append(*h, x.(pointplus)) }
func (h *pointheap) Pop() interface{} {
	item := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return item
}
