package day_3

import (
	"Advent_Of_Code_2024/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
)


func Part1(){
	defer utils.Timer("Part 1")()

	file, err := os.Open("./data/day_3.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum:=0
	for scanner.Scan(){
		var line = scanner.Text()

		re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
		matches := re.FindAllStringSubmatch(line, -1)

	for _, match := range matches {
		num1 := utils.ConvertToInt(match[1])
		num2 := utils.ConvertToInt(match[2])

		mult:=num1*num2
		sum = sum+mult
	}





	}


	fmt.Printf("SUM: %d \n",sum)
}


