package day_5

import (
	"Advent_Of_Code_2023/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Groups struct {
	dest int
	source int
	rangeToEnd int
}


func Part1(){
	defer utils.Timer("Part 1")()
	file, err := os.Open("./data/day5.txt")
	utils.Check(err)
	defer file.Close()

	var keyword string 

	var seedArr []int
	var seedToSoil []Groups
	var soilToFert []Groups
	var FertToWat []Groups
	var WatToLight []Groups
	var LightToTemp []Groups
	var TempToHumid []Groups
	var HumidToLoc []Groups

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
			for i := 0; i < len(numbers); i++ {
				seedNum := utils.ConvertToInt(numbers[i])
				seedArr = append(seedArr, seedNum)
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
				FertToWat = append(FertToWat, Groups{dest: numArr[0], source: numArr[1], rangeToEnd: numArr[2]})
			}
		} else if(keyword == "water"){
			if(len(numbers) > 0){
				numArr := utils.ConvertStringsToNums(numbers)
				WatToLight =append(WatToLight, Groups{dest: numArr[0], source: numArr[1], rangeToEnd: numArr[2]})
			}
		} else if(keyword == "light"){
			if(len(numbers) > 0){
				numArr := utils.ConvertStringsToNums(numbers)
				LightToTemp =append(LightToTemp, Groups{dest: numArr[0], source: numArr[1], rangeToEnd: numArr[2]})
			}
		} else if(keyword == "temperature"){
			if(len(numbers) > 0){
				numArr := utils.ConvertStringsToNums(numbers)
				TempToHumid =append(TempToHumid, Groups{dest: numArr[0], source: numArr[1], rangeToEnd: numArr[2]})
			}
		} else if(keyword == "humidity"){
			if(len(numbers) > 0){
				numArr := utils.ConvertStringsToNums(numbers)
				HumidToLoc =append(HumidToLoc, Groups{dest: numArr[0], source: numArr[1], rangeToEnd: numArr[2]})
			}
		}
	}


	lowestLoc := -1

	for i := 0; i < len(seedArr); i++ {
		seed := seedArr[i]
		dest := findDest(seedToSoil, seed)
		dest = findDest(soilToFert, dest)
		dest = findDest(FertToWat, dest)
		dest = findDest(WatToLight, dest)
		dest = findDest(LightToTemp, dest)
		dest = findDest(TempToHumid, dest)
		dest = findDest(HumidToLoc, dest)
		if(lowestLoc == -1){
			lowestLoc = dest
		}
		if(dest < lowestLoc){
			lowestLoc = dest
		}
	
	}

	fmt.Printf("Lowest Loc: %d  \n", lowestLoc)
	

}

func findDest(arr []Groups, source int) int{
	destFound  := -999
	for i:= 0 ; i < len(arr); i++ {
		currentArrVal := arr[i]
		start := currentArrVal.source
		end := currentArrVal.source + currentArrVal.rangeToEnd-1
		if(source >= start && source <= end){
			destFound = currentArrVal.dest + (source - start)
		}

	}
	
	if(destFound == -999){
		destFound = source
		
	}
	return destFound
}