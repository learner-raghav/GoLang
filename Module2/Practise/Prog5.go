package  main

import "time"
import "fmt"

func main(){
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Day()," ",t.Month()," ",t.Year())

	var a int = 10
	var intP *int
	intP = &a
	fmt.Println(intP)
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(*intP)

	*intP = 20
	fmt.Println(intP)
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(*intP)

	const i = 5
	// ptr1 := &i - This is worng. We cannot take the address of a literal or a constant
	// fmt.Println(*ptr1)



}