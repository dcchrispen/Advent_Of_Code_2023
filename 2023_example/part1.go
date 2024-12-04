package example_1

import (
	"Advent_Of_Code_2024/utils"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)




func Part1(){
	file, err := os.Open("./data/example_1.txt")
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	var cals []int
	for scanner.Scan(){
		var line = scanner.Text()

		
		if(len(line)<1){
			fmt.Printf("SUM: %d \n", sum)
			cals = append(cals, sum)
			sum = 0
		} else {
			num, err := strconv.Atoi(line)
			utils.Check(err)

			sum = sum + num
		}		
	}
	if(sum > 0){
		cals = append(cals, sum)
			sum = 0
	}
	sort.SliceStable(cals, func(i,j int) bool {
		return cals[i] > cals[j]
	})

	tot := cals[0] + cals[1] +cals[2]

	fmt.Printf("Total: %d \n", tot)

}

