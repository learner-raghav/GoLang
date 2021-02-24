package main

import "fmt"

func receiveMarks(mark1 int,marks ...int){
	for i:=0;i<len(marks);i++ {
		fmt.Println(marks[i])
	}
}

func receiveNames(name1 string,name2 string,name3 ...string) {
	fmt.Println(name1)
	fmt.Println(name2)
	//The name3 here is received as an array and it contains all the
	// remaining parameter values.
	for i:=0;i<len(name3);i++{
		fmt.Println(name3[i])
	}
	receiveMarks(100,93,45,56,76,78)
}

func main(){
	receiveNames("Raghav","Sunil","Akash","Praveen","Pushpraj")
}
