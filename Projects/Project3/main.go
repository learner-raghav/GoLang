package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"errors"
	"fmt"
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
func handleGetEmployee(res http.ResponseWriter, req *http.Request){

	if req.Method != "GET" {
		fmt.Fprintf(res,"Incorrect Method")
		return
	}

	dbObject := DBObject{}
	dbObject.connectToDB()
	defer dbObject.db.Close()

	employeeId,err := strconv.Atoi(req.FormValue("id"))
	emp := Employee{EmployeeId: employeeId}
	if err != nil {
		fmt.Fprintf(res, "Invalid Id Format")
		return
	}
	employee,err := GetById(dbObject,emp)
	if err != nil {
		fmt.Fprintf(res,err.Error())
		return
	}

	res.Header().Set("Content-Type","application/json")
	json.NewEncoder(res).Encode(employee)

}

func handleCreateEmployee(res http.ResponseWriter, req *http.Request){

	if req.Method != "POST" {
		fmt.Fprintf(res,"Incorrect Method")
		return
	}

	dbObject := DBObject{}
	dbObject.connectToDB()
	defer dbObject.db.Close()
	emp := Employee{}
	emp.Name = req.FormValue("name")
	emp.Email = req.FormValue("email")
	emp.Phone = req.FormValue("phone")
	emp.RoleId, _ = strconv.Atoi(req.FormValue("roleId"))
	fmt.Println(emp)
	err := CreateEmployee(dbObject,emp)
	if err != nil {
		fmt.Fprintf(res,err.Error())
		return
	}

	fmt.Fprintf(res,"Created Employee Successfully.")
}

func handleUpdateEmployee(res http.ResponseWriter, req *http.Request) {

	if req.Method != "PUT" {
		fmt.Fprintf(res,"Incorrect Method")
		return
	}

	dbObject := DBObject{}
	dbObject.connectToDB()
	defer dbObject.db.Close()

	emp := Employee{}
	var err error
	emp.EmployeeId ,err= strconv.Atoi(req.FormValue("id"))

	if err != nil {
		fmt.Fprintf(res,"Id type not proper")
	}

	emp.Name = req.FormValue("name")
	emp.Email = req.FormValue("email")
	emp.Phone = req.FormValue("phone")
	emp.RoleId, err = strconv.Atoi(req.FormValue("roleId"))
	if err != nil {
		fmt.Fprintf(res,"Role Id not proper")
	}
	err = UpdateEmployee(dbObject,emp)
	if err != nil {
		fmt.Fprintf(res,err.Error())
	} else{
		fmt.Fprintf(res,"Employee Updated Successfully!!")
	}
}

func handleDeleteEmployee(res http.ResponseWriter, req * http.Request){
	if req.Method != "DELETE" {
		fmt.Fprintf(res,"Incorrect Method")
		return
	}
	dbObject := DBObject{}
	dbObject.connectToDB()
	defer dbObject.db.Close()

	employeeId,err := strconv.Atoi(req.FormValue("id"))
	if err != nil {
		fmt.Fprintf(res,err.Error())
		return
	}
	emp := Employee{EmployeeId: employeeId}
	err = DeleteEmployee(dbObject,emp)

	if err != nil {
		fmt.Fprintf(res,err.Error())
		return
	}

	fmt.Fprintf(res,"Employee Deleted Successfully")

}

func main(){

	mux := http.NewServeMux()
	mux.HandleFunc("/getEmployeeById",handleGetEmployee)
	mux.HandleFunc("/create",handleCreateEmployee)
	mux.HandleFunc("/update",handleUpdateEmployee)
	mux.HandleFunc("/delete",handleDeleteEmployee)

	http.ListenAndServe("0.0.0.0:3000",mux)
}
