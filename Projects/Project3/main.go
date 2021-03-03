package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
)

/**
	Type definitions
 */

type DatabaseHandler interface {
	connectToDB()
}

type DBObject struct {
	db *sql.DB
	err error
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

/**
	Methods needed for testing.
 */

func GetById(object DBObject,employee Employee) (Employee,error){
	db := object.db
	if db == nil {
		return Employee{},errors.New("DB not configured properly")
	}

	viewQuery, err := db.Query("SELECT * FROM employee where employeeid = ?",employee.EmployeeId)

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
	return emp,nil
}


func CreateEmployee(object DBObject,employee Employee) error {
	db := object.db

	if db == nil {
		return errors.New("DB not configured properly")
	}

	insertQuery, err := db.Prepare("INSERT INTO employee(name,email,phone,roleId) values(?,?,?,?)")
	if err != nil {
		return err
	}
	_, error := insertQuery.Exec(employee.Name,employee.Email,employee.Phone,employee.RoleId)

	if error != nil {
		return error
	}
	return nil
}

func UpdateEmployee(object DBObject, employee Employee) error {
	db := object.db
	if db == nil {
		return errors.New("DB not configured properly")
	}

	updateQuery, err := db.Prepare("UPDATE employee SET name=?,email=?,phone=?,roleId=? where employeeid = ?")
	if err != nil {
		return err
	}
	res1,_ := updateQuery.Exec(employee.Name,employee.Email,employee.Phone,employee.RoleId,employee.EmployeeId)
	rowsAffected,_ := res1.RowsAffected()
	if rowsAffected == 0 {
		error := errors.New("No modifications done")
		return error
	}
	return nil
}

func DeleteEmployee(object DBObject,employee Employee) error {
	db := object.db

	if db == nil {
		return errors.New("DB not configured properly")
	}

	deleteStatement,err := db.Prepare("DELETE from employee where employeeid=?")
	if err != nil {
		return err
	}

	result,err := deleteStatement.Exec(employee.EmployeeId)
	if err != nil {
		return err
	}
	rowsAffected,_ := result.RowsAffected()
	if rowsAffected == 0 {
		err := errors.New("The id does not exist")
		return err
	}
	return nil

}

/**
	Handlers
 */

func (dbObject *DBObject) handleEmployee(res http.ResponseWriter, req *http.Request){
	switch req.Method {
		case "GET":
			dbObject.handleGetEmployee(res, req)
		case "POST":
			dbObject.handleCreateEmployee(res,req)
		case "PUT":
			dbObject.handleUpdateEmployee(res,req)
		case "DELETE":
			dbObject.handleDeleteEmployee(res,req)
	}
}

func (dbObject *DBObject) handleGetEmployee(res http.ResponseWriter, req *http.Request){

	employeeId,err := strconv.Atoi(req.FormValue("id"))
	emp := Employee{EmployeeId: employeeId}
	if err != nil {
		fmt.Fprintf(res, "Invalid Id Format")
		return
	}
	employee,err := GetById(*dbObject,emp)
	if err != nil {
		fmt.Fprintf(res,err.Error())
		return
	}

	res.Header().Set("Content-Type","application/json")
	json.NewEncoder(res).Encode(employee)

}

func (dbObject *DBObject) handleCreateEmployee(res http.ResponseWriter, req *http.Request){

	emp := Employee{}
	err := json.NewDecoder(req.Body).Decode(&emp)
	if err != nil {
		fmt.Fprintf(res,"Error UnMarshalling the Body")
	}

	err = CreateEmployee(*dbObject,emp)
	if err != nil {
		fmt.Fprintf(res,err.Error())
		return
	}

	fmt.Fprintf(res,"Created Employee Successfully.")
}

func (dbObject *DBObject) handleUpdateEmployee(res http.ResponseWriter, req *http.Request) {

	employeeId,err := strconv.Atoi(req.FormValue("id"))

	if err != nil {
		fmt.Fprintf(res, "Invalid Id Format")
		return
	}

	var emp Employee
	err = json.NewDecoder(req.Body).Decode(&emp)
	if err != nil {
		fmt.Fprintf(res,"Error UnMarshalling the Body")
	}
	emp.EmployeeId = employeeId //Assigning the id to the object.
	err = UpdateEmployee(*dbObject,emp)
	if err != nil {
		fmt.Fprintf(res,err.Error())
	} else{
		fmt.Fprintf(res,"Employee Updated Successfully!!")
	}
}

func (dbObject *DBObject) handleDeleteEmployee(res http.ResponseWriter, req * http.Request){

	employeeId,err := strconv.Atoi(req.FormValue("id"))
	if err != nil {
		fmt.Fprintf(res,err.Error())
		return
	}
	emp := Employee{EmployeeId: employeeId}
	err = DeleteEmployee(*dbObject,emp)

	if err != nil {
		fmt.Fprintf(res,err.Error())
		return
	}

	fmt.Fprintf(res,"Employee Deleted Successfully")

}

func main(){

	var dbObject DBObject
	dbObject.connectToDB()
	defer dbObject.db.Close()

	mux := http.NewServeMux()

	mux.Handle("/employee",http.HandlerFunc(dbObject.handleEmployee))
	http.ListenAndServe("0.0.0.0:3000",mux)
}
