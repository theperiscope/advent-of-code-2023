//go:build ignore

package main

import (
	"AOC/lib"
	"fmt"
)

func main() {
	lib.AssertArgs()
	rows := lib.AssertInput()

	patterns := [][]string{}
	p := []string{}
	for i := 0; i < len(rows); i++ {
		if rows[i] != "" {
			p = append(p, rows[i])
		}
		if rows[i] == "" || i == len(rows)-1 {
			patterns = append(patterns, p)
			p = []string{}
		}
	}

	ss := 0
	for i := 0; i < len(patterns); i++ {
		p := patterns[i]
		h, v := checkGridMirrored(p, 0)
		if h != -1 {
			ss += (h + 1) * 100
		}
		if v != -1 {
			ss += (v + 1)
		}
	}
	fmt.Println("Part 1:", ss)

	ss2 := 0
	for i := 0; i < len(patterns); i++ {
		p := patterns[i]
		h, v := checkGridMirrored(p, 1)
		if h != -1 {
			ss2 += (h + 1) * 100
		}
		if v != -1 {
			ss2 += v + 1
		}
	}
	fmt.Println("Part 2:", ss2)
}

func checkGridMirrored(rows []string, allowedErrors int) (h int, v int) {
	rowCount, columnCount := len(rows), len(rows[0])
	for y := 0; y < rowCount-1; y++ {
		errors := 0
		for i := 0; i < min(y+1, rowCount-y-1); i++ {
			start, end := y-i, y+i+1
			for j := 0; j < columnCount; j++ {
				if rows[start][j] != rows[end][j] {
					errors++
				}
			}
		}
		if allowedErrors == errors {
			return y, -1
		}
	}
	for x := 0; x < columnCount-1; x++ {
		errors := 0
		for i := 0; i < min(x+1, columnCount-x-1); i++ {
			start, end := x-i, x+i+1
			for i := 0; i < rowCount; i++ {
				if rows[i][start] != rows[i][end] {
					errors++
				}
			}
		}
		if allowedErrors == errors {
			return -1, x
		}
	}
	return -1, -1
}
