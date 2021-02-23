package main

import "fmt"

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main(){
	fmt.Println(Abs(-10))
	LABEL1:
		for i:=0;i<5;i++ {
			for j:=0 ; j < 5; j++ {
				fmt.Printf("i: %d , j:%d\n",i,j)
				if j == 4{
					//Continue label1 does not do the re initialization again
					//Goto label reinitializes the entire code.
					// A better practise is to use go to whenever label is after the current line goto is being used.
					//goto LABEL1
					continue LABEL1
				}
			}
		}

	//LABEL2:
		fmt.Println("Hello World")
}