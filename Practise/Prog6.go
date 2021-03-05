package main

import "fmt"

type Vehicle struct {
	company string
	model string
}

type Car struct {
	Vehicle //This implies embedding!
	airbags string
}

func (v *Vehicle) details(){
	fmt.Println("The Vehicle is of company "+v.company+" and the model is "+v.model)
}

//func (c *Car) details(){
//	fmt.Println("The car has "+c.airbags+" airbags")
//}

func main(){
	hondaCity := Car{airbags: "4"}
	hondaCity.model = "2032"
	hondaCity.company = "Honda"
	hondaCity.details()
}