package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func timer(name string) func() {
    start := time.Now()
    return func() {
        fmt.Printf("%s took %v\n", name, time.Since(start))
    }
}

type Card struct {
	numberOfCards int
	winnindNums []int
	possibleNums []int
}

var cards []Card

func addCard(card Card) {
	cards = append(cards, card)
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

func main(){
	defer timer("main")()
	file, err := os.Open("../input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	for scanner.Scan(){
		var line = scanner.Text()
		// get nums and store in two arrays
		numbers := strings.Split(line, ":")[1]
		winningNumStrings := strings.Split(numbers, "|")[0]
		possibleNumStrings := strings.Split(numbers, "|")[1]

		re := regexp.MustCompile(`\d+`)
		winningArr := convertStringsToNums(re.FindAllString(winningNumStrings, -1))
		possibleArr := convertStringsToNums(re.FindAllString(possibleNumStrings, -1))

		winningArr = sortNums(winningArr)
		possibleArr = sortNums(possibleArr)


		addCard(Card{numberOfCards: 1, winnindNums: winningArr, possibleNums: possibleArr})

	}

	for i:= 0 ; i < len(cards); i++ {
		currCard := cards[i]
		matches := getMatchCount(currCard.winnindNums, currCard.possibleNums)

		for j := 0 ; j < currCard.numberOfCards; j++ {
			for k := 0 ; k < matches ; k++ {
				cards[i+k+1].numberOfCards ++
			}
		}

	}

	sum:=0 
	for i:= 0 ; i < len(cards); i++ {
		sum = sum + cards[i].numberOfCards
	}
	fmt.Printf("Final Sum: %d  \n", sum)

}

func getMatchCount(winningArr []int, possibleArr []int) int {
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

	return matchCount

}