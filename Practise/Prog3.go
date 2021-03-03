package main

import (
	"fmt"
	"sort"
	"unsafe"
)

type Orange struct {
	Quantity int
}

func (o *Orange) Increase(n int) {
	o.Quantity += n
}

func (o *Orange) Decrease(n int) {
	o.Quantity -= n
}

func (o *Orange) String() string {
	return fmt.Sprintf("%v", o.Quantity)
}

func main() {
	orange := &Orange{}
	fmt.Println(orange)
	orange.Increase(10)
	fmt.Println(orange)
	orange.Decrease(5)
	fmt.Println(orange,orange.Quantity)

	fruits := map[string]int{
		"oranges": 100,
		"apples":  200,
		"bananas": 300,
	}

	a := []string{}
	for k,_ := range fruits {
		a = append(a,k)
	}
	fmt.Println(a)
	sort.Strings(a)
	fmt.Println(a)

	var n1 []int
	n2 := []int{}
	fmt.Println(unsafe.Sizeof(n1),unsafe.Sizeof(n2))
	fmt.Println(n1)
	n1 = append(n1,2)
	fmt.Println(n1)
	fmt.Println(n2)
	n2 = append(n2,2)
	fmt.Println(n2)

	m3 := make(map[string]int)
	m3["h"] = 23

	if ans,ok := m3["h"]; ok {
		fmt.Println(ans)
	}

	m4 := m3
	m3["g"] = 34
	fmt.Println(m3,m4)
	//To copy a map in go, we have to create new maps and then assign key value pairs one by one
	//Slices are passed by refrrnce, arrays are passed by value
	v := []int{1,2,3,4,5}
	fmt.Println(v)
	reverse(v)
	fmt.Println(v)
}

func reverse(a []int){
	for i:=0;i<len(a)/2;i++ {
		a[i],a[len(a)-i-1] = a[len(a)-i-1],a[i]
	}
}

/**
	1. structs - Structs are user defined types, They serve very similar purpose to classes
		in other programming languages.
	2. Methods - Methods are functions that operate on particular types.
	3. Encapsulation  - Exported and unexported fields.
	4. Polymorphism - A variable of type interface can hold any value which implements the
		interface. Thisproperty of interfaces is used to achieve polymorphism in Golang
	5. rune type in golang -

 */