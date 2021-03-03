package main

import (
	"encoding/json"
	"fmt"
)

/**
	JSON - Javascript object Notation
	Marshal - Generating a JSON string from a data structure
	UnMarshal - The act of parsing JSON to a data structure.

	JSON is a language independent data format. The golang is providing `encoding/json` package for
	json related operation. It has many inbuilt methods for processing json data.

	Marshal used to convert Go object to json and Unmarchal is vice versa

 */

type Employee struct {
	Name string `json:"name"`
	Phone string `json:"phone"`
	Email string  `json:"email"`
	Id int `json:"id"`
	Age int `json:"age"`
}

func main(){
	employee := Employee{Id: 1,Name: "Raghav",Age: 23,Email: "raghav.ddps2@gmail.com",Phone: "8384852943"}

	//Marshalling
	emp, _ := json.Marshal(employee)
	fmt.Println(string(emp))

	//Unmarshalling, we take the byte data and connvert it to string'
	var employee1 Employee
	_ = json.Unmarshal(emp, &employee1)
	fmt.Println(employee1)

}
