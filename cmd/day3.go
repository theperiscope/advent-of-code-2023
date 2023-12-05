package main

import (
	"AOC/lib"
	"fmt"
)

func main() {
	lib.AssertArgs()
	rows := lib.AssertInput()

	numbers, symbols := []number{}, []symbol{}
	for index, row := range rows {
		currentNumber := 0
		currentNumberStartIndex, currentNumberEndIndex := -1, -1
		for pos, char := range row {
			if char >= '0' && char <= '9' {
				if currentNumberStartIndex == -1 {
					currentNumberStartIndex = pos
				}
				currentNumber = currentNumber*10 + int(char-'0')
			} else {
				currentNumberEndIndex = pos - 1
				if currentNumberStartIndex != -1 {
					//fmt.Printf("Row %d: %d from %d to %d\n", index, currentNumber, currentNumberStartIndex, currentNumberEndIndex)
					numbers = append(numbers, number{n: currentNumber, row: index, from: currentNumberStartIndex, to: currentNumberEndIndex})
				}
				currentNumber = 0
				currentNumberStartIndex, currentNumberEndIndex = -1, -1

				if char != '.' {
					//fmt.Printf("Row %d: %c at position %d\n", index, char, pos)
					symbols = append(symbols, symbol{s: char, row: index, pos: pos})
				}
			}
		}
		if currentNumberStartIndex != -1 {
			//fmt.Printf("Row %d: %d from %d to %d\n", index, currentNumber, currentNumberStartIndex, len(rows[0])-1)
			numbers = append(numbers, number{n: currentNumber, row: index, from: currentNumberStartIndex, to: len(rows[0]) - 1})
			currentNumber = 0
			currentNumberStartIndex, currentNumberEndIndex = -1, -1
		}
	}

	data := dataRows{}
	for _, n := range numbers {
		ns := data[n.row]
		ns.n = append(ns.n, n)
		data[n.row] = ns
	}
	for row, _ := range data {
		for _, s := range symbols {
			if !(lib.AbsDiff(row, s.row) <= 1) {
				continue
			}
			ns := data[row]
			ns.s = append(ns.s, s)
			data[row] = ns
		}
	}

	sum := 0
	for _, row := range data {
		for _, n := range row.n {
			for _, s := range row.s {
				if isValidEnginePart(n, s) {
					//fmt.Printf("Engine part: %d%c\n", n.n, s.s)
					sum += n.n
				}
			}
		}
	}

	fmt.Printf("Sum Part 1: %d\n", sum)

	sum2 := 0
	for _, s := range symbols {
		if s.s != '*' {
			continue
		}

		nn := []int{}
		for _, n := range numbers {
			if isValidEnginePart(n, s) {
				nn = append(nn, n.n)
			}
		}

		if len(nn) == 2 {
			sum2 += nn[0] * nn[1]
			//fmt.Printf("%d*%d = %d\n", nn[0], nn[1], nn[0]*nn[1])
		}
	}
	fmt.Printf("Sum Part 2: %d\n", sum2)
}

type number struct {
	n    int
	row  int
	from int
	to   int
}

type symbol struct {
	s   rune
	row int
	pos int
}

type numberSymbol struct {
	n []number
	s []symbol
}

type dataRows map[int]numberSymbol

func isValidEnginePart(n number, s symbol) bool {
	if !(lib.AbsDiff(n.row, s.row) <= 1) { // not on same or neighboring row
		return false
	}
	if !(s.pos >= n.from-1 && s.pos <= n.to+1) { // outside expected column index range
		return false
	}
	return true
}
