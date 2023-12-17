//go:build ignore

package main

import (
	"AOC/lib"
	"bytes"
	"crypto/sha256"
	"fmt"
)

func main() {
	lib.AssertArgs()
	grid := lib.AssertInputByteGrid()
	dish := Dish(grid)
	dish.TiltNorth()
	fmt.Println(dish.NorthSupportBeamsLoad())

	hashMap := map[string][]int{} // values are loop indexes when particular hash occurred (last only used)
	dish2 := Dish(grid)
	N := 1_000_000_000
	for i := 0; i < N; i++ {
		dish2.TiltNorth()
		dish2.TiltWest()
		dish2.TiltSouth()
		dish2.TiltEast()
		sha := sha256.New()
		sha.Write(bytes.Join(dish2, []byte{}))
		h := fmt.Sprintf("%x", sha.Sum(nil))
		if _, ok := hashMap[h]; ok {
			remaining := N - i
			repeatCycle := i - hashMap[h][len(hashMap[h])-1]
			i = N - remaining%repeatCycle
		}
		hashMap[h] = append(hashMap[h], i)
	}
	/*
		for k, v := range hashMap {
			fmt.Printf("%s %d\n", string(k), v)
		}
	*/
	fmt.Println(dish2.NorthSupportBeamsLoad())
}

type Dish [][]byte

func (d Dish) TiltNorth() {
	rows, columns := len(d), len(d[0])
	for column := 0; column < columns; column++ {
		firstAvailableRow := 0
		for row := 0; row < rows; row++ {
			switch d[row][column] {
			case '#':
				firstAvailableRow = row + 1
			case 'O':
				if firstAvailableRow < row {
					d[firstAvailableRow][column] = 'O'
					d[row][column] = '.'
				}
				firstAvailableRow++
			}
		}
	}
}

func (d Dish) TiltWest() {
	rows, columns := len(d), len(d[0])
	for row := 0; row < rows; row++ {
		firstAvailableColumn := 0
		for column := 0; column < columns; column++ {
			switch d[row][column] {
			case '#':
				firstAvailableColumn = column + 1
			case 'O':
				if firstAvailableColumn < column {
					d[row][firstAvailableColumn] = 'O'
					d[row][column] = '.'
				}
				firstAvailableColumn++
			}
		}
	}
}

func (d Dish) TiltSouth() {
	rows, columns := len(d), len(d[0])
	for column := 0; column < columns; column++ {
		firstAvailableRow := rows - 1
		for row := rows - 1; row >= 0; row-- {
			switch d[row][column] {
			case '#':
				firstAvailableRow = row - 1
			case 'O':
				if firstAvailableRow > row {
					d[firstAvailableRow][column] = 'O'
					d[row][column] = '.'
				}
				firstAvailableRow--
			}
		}
	}
}

func (d Dish) TiltEast() {
	rows, columns := len(d), len(d[0])
	for row := 0; row < rows; row++ {
		firstAvailableColumn := columns - 1
		for column := columns - 1; column >= 0; column-- {
			switch d[row][column] {
			case '#':
				firstAvailableColumn = column - 1
			case 'O':
				if firstAvailableColumn > column {
					d[row][firstAvailableColumn] = 'O'
					d[row][column] = '.'
				}
				firstAvailableColumn--
			}
		}
	}
}

func (d Dish) NorthSupportBeamsLoad() int {
	sum, rows, columns := 0, len(d), len(d[0])
	for column := 0; column < columns; column++ {
		for row := 0; row < rows; row++ {
			if d[row][column] == 'O' {
				sum += rows - row
			}
		}
	}
	return sum
}

func (d Dish) Print() {
	rows, columns := len(d), len(d[0])
	for row := 0; row < rows; row++ {
		for x := 0; x < columns; x++ {
			fmt.Printf("%c", d[row][x])
		}
		fmt.Println()
	}
}
