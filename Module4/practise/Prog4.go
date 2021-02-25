package main

import "fmt"

/**
	Basically the task is to write recursive fibonacci in Golang
 */

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1{
		return 1
	}
	return fib(n-1)+fib(n-2)
}

func factorial(n int) int {
	if n <= 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func main(){
	n := 10
	fib10 := fib(n)
	fmt.Println(fib10)

	n2 := 5
	factorial := factorial(n2)
	fmt.Println(factorial)
}