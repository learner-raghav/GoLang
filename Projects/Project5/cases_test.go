package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func (a *api) assertJSON(actual []byte, data interface{}, t *testing.T) {
	expected, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}
	fmt.Println(string(expected),string(actual))
	if bytes.Compare(expected, actual) != 0 {
		t.Errorf("the expected json: %s is different from actual %s", expected, actual)
	}
}


func TestCreateEmployee(t *testing.T){
	db,mock,err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error occured : %v",err)
	}
	defer db.Close()

	app := &api{db}
	data := Employee{
		EmployeeId: 1,
		Name:       "Raghav",
		Email:      "raghav@ZopSmart.com",
		Phone:      "8384852943",
		RoleId:     4,
	}
	body, _ := json.Marshal(data)
	req := httptest.NewRequest("POST","localhost:3000/EmployeeService",bytes.NewBuffer(body))

	w := httptest.NewRecorder()

	mock.ExpectPrepare("INSERT INTO EmployeeService(employeeId,name,email,phone,roleId) values(?,?,?,?,?)").
		ExpectExec().
		WithArgs(1,data.Name,data.Email,data.Phone,data.RoleId).
		WillReturnResult(sqlmock.NewResult(1,1))


	app.createEmployee(w,req)
	app.assertJSON(w.Body.Bytes(),data,t)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("All expectations were not fulfilled: %v",err)
	}
}

func TestShouldGetEmployee(t *testing.T){
	db,mock,err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error occured : %v",err)
	}
	defer db.Close()

	app := &api{db}
	req,err := http.NewRequest("GET","http://localhost/getEmployeeById?id="+strconv.Itoa(34),nil)

	if err != nil {
		t.Fatalf("Error Occured: %v",err)
	}

	w := httptest.NewRecorder()

	rows := sqlmock.NewRows([]string{"employeeId","name","email","phone","roleId"}).
		AddRow(34,"Raghav","raghav@ZopSmart.com","8384852943",4)

	mock.ExpectQuery("SELECT (.+) FROM sales WHERE id = ?").
		WithArgs(34).
		WillReturnRows(rows)

	data := Employee{
		EmployeeId: 34,
		Name:       "Raghav",
		Email:      "raghav@ZopSmart.com",
		Phone:      "8384852943",
		RoleId:     4,
	}
	app.getEmployeeById(w,req)

	app.assertJSON(w.Body.Bytes(),data,t)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("All expectations were not fulfilled: %v",err)
	}
}

func TestUpdateEmployee(t *testing.T){
	db,mock,err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error occured : %v",err)
	}
	defer db.Close()

	app := &api{db}
	data := Employee{
		EmployeeId: 1,
		Name:       "Raghav",
		Email:      "raghav@ZopSmart.com",
		Phone:      "8384852943",
		RoleId:     4,
	}
	body, _ := json.Marshal(data)
	req := httptest.NewRequest("POST","localhost:3000/EmployeeService?id="+strconv.Itoa(1),bytes.NewBuffer(body))

	w := httptest.NewRecorder()

	mock.ExpectPrepare("UPDATE EmployeeService SET name=?,email=?,phone=?,roleId=? where employeeid = ?").
		ExpectExec().
		WithArgs(data.Name,data.Email,data.Phone,data.RoleId,1).
		WillReturnResult(sqlmock.NewResult(1,1))


	app.UpdateEmployee(w,req)
	app.assertJSON(w.Body.Bytes(),data,t)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("All expectations were not fulfilled: %v",err)
	}
}

func TestDeleteEmployee(t *testing.T){
	db,mock,err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error occured : %v",err)
	}
	defer db.Close()

	app := &api{db}

	req := httptest.NewRequest("POST","localhost:3000/EmployeeService?id="+strconv.Itoa(1),nil)

	w := httptest.NewRecorder()

	mock.ExpectPrepare("DELETE from EmployeeService where employeeid=?").
		ExpectExec().
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1,1))


	app.deleteEmployee(w,req)
	op := struct{
		msg string
	}{"SUCCESS"}
	app.assertJSON(w.Body.Bytes(),op,t)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("All expectations were not fulfilled: %v",err)
	}
}