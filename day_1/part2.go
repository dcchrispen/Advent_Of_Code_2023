package day_1

import (
	"Advent_Of_Code_2023/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)



func Part2(){
	file, err := os.Open("./data/day1.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan(){
		var line = scanner.Text()
		var numsReplaced string
		for i := 0; i < 10; i++ {
			switch i {	
			case 1:
				numsReplaced = strings.ReplaceAll(numsReplaced, "one", "o1ne")
			case 2:
				numsReplaced = strings.ReplaceAll(numsReplaced, "two", "tw2o")
			case 3:
				numsReplaced = strings.ReplaceAll(numsReplaced, "three", "th3ree")
			case 4:
				numsReplaced = strings.ReplaceAll(numsReplaced, "four", "fo4ur")
			case 5:
				numsReplaced = strings.ReplaceAll(numsReplaced, "five", "fi5ve")
			case 6:
				numsReplaced = strings.ReplaceAll(numsReplaced, "six", "si6x")
			case 7:
				numsReplaced = strings.ReplaceAll(numsReplaced, "seven", "sev7en")
			case 8:
				numsReplaced = strings.ReplaceAll(numsReplaced, "eight", "ei8ght")
			case 9:
				numsReplaced = strings.ReplaceAll(numsReplaced, "nine", "ni9ne")
			default: 
				numsReplaced = strings.ReplaceAll(line, "zero", "ze0ro")
			}
		}
		reg := regexp.MustCompile(`\d+`)
		nums := reg.FindAllString(numsReplaced, -1)


		var allDigits string
		for i := 0; i < len(nums); i++ {
			allDigits = allDigits + nums[i]
		}

		finalStrArr := strings.Split(allDigits, "")
		finalStr := finalStrArr[0] + finalStrArr[len(finalStrArr)-1]
		num, err := strconv.Atoi(finalStr)
		utils.Check(err)
		sum = sum + num
	}

	fmt.Printf("Final Sum: %d  \n", sum)
}

