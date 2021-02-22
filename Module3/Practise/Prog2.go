package main

import (
	"fmt"
)

func main() {


	var1 := 100
	switch {

	case var1 == 100:
		fmt.Println("Value if less than 100")
		fmt.Println("Hola")

	case var1 == 100:
		fmt.Println("Value is equal to 100")
	
	case var1 > 100:
		fmt.Println("Value is nore than 100")
	
	default:
		fmt.Println("Lol")
	
	}

}