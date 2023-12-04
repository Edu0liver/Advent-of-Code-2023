package day1

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"

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

// Wrong Answer:54547
// Right Answer: 54530

func Exec2() *int {
	dataInput := pkg.OpenFile("day1/input.txt")
	defer dataInput.Close()

	return part2(dataInput)
}

func part2(input *os.File) *int {
	fileLines := pkg.FileLines(input)
	var numsLines []int

	for _, line := range *fileLines {
		numsLineMap := make(map[int]string)

		for k, number := range numbers {
			r := strings.Split(line, k)

			if len(r) > 1 {
			inner:
				for count := 0; count < len(r); count++ {
					if _, ok := numsLineMap[len(r[count])]; ok {
						continue inner
					}

					numsLineMap[len(r[count])] = strconv.Itoa(number)
					break
				}
			}
		}

		for i, v := range line {
			if unicode.IsDigit(v) {
				numsLineMap[i] = string(v)
			}
		}

		numsLineMapKeys := make([]int, 0, len(numsLineMap))

		for k := range numsLineMap {
			numsLineMapKeys = append(numsLineMapKeys, k)
		}

		sort.Ints(numsLineMapKeys)

		numsLine := ""

		for _, key := range numsLineMapKeys {
			if val, ok := numsLineMap[key]; ok {
				numsLine += val
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
