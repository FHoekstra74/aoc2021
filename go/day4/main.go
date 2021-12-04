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
					nr, _ := strconv.Atoi(strings.Trim((input[i+1+counter][counter2*3 : (counter2*3)+2]), " "))
					bingocard.numbers[counter][counter2] = nr
					bingocard.remainer += nr
				}
			}
			for run, draw := range draws {
				for counter := 0; counter < 5; counter++ {
					for counter2 := 0; counter2 < 5; counter2++ {
						if bingocard.numbers[counter][counter2] == draw {
							bingocard.remainer -= draw
							bingocard.numbers[counter][counter2] = -1
						}
					}
				}
				bingo := false
				for counter := 0; counter < 5; counter++ {
					if bingocard.numbers[counter][0] == -1 && bingocard.numbers[counter][1] == -1 && bingocard.numbers[counter][2] == -1 && bingocard.numbers[counter][3] == -1 && bingocard.numbers[counter][4] == -1 {
						bingo = true
					}
					if bingocard.numbers[0][counter] == -1 && bingocard.numbers[1][counter] == -1 && bingocard.numbers[2][counter] == -1 && bingocard.numbers[3][counter] == -1 && bingocard.numbers[4][counter] == -1 {
						bingo = true
					}
				}
				if bingo {
					bingocard.runs = run
					bingocard.last = draw
					boards = append(boards, bingocard)
					break
				}
			}
		}
	}
	lowest, highest, resulta, resultb := len(draws)+1, 0, 0, 0
	for _, bingocard := range boards {
		if bingocard.runs < lowest {
			lowest = bingocard.runs
			resulta = bingocard.remainer * bingocard.last
		}
		if bingocard.runs > highest {
			highest = bingocard.runs
			resultb = bingocard.remainer * bingocard.last
		}
	}
	fmt.Println(resulta)
	fmt.Println(resultb)
}
