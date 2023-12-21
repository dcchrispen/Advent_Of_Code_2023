package utils

import (
	"fmt"
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

