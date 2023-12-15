//go:build ignore

package main

import (
	"AOC/lib"
	"fmt"
	"strconv"
	"strings"
)

func count(lava string, damagedSpringGroups []int) int {
	pos := 0
	currentState := map[[4]int]int{{0, 0, 0, 0}: 1}
	nextState := map[[4]int]int{}
	for len(currentState) > 0 {
		for state, num := range currentState {
			lavaIndex, damagedSpringGroupIndex, currentDamagedSpringGroupCount, expectOperational := state[0], state[1], state[2], state[3]
			if lavaIndex == len(lava) {
				if damagedSpringGroupIndex == len(damagedSpringGroups) {
					pos += num
				}
				continue
			}
			switch {
			case (lava[lavaIndex] == '#' || lava[lavaIndex] == '?') && damagedSpringGroupIndex < len(damagedSpringGroups) && expectOperational == 0:
				// we are still searching for broken springs
				if lava[lavaIndex] == '?' && currentDamagedSpringGroupCount == 0 {
					// not in a group of broken springs, so ? can be a .
					nextState[[4]int{lavaIndex + 1, damagedSpringGroupIndex, currentDamagedSpringGroupCount, expectOperational}] += num
				}
				currentDamagedSpringGroupCount++
				if currentDamagedSpringGroupCount == damagedSpringGroups[damagedSpringGroupIndex] {
					// we filled the damaged spring group
					damagedSpringGroupIndex++
					currentDamagedSpringGroupCount = 0
					expectOperational = 1 // we expect only working spring next
				}
				nextState[[4]int{lavaIndex + 1, damagedSpringGroupIndex, currentDamagedSpringGroupCount, expectOperational}] += num
			case (lava[lavaIndex] == '.' || lava[lavaIndex] == '?') && currentDamagedSpringGroupCount == 0:
				// not in a group of broken springs
				expectOperational = 0
				nextState[[4]int{lavaIndex + 1, damagedSpringGroupIndex, currentDamagedSpringGroupCount, expectOperational}] += num
			}
		}
		currentState, nextState = nextState, currentState
		nextState = map[[4]int]int{}
	}
	return pos
}

func main() {
	lib.AssertArgs()
	rows := lib.AssertInput()

	c := 0
	for _, row := range rows {
		before, after, _ := strings.Cut(row, " ")
		damagedSprings, _ := lib.Convert(strings.Split(after, ","), strconv.Atoi)
		c += count(before, damagedSprings)
	}
	fmt.Println(c)

	c = 0
	for _, row := range rows {
		before, after, _ := strings.Cut(row, " ")
		beforeNew, afterNew := "", ""
		for i := 0; i < 5; i++ {
			beforeNew, afterNew = beforeNew+before+"?", afterNew+after+","
		}
		before, after = strings.TrimSuffix(beforeNew, "?"), strings.TrimSuffix(afterNew, ",")
		damagedSprings, _ := lib.Convert(strings.Split(after, ","), strconv.Atoi)
		c += count(before, damagedSprings)
	}
	fmt.Println(c)
}
