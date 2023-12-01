package lib

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func GetProgramName() string {
	return filepath.Base(os.Args[0])
}

func AssertArgs() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 1 {
		fmt.Printf("Usage: %s <inputfile>\n", GetProgramName())
		os.Exit(1)
	}
}

func AssertInput() []string {
	lines, err := ReadInput(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	return lines
}

func Filter(input []string, f func(string) bool) []string {
	filtered := make([]string, 0)
	for _, v := range input {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

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
