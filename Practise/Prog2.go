package main

import (
	"fmt"
	"unsafe"
)

func main(){
	a := struct{}{}
	fmt.Println(unsafe.Sizeof(a),a) // size is 0
	b := map[string]struct{}{}
	fmt.Println(unsafe.Sizeof(b),b)

	c := [5]int{1,2,3}
	c[3] = 2
	fmt.Println(len(c),cap(c),c)
	//c = append(c,10)
	fmt.Println(len(c),cap(c),c)

	d := make([]int,5,10)
	fmt.Println(d)
	d = append(d,10)
	fmt.Println(d)

	e := []int{1,2,3}
	f := e
	f[1] = 3
	fmt.Println(e,f,&e[0],&f[0])



	g := []int{1, 2}
	h := []int{3, 4}
	check := g
	fmt.Println(&check[0],&g[0],&h[0])
	g = h
	fmt.Println(&check[0],&g[0],&h[0])
	fmt.Println(g,h,check)

	//We can compare two structs with the == operator, as we would do with other simple types.
	// We need to make sure that they do not contain slices, maps, functions etc,


}
