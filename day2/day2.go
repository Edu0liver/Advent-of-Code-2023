package day2

import (
	"os"
	"strconv"
	"strings"

	"github.com/Edu0liver/Advent-of-Code-2023/pkg"
)

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

func Exec() *int {
	dataInput := pkg.OpenFile("day2/input.txt")
	defer dataInput.Close()

	return part1(dataInput)
}

func part1(input *os.File) *int {
	fileLines := pkg.FileLines(input)

	gamesValids := []int{}

	for _, line := range *fileLines {
		games := strings.Split(line, ": ")
		gameStats := make(map[string]int)
		gameNum, err := strconv.Atoi(strings.Split(games[0], " ")[1])
		if err != nil {
			panic(err)
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

		if gameStats["blue"] <= MAX_BLUE && gameStats["green"] <= MAX_GREEN && gameStats["red"] <= MAX_RED {
			gamesValids = append(gamesValids, gameNum)
		}
	}

	result := 0

	for _, v := range gamesValids {
		result += v
	}

	return &result
}
