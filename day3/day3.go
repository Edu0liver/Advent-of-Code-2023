package day3

import (
	"os"

	"github.com/Edu0liver/Advent-of-Code-2023/pkg"
)

func Exec() *int {
	dataInput := pkg.OpenFile("day3/input.txt")
	defer dataInput.Close()

	return partsSum(dataInput)
}

func partsSum(input *os.File) *int {
	a := 1

	return &a
}