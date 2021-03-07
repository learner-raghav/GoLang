package main

import "fmt"

func describe1(i interface{}){
	switch i.(type) {
	case int:
		fmt.Println("Integer")
	case string:
		fmt.Println("String")
	default:
		fmt.Println("Unknown")
	}
}
func main(){
	var i interface{}
	describe1(i)

	i = 42
	describe1(i)

	i = "Raghav"
	describe1(i)
}

/**
	type interface Reader {
		func (T) Read(b []byte) (n int,err error)
	}
 */