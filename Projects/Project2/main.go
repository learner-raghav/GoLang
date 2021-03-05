package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
)
type Employee struct {
	EmployeeId int
	Name string
	Email string
	Phone string
	RoleId int
}

/**
		Purpose : Setting up the DB Connection and returning the sql db object.
		sql.Open : https://golang.org/pkg/database/sql/#Open
		Open may not necessarily connect to DB, instead it might just validate the arguments.
		Data source name format : DSN Format : username:password@protocol(address)/dbname

		The documentation says the sql package should be used in conjunction with the database drivers.

*/
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "raghav"
	dbPass := "raghav@123M"
	dbName := "testDB"

	db, err := sql.Open(dbDriver,dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db

}
/**
	db.Query - Basically takes the Query and returns the rows and error
	Next prepares the next result row for reading with the Scan option.
	Scan copies the column values into variables specified
 */
func ViewAllEmployees(w http.ResponseWriter, r *http.Request){

	db := dbConn()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM EmployeeService ORDER BY employeeId DESC")

	if err != nil {
		panic(err.Error())
	}

	emp := Employee{}
	employees := []Employee{}
	for rows.Next() {
		var id,roleId int
		var name,email,phone string
		err := rows.Scan(&id,&name,&email,&phone,&roleId)
		if err != nil {
			panic(err.Error())
		}

		emp.EmployeeId = id
		emp.Name = name
		emp.Email = email
		emp.Phone = phone
		emp.RoleId = roleId
		employees = append(employees,emp)
	}

	fmt.Println(employees)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(employees)
}

/**
	This is to insert an EmployeeService in the database
 */
func Insert(res http.ResponseWriter,req *http.Request){
	db := dbConn()
	defer db.Close()
	name := req.FormValue("name")
	email := req.FormValue("email")
	phone := req.FormValue("phone")
	roleId, _ := strconv.Atoi(req.FormValue("roleid"))

	insertQuery, err := db.Prepare("INSERT INTO EmployeeService(name,email,phone,roleId) values(?,?,?,?)")
	if err != nil {
		fmt.Println("Error occured")
	}
	insertQuery.Exec(name,email,phone,roleId)
	fmt.Fprintf(res,"Inserted EmployeeService record ino the Database")
}

func ViewAnEmployee(res http.ResponseWriter,req *http.Request){
	db := dbConn()
	defer db.Close()
	employeeId := req.FormValue("id")
	viewQuery, err := db.Query("SELECT * FROM EmployeeService where employeeid = ?",employeeId)
	emp := Employee{}
	if viewQuery.Next() {
		var empId,roleId int
		var name,phone,email string
		err = viewQuery.Scan(&empId,&name,&email,&phone,&roleId)
		if err != nil {
			fmt.Println("Error")
		}
		emp = Employee{empId,name,email,phone,roleId}
	} else{
		fmt.Fprintf(res,"The EmployeeService record does not exist!")
		return
	}
	res.Header().Set("Content-Type","application/json")
	json.NewEncoder(res).Encode(emp)
}

func updateAnEmployee(res http.ResponseWriter,req *http.Request){
	db := dbConn()
	defer db.Close()
	id ,_:= strconv.Atoi(req.FormValue("id"))
	name := req.FormValue("name")
	email := req.FormValue("email")
	phone := req.FormValue("phone")
	roleId, _ := strconv.Atoi(req.FormValue("roleId"))

	updateQuery, err := db.Prepare("UPDATE EmployeeService SET name=?,email=?,phone=?,roleId=? where employeeid = ?")
	if err != nil {
		fmt.Println("An error occured!!")
	}
	res1,_ := updateQuery.Exec(name,email,phone,roleId,id)
	rowsAffected,_ := res1.RowsAffected()
	if rowsAffected == 0 {
		fmt.Println("This record does not exist")
		fmt.Fprintf(res,"Record does not exist")
		return
	}

	fmt.Fprintf(res,"Record updated successfully!!")

}


func deleteAnEmployee(res http.ResponseWriter,req *http.Request){
	db := dbConn()
	defer db.Close()
	idToDelete,_ := strconv.Atoi(req.FormValue("id"))

	deleteStatement,err := db.Prepare("DELETE from EmployeeService where employeeid=?")
	if err != nil {
		fmt.Println("Some error occured")
	}

	result,err := deleteStatement.Exec(idToDelete)

	if rowsAffected,_ := result.RowsAffected(); rowsAffected == 0 {
		fmt.Fprintf(res,"Cannot delete or does not exist")
		return
	}

	fmt.Fprintf(res,"Deleted the EmployeeService with id %d",idToDelete)
}

/**
	Select 1/all - db.Query, stmt.Next, stmt.Scan
	Insert/Update/Delete - db.Prepare and then stmt.Exec

 */

func home(res http.ResponseWriter,req *http.Request){
	fmt.Fprintf(res,"Hello World\n")
}

func main(){
	log.Println("Server started at PORT 5000")

	mux := http.NewServeMux()

	mux.HandleFunc("/",home)
	mux.HandleFunc("/employees",ViewAllEmployees)
	mux.HandleFunc("/insert",Insert)
	mux.HandleFunc("/EmployeeService",ViewAnEmployee)
	mux.HandleFunc("/update",updateAnEmployee)
	mux.HandleFunc("/delete",deleteAnEmployee)
	http.ListenAndServe("0.0.0.0:3000",mux)
}
