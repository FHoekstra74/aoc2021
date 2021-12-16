package main

import (
	"fmt"
	"strconv"
)

func main() {
	_, _ = input, testinput
	//input = testinput
	hx := map[string]string{"0": "0000", "1": "0001", "2": "0010", "3": "0011", "4": "0100", "5": "0101", "6": "0110", "7": "0111", "8": "1000", "9": "1001", "A": "1010", "B": "1011", "C": "1100", "D": "1101", "E": "1110", "F": "1111"}
	bits := ""
	for _, v := range input[0] {
		bits = bits + hx[string(v)]
	}
	_, a, b := parse(bits)
	fmt.Println(a)
	fmt.Println(b)
}

func parse(data string) (dataout string, sumversion int64, result int64) {
	version, _ := strconv.ParseInt(data[:3], 2, 64)
	sumversion += version
	data = data[3:]
	typeid, _ := strconv.ParseInt(data[:3], 2, 64)
	data = data[3:]
	if typeid == 4 {
		literal := ""
		for {
			fivebits := data[:5]
			data = data[5:]
			literal += fivebits[1:]
			if string(fivebits[0]) == "0" {
				break
			}
		}
		result, _ = strconv.ParseInt(literal, 2, 64)
	} else {
		lengthtype := data[:1]
		data = data[1:]
		var results []int64

		if string(lengthtype[0]) == "0" {
			sub := data[:15]
			data = data[15:]
			length, _ := strconv.ParseInt(sub, 2, 64)
			sub = data[:length]
			data = data[length:]
			for {
				var v, res int64
				sub, v, res = parse(sub)
				results = append(results, res)
				sumversion += v
				if len(sub) == 0 {
					break
				}
			}
		} else {
			sub := data[:11]
			data = data[11:]
			nr, _ := strconv.ParseInt(sub, 2, 64)
			for i := 0; int64(i) < nr; i++ {
				var v, res int64
				data, v, res = parse(data)
				results = append(results, res)
				sumversion += v
			}
		}
		switch typeid {
		case 0:
			for _, v := range results {
				result += v
			}
		case 1:
			result = 1
			for _, v := range results {
				result *= v
			}
		case 2:
			result = results[0]
			for _, v := range results {
				if v < result {
					result = v
				}
			}
		case 3:
			result = results[0]
			for _, v := range results {
				if v > result {
					result = v
				}
			}
		case 5:
			if results[0] > results[1] {
				result = 1
			}
		case 6:
			if results[0] < results[1] {
				result = 1
			}
		case 7:
			if results[0] == results[1] {
				result = 1
			}
		}
	}
	return data, sumversion, result
}
