package main

import "fmt"

func main(){

	//This is the condition controlled loop.
	var str string = "Raghav is a boy"
	for i:=0;i<len(str);i++ {
		fmt.Printf("%c",str[i])
	}
	fmt.Println()

	//Another construct that comes is the condition controled loop
	i := 0
	for i < len(str) {
		fmt.Printf("%c",str[i])
		i += 1
	}
	fmt.Println()

	//In go lang we have another kind of loops, they are nothing but the for range loops.
	for _,j := range str {
		fmt.Printf("%c",j)
	}
	fmt.Println()
	j  := 0
	for {
		if j == 5 {
			//We need to do j += 1 here because it will continue and
			//wont execute the remaining statements in the loop.
			j += 1
			continue
		} else if j == 9 {
			break
		} else{
			fmt.Println("Hello!!")
		}
		j += 1
	}
}