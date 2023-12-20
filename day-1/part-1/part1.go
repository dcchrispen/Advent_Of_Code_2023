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
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan(){
		var line = scanner.Text()
		fmt.Printf("Line: %s  \n", line)
		
		reg := regexp.MustCompile(`\d+`)
		nums := reg.FindAllString(line, -1)

		fmt.Printf("Nums: %s  \n", nums)

		var allDigits string
		for i := 0; i < len(nums); i++ {
			allDigits = allDigits + nums[i]
		}

		finalStrArr := strings.Split(allDigits, "")
		finalStr := finalStrArr[0] + finalStrArr[len(finalStrArr)-1]
		fmt.Println(finalStr)
		num, err := strconv.Atoi(finalStr)
		check(err)
		sum = sum + num
	}

	fmt.Printf("Final Sum: %d  \n", sum)
}

