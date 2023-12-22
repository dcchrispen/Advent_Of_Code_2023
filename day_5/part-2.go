package day_5

import (
	"Advent_Of_Code_2023/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sync"
)

type Seed struct {
	seedNum int
	rangeToEnd int
	
}

var lowestArr []int

func addLowest(lowest int) {
	lowestArr = append(lowestArr, lowest)
}

func findLowest() int {
	lowest := lowestArr[0]
	for i := 0; i < len(lowestArr); i++ {
		if(lowestArr[i] < lowest){
			lowest = lowestArr[i]
		}
	}
	return lowest
}

func Part2(){
	defer utils.Timer("Part 2")()
	file, err := os.Open("./data/day5.txt")
	utils.Check(err)
	defer file.Close()

	var keyword string 

	var seedArr []Seed
	var seedToSoil []Groups
	var soilToFert []Groups
	var fertToWat []Groups
	var watToLight []Groups
	var lightToTemp []Groups
	var tempToHumid []Groups
	var humidToLoc []Groups

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		var line = scanner.Text()
		
		nums := regexp.MustCompile(`\d+`)
		alph := regexp.MustCompile(`[a-z]+`)
		words :=alph.FindAllString(line, -1)
		numbers := nums.FindAllString(line, -1)

		if(len(words) > 0){
			keyword = words[0]
		}
		if(keyword == "seeds"){
			for i := 0; i < len(numbers); i = i+2 {
				seedNum := utils.ConvertToInt(numbers[i])
				rangeToEnd := utils.ConvertToInt(numbers[i+1])
				seedArr = append(seedArr, Seed{seedNum: seedNum, rangeToEnd: rangeToEnd})
			}

		} else if(keyword == "seed"){
			if(len(numbers) > 0){
				numArr := utils.ConvertStringsToNums(numbers)
				seedToSoil =append(seedToSoil, Groups{dest: numArr[0], source: numArr[1], rangeToEnd: numArr[2]})
			}
		} else if(keyword == "soil"){
			if(len(numbers) > 0){
				numArr := utils.ConvertStringsToNums(numbers)
				soilToFert = append(soilToFert, Groups{dest: numArr[0], source: numArr[1], rangeToEnd: numArr[2]})
			}
		} else if(keyword == "fertilizer"){
			if(len(numbers) > 0){
				numArr := utils.ConvertStringsToNums(numbers)
				fertToWat = append(fertToWat, Groups{dest: numArr[0], source: numArr[1], rangeToEnd: numArr[2]})
			}
		} else if(keyword == "water"){
			if(len(numbers) > 0){
				numArr := utils.ConvertStringsToNums(numbers)
				watToLight =append(watToLight, Groups{dest: numArr[0], source: numArr[1], rangeToEnd: numArr[2]})
			}
		} else if(keyword == "light"){
			if(len(numbers) > 0){
				numArr := utils.ConvertStringsToNums(numbers)
				lightToTemp =append(lightToTemp, Groups{dest: numArr[0], source: numArr[1], rangeToEnd: numArr[2]})
			}
		} else if(keyword == "temperature"){
			if(len(numbers) > 0){
				numArr := utils.ConvertStringsToNums(numbers)
				tempToHumid =append(tempToHumid, Groups{dest: numArr[0], source: numArr[1], rangeToEnd: numArr[2]})
			}
		} else if(keyword == "humidity"){
			if(len(numbers) > 0){
				numArr := utils.ConvertStringsToNums(numbers)
				humidToLoc =append(humidToLoc, Groups{dest: numArr[0], source: numArr[1], rangeToEnd: numArr[2]})
			}
		}
	}



	wg := &sync.WaitGroup{}

	for i := 0; i < len(seedArr); i++ {
		wg.Add(1)
		go findLow(seedArr[i], seedToSoil, soilToFert, fertToWat, watToLight, lightToTemp, tempToHumid, humidToLoc, wg)
	}
	
	wg.Wait()

	fmt.Printf("Lowest Loc: %d  \n", findLowest())
	

}


func findLow(seedsToEval Seed, seedToSoil []Groups, soilToFert []Groups, fertToWat []Groups, watToLight []Groups, lightToTemp []Groups, tempToHumid []Groups, humidToLoc []Groups, wg *sync.WaitGroup ) {
	defer wg.Done()
	fmt.Printf("Processing Seed: %d  \n", seedsToEval.seedNum)

	lowestLoc := -1
	rangeOfSeeds := seedsToEval.rangeToEnd
	seedStart := seedsToEval.seedNum
	seedEnd := seedStart + rangeOfSeeds-1
	
	for j:= seedStart ;  j < seedEnd; j++ {

		dest := findDest(seedToSoil, j)
		dest = findDest(soilToFert, dest)
		dest = findDest(fertToWat, dest)
		dest = findDest(watToLight, dest)
		dest = findDest(lightToTemp, dest)
		dest = findDest(tempToHumid, dest)
		dest = findDest(humidToLoc, dest)
		if(lowestLoc == -1){
			lowestLoc = dest
		}
		if(dest < lowestLoc){
			lowestLoc = dest
	}
	
}

addLowest(lowestLoc)

	fmt.Printf("Finished Processing Seed: %d  \n", seedsToEval.seedNum)
}