package day_3

import (
	"Advent_Of_Code_2023/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)



func Part2(){

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
		re := regexp.MustCompile(`[^0-9]`)
		slice2[i] = re.ReplaceAllString(slice2[i], "F")
	}

	for i := 0; i < len(slice1); i++ {
		stringToEval := slice1[i]
		indexes := getIndexs(stringToEval)
		for j := 0; j < len(indexes); j++ {
			currentIndexes := indexes[j]
			num := getNums(slice2, i, currentIndexes[0])
			sum += num			
		}
	}
	fmt.Printf("Sum: %d \n", sum)

}


func getNums(sliceToCheck []string, indexOfSlice int, indexOfChar int) int{
	var nums []int

	re := regexp.MustCompile(`[0-9]+`)
	in := regexp.MustCompile(`[0-9]+`)

	stringNums := re.FindAllString(sliceToCheck[indexOfSlice], -1)
	stringIndexes :=in.FindAllStringIndex(sliceToCheck[indexOfSlice], -1)

	for i := 0; i < len(stringIndexes); i++ {
		currentIndexes := stringIndexes[i]
		currentNum := stringNums[i]

		if(currentIndexes[0]-1 == indexOfChar || currentIndexes[1] == indexOfChar){
			currentNumToInt, err := strconv.Atoi(currentNum)
			utils.Check(err)
			nums = append(nums, currentNumToInt)
		
		}
	}

	var topNums []int
	var bottomNums []int
	if(indexOfSlice != 0){
		topNums = getTopBottomNums(sliceToCheck, indexOfSlice - 1, indexOfChar)
	}
	if (indexOfSlice != len(sliceToCheck) - 1){
		bottomNums = getTopBottomNums(sliceToCheck, indexOfSlice + 1, indexOfChar)
	}

	if(len(topNums) > 0){
		nums = append(nums, topNums...)
	}
	if(len(bottomNums) > 0){
		nums = append(nums, bottomNums...)
	}
	finalNum := 0
	if(len(nums) > 1){
		for i := 0; i < len(nums); i++{
			if(i != len(nums) - 1){
				finalNum += nums[i] * nums[i+1]
			}
	}

	
}
return finalNum
}

func getTopBottomNums(sliceToCheck []string, indexOfSlice int, indexOfChar int) []int{
	var nums []int
	re := regexp.MustCompile(`\d+`)
	stringToEval := sliceToCheck[indexOfSlice]

	content := []byte(stringToEval)
	indexes := re.FindAllIndex(content, -1)
	for i := 0; i < len(indexes); i++ {
		currentIndexes := indexes[i]
		for j := currentIndexes[0]; j < currentIndexes[1]; j++{
			if(j == indexOfChar || j == indexOfChar - 1 || j == indexOfChar + 1){
				currentNum := string(stringToEval[currentIndexes[0]:currentIndexes[1]])
				currentNumToInt, err := strconv.Atoi(currentNum)
				utils.Check(err)
				nums = append(nums, currentNumToInt)
				
				break
			}
		}
	}

	return nums
}

func getIndexs (stringToEval string) [][]int {
	re := regexp.MustCompile(`[*]`)
	content := []byte(stringToEval)
	indexes := re.FindAllIndex(content, -1)
	return indexes
}