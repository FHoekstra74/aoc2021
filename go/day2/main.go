package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	_, _ = input, testinput
	pos, deptha, depthb := 0, 0, 0
	for _, v := range input {
		command := strings.Split(v, " ")
		val, _ := strconv.Atoi(command[1])
		switch command[0] {
		case "forward":
			pos += val
			depthb += (val * deptha)
		case "up":
			deptha -= val
		case "down":
			deptha += val
		}
	}
	fmt.Println(deptha * pos)
	fmt.Println(depthb * pos)
}
