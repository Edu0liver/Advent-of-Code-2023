package day1

import (
	"fmt"
	"os"
	"strconv"
	"unicode"

	"github.com/Edu0liver/Advent-of-Code-2023/pkg"
)

func Exec() *int {
	dataInput := pkg.OpenFile("day1/input.txt")
	defer dataInput.Close()

	return textSum(dataInput)
}

func textSum(input *os.File) *int {
	fileLines := pkg.FileLines(input)
	var numsLines []int

	for _, line := range *fileLines {
		numsLine := ""

		for _, v := range line {
			if unicode.IsDigit(v) {
				numsLine += string(v)
			}
		}

		if len(numsLine) > 2 {
			numsLine = fmt.Sprintf("%v%v", string(numsLine[0]), string(numsLine[len(numsLine)-1]))
		}

		if len(numsLine) < 2 {
			numsLine += string(numsLine[0])
		}

		num, err := strconv.Atoi(numsLine)
		if err != nil {
			panic(err)
		}

		numsLines = append(numsLines, num)
	}

	result := 0

	for _, v := range numsLines {
		result += v
	}

	return &result
}
