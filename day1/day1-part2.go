package day1

import (
	"os"
	"strconv"
	"strings"

	"github.com/Edu0liver/Advent-of-Code-2023/pkg"
)

var numbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func Exec2() *int {
	dataInput := pkg.OpenFile("day1/input.txt")
	defer dataInput.Close()

	return part2(dataInput)
}

func part2(input *os.File) *int {
	fileLines := pkg.FileLines(input)
	var numsLines []int

	for _, line := range *fileLines {
		digits := []int{}
		answer := 0

		for i, char := range line {
			if digit, err := strconv.Atoi(string(char)); err == nil {
				digits = append(digits, digit)
			} else {
				for spelling, number := range numbers {
					if strings.HasPrefix(line[i:], spelling) {
						digits = append(digits, number)
						break
					}
				}
			}
		}

		first, last := digits[0], digits[len(digits)-1]
		answer += first*10 + last

		numsLines = append(numsLines, answer)
	}

	result := 0

	for _, v := range numsLines {
		result += v
	}

	return &result
}
