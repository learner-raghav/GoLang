package main

import "fmt"

/**
	Function used as a value in go
	In golang ,,amy a time we might be assigning functions to variables.
 */

//This is how we can use functions as arguments.
func useFunctionAsParameter(x int,f func(int,int) int) int{
	return x + f(2,3)
}
func main(){
	//Here we basically write an anonymous function
	add := func(a int,b int) int{
		return a + b
	}

	//Here basically the function is assigned to the add variable.
	add2And3 := add(2,3)
	fmt.Println(add2And3)

	add2And3And4 := useFunctionAsParameter(4,add)
	fmt.Println(add2And3And4)
}
