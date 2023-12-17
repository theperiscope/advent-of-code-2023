//go:build ignore

package main

import (
	"AOC/lib"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	lib.AssertArgs()
	rows := lib.AssertInput()

	r := regexp.MustCompile("\\#")
	galaxies := []lib.Point{}

	verticals := []int{}
	for columnIndex := 0; columnIndex < len(rows[0]); columnIndex++ {
		if isEmptyColumn(columnIndex, rows) {
			verticals = append(verticals, columnIndex)
		}
	}
	//fmt.Println(verticals)

	horizontals := []int{}
	for rowIndex, row := range rows {
		if isEmptyRow(rowIndex, rows) {
			horizontals = append(horizontals, rowIndex)
		} else {
			g := r.FindAllIndex([]byte(row), -1)
			for i := 0; i < len(g); i++ {
				galaxies = append(galaxies, lib.Point{X: g[i][0], Y: rowIndex})
			}
		}
	}
	//fmt.Println(horizontals)
	//fmt.Println(galaxies)

	sum := 0
	for i, from := range galaxies {
		for j, to := range galaxies {
			if i >= j {
				continue
			}
			nh := len(lib.Filter(horizontals, func(n int) bool { return n >= min(from.Y, to.Y) && n <= max(from.Y, to.Y) }))
			nv := len(lib.Filter(verticals, func(n int) bool { return n >= min(from.X, to.X) && n <= max(from.X, to.X) }))
			sum += from.ManhattanDistance(to) + nh + nv
		}
	}
	fmt.Println(sum)

	sum2 := 0
	for i, from := range galaxies {
		for j, to := range galaxies {
			if i >= j {
				continue
			}
			nh := len(lib.Filter(horizontals, func(n int) bool { return n >= min(from.Y, to.Y) && n <= max(from.Y, to.Y) }))
			nv := len(lib.Filter(verticals, func(n int) bool { return n >= min(from.X, to.X) && n <= max(from.X, to.X) }))
			sum2 += from.ManhattanDistance(to) + (nh * 999999) + (nv * 999999)
		}
	}
	fmt.Println(sum2)
}

func isEmptyRow(row int, rows []string) bool {
	return strings.Index(rows[row], "#") < 0
}

func isEmptyColumn(column int, rows []string) bool {
	for i := 0; i < len(rows); i++ {
		if rows[i][column] == '#' {
			return false
		}
	}

	return true
}
