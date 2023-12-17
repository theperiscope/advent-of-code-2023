//go:build ignore

package main

import (
	"AOC/lib"
	"fmt"
)

func main() {
	lib.AssertArgs()
	grid := TileGrid(lib.AssertInputByteGrid())
	fmt.Println(grid.TraceBeam(position{x: 0, y: 0, d: lib.RIGHT}))

	possibleStarts := []position{}
	for y := 0; y <= len(grid)-1; y++ {
		possibleStarts = append(possibleStarts, position{x: 0, y: int8(y), d: lib.RIGHT})
		possibleStarts = append(possibleStarts, position{x: int8(len(grid[0]) - 1), y: int8(y), d: lib.LEFT})
	}
	for x := 0; x <= len(grid[0])-1; x++ {
		possibleStarts = append(possibleStarts, position{x: int8(x), y: 0, d: lib.DOWN})
		possibleStarts = append(possibleStarts, position{x: int8(x), y: int8(len(grid) - 1), d: lib.UP})
	}

	maxE := 0
	maxP := position{x: -1, y: -1, d: lib.UP}
	for _, start := range possibleStarts {
		e := grid.TraceBeam(start)
		if e > maxE {
			maxE = e
			maxP = start
		}
	}
	fmt.Println(maxP, maxE)
}

type position struct {
	x, y int8
	d    lib.OrthogonalDirection
}
type TileGrid [][]byte
type VisitedTileGrid [][][4]bool // visit map: each tile can be visited from 4 directions

func (tileGrid TileGrid) TraceBeam(p position) int {
	startX, startY, d := p.x, p.y, p.d
	v := make(VisitedTileGrid, len(tileGrid))
	for y := range v {
		v[y] = make([][4]bool, len(tileGrid[y]))
	}
	return tileGrid.traceBeam(startX, startY, d, v)
}

func (tileGrid TileGrid) traceBeam(x, y int8, d lib.OrthogonalDirection, v VisitedTileGrid) (count int) { // int8 for less memory usage during recursive calls
	if x < 0 || y < 0 || y >= int8(len(tileGrid)) || x >= int8(len(tileGrid[y])) { // out of bounds
		return
	}
	if v[y][x][d] { // already visited
		return
	}

	if !v[y][x][0] && !v[y][x][1] && !v[y][x][2] && !v[y][x][3] { // never energized in any direction
		count++
	}
	v[y][x][d] = true

	switch tileGrid[y][x] {
	case '-':
		if d == lib.UP || d == lib.DOWN {
			count += tileGrid.traceBeam(x-1, y, lib.LEFT, v) + tileGrid.traceBeam(x+1, y, lib.RIGHT, v)
			return
		}
	case '|':
		if d == lib.LEFT || d == lib.RIGHT {
			count += tileGrid.traceBeam(x, y-1, lib.UP, v) + tileGrid.traceBeam(x, y+1, lib.DOWN, v)
			return
		}
	case '\\':
		switch d {
		case lib.UP:
			d = lib.LEFT
		case lib.RIGHT:
			d = lib.DOWN
		case lib.DOWN:
			d = lib.RIGHT
		case lib.LEFT:
			d = lib.UP
		}
	case '/':
		switch d {
		case lib.UP:
			d = lib.RIGHT
		case lib.RIGHT:
			d = lib.UP
		case lib.DOWN:
			d = lib.LEFT
		case lib.LEFT:
			d = lib.DOWN
		}
	}

	switch d {
	case lib.UP:
		y--
	case lib.RIGHT:
		x++
	case lib.DOWN:
		y++
	case lib.LEFT:
		x--
	}

	return count + tileGrid.traceBeam(x, y, d, v)
}
