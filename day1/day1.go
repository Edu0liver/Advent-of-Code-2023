package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func Exec() *int {
    dataInput, err := os.Open(os.Getenv("ABSOLUTE_PATH") + "day1/input.txt")
    if err != nil {
        panic(err)
    }
    defer dataInput.Close()
    
    return textSum(dataInput)
}

func textSum(input *os.File) *int {
    scanner := bufio.NewScanner(input)
    scanner.Split(bufio.ScanLines)

    var fileLines []string
    var numsLines []int

    for scanner.Scan() {
        fileLines = append(fileLines, scanner.Text())
    }

    for _, line := range fileLines {
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
