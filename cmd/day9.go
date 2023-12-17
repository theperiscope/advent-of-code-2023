//go:build ignore

package main

import (
	"AOC/lib"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lib.AssertArgs()
	rows := lib.AssertInput()

	sum := 0
	for _, row := range rows {
		v, _ := lib.Convert[string, int](strings.Split(row, " "), strconv.Atoi)
		next := v[len(v)-1] + solveDay9(v, 0)
		sum += next
	}
	fmt.Println("Sum", sum)

	sum2 := 0
	for _, row := range rows {
		v, _ := lib.Convert[string, int](strings.Split(row, " "), strconv.Atoi)
		x := solveDay9Part2(v, 0)
		next := v[0] - x
		sum2 += next
	}
	fmt.Println("Sum Part 2", sum2)
}

func solveDay9(numbers []int, level int) int {
	allZero := true
	result := []int{}
	for i := 1; i < len(numbers); i++ {
		a, b := numbers[i-1], numbers[i]
		result = append(result, b-a)
		if b-a != 0 {
			allZero = false
		}
	}

	if allZero {
		return 0
	}

	v1 := solveDay9(result, level+1)
	return v1 + result[len(result)-1]
}

func solveDay9Part2(numbers []int, level int) int {
	allZero := true
	result := []int{}
	for i := 1; i < len(numbers); i++ {
		a, b := numbers[i-1], numbers[i]
		result = append(result, b-a)
		if b-a != 0 {
			allZero = false
		}
	}

	if allZero {
		return 0
	}

	v1 := solveDay9Part2(result, level+1)
	return result[0] - v1
}
