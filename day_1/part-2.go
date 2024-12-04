package day_1

import (
	"Advent_Of_Code_2024/utils"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)


func Part2(){
	defer utils.Timer("Part 1")()

	file, err := os.Open("./data/day_1.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var firstArr []int
	var secArr []int


	for scanner.Scan(){
		var line = scanner.Text()

		strSplits := strings.Split(line, " ")

		firstNum, err := strconv.Atoi(strSplits[0])
			utils.Check(err)

		firstArr = append(firstArr, firstNum)
		secNum, err := strconv.Atoi(strSplits[3])
			utils.Check(err)
		secArr = append(secArr, secNum)
	
	}

	sort.SliceStable(firstArr, func(i,j int) bool {
		return firstArr[i] < firstArr[j]
	})
	sort.SliceStable(secArr, func(i,j int) bool {
		return secArr[i] < secArr[j]
	})


	simScore := 0

	for i:=0; i<len(firstArr); i++{
		numCheck := firstArr[i]
		timesFound := 0

		for j:=0; j<len(secArr); j++{
			if(secArr[j]==numCheck){
				timesFound ++
			} else if (secArr[j] > numCheck) {
				simScore = simScore + (numCheck*timesFound)
				timesFound=0
				continue
			}
			
		}

	}

	fmt.Printf("SUM: %d \n", simScore)
}

