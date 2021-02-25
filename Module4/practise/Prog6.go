package main

import "fmt"

/**
	Now, we will use a filter function
 */

//We define the usual signature of the function
type flt func(int) bool

//Filter1 - isOdd
func isOdd(n int) bool {
	if n%2 != 0 {
		return true
	}
	return false
}

//Filter2 - isEven
func isEven(n int) bool {
	if n%2 == 0 {
		return true
	} else{
		return false
	}
}

func filter(slice []int,filter flt)	[]int {
	var res[] int
	//We check on a specific filter.
	for _,val := range slice {
		if filter(val) {
			res = append(res,val)
		}
	}
	return res
}

func main(){
	slice := []int{1,2,3,4,5,6,7}
	fmt.Println(slice)
	e := []int{}
	e = append(e,1)
	fmt.Println(e)
	odd := filter(slice,isOdd)
	even := filter(slice,isEven)

	fmt.Println("Even: ",even)
	fmt.Println("Odd: ",odd)
}