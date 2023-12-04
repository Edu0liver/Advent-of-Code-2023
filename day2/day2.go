package day2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

func Exec() *[]string {
	dataInput, err := os.Open(os.Getenv("ABSOLUTE_PATH") + "day2/input.txt")
	if err != nil {
		panic(err)
	}
	defer dataInput.Close()

	return possibleGames(dataInput)
}

func possibleGames(input *os.File) *[]string {
    scanner := bufio.NewScanner(input)
    scanner.Split(bufio.ScanLines)

    var fileLines []string

    for scanner.Scan() {
        fileLines = append(fileLines, scanner.Text())
    }

	gamesValids := []string{}

    for _, line := range fileLines {
		games := strings.Split(line, ": ")
		gameStats := make(map[string]int)
		gameNum := games[0]
		
		for _, game := range games[1:] {
			sets := strings.Split(game, "; ")

			for _, set := range sets{
				colors := strings.Split(set, ", ")
				
				for _, numColor := range colors {
					numAndColor := strings.Split(numColor, " ")

					color := strings.TrimSpace(numAndColor[1])
					num, err := strconv.Atoi(strings.TrimSpace(numAndColor[0]))
					if err != nil {
						panic(err)
					}

					gameStats[color] += num
				}
			}
		}

		if gameStats["blue"] <= MAX_BLUE && gameStats["green"] <= MAX_GREEN && gameStats["red"] <= MAX_RED {
			gamesValids = append(gamesValids, gameNum)
		}
    }

	return &gamesValids
}
