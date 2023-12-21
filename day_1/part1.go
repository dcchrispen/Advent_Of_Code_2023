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




func Part1(){
	file, err := os.Open("./data/day1.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan(){
		var line = scanner.Text()
		
		reg := regexp.MustCompile(`\d+`)
		nums := reg.FindAllString(line, -1)


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

