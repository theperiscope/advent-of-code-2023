package main

import (
	"AOC/lib"
	"fmt"
	"strings"
)

func main() {
	lib.AssertArgs()
	rows := lib.AssertInput()

	instructions := strings.Split(rows[0], "")
	nodes := map[string][]string{}

	for _, row := range rows[2:] {
		key, left, right := row[0:3], row[7:10], row[12:15]
		nodes[key] = []string{left, right}
	}

	// part 1
	steps := solveDay8("AAA", instructions, nodes)
	fmt.Println(steps)

	// part 2
	As := lib.Filter(lib.Keys(nodes), func(s string) bool {
		return s[2] == 'A'
	})

	stepCounts := []int{}
	for _, a := range As {
		s := solveDay8Part2(a, instructions, nodes)
		stepCounts = append(stepCounts, s)
	}
	fmt.Println(lib.LCM(stepCounts[0], stepCounts[1], stepCounts[2:]...))
}

func solveDay8(currentNode string, instructions []string, nodes map[string][]string) int {
	steps := 0
	instructionPointer := -1
	for {
		steps++
		instructionPointer++
		instruction := instructions[instructionPointer%len(instructions)]

		switch instruction {
		case "L":
			currentNode = nodes[currentNode][0]
		case "R":
			currentNode = nodes[currentNode][1]
		default:
			panic("ohno")
		}

		if currentNode == "ZZZ" {
			break
		}
	}

	return steps
}

func solveDay8Part2(currentNode string, instructions []string, nodes map[string][]string) int {
	steps := 0
	instructionPointer := -1
	for {
		steps++
		instructionPointer++
		instruction := instructions[instructionPointer%len(instructions)]

		switch instruction {
		case "L":
			currentNode = nodes[currentNode][0]
		case "R":
			currentNode = nodes[currentNode][1]
		default:
			panic("ohno")
		}

		if currentNode[2] == 'Z' {
			break
		}
	}

	return steps
}
