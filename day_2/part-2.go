package day_2

import (
	"Advent_Of_Code_2024/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)



func Part2(){
	defer utils.Timer("Part 1")()
	numValid:=0

	file, err := os.Open("./data/day_2.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		var line = scanner.Text()

		strSplits := strings.Split(line, " ")

		var numArr []int
		for i:=0; i<len(strSplits); i++ {
			num, err := strconv.Atoi(strSplits[i])
			utils.Check(err)

			numArr = append(numArr, num)
		}
	
	
		if(checkArr(numArr, false)){
			numValid++
		}
	}


		

		fmt.Printf("Num Valid: %d \n", numValid)
	}

	func checkArr(arr []int, isFiltered bool) bool{
		isValid:=true

		increaseCheck := checkIncrease(arr)
		decreaseCheck := checkDecrease(arr)


		
		if(!increaseCheck && !decreaseCheck){
			isValid = false
		}

		if(!checkSpacing(arr)){
			isValid = false
		}

		if(!isFiltered && !isValid){
			for k:=0;k<len(arr);k++{
				newArr:=utils.FilterOutAtIndex(arr,k)

				if(checkArr(newArr, true)){
					return true

			}


			}
		}

		return isValid



	}

	func checkIncrease(arr []int) bool {
		
		isValid:=true

		for i:=0;i<len(arr);i++{
			if(i==0){ 
				continue;
			} else if(arr[i]<arr[i-1]){
				isValid=false
			}


		}



		return isValid

	}

	func checkDecrease(arr []int) bool {
		
		isValid:=true

		for i:=0;i<len(arr);i++{
			if(i==0){ 
				continue;
			}else if(arr[i]>arr[i-1]){
				isValid=false
			}


		}



		return isValid

	}

	func checkSpacing(arr []int) bool {
		
		isValid:=true

		for i:=0;i<len(arr);i++{
			if(i==0){ 
				continue;
			}
			
			numCheck :=utils.AbsDiffInt(arr[i-1],arr[i])


			if(numCheck<1 || numCheck>3){
				isValid=false

			}
		}



		return isValid

	}