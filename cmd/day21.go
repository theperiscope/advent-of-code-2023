//go:build ignore

package main

import (
	"AOC/lib"
	"fmt"
)

func main() {
	lib.AssertArgs()
	g := grid{
		grid: lib.AssertInputByteGrid(),
	}

	start := lib.Point{X: -1, Y: -1}
	for r := 0; r < len(g.grid); r++ {
		for c := 0; c < len(g.grid[0]); c++ {
			if g.grid[r][c] == 'S' {
				start = lib.Point{X: c, Y: r}
				break
			}
		}
		if start.X != -1 {
			break
		}
	}

	g.start = map[lib.Point]bool{}
	g.start[start] = true

	//g.print()
	for i := 0; i < 64; i++ {
		g.step()
		//g.print()
	}

	fmt.Println("Part 1", len(g.start))
}

type grid struct {
	grid    [][]byte
	start   map[lib.Point]bool
	visited map[lib.Point]bool
}

func (g *grid) print() {
	for r := 0; r < len(g.grid); r++ {
		for c := 0; c < len(g.grid[0]); c++ {
			if _, ok := g.start[lib.Point{X: c, Y: r}]; ok {
				fmt.Printf("O")
			} else {
				fmt.Printf("%c", g.grid[r][c])
			}
		}
		fmt.Println()
	}
	fmt.Println("Next start points", len(g.start))
	fmt.Println("======")
}

func (g *grid) step() {
	nextStartPoints := map[lib.Point]bool{}
	g.visited = map[lib.Point]bool{}

	for p, _ := range g.start {
		possibleNextPoints := []lib.Point{
			{X: p.X - 1, Y: p.Y},
			{X: p.X + 1, Y: p.Y},
			{X: p.X, Y: p.Y - 1},
			{X: p.X, Y: p.Y + 1},
		}

		for _, n := range possibleNextPoints {
			if n.X < 0 || n.Y < 0 || n.X > int(len(g.grid[0])-1) || n.Y > int(len(g.grid)-1) { // out of bounds
				continue
			}
			if g.grid[n.Y][n.X] == '#' { // rock
				continue
			}
			if _, ok := g.visited[n]; ok { // visited before
				continue
			}
			g.visited[n] = true
			nextStartPoints[n] = true
		}
	}

	g.start = nextStartPoints
}
