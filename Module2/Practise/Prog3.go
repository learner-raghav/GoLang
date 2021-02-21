package main

import "fmt"


//We can also define type aliases for this input
var Celcius float32
var Fahrenheit float32

func toFahrenheit(celTemp float64) float64 {
	f := (9*celTemp)/5 + 32
	return f
}

func main4(){
	celcius := 112.00
	fahrenheit := toFahrenheit(celcius)
	fmt.Println(fahrenheit)	
}