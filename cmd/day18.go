//go:build ignore

package main

import (
	"AOC/lib"
	"fmt"
	"strconv"
)

func main() {
	lib.AssertArgs()
	rows := lib.AssertInput()
	part1(rows)
	part2(rows)
}

// calculate area of rectilinear figure by using shoelace algorithm https://www.youtube.com/watch?v=ATt3NFFxtzU
func shoelace(points []lib.Point) int {
	area := 0
	for i := 0; i < len(points)-1; i++ {
		j := i + 1
		area += (points[j].X + points[i].X) * (points[j].Y - points[i].Y)
	}
	return area / 2
}

func part1(rows []string) {
	c := 2
	p := lib.Point{X: 0, Y: 0}
	points := []lib.Point{p}
	for _, row := range rows {
		d := lib.Atoi(row[2 : len(row)-10])
		c += d
		switch row[0] {
		case 'U':
			p.Y -= d
		case 'D':
			p.Y += d
		case 'L':
			p.X -= d
		case 'R':
			p.X += d
		}
		points = append(points, p)
	}
	fmt.Println(shoelace(points) + c/2)
}

func part2(rows []string) {
	c := 2
	p := lib.Point{X: 0, Y: 0}
	points := []lib.Point{p}
	for _, row := range rows {
		hex := row[len(row)-7 : len(row)-2]
		d, _ := strconv.ParseInt(hex, 16, 64)
		c += int(d)
		switch row[len(row)-2] {
		case '3': // U
			p.Y -= int(d)
		case '1': // D
			p.Y += int(d)
		case '2': // L
			p.X -= int(d)
		case '0': // R
			p.X += int(d)
		}
		points = append(points, p)
	}
	fmt.Println(shoelace(points) + c/2)
}
