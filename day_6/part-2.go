package day_6

import (
	"Advent_Of_Code_2023/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)


func Part2(){
	defer utils.Timer("Part 1")()

	file, err := os.Open("./data/day6.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line1Grabbed := false
	var time int
	var distance int




	for scanner.Scan(){
		var line = scanner.Text()
		trimmedString := strings.ReplaceAll(line, " ", "")
		re := regexp.MustCompile(`(\d+)`)
		if !line1Grabbed {
			line1Grabbed = true
			timeAsString := re.FindAllString(trimmedString, -1)
			time = utils.ConvertToInt(timeAsString[0])
		} else {
			distAsString := re.FindAllString(trimmedString, -1)
			distance = utils.ConvertToInt(distAsString[0])
	}

	fmt.Printf("Wins Number: %v \n", getPossibleWins2(time, distance))
	
}
}
func getPossibleWins2(time int, dist int) int {
		
		

	// 	fmt.Printf("searchCount1: %v \n", searchCount1+searchCount2)
		upperBound := getUpperBound(time, dist, time, 0)
		lowerBound := getLowerBound(time, dist, time, 0)
		if(lowerBound == -1 || upperBound == -1){
			panic("Something went wrong")
		}

	
	return upperBound - lowerBound
}

func getUpperBound(time int, dist int, high int, low int) (results int) {
	if(low > high){
		return -1
	}
	
	mid := low + (high-low) / 2
	if(getDistance2(time, mid) > dist){
	   if(getDistance2(time, mid+1) <= dist){
		   return mid+1
	   } else {
		   return getUpperBound(time, dist, high, mid+1)
	   }
	} else {
		return getUpperBound(time, dist, mid-1, low)
	}


}

func getLowerBound(time int, dist int, high int, low int) (results int) {
	if(low > high){
		return -1
	}
	
	mid := low + (high-low) / 2
	 if(getDistance2(time, mid) > dist){
		if(getDistance2(time, mid-1) <= dist){
			return mid
		} else {
			return getLowerBound(time, dist, mid-1, low)
		}
	 } else {
		 return getLowerBound(time, dist, high, mid+1)
	 }



}


func getDistance2(time int, speed int) int {
	timeLeft := time - speed
	distance := 0
	for i := timeLeft; i > 0; i-- {
		distance += speed
	}

	return distance

}
