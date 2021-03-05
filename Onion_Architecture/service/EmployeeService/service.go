package EmployeeService

import (
	"../../entity"
	"../../service"
	"../../store"
	"fmt"
	"strings"
)
type EmpService struct {
	empStore store.EmployeeStoreHandler
}

func New(empStore store.EmployeeStoreHandler) service.EmployeeServicer {
	return EmpService{empStore: empStore}
}

func (empService EmpService) GetById(id int) (entity.Employee,error) {

	employee,err := empService.empStore.GetById(id)

	//Business Logic
	if err == nil {
		splitNames := strings.Split(employee.FullName," ")
		employee.FirstName = splitNames[0]
		employee.LastName = splitNames[1]
	}

	return employee,err

}

func (empService EmpService) Create(employee entity.Employee) (entity.Employee,error){

	//Business Logic
	if employee.LastName != "" && employee.FirstName != "" {
		employee.FullName = employee.FirstName + " " + employee.LastName
	}
	emp,err := empService.empStore.Create(employee)
	return emp,err
}

func (empService EmpService) Delete(employeeId int) (entity.Employee,error){

	fmt.Println("Service Layer: No function for delete")
	emp,err := empService.empStore.Delete(employeeId)
	return emp,err
}

func (empService EmpService) Update(employee entity.Employee) (entity.Employee,error){

	//Business Logic
	if employee.LastName != "" && employee.FirstName != "" {
		employee.FullName = employee.FirstName + " " + employee.LastName
	}
	emp,err := empService.empStore.Update(employee)
	return emp,err
}


