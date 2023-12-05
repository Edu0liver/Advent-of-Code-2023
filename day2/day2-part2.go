package day2

import (
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/Edu0liver/Advent-of-Code-2023/pkg"
)

func Exec2() *int {
	dataInput := pkg.OpenFile("day2/input.txt")
	defer dataInput.Close()

	return part2(dataInput)
}

func part2(input *os.File) *int {
	fileLines := pkg.FileLines(input)

	gamesValids := []int{}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for _, line := range *fileLines {
			games := strings.Split(line, ": ")

			gameStats := map[string]int{
				"red":   0,
				"green": 0,
				"blue":  0,
			}

			for _, game := range games[1:] {
				sets := strings.Split(game, "; ")

				for _, set := range sets {
					colors := strings.Split(set, ", ")

					for _, numColor := range colors {
						numAndColor := strings.Split(numColor, " ")

						color := strings.TrimSpace(numAndColor[1])
						num, err := strconv.Atoi(strings.TrimSpace(numAndColor[0]))
						if err != nil {
							panic(err)
						}

						if gameStats[color] < num {
							gameStats[color] = num
						}
					}
				}
			}

			power := 1
			for _, v := range gameStats {
				power *= v
			}

			gamesValids = append(gamesValids, power)
		}
		wg.Done()
	}()
	wg.Wait()

	result := 0
	for _, v := range gamesValids {
		result += v
	}

	return &result
}
