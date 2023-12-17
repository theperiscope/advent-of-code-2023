//go:build ignore

package main

import (
	"AOC/lib"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	lib.AssertArgs()
	lines := lib.AssertInput()

	fmt.Printf("Lines: %d\n", len(lines))

	part1Sum := solve(lines, getDigitPart1)
	fmt.Println("Sum:  ", part1Sum)

	part2Sum := solve(lines, getDigitPart2)
	fmt.Println("Sum:  ", part2Sum)
}

func solve(lines []string, getDigit func(string, int) int) int32 {
	start := time.Now()
	// Then we calculate the execution time duration
	// inside a deferred function, since it will be execute after f() returns
	defer func(start time.Time) {
		dur := time.Since(start)
		fmt.Printf("Solve: %s\n", dur)
	}(start)

	var wg sync.WaitGroup
	var sum int32

	for i, line := range lines {
		wg.Add(1)

		go func(i int, s string) {
			defer wg.Done()

			// processing
			first := -1
			last := -1
			// start scanning ---> and <--- simultaneously until we hit the first digit
			for a, b := 0, len(s)-1; a <= len(s)-1 && b >= 0; a, b = a+1, b-1 {
				if first == -1 {
					first = getDigit(s, a)
				}
				if last == -1 {
					last = getDigit(s, b)
				}
				if first >= 0 && last >= 0 {
					break
				}
			}
			_ = atomic.AddInt32(&sum, int32(first*10+last)) // critical for correctly adding up sum

		}(i, line)
	}

	wg.Wait()
	return sum
}

func getDigitPart1(s string, startIndex int) int {
	if s[startIndex] >= '0' && s[startIndex] <= '9' {
		return int(s[startIndex] - '0')
	}

	return -1
}

func getDigitPart2(s string, startIndex int) int {
	if s[startIndex] >= '0' && s[startIndex] <= '9' {
		return int(s[startIndex] - '0')
	}

	letters := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} // no zero
	for i, l := range letters {
		res := strings.Index(s[startIndex:], l)
		if res == 0 {
			return i + 1
		}
	}

	return -1
}
