package main

import (
	"fmt"
	"log"

	"github.com/Edu0liver/Advent-of-Code-2023/day1"
	"github.com/Edu0liver/Advent-of-Code-2023/day2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	fmt.Println("Day 1:", *day1.Exec())
	fmt.Println("Day 1, part 2:", *day1.Exec2())
	fmt.Println("Day 2:", *day2.Exec())
	fmt.Println("Day 2, part 2:", *day2.Exec2())
	// fmt.Println("Day 3:", *day3.Exec())
}
