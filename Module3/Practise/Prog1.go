package main

import (
	"fmt"
)

func main1() {
	var n int = 18

	if n%15 == 0 {
		fmt.Println("FizzBuzz")
	} else if n%5 == 0 {
		fmt.Println("Fizz")
	} else{
		fmt.Println("Buzz")
	}


	if a:=40; a < 20 {
		fmt.Println("Hello world")
	} else if a < 30 {
		fmt.Println("Hello World by Raghav")
	} else {
		fmt.Println("Namaskara Bangalore!!")
	}

	//if you try accessing a here, it will be an error.
	// fmt.Println(a)
}