package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func main(){
	file, err := os.Open("input.txt")
	check(err)
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
			
				check(err)
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

