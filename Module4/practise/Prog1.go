package main

import "fmt"

func calc(x,y int) (sum1 int,product1 int)  {
	sum1 = x+y
	product1 = x*y

	return
}

func main(){

	x := 10
	y := 20

	sum,product := calc(x,y)
	fmt.Println(sum," ",product)
}
