package main

import (
	"AOC/lib"
	"fmt"
	"math"
	"regexp"
	"strings"
)

func main() {
	lib.AssertArgs()
	rows := lib.AssertInput()

	sum := 0
	copies := make([]int, len(rows)) // for part 2
	for i := 0; i < len(copies); i++ {
		copies[i] = 1
	}
	for i, row := range rows {
		colonIndex, pipeIndex := strings.Index(row, ": "), strings.Index(row, " | ")

		// winning numbers are after the colon; trim spaces
		winningNumbers := strings.Trim(row[colonIndex+2:pipeIndex], " ")
		// trim out extra space in any single-digit numbers
		winningNumbers = strings.ReplaceAll(winningNumbers, "  ", " ")
		// turn it into a regex string
		winningNumbers = " " + strings.ReplaceAll(winningNumbers, " ", " | ") + " "

		// pad with space, and replace " " with "  " to match correctly regex without backtracking
		ownNumbers := " " + strings.Replace(strings.Trim(row[pipeIndex+3:], " ")+" ", " ", "  ", -1)

		r, _ := regexp.Compile(winningNumbers)
		matches := r.FindAllString(ownNumbers, -1)

		// next len(matches) get copies of current (i.e. i) card
		for j := 0; j < len(matches); j++ {
			copies[i+j+1] += copies[i]
		}

		if len(matches) > 0 {
			sum += int(math.Pow(2, float64(len(matches)-1)))
		}
	}
	fmt.Printf("Sum Part 1: %d\n", sum)

	sum2 := 0
	for _, v := range copies {
		sum2 += v
	}
	fmt.Printf("Sum Part 2: %d\n", sum2)
}
