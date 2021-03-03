package main

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

type DatabaseHandler interface {
	connectToDB()
}

type DBObject struct {
	db *sql.DB
	err error
}

type Employee struct {
	EmployeeId int `json:"employee_id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	RoleId int `json:"role_id"`
}


func (dbObject *DBObject) connectToDB(){
	dbDriver := "mysql"
	dbUser := "raghav"
	dbPass := "raghav@123M"
	dbName := "testDB"

	dbObject.db,dbObject.err = sql.Open(dbDriver,dbUser+":"+dbPass+"@/"+dbName)

	if dbObject.err != nil {
		panic(dbObject.err.Error())
	}
}

func GetById(object DBObject,employee Employee) (Employee,error){
	db := object.db
	if db == nil {
		return Employee{},errors.New("DB not configured properly")
	}
	viewQuery, err := db.Query("SELECT * FROM employee where employeeId = ?",employee.EmployeeId)
	if err != nil {
		return Employee{},err
	}

	emp := Employee{}

	if viewQuery.Next() {
		var empId,roleId int
		var name,phone,email string
		err = viewQuery.Scan(&empId,&name,&email,&phone,&roleId)
		if err != nil {
			return Employee{},err
		}
		emp = Employee{empId,name,email,phone,roleId}
	} else{
		err := errors.New("Employee record does not exist")
		return Employee{},err
	}
	return emp, errors.New("Success")
}

func CreateEmployee(object DBObject,employee Employee) (Employee, error) {
	db := object.db

	if db == nil {
		return Employee{}, errors.New("DB not configured properly")
	}

	insertQuery, err := db.Prepare("INSERT INTO employee(name,email,phone,roleId) values(?,?,?,?)")
	if err != nil {
		return Employee{},err
	}
	_, error := insertQuery.Exec(employee.Name,employee.Email,employee.Phone,employee.RoleId)

	if error != nil {
		return Employee{},error
	}
	return employee,errors.New("Success")
}

func UpdateEmployee(object DBObject, employee Employee) (Employee,error) {
	db := object.db
	if db == nil {
		return Employee{},errors.New("DB not configured properly")
	}
	updateQuery, err := db.Prepare("UPDATE employee SET name=?,email=?,phone=?,roleId=? where employeeid = ?")
	if err != nil {
		return Employee{}, err
	}
	res1,err := updateQuery.Exec(employee.Name,employee.Email,employee.Phone,employee.RoleId,employee.EmployeeId)
	if err != nil {
		error := errors.New("Error occured")
		return Employee{},error
	}
	rowsAffected,_ := res1.RowsAffected()

	if rowsAffected == 0 {
		return Employee{},err
	}
	return employee,errors.New("Success")
}

func DeleteEmployee(object DBObject,employee Employee) (Employee,error) {

	db := object.db
	emp, err := GetById(object,employee)

	if err.Error() != "Success" {
		return Employee{},errors.New(err.Error())
	}
	if db == nil {
		return Employee{},errors.New("DB not configured properly")
	}

	deleteStatement,err := db.Prepare("DELETE from employee where employeeid=?")
	if err != nil {
		return Employee{}, err
	}

	result,err := deleteStatement.Exec(employee.EmployeeId)
	if err != nil {

		return Employee{},err
	}
	rowsAffected,_ := result.RowsAffected()
	if rowsAffected == 0 {
		err := errors.New("The id does not exist")
		return Employee{},err
	}
	return emp,errors.New("Success")
}

