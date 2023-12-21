package main

import (
	"bufio"
	"fmt"
	"os"
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

func main(){
	defer timer("main")()

	file, err := os.Open("../input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		var line = scanner.Text()
		fmt.Printf("Line: %s  \n", line)
	}

}