package main

import (
	"testing"
)

func TestApis(t *testing.T){
	dbObject := DBObject{}
	dbObject.connectToDB()
	defer dbObject.db.Close()
	for i,tc := range testCases {
		t.Run(tc.desc,func(t *testing.T){

			if tc.route == "/getEmployeeById" {
				emp,err := GetById(dbObject,tc.employee)
				if (err != nil) && (err.Error() != tc.output.err.Error() || tc.output.employee != emp){
					t.Errorf("Error Occured")
				}
			} else if tc.route == "/create" {
				err := CreateEmployee(dbObject,tc.employee)
				if (err != nil) && (err.Error() != tc.output.err.Error()) {
					t.Errorf(err.Error())
				}
			} else if tc.route == "/update" {
				err := UpdateEmployee(dbObject,tc.employee)
				if (err != nil) && (err.Error() != tc.output.err.Error()) {
					t.Errorf(err.Error())
				}
			} else if tc.route == "/delete" {
				err := DeleteEmployee(dbObject,tc.employee)
				if (err != nil) && (err.Error() != tc.output.err.Error()) {
					t.Errorf(err.Error())
				}
			} else{
				t.Errorf("Invalid Test ID: %d",i)
			}
		})
	}
}

