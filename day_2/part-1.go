package day_2

import (
	"Advent_Of_Code_2024/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func Part1(){
	defer utils.Timer("Part 1")()
	numValid:=0

	file, err := os.Open("./data/day_2.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		var line = scanner.Text()

		strSplits := strings.Split(line, " ")

		lastNum:=0
		increasingOrDecreasing:="initial"
		isValid:=true

		for i:=0; i<len(strSplits); i++ {
			num, err := strconv.Atoi(strSplits[i])
			utils.Check(err)
			if(i==0){
				lastNum=num
				continue;
			}
			if(increasingOrDecreasing=="initial"){
				if(lastNum-num >0){
					increasingOrDecreasing="increase"
				} else {
					increasingOrDecreasing="decrease"
				}

			} else if (increasingOrDecreasing=="increase" && lastNum-num <0){
				isValid=false
			} else if (increasingOrDecreasing=="decrease" && lastNum-num >0){
				isValid=false

			}
			numCheck :=utils.AbsDiffInt(lastNum,num)

			if(numCheck<1 || numCheck>3){
				isValid=false

			}

			lastNum = num
		}

		if(isValid){
			numValid++
		}
	}

	fmt.Printf("num valid: %d \n", numValid)

}

