package main

import "fmt"

//defer keyword basically is used to execute a statement after the function returns.
//If there are multiple defer calls they are executed in LIFO (Stack) fashion.

var k int
func main(){
	fmt.Println(k)
	hello()
	fmt.Println(k)
}

func hello(){

	for i:=0 ; i<=5;i++ {
		defer  fmt.Println("Hello World", k); k += 2 //only one statement deferred
		fmt.Println(k)
	}
}