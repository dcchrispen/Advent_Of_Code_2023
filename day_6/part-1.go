package day_6

import (
	"Advent_Of_Code_2023/utils"
	"bufio"
	"fmt"
	"os"
)


func Part1(){
	defer utils.Timer("Part 1")()

	file, err := os.Open("./data/day6.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line1Grabbed := false
	var times []int
	var distances []int




	for scanner.Scan(){
		var line = scanner.Text()
		if !line1Grabbed {
			line1Grabbed = true
			times = utils.GetNumArrayFromString(line)
		} else {
			distances = utils.GetNumArrayFromString(line)
		}
	}

	fmt.Printf("Wins Number: %v \n", getPossibleWins(times, distances))

}

func getPossibleWins(times []int, distances []int) int {
	var winsArray []int	
	for i := 0 ; i < len(times); i++ {
		possibleWins :=0
		
		for j := 0 ; j < times[i]; j++ {
			distance := getDistance(times[i], j)
			if distance > distances[i] {
					
				possibleWins++
			}
		}
		winsArray = append(winsArray, possibleWins)
	}
	totalWins := 1
	for k := 0 ; k < len(winsArray); k++ {
			totalWins = totalWins * winsArray[k]
	}
	return totalWins
}


func getDistance(time int, speed int) int {
	timeLeft := time - speed
	distance := 0
	for i := timeLeft; i > 0; i-- {
		distance += speed
	}
	return distance

}