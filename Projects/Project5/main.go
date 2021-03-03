package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type api struct {
	db *sql.DB
}

type Employee struct {
	EmployeeId int
	Name string
	Email string
	Phone string
	RoleId int
}

func (a *api) fail(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "application/json")

	data := struct {
		Error string
	}{Error: msg}

	resp, _ := json.Marshal(data)
	w.WriteHeader(status)
	w.Write(resp)
}

func (a *api) ok(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	resp, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.fail(w, "oops something evil has happened", 500)
		return
	}

	w.Write(resp)
}

func (a *api) getEmployeeById(w http.ResponseWriter,req *http.Request){
	val,err := strconv.Atoi(req.FormValue("id"))
	rows, err := a.db.Query("SELECT (.+) FROM sales WHERE id = ?",val)
	if err != nil {
		fmt.Println("")
		a.fail(w, "failed to fetch posts: "+err.Error(), 500)
		return
	}
	defer rows.Close()

	var emp Employee
	if rows.Next() {
		var empId,roleId int
		var name,phone,email string
		err = rows.Scan(&empId,&name,&email,&phone,&roleId)
		if err != nil {
			a.fail(w, "failed to scan employee: "+err.Error(), 500)
			return
		}
		emp = Employee{empId,name,email,phone,roleId}
		fmt.Println(emp)
	} else {
		a.fail(w,"Record does not exist",500)
	}

	if rows.Err() != nil {
		a.fail(w, "failed to read all posts: "+rows.Err().Error(), 500)
		return
	}

	a.ok(w,emp)
}

func (a *api) createEmployee(w http.ResponseWriter, req *http.Request){
	db := a.db
	if db == nil {
		a.fail(w, "DB not configured properly", 500)
		return
	}
	employee := Employee{}
	err := json.NewDecoder(req.Body).Decode(&employee)

	insertQuery, err := db.Prepare("INSERT INTO employee(employeeId,name,email,phone,roleId) values(?,?,?,?,?)")
	if err != nil {
		fmt.Println("hiii",err.Error())
		a.fail(w,err.Error(),500)
	}
	_, error := insertQuery.Exec(employee.EmployeeId,employee.Name,employee.Email,employee.Phone,employee.RoleId)

	if error != nil {
		a.fail(w,error.Error(),500)
	}

	a.ok(w,employee)

}

func (a *api) UpdateEmployee(w http.ResponseWriter,req *http.Request) {
	db := a.db

	if db == nil {
		a.fail(w,"DB not configured properly",500)
	}
	empId,_ := strconv.Atoi(req.FormValue("id"))
	employee := Employee{}
	err := json.NewDecoder(req.Body).Decode(&employee)
	employee.EmployeeId = empId

	updateQuery, err := db.Prepare("UPDATE employee SET name=?,email=?,phone=?,roleId=? where employeeid = ?")

	if err != nil {
		a.fail(w,err.Error(),500)
	}

	res1,_ := updateQuery.Exec(employee.Name,employee.Email,employee.Phone,employee.RoleId,employee.EmployeeId)
	rowsAffected,_ := res1.RowsAffected()

	if rowsAffected == 0 {
		error := errors.New("No modifications done")
		a.fail(w,error.Error(),500)
	}
	a.ok(w,employee)
}

func (a *api) deleteEmployee(w http.ResponseWriter,req *http.Request){
	db := a.db
	empId,_ := strconv.Atoi(req.FormValue("id"))
	if db == nil {
		a.fail(w,"DB not configured properly",500)
	}

	deleteStatement,err := db.Prepare("DELETE from employee where employeeid=?")
	if err != nil {
		a.fail(w,err.Error(),500)
	}

	result,err := deleteStatement.Exec(empId)
	if err != nil {
		a.fail(w,err.Error(),500)
	}
	rowsAffected,_ := result.RowsAffected()
	if rowsAffected == 0 {
		err := errors.New("The id does not exist")
		a.fail(w,err.Error(),500)
	}
	op := struct{
		msg string
	}{"SUCCESS"}
	a.ok(w,op)

}