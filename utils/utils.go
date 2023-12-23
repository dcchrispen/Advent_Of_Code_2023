package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
func ConvertToInt(num string) int {

	numInt, err := strconv.Atoi(num)
	Check(err)
	return numInt
}

func ConvertStringsToNums(stringArr []string) []int {
	var numArr []int
	for i := 0; i < len(stringArr); i++ {
		num, err := strconv.Atoi(stringArr[i])
		Check(err)
		numArr = append(numArr, num)
	}
	return numArr
}

func SortNums(numArr []int) []int {
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

func GetNumArrayFromString(str string) []int {
	var numArr []int
	re := regexp.MustCompile(`(\d+)`)
	numStrings := re.FindAllString(str, -1)
	for i := 0; i < len(numStrings); i++ {
		numArr = append(numArr, ConvertToInt(string(numStrings[i])))
	}
	return numArr
}

func FindMod(a int) int {
	mod :=0
	for i := 1; i < 100000; i++ {
		if(a % i == 0){
			mod = i
			
		}
	}	
	return mod
}