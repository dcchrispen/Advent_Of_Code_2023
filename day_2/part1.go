package day_2

import (
	"Advent_Of_Code_2023/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)



func Part1(){
	file, err := os.Open("./data/day2.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	currentId := 1
	for scanner.Scan(){
		
		valid := true
		var line = scanner.Text()
		//extract id
	
		stringWithoutId  := strings.Split(line, ":")[1]
		// extract games
		games := strings.Split(stringWithoutId, ";")
	
		for i := 0; i < len(games); i++ {
			cubeGroups := strings.Split(games[i], ", ")
			for j :=0 ; j < len(cubeGroups); j++ {
				trimmedCubes := strings.TrimSpace(cubeGroups[j])
				cubes := strings.Split(trimmedCubes, " ")
				color := cubes[1]
				number, err := strconv.Atoi(cubes[0])
			
				utils.Check(err)
				switch color {
					case "red":
						if number > 12 {
							valid = false
						}
					case "green":
						if number > 13 {
							valid = false
						}
					case "blue":
						if number > 14 {
							valid = false
						}
				}

				if(!valid){
					break
				}
				
				
			}
			if(!valid){
				break
			}
			
		}
		
		if valid{sum = sum + currentId}
		currentId = currentId + 1
	}

	fmt.Printf("Final Sum: %d  \n", sum)
}

