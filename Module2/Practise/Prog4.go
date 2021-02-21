package main


import (
	"strings"
	"fmt"
	"strconv"
)
/**
	1. Strings are a basic data structure and every language has a predefined function for manipulating strings.
	   In Go, these are gathered in a package.
	
	2. 
*/

func main5(){

	// strings.HasPrefix(str,prefix) function checks if the string has a specified prefix.
	input := "Raghav is a fun loving boy!"
	ans1 := strings.HasPrefix(input,"Rag")
	ans2 := strings.HasPrefix(input,"rag")

	fmt.Println(ans1)
	fmt.Println(ans2)


	//strings.HasSuffix is a function that checks if a string has a particular suffix or not
	input2 := "Hello World!"
	ans3 := strings.HasSuffix(input2,"ld!")
	ans4 := strings.HasSuffix(input2,"l!")
	fmt.Println(ans3)
	fmt.Println(ans4)

	//strings.Contains(s) this is used to check whether a string has a particular substring or not.
	input3 := "Raghav loves to cook and he is a cook"
	ans5 := strings.Contains(input3,"cook")
	fmt.Println(ans5)

	//checking the indes of a pafrticular substring
	ans6 := strings.Index(input3,"cook")
	fmt.Println(ans6)

	ans7 := strings.LastIndex(input3,"cook")
	fmt.Println(ans7)

	//Replacing a substring, -1 as n implies all occureneces are replaced.
	input4 := "raghav mohan rahul tanisha nimish rohit"
	input4 = strings.Replace(input4," ",",",-1)
	fmt.Println(input4)

	//Counting occurences of a substring - We simply have the Count function that is used to count the occurences of 
	// a substring.

	input5 := "Raghav is an amazing cook and he is a cook and he wants to be a cook for his wife"
	ans8 := strings.Count(input5,"cook")
	fmt.Println(ans8)

	input6 := input
	input6 = strings.ToUpper(input6)
	fmt.Println(input6)

	input6 = strings.ToLower(input6)
	fmt.Println(input6)

	input1 := "Hello "
	ans9 := strings.Repeat(input1,10)
	fmt.Println(ans9)

	//If you want to trim all the leading and trailing white spaces, all we need to do is
	input7 := "   Hello world   "
	input7 = strings.TrimSpace(input7)
	fmt.Println(input7)

	//Suppose we want to split the string atw hitespaces
	splittedString := strings.Fields(input)
	fmt.Println(splittedString)

	//suppose we wnat to join this string back
	ans10 := strings.Join(splittedString," ")
	fmt.Println(ans10)

	var orig string  = "666"
	var an int
	var newS string


	fmt.Printf("The size of the int is %d\n",strconv.IntSize)
	//converting ASCII to integer
	an, _ = strconv.Atoi(orig)
	fmt.Printf("The converted value is %d\n",an)

	an = an + 5
	//converting integer to ASCII
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is %s\n",newS)

}