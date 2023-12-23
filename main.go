package main

import (
	"Advent_Of_Code_2023/day_1"
	"Advent_Of_Code_2023/day_2"
	"Advent_Of_Code_2023/day_3"
	"Advent_Of_Code_2023/day_4"
	"Advent_Of_Code_2023/day_5"
	"Advent_Of_Code_2023/day_6"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalf("You need to provide the day !\n\t%s <day>\n", os.Args[0])
	}

	day := os.Args[1]

	switch day {
	case "1":
		runDay(day, day_1.Part1, day_1.Part2)
	case "2":
		runDay(day, day_2.Part1, day_2.Part2)
	case "3":
		runDay(day, day_3.Part1, day_3.Part2)
	case "4":
		runDay(day, day_4.Part1, day_4.Part2)
	case "5":
		runDay(day, day_5.Part1, day_5.Part2)
	case "6":
		runDay(day, day_6.Part1, day_6.Part2)
	default:
		log.Fatalf("Unknown day: %s\n", day)
	}
}

func runDay(day string, part1, part2 func()) {
	fmt.Printf("===== Day - %s =====\n", day)

	if part1 != nil {
		
		fmt.Printf("===== Part - 1 =====\n")
		part1()
		
	}

	if part2 != nil {
		
		fmt.Printf("===== Part - 2 =====\n")

		part2()
	}

	fmt.Println()
}