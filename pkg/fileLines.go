package pkg

import (
	"bufio"
	"os"
)

func FileLines(input *os.File) *[]string {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var fileLines []string

	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	return &fileLines
}
