package main

type Employee struct {
	EmployeeId int `json:"employee_id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	RoleId int `json:"role_id"`
}
