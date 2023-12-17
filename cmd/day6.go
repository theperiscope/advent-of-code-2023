//go:build ignore

package main

import (
	"AOC/lib"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lib.AssertArgs()
	rows := lib.AssertInput()

	// part 1
	time, distance := part1(rows)
	fmt.Println(time)
	fmt.Println(distance)

	solve(time, distance)

	// part 2
	time, distance = part2(rows)
	fmt.Println(time)
	fmt.Println(distance)

	solve(time, distance)
}

func part1(rows []string) ([]int, []int) {
	digits, _ := regexp.Compile("\\d+")
	time, _ := lib.Convert[string, int](digits.FindAllString(rows[0], -1), strconv.Atoi)
	distance, _ := lib.Convert[string, int](digits.FindAllString(rows[1], -1), strconv.Atoi)
	return time, distance
}

func part2(rows []string) ([]int, []int) {
	digits, _ := regexp.Compile("\\d+")
	time, _ := lib.Convert[string, int]([]string{strings.Join(digits.FindAllString(rows[0], -1), "")}, strconv.Atoi)
	distance, _ := lib.Convert[string, int]([]string{strings.Join(digits.FindAllString(rows[1], -1), "")}, strconv.Atoi)
	return time, distance
}

func solve(time, distance []int) int {
	// scoringTime = raceTime-holdTime
	// scoringTime*holdTime = distanceToBeat
	// (raceTime-holdTime)*holdTime = distanceToBeat
	// -holdTime^2 + raceTime*holdTime - distanceToBeat = 0
	// quadratic to solve for hold time: A=-1, B=raceTime, C=-distanceToBeat

	mult := 1
	for i := 0; i < len(time); i++ {
		raceTime := float64(time[i])
		distanceToBeat := float64(distance[i]) + 1
		fmt.Println("raceTime=", raceTime, "distanceToBeat=", distanceToBeat)

		b1, b2, _ := quadratic(-1, raceTime, -distanceToBeat)
		fmt.Println(b1, b2) // b2 is always bigger
		b := math.Ceil(math.Min(b1, b2))
		possibleSolutions := raceTime + 1 - 2*b // another way to calculate/verify

		// x^2-7x-9=0
		// x^2-15x-40=0
		// x1+x2 = -b/a ... for a=1,b=-7  x1=8.11 x2=-1.11
		// x1x2 = c/a
		// also, the 1/2*(x1+x2) = -b/2*a (root average)
		// (x−x1)(x−x2)=x^2−(x1+x2)+x1x2
		// |x1-x2| = sqrt(b^2-4*a*c)/a
		// means function mininum is for x=3.5

		fmt.Println(math.Floor(b2) - math.Ceil(b1) + 1)
		fmt.Println(possibleSolutions)
		fmt.Println()
		mult *= int(possibleSolutions)
	}
	fmt.Println(mult)
	return mult
}

// B^2 - T*B - D = 0
func quadratic(a, b, c float64) (float64, float64, bool) {
	d := b*b - 4*a*c
	if d > 0 {
		root1 := (-b + math.Sqrt(d)) / (2 * a)
		root2 := (-b - math.Sqrt(d)) / (2 * a)
		return root1, root2, true
	} else if d == 0 {
		root1 := (-b + math.Sqrt(d)) / (2 * a)
		return root1, root1, true
	} else {
		realPart := -b / (2 * a)
		imaginaryPart := math.Sqrt(math.Abs(d)) / (2 * a)
		return realPart, imaginaryPart, false
	}
}
