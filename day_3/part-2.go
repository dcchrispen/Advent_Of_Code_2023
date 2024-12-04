package day_3

import (
	"Advent_Of_Code_2024/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
)


func Part2(){
	defer utils.Timer("Part 2")()

	file, err := os.Open("./data/day_3.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum:=0

	// enabledString:=""
	// enabled:=true

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	reNum:= regexp.MustCompile(`\d{1,3}`)
	enabled:=true
	stringToBeParsed :=""
	for scanner.Scan(){
		var line = scanner.Text()
		stringToBeParsed = stringToBeParsed+line

	}
	matches := re.FindAllString(stringToBeParsed, -1)
	for _, match := range matches {
		fmt.Printf("Match: %s \n", match)
		if match=="do()"{
			enabled=true
		} else if(match=="don't()"){
			enabled=false
		} else {
			if (enabled){
				numbers := reNum.FindAllString(match, -1)
				num1 := utils.ConvertToInt(numbers[0])
				num2 := utils.ConvertToInt(numbers[1])
		
				mult:=num1*num2
				sum = sum+mult

			}

		}
	}

	fmt.Printf("SUM: %d \n",sum)
}


