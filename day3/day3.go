package day3

import (
	"fmt"
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
	fileLines := pkg.FileLines(input)

	validNums := make(map[int][]int)
	mappedInput := [][]string{}

	for _, line := range *fileLines {
		lineMap := []string{}

		for _, c := range line {
			lineMap = append(lineMap, string(c))
		}

		mappedInput = append(mappedInput, lineMap)
	}

	for iarr, arr := range mappedInput {
		prevArr := arr
		nextArr := arr

		if iarr > 0 {
			prevArr = mappedInput[iarr-1]
		}

		if iarr < (len(arr) - 1) {
			nextArr = mappedInput[iarr+1]
		}

	innerArray:
		for i, v := range arr {
			if !unicode.IsDigit(toRune(&v)) {
				continue innerArray
			}

			if HaveSymbolsAround(&arr, i) ||
				HaveSymbolsAround(&prevArr, i) ||
				HaveSymbolsAround(&nextArr, i) {
				validNums[iarr] = append(validNums[iarr], i)
			}
		}
	}

	fmt.Println(validNums)

	result := 0

	// for _, v := range validNums {
	// 	result += v
	// }

	return &result
}

func HaveSymbolsAround(currentArray *[]string, currentIndex int) bool {
	currentChar := (*currentArray)[currentIndex]
	prevChar := (*currentArray)[currentIndex]
	nextChar := (*currentArray)[len(*currentArray)-1]

	if currentIndex > 0 {
		prevChar = (*currentArray)[currentIndex-1]
	}

	if currentIndex < (len(*currentArray) - 1) {
		nextChar = (*currentArray)[currentIndex+1]
	}

	if (!unicode.IsDigit(toRune(&currentChar)) && !unicode.IsLetter(toRune(&currentChar)) && currentChar != ".") ||
		(!unicode.IsDigit(toRune(&prevChar)) && !unicode.IsLetter(toRune(&prevChar)) && prevChar != ".") ||
		(!unicode.IsDigit(toRune(&nextChar)) && !unicode.IsLetter(toRune(&nextChar)) && nextChar != ".") {
		return true
	}

	return false
}

func toRune(c *string) rune {
	return []rune(*c)[0]
}
