package main

import (
	"fmt"
	"strconv"
)

/**
	1. What is Go programming language ?

	Ans: Go is an open source programming language, developed by Google, It is also called as Go
		lang. This language is primarily used for system programming. It is very easy to build simple,
		efficient and reliable software in Golang

	2. What is the default value of type bool in Go programming language ?

	Ans: The default value of bool variable in Go programming language is false

	3. What are the advantagges/benefits of the Go programming language ?

	Ans: Go is fast and compiles very quickly
		It supports concurrency at the language level.
		It has Garbage collection
		It supports various safety features and CSP style concurrent programming language.
		Strings and maps are built in into the language
		Functions are first class objects in this language.

	4. Structs, int, float, string, bool are by default passed by value.
	5. map is by default passed by reference
	6. In Go, Interfaces is a way to identify the behaviour of an object. An interface is created
		by using the type word, followed by the name of the type and the keyword interface.
	7. Type assertion in golang takes an interface value and retrieves from it a value
		of a specified explicit type
	8. In methods, we have receivers that are basically used to call the functions defined on a type/

 */
type Ninja struct{
	name,gadget string
	power int
}
func changeValue(n *int){
	*n = 20
}

const (
	i = 7
	j
	k
)

type ReaderWriter interface{
	Writer(name string,age int,occupation string,salary float64)
	Reader() string
	Change(name string)
}

type Person struct {
	name string
	age int
	occupation string
	salary float64
}

func (p *Person) Reader() string {
	return p.name + "," + strconv.Itoa(p.age)
}

func (p *Person) Writer(name string,age int,occupation string,salary float64){
	p.age = age
	p.name = name
	p.occupation = occupation
	p.salary = salary
}

func (p *Person) Change(name string){
	p.name = name
}

func main(){

	var a bool //false
	var b int //0
	var c float64 //0
	var d string // ""
	var m map[string]int //This is nil
	//The zero value will be the zero value of the individual types of the struct
	var s struct{
		name string
		age int
	}
	fmt.Println(m == nil)
//	m["abc"] = 1 //This will be an error telling assign,ent to entry in nil map

	// In make specifying map size does not really matter!
	//We should not try to read from or write to a zero values map.
	//It is necessary to implement all the methods for implementing an interface
	var m1 = make(map[string]int,10)
	fmt.Println(m1,len(m1))
	m1["Hello"] = 1
	m1["Wow"] = 2
	fmt.Println(m1,len(m1))
	fmt.Println(a,b,c,d,m,s)

	//var n = 10
	//fmt.Println(n)
	//changeValue(&n)
	//fmt.Println(n)



	ninja := Ninja{name: "Raghav",gadget: "Bey Blade"}
	fmt.Println(ninja)
	change(&ninja)
	fmt.Println(ninja)

	var temp = make(map[int]string)
	temp[1] = "Raghav"
	fmt.Println(temp)
	change1(temp)
	fmt.Println(temp)
	fmt.Println(i,j,k)

	var t ReaderWriter
	fmt.Printf("Type: %T",t)
	var p Person //As soon as we defined it, we can see it implements the reader and the writer methods.
	t = &p
	fmt.Printf("Type : %T",t)
	// Basically implies it implements the ReaderWriter interface.
	p.Writer("Raghav",23,"Student",250000)
	fmt.Println(p.Reader())
	p.Change("")

	var rd ReaderWriter
	fmt.Println(rd)

	switch t.(type) {
	case *Person:
		fmt.Println("Hello")

	}

	myFunc("S")
	myFunc(0.0)

}
func change1(m map[int]string){
	m[2] = "Rahul"
}
//All types implement the empty intyerface
func change(n *Ninja){
	n.gadget = "Albatross!"
	fmt.Println(*n)
}

func myFunc(a interface{}){
	switch a.(type) {
	case int:
		fmt.Println("Type: int",a.(int))
	case string:
		fmt.Println("Type: string",a.(string))
	default:
		fmt.Println("Type: Default")

	}
}