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
	rows := lib.AssertInput()
	fmt.Println(part1(rows[0]))
	fmt.Println(part2(rows[0]))
}

func part1(input string) int {
	steps := strings.Split(input, ",")
	initializationSequenceSum := 0
	for _, step := range steps {
		initializationSequenceSum += getBox(step)
	}
	return initializationSequenceSum
}

func getBox(label string) int {
	sum := 0
	for _, ch := range label {
		sum = (sum + int(ch)) * 17 % 256
	}
	return sum
}

type lens struct {
	label       string
	focalLength int
}

func part2(input string) int {
	steps := strings.Split(input, ",")
	boxes := [256][]lens{}
	for _, step := range steps {
		if step[len(step)-1] == '-' {
			label := step[:len(step)-1]
			box := getBox(label)
			index := slices.IndexFunc(boxes[box], func(l lens) bool { return l.label == label })
			if index != -1 {
				boxes[box] = lib.Remove(boxes[box], index)
			}
		} else {
			label := step[:len(step)-2]
			focalLength, _ := strconv.Atoi(string(step[len(step)-1]))
			box := getBox(label)
			index := slices.IndexFunc(boxes[box], func(l lens) bool { return l.label == label })
			if index == -1 {
				boxes[box] = append(boxes[box], lens{label, focalLength})
			} else {
				boxes[box][index].focalLength = focalLength
			}
		}
	}

	sum := 0
	for box, lenses := range boxes {
		for i, lens := range lenses {
			sum += (box + 1) * (i + 1) * lens.focalLength
		}
	}
	return sum
}
