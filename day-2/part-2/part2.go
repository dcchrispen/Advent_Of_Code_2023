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
	for scanner.Scan(){
		power := 0
		
		var line = scanner.Text()
		fmt.Printf("Line: %s  \n", line)
		//extract id
	
		stringWithoutId  := strings.Split(line, ":")[1]
		// extract games
		games := strings.Split(stringWithoutId, ";")
		red := 0
		green := 0
		blue := 0
		
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
						if number > red {
							red = number
						}
					case "green":
						if number > green {
							green = number
						}
					case "blue":
						if number > blue {
							blue = number
						}
				}

				
				
				
			}
		
			
		}

		power = red * green * blue
		
		sum = sum + power
	}
	fmt.Printf("Sum: %d  \n", sum)

}

