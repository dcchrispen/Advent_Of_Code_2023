package day_5

import (
	"Advent_Of_Code_2023/utils"
	"bufio"
	"fmt"
	"os"
)

type Element struct {
	elementType string
	source int
	destination int
}



func Part1(){

	file, err := os.Open("./data/day5.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		var line = scanner.Text()
		fmt.Printf("Line: %s  \n", line)
	}

}

func parseFile() {

}