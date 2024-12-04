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

func InsertIntoSliceAtIndex[T any](destination []T, element T, index int) []T {
	if len(destination) == index {
		return append(destination, element)
	}

	destination = append(destination[:index+1], destination[index:]...) // index < len(a)
	destination[index] = element

	return destination
}

func AbsDiffInt(x, y int) int {
	if x < y {
	   return y - x
	}
	return x - y
 }
 
 func FilterOutAtIndex(arr []int, index int) []int {
    // Create a new slice with the same capacity as the original array
    result := make([]int, 0, len(arr))

    // Iterate over the original array
    for i, v := range arr {
        // If the current index is not the one to exclude, append to the result
        if i != index {
            result = append(result, v)
        }
    }

    return result
}

