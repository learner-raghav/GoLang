package main

import "fmt"

func main1(){
	a := 10
	b := a //It will create a copy of the variable
	b = 20
	fmt.Println(&a)
	fmt.Println(&b)
	fmt.Println(a)
	fmt.Println(b)

	c := []int {1,2,3,4,5}
	d := c //it creates a reference to the array in the memory

	fmt.Println(c)
	d[0] = 2
	fmt.Println(c)
	fmt.Println(d)

	var complex1 complex128 = 3 + 2i
	fmt.Println(complex1)
	fmt.Println(real(complex1))

	var charType uint8 = 'A'
	fmt.Println(string(charType))
}