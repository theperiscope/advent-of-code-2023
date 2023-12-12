package main

import (
	"AOC/lib"
	"fmt"
	"strings"
)

func main() {
	lib.AssertArgs()
	rows := lib.AssertInput()

	var g [][]string
	for _, r := range rows {
		r = strings.Replace(r, ".", ".", -1)
		r = strings.Replace(r, "|", "│", -1)
		r = strings.Replace(r, "-", "─", -1)
		r = strings.Replace(r, "L", "└", -1)
		r = strings.Replace(r, "J", "┘", -1)
		r = strings.Replace(r, "7", "┐", -1)
		r = strings.Replace(r, "F", "┌", -1)

		fmt.Println(r)
		g = append(g, strings.Split(strings.Trim(r, " "), ""))
	}

	maxX, maxY := len(rows[0])-1, len(rows)-1
	//sx, sy, x, y, nn := 12, 4, 12, 4, 0
	sx, sy, x, y, nn := 25, 32, 25, 32, 0 // manually entered "S" position to avoid having to calculate missing piece
	loop := map[point]bool{}
	visited := map[point]bool{}
	for {
		visited[point{X: x, Y: y}] = true
		piece := g[y][x]
		nn++
		n, s, e, w := "", "", "", ""
		neighbors := []point{}
		switch piece {
		case "│":
			if y > 0 {
				neighbors = append(neighbors, point{X: 0, Y: -1})
				n = g[y-1][x]
			}
			if y <= maxY-1 {
				neighbors = append(neighbors, point{X: 0, Y: 1})
				s = g[y+1][x]
			}
		case "─":
			if x > 0 {
				neighbors = append(neighbors, point{X: -1, Y: 0})
				w = g[y][x-1]
			}
			if x <= maxX-1 {
				neighbors = append(neighbors, point{X: 1, Y: 0})
				e = g[y][x+1]
			}
		case "└":
			if y > 0 {
				neighbors = append(neighbors, point{X: 0, Y: -1})
				n = g[y-1][x]
			}
			if x <= maxX-1 {
				neighbors = append(neighbors, point{X: 1, Y: 0})
				e = g[y][x+1]
			}
		case "┘":
			if y > 0 {
				neighbors = append(neighbors, point{X: 0, Y: -1})
				n = g[y-1][x]
			}
			if x > 0 {
				neighbors = append(neighbors, point{X: -1, Y: 0})
				w = g[y][x-1]
			}
		case "┌":
			if y <= maxY-1 {
				neighbors = append(neighbors, point{X: 0, Y: 1})
				s = g[y+1][x]
			}
			if x <= maxX-1 {
				neighbors = append(neighbors, point{X: 1, Y: 0})
				e = g[y][x+1]
			}
		case "┐":
			if y <= maxY-1 {
				neighbors = append(neighbors, point{X: 0, Y: 1})
				s = g[y+1][x]
			}
			if x > 0 {
				neighbors = append(neighbors, point{X: -1, Y: 0})
				w = g[y][x-1]
			}
		}

		isPartOfLoop := f(piece, n, s, e, w)
		//fmt.Println(x, y, isPartOfLoop)

		if isPartOfLoop {
			loop[point{X: x, Y: y}] = true
			if !visited[point{X: x + neighbors[0].X, Y: y + neighbors[0].Y}] {
				x += neighbors[0].X
				y += neighbors[0].Y
			} else {
				x += neighbors[1].X
				y += neighbors[1].Y
			}
			if sx == x && sy == y {
				break
			}
		} else {
			break
		}
	}
	fmt.Println(nn)
	//fmt.Println(loop)

	allEmpty := map[point]bool{}
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if !loop[point{X: x, Y: y}] {
				allEmpty[point{X: x, Y: y}] = true
				g[y][x] = "."
			}
		}
		fmt.Println(y, "\t", g[y])
	}

	pointsInside := 0
	for p, _ := range allEmpty {

		x := p.X
		if x == 0 || x == maxX || y == 0 || y == maxY {
			//fmt.Println(p, "outside for sure")
			continue
		}
		count := 0
		start, end := "", ""
		for i := x + 1; i <= maxX; i++ {
			curr := g[p.Y][i]
			if curr == "." || curr == "─" {
				continue
			} else if curr == "│" {
				count++
			} else {
				if start == "" {
					start = curr
				} else {
					end = curr
					// process
					if start+end == "└┘" { // don't count
						start, end = "", ""
						continue
					} else if start+end == "└┐" {
						count++
					} else if start+end == "┌┘" {
						count++
					} else if start+end == "┌┐" { // don't count
						count += 2
					}
					start, end = "", ""
				}
			}
		}
		if count%2 == 1 {
			//fmt.Println(p, "=>", count)
			pointsInside++
		}
	}

	fmt.Println("Points inside", pointsInside)
}

type point struct {
	X int
	Y int
}

func (p point) String() string {
	return fmt.Sprintf("[X=%d, Y=%d]", p.X, p.Y)
}

func f(piece string, nn, ss, ee, ww string) bool {
	m := map[string]string{
		"│": "NS",
		"─": "EW",
		"└": "NE",
		"┘": "NW",
		"┐": "SW",
		"┌": "SE",
	}
	n, s, e, w := m[nn], m[ss], m[ee], m[ww]
	switch piece {
	case "│":
		// North-South piece: part of loop if N neighbor has S piece and S neighbor has N piece
		return n != "" && (n[0] == 'S' || n[1] == 'S') && s != "" && (s[0] == 'N' || s[1] == 'N')
	case "─":
		return e != "" && (e[0] == 'W' || e[1] == 'W') && w != "" && (w[0] == 'E' || w[1] == 'E')
	case "└":
		return n != "" && (n[0] == 'S' || n[1] == 'S') && e != "" && (e[0] == 'W' || e[1] == 'W')
	case "┘":
		return n != "" && (n[0] == 'S' || n[1] == 'S') && w != "" && (w[0] == 'E' || w[1] == 'E')
	case "┌":
		return s != "" && (s[0] == 'N' || s[1] == 'N') && e != "" && (e[0] == 'W' || e[1] == 'W')
	case "┐":
		return s != "" && (s[0] == 'N' || s[1] == 'N') && w != "" && (w[0] == 'E' || w[1] == 'E')
	}
	return false
}
