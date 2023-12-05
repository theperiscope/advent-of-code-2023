package lib

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// returns program name
func GetProgramName() string {
	return filepath.Base(os.Args[0])
}

// asserts exactly one program argument is used and exits otherwise
func AssertArgs() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 1 {
		fmt.Printf("Usage: %s <inputfile>\n", GetProgramName())
		os.Exit(1)
	}
}

// asserts input lines are read from first command line argument and logs/exit otherwise
func AssertInput() []string {
	lines, err := ReadInput(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	return lines
}

// reads line-by-line input from specified file
func ReadInput(fileName string) (input []string, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input, nil
}

// converts string to int and panics on error (useful for slice conversions)
func Atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
