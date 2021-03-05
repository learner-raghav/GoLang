package store

import "../entity"

type EmployeeStoreHandler interface{
	GetById(id int) (entity.Employee,error)
	Create(employee entity.Employee) (entity.Employee,error)
	Update(employee entity.Employee) (entity.Employee,error)
	Delete(id int) (entity.Employee,error)
}

