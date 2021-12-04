package main

import (
	"fmt"
	"strconv"
	"strings"
)

type board struct {
	numbers  [5][5]int
	runs     int
	last     int
	remainer int
}

func (b *board) hasbingo() bool {
	for counter := 0; counter < 5; counter++ {
		xok, yok := true, true
		for counter2 := 0; counter2 < 5; counter2++ {
			if b.numbers[counter][counter2] > -1 {
				xok = false
			}
			if b.numbers[counter2][counter] > -1 {
				yok = false
			}
		}
		if xok || yok {
			return true
		}
	}
	return false
}

func (b *board) play(draw int) (bingo bool) {
	b.runs += 1
	b.last = draw
	for counter := 0; counter < 5; counter++ {
		for counter2 := 0; counter2 < 5; counter2++ {
			if b.numbers[counter][counter2] == draw {
				b.remainer -= draw
				b.numbers[counter][counter2] = -1
			}
		}
	}
	return b.hasbingo()
}

func main() {
	_, _ = input, testinput
	var boards = []board{}
	var draws = []int{}
	for i, v := range input {
		if i == 0 {
			for _, j := range strings.Split(v, ",") {
				j, _ := strconv.Atoi(j)
				draws = append(draws, j)
			}
		}
		if len(v) == 0 {
			bingocard := board{}
			for counter := 0; counter < 5; counter++ {
				for counter2 := 0; counter2 < 5; counter2++ {
					nr, _ := strconv.Atoi(strings.Trim(input[i+1+counter][counter2*3:(counter2*3)+2], " "))
					bingocard.numbers[counter][counter2] = nr
					bingocard.remainer += nr
				}
			}
			for _, draw := range draws {
				if bingocard.play(draw) {
					boards = append(boards, bingocard)
					break
				}
			}
		}
	}
	a, b := boards[0], boards[0]
	for _, bingocard := range boards {
		if bingocard.runs < a.runs {
			a = bingocard
		}
		if bingocard.runs > b.runs {
			b = bingocard
		}
	}
	fmt.Println(a.remainer * a.last)
	fmt.Println(b.remainer * b.last)
}
