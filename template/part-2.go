package day_

import (
	"Advent_Of_Code_2023/utils"
	"bufio"
	"fmt"
	"os"
)


func Part2(){

	file, err := os.Open("./data/day.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		var line = scanner.Text()
		fmt.Printf("Line: %s  \n", line)
	}

}

