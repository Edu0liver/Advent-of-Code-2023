package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func Exec() {
    dataInput, err := os.Open(os.Getenv("ABSOLUTE_PATH") + "day1/input.txt")
    if err != nil {
        panic(err)
    }
    defer dataInput.Close()
    
    result := textSum(dataInput)

    fmt.Println(*result)
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
            numsLine = fmt.Sprintf("%v%v", string(numsLine[0]), string(numsLine[len(numsLines)-1]))

            fmt.Println(string(numsLine[len(numsLines)-1]))
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

    fmt.Println(numsLines)

    result := sum(numsLines)

    return &result 
}

func sum(array []int) int {
    result := 0

    for _, v := range array {
        result += v
    }

    return result
}