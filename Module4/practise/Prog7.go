package main

import "fmt"

func main(){
	counter := func() (func() int ){
		x := 1
		y := 2
		//We can maintain the state of the above variables using closures!!
		fmt.Println(x,y)
		return func() int {
			z := x + y
			x = y
			y = z
			return z
		}
	}

	/**
		SQL Adapter in Go
		Database/sql
	 */
	//Here we are doing our function initialization and then we will just cal that method.
	count := counter()

	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())

}
