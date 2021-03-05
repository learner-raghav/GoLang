package entity

type Employee struct {
	EmployeeId int `json:"employee_id"`
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	RoleId int `json:"role_id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}
