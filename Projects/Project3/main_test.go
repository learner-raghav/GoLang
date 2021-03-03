package main

import (
	"os"
	"testing"
)
var dbObject DBObject

type createOutput struct {
	employee Employee
	err error
}

type testStruct struct{
	desc string
	route string
	method string
	dbObj DBObject
	employee Employee
	output createOutput
}

var testCases []testStruct


func TestMain(m *testing.M){

	dbObject.connectToDB()
	defer dbObject.db.Close()

	testCases = getTestCases(dbObject)

	//If success - returns 0, else returns something else.
	executionCode := m.Run()
	os.Exit(executionCode)
}

func TestApis(t *testing.T){

	for i,tc := range testCases {
		t.Run(tc.desc,func(t *testing.T){
			if tc.route == "/getEmployeeById" {
				emp,err := GetById(tc.dbObj,tc.employee)

				if (err != nil) && (err.Error() != tc.output.err.Error() || tc.output.employee != emp){
					t.Errorf(err.Error())
				}
			} else if tc.route == "/create" {
				err := CreateEmployee(tc.dbObj,tc.employee)
				if (err != nil) && (err.Error() != tc.output.err.Error()) {
					t.Errorf(err.Error())
				}
			} else if tc.route == "/update" {
				err := UpdateEmployee(tc.dbObj,tc.employee)
				if (err != nil) && (err.Error() != tc.output.err.Error()) {
					t.Errorf(err.Error())
				}
			} else if tc.route == "/delete" {
				err := DeleteEmployee(tc.dbObj,tc.employee)
				if (err != nil) && (err.Error() != tc.output.err.Error()) {
					t.Errorf(err.Error())
				}
			} else{
				t.Errorf("Invalid Test ID: %d",i)
			}
		})
	}
}

