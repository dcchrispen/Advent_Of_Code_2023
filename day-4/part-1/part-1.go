package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func main(){
	file, err := os.Open("../input.txt")
	check(err)
	defer file.Close()

	sum :=0

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		var line = scanner.Text()
		fmt.Printf("Line: %s  \n", line)
		// get nums and store in two arrays
		numbers := strings.Split(line, ":")[1]
		winningNumStrings := strings.Split(numbers, "|")[0]
		possibleNumStrings := strings.Split(numbers, "|")[1]
		re := regexp.MustCompile(`\d+`)
		winningArr := convertStringsToNums(re.FindAllString(winningNumStrings, -1))
		possibleArr := convertStringsToNums(re.FindAllString(possibleNumStrings, -1))

		winningArr = sortNums(winningArr)
		possibleArr = sortNums(possibleArr)

		
		sum = sum+ calculateNums(winningArr, possibleArr)
	}

	fmt.Printf("Final Sum: %d  \n", sum)

}

func calculateNums(winningArr []int, possibleArr []int) int {
	matchCount := 0

	for i := 0; i < len(winningArr); i++ {
		for j := 0; j < len(possibleArr); j++ {
			if(winningArr[i] < possibleArr[j]){
				break
			}
			if winningArr[i] == possibleArr[j] {
				matchCount++
			}
		}
	
	}

	if(matchCount <2){
		return matchCount
	} else {
		
		sum :=1
		for i := 1 ; i < matchCount; i++ {
			sum = sum * 2
		
		}
		return sum
	}

}

func convertStringsToNums(stringArr []string) []int {
	var numArr []int
	for i := 0; i < len(stringArr); i++ {
		num, err := strconv.Atoi(stringArr[i])
		check(err)
		numArr = append(numArr, num)
	}
	return numArr
}

func sortNums(numArr []int) []int {
	for i := 0; i < len(numArr); i++ {
		for j := 0; j < len(numArr); j++ {
			if numArr[i] < numArr[j] {
				temp := numArr[i]
				numArr[i] = numArr[j]
				numArr[j] = temp
			}
		}
	}
	return numArr
}