package service

import "../entity"

type EmployeeServicer interface {
	GetById(id int) (entity.Employee,error)
	Create(emp entity.Employee) (entity.Employee,error)
	Update(emp entity.Employee) (entity.Employee,error)
	Delete(employeeId int) (entity.Employee,error)
}
