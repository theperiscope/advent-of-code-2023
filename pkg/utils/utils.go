package utils

import (
	"bufio"
	"os"
)

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
