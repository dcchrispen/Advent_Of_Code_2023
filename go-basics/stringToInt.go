package main

import (
	"fmt"
	"strconv"
)



func main(){
	var myNum string ="41"

	myInt, err := strconv.Atoi(myNum)

	fmt.Println(myInt)

	if(err != nil){
		fmt.Println("Error")
	}


}