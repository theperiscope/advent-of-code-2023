//go:build ignore

package main

import (
	"AOC/lib"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lib.AssertArgs()
	input := lib.AssertInput()
	games := parseInput(input)

	sum := 0
	sum2 := 0
	for _, g := range games {
		if g.isPossiblePart1() {
			sum += g.id
		}
		mr, mg, mb := g.minimumSetPart2()
		sum2 += mr * mg * mb
	}
	fmt.Println(sum)
	fmt.Println(sum2)
}

type game struct {
	id    int
	red   []int
	green []int
	blue  []int
}

func (g *game) isPossiblePart1() bool {
	if slices.Max(g.red) <= 12 && slices.Max(g.green) <= 13 && slices.Max(g.blue) <= 14 {
		return true
	}
	return false
}

func (g *game) minimumSetPart2() (int, int, int) {
	return slices.Max(g.red), slices.Max(g.green), slices.Max(g.blue)
}

func parseInput(rows []string) []game {
	result := []game{}
	for x := 0; x < len(rows); x++ {
		row := rows[x]
		g := game{id: x + 1, red: []int{}, green: []int{}, blue: []int{}}
		turns := strings.Split(row[strings.Index(row, ":")+2:], "; ")
		for _, turn := range turns {
			red, green, blue := 0, 0, 0
			for _, color := range strings.Split(turn, ", ") {
				if strings.Index(color, "red") > 0 {
					red, _ = strconv.Atoi(color[:len(color)-4])
				} else if strings.Index(color, "green") > 0 {
					green, _ = strconv.Atoi(color[:len(color)-6])
				} else if strings.Index(color, "blue") > 0 {
					blue, _ = strconv.Atoi(color[:len(color)-5])
				}
			}
			g.red = append(g.red, red)
			g.green = append(g.green, green)
			g.blue = append(g.blue, blue)
		}
		result = append(result, g)
	}
	return result
}
