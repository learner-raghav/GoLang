/**
	This is a part of the store layer - It just interacts with the Database.
 */
package EmployeeStore

import (
	"database/sql"
	"errors"
)
import "../../entity"
type EmployeeStore struct {
	db *sql.DB //should be unexported.
}

func New(db *sql.DB) EmployeeStore {
	return EmployeeStore{db: db}
}

func (empStore EmployeeStore) GetById(id int) (entity.Employee,error){
	db := empStore.db
	if db == nil {
		return entity.Employee{},errors.New("DB not configured properly")
	}
	viewQuery, err := db.Query("SELECT * FROM employee where employeeId = ?",id)
	if err != nil {
		return entity.Employee{},err
	}
	var emp entity.Employee
	if viewQuery.Next() {
		var empId,roleId int
		var fullName,phone,email string
		err = viewQuery.Scan(&empId,&fullName,&email,&phone,&roleId)
		if err != nil {
			return entity.Employee{},err
		}
		emp = entity.Employee{EmployeeId: empId,FullName: fullName,Email: email,Phone: phone,RoleId: roleId}

	} else{
		err := errors.New("EmployeeStore record does not exist")
		return entity.Employee{},err
	}
	return emp, nil
}

func (empStore EmployeeStore) Create(employee entity.Employee) (entity.Employee,error){
	db := empStore.db

	if db == nil {
		return entity.Employee{}, errors.New("DB not configured properly")
	}

	insertQuery, err := db.Prepare("INSERT INTO employee(name,email,phone,roleId) values(?,?,?,?)")
	if err != nil {
		return entity.Employee{},err
	}
	_, error := insertQuery.Exec(employee.FullName,employee.Email,employee.Phone,employee.RoleId)

	if error != nil {
		return entity.Employee{},error
	}
	return employee,nil
}

func (emp EmployeeStore) Update(employee entity.Employee) (entity.Employee,error){
	db := emp.db
	if db == nil {
		return entity.Employee{},errors.New("DB not configured properly")
	}
	updateQuery, err := db.Prepare("UPDATE employee SET name=?,email=?,phone=?,roleId=? where employeeid = ?")
	if err != nil {
		return entity.Employee{}, err
	}
	res1,err := updateQuery.Exec(employee.FullName,employee.Email,employee.Phone,employee.RoleId,employee.EmployeeId)
	if err != nil {
		error := errors.New("Error occured")
		return entity.Employee{},error
	}
	rowsAffected,_ := res1.RowsAffected()

	if rowsAffected == 0 {
		return entity.Employee{},err
	}
	return employee,nil
}

func (emp EmployeeStore) Delete(employeeId int) (entity.Employee,error){

	db := emp.db
	employee, err := emp.GetById(employeeId)

	if err != nil {
		return entity.Employee{},errors.New(err.Error())
	}
	if db == nil {
		return entity.Employee{},errors.New("DB not configured properly")
	}

	deleteStatement,err := db.Prepare("DELETE from employee where employeeid=?")
	if err != nil {
		return entity.Employee{}, err
	}

	result,err := deleteStatement.Exec(employee.EmployeeId)
	if err != nil {
		return entity.Employee{},err
	}
	rowsAffected,_ := result.RowsAffected()
	if rowsAffected == 0 {
		err := errors.New("The id does not exist")
		return entity.Employee{},err
	}
	return employee,nil
}