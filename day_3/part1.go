package day_3

import (
	"Advent_Of_Code_2023/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)




func Part1(){

	file, err := os.Open("./data/day3.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	slice1 := make([]string, 0)
	slice2 := make([]string, 0)
	sum := 0
	for scanner.Scan(){
		var line = scanner.Text()
		slice1 = append(slice1, line)
		slice2 = append(slice2, line)
	}
	for i := 0; i < len(slice1); i++ {
		slice2[i] = strings.ReplaceAll(slice2[i], ".", "F")
		re := regexp.MustCompile(`\d`)
		slice2[i] = re.ReplaceAllString(slice2[i], "F")

	}

	for i := 0; i < len(slice1); i++ {
		stringToEval := slice1[i]
		indexes := getIndexes(stringToEval)
		for j := 0; j < len(indexes); j++ {
			currentIndexes := indexes[j]
			isPart := checkIndexes(currentIndexes, slice2, i)
			if(isPart){
				currentNum := stringToEval[currentIndexes[0]:currentIndexes[1]]
			
				currentNumToInt, err := strconv.Atoi(currentNum)
				utils.Check(err)
				sum += currentNumToInt
			}		
		}
	}

	fmt.Printf("Sum: %d \n", sum)
}

func checkIndexes(indexes [] int, sliceToCheck []string, currentSliceIndex int) bool { 
	isPartOfSlice := false
	if(currentSliceIndex != 0 && currentSliceIndex != len(sliceToCheck)-1){
			if(checkSliceForSymbol(sliceToCheck[currentSliceIndex-1], indexes[0], indexes[1]-1) || checkSliceForSymbol(sliceToCheck[currentSliceIndex+1], indexes[0], indexes[1]-1) || checkSidesOfSlices(sliceToCheck[currentSliceIndex], indexes)){
				isPartOfSlice = true
			}
	}
	if(currentSliceIndex == 0){	
		if(checkSliceForSymbol(sliceToCheck[currentSliceIndex+1], indexes[0], indexes[1]-1) || checkSidesOfSlices(sliceToCheck[currentSliceIndex], indexes)){
				isPartOfSlice = true
		}
	}
	if(currentSliceIndex == len(sliceToCheck)-1){
		if(checkSliceForSymbol(sliceToCheck[currentSliceIndex-1], indexes[0], indexes[1]-1) || checkSidesOfSlices(sliceToCheck[currentSliceIndex], indexes)){
				isPartOfSlice = true
		}
	}
	return isPartOfSlice
}

func checkSidesOfSlices (sliceToCheck string, indexes [] int) bool {

	if(indexes[0] >0 && indexes[1] < len(sliceToCheck)){
		firstChar := string(sliceToCheck[indexes[0]-1])
		secondChar := string(sliceToCheck[indexes[1]])
		checkChar := "F"
		if(firstChar == checkChar && secondChar == checkChar){
			return false 
		} else {
			return true
		}
	}

	return false
}

func checkSliceForSymbol(stringToEval string, firstInd int , secondInd int) bool {
	if(firstInd <1){
		return checkForSymbols(stringToEval, firstInd, secondInd+1)

	} else if (secondInd == 139){
		return checkForSymbols(stringToEval, firstInd-1, secondInd)

	} else {
		return checkForSymbols(stringToEval, firstInd-1, secondInd+1)
	}

}
func checkForSymbols(stringToEval string, start int, end int) bool {
	re := regexp.MustCompile(`[^F]`)
	for i := start; i <= end; i++ {
		if(re.MatchString(string(stringToEval[i]))){
			return true
		}
	
	}
	return false
}

func getIndexes (stringToEval string) [][]int {
	re := regexp.MustCompile(`\d+`)
	content := []byte(stringToEval)
	indexes := re.FindAllIndex(content, -1)
	return indexes
}