package day3

import (
	"bufio"
	"os"
	"unicode"

	"github.com/Edu0liver/Advent-of-Code-2023/pkg"
)

func Exec() *int {
	dataInput := pkg.OpenFile("day3/input.txt")
	defer dataInput.Close()

	return partsSum(dataInput)
}

func partsSum(input *os.File) *int {
    scanner := bufio.NewScanner(input)
    scanner.Split(bufio.ScanLines)

    var fileLines []string

    for scanner.Scan() {
        fileLines = append(fileLines, scanner.Text())
    }

	validNums := []int{}
	mappedInput := [][]string{}

    for _, line := range fileLines {
		lineMap := []string{}

		for _, c := range line {
			lineMap = append(lineMap, string(c))
		}

		mappedInput = append(mappedInput, lineMap)
    }

	for _, arr := range mappedInput {
		valueArr:
		for i, v := range arr {
            if !unicode.IsDigit([]rune(v)[0]) {
				continue valueArr
            }

			
		}
	}

    result := 0

    for _, v := range validNums {
        result += v
    }

    return &result
}