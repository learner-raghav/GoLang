package main

import "fmt"

type I interface {
	Work()
}

type Plumber struct {
	hrs int
}

type Carpenter struct {
	hrs int
}

func (p *Plumber) Work(){
	fmt.Println(p.hrs)
}

func (c Carpenter) Work(){
	fmt.Println(c.hrs)
}
//Calling a method on a nil interface is a run type error.
func describe(i I){
	i.Work()
	fmt.Printf("%v %T",i,i)
	fmt.Println()
}

func main(){
	var c Carpenter
	var p Plumber

	c = Carpenter{hrs: 10}
	p = Plumber{hrs: 20}

	describe(c)
	describe(&p)
}

