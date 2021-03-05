package main

import (
	"./driver"
	"log"
	"./entity"
	"./service/EmployeeService"
	"./store/EmployeeStore"
	"./delivery/EmployeeDelivery"
	"net/http"
)
func main(){
	conf := entity.MySQLConfig{
		DbName: "testDB",
		DbUser: "raghav",
		DbPass: "raghav@123M",
	}

	db,err := driver.ConnectToDB(conf)
	if err != nil {
		log.Println("Could not connect to server ",err)
		return
	}
	empStore := EmployeeStore.New(db)
	empService := EmployeeService.New(empStore)
	handler := EmployeeDelivery.New(empService)

	http.HandleFunc("/employee",handler.Handle)
	log.Println("Starting server at PORT 8000")
	log.Fatal(http.ListenAndServe(":8000",nil))
}
