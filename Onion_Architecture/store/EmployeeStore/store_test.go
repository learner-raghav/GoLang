package EmployeeStore

import (
	"../../entity"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)
type Output struct{
	emp entity.Employee
	err error
}
func TestShouldGetEmployee(t *testing.T){
	var testCases = []struct{
		desc string
		input entity.Employee
		output Output
	}{
		{
			desc : "EmployeeStore Exists",
			input: entity.Employee{
				EmployeeId: 34,
			},
			output: Output{
				entity.Employee{
					EmployeeId: 34,
					FullName: "Raghav",
					Email: "raghav@ZopSmart.com",
					Phone: "8384852943",
					RoleId: 4,
				},
				errors.New("Success"),
			},
		},
		{
			desc: "EmployeeStore doesn't exist",
			input: entity.Employee{
				EmployeeId: 31,
			},
			output: Output{
				entity.Employee{
				},
				errors.New("EmployeeStore record does not exist"),
			},
		},
	}

	db,mock,err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error occured : %v",err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"employeeId","name","email","phone","roleId"}).
		AddRow(34,"Raghav","raghav@ZopSmart.com","8384852943",4)

	mock.ExpectQuery("SELECT * FROM EmployeeService where employeeId = ?").
		WithArgs(34).
		WillReturnRows(rows)
	empStore := New(db)
	for _,tc := range testCases {
		t.Run(tc.desc,func(t *testing.T){
			emp,error := empStore.GetById(tc.output.emp.EmployeeId)
			if tc.output.emp != emp && err != nil && error.Error() != tc.output.err.Error(){
				t.Errorf(err.Error()+"Expected: %v, Found: %v",tc.output.emp,emp)
			}
		})
	}
}

func TestShouldCreateEmployee(t *testing.T){
	var testCases = []struct{
		desc string
		input entity.Employee
		output Output
	}{
		{
			desc : "EmployeeStore Exists",
			input: entity.Employee{
				FullName:       "Raghav",
				Email:      "raghav@ZopSmart.com",
				Phone:      "8384852943",
				RoleId:     4,
			},
			output: Output{
				entity.Employee{
					FullName:  "Raghav",
					Email:      "raghav@ZopSmart.com",
					Phone:      "8384852943",
					RoleId:     4,
				},
				errors.New("Success"),
			},
		},
		{
			desc: "No modifications done",
			input: entity.Employee{
				EmployeeId: 2,
			},
			output: Output{
				entity.Employee{
				},
				errors.New("Error occured"),
			},
		},
	}
	db,mock,err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	empStore := New(db)
	if err != nil {
		t.Fatalf("Error occured : %v",err)
	}
	defer db.Close()

	for _,tc := range testCases {
		t.Run(tc.desc,func(t *testing.T){
			prepared := mock.ExpectPrepare("INSERT INTO EmployeeService(name,email,phone,roleId) values(?,?,?,?)").
				ExpectExec().
				WithArgs(tc.input.FullName,tc.input.Email,tc.input.Phone,tc.input.RoleId)

			if tc.output.err.Error() != "Success"{
				prepared.WillReturnError(errors.New("Error occured"))
			} else{
				prepared.WillReturnResult(sqlmock.NewResult(1,1))
			}
			emp,error := empStore.Create(tc.input)
			if tc.output.emp != emp && err != nil && error.Error() != tc.output.err.Error(){
				t.Errorf(err.Error()+"Expected: %v, Found: %v",tc.output.emp,emp)
			}
		})

	}
}

func TestShouldUpdateEmployee(t *testing.T){
	var testCases = []struct{
		desc string
		input entity.Employee
		output Output
	}{
		{
			desc : "EmployeeStore Exists",
			input: entity.Employee{
				EmployeeId: 1,
				FullName: "Raghavs",
				Email: "raghavs@ZopSmart.com",
				Phone: "8384852943",
				RoleId: 4,
			},
			output: Output{
				entity.Employee{
					EmployeeId: 1,
					FullName: "Raghavs",
					Email: "raghavs@ZopSmart.com",
					Phone: "8384852943",
					RoleId: 4,
				},
				errors.New("Success"),
			},
		},
		{
			desc: "No modifications done",
			input: entity.Employee{
				EmployeeId: 2,
			},
			output: Output{
				entity.Employee{
				},
				errors.New("Error occured"),
			},
		},
	}
	db,mock,err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	empStore := New(db)
	if err != nil {
		t.Fatalf("Error occured : %v",err)
	}
	defer db.Close()

	for _,tc := range testCases {
		t.Run(tc.desc,func(t *testing.T){
			prepared := mock.ExpectPrepare("UPDATE EmployeeService SET name=?,email=?,phone=?,roleId=? where employeeid = ?").
				ExpectExec().
				WithArgs(tc.input.FullName,tc.input.Email,tc.input.Phone,tc.input.RoleId,tc.input.EmployeeId)

			if tc.output.err.Error() != "Success"{
				prepared.WillReturnError(errors.New("Error occured"))
			} else{
				prepared.WillReturnResult(sqlmock.NewResult(0,1))
			}
			emp,error := empStore.Update(tc.input)

			if tc.output.emp != emp && err != nil && error.Error() != tc.output.err.Error(){
				t.Errorf(err.Error()+"Expected: %v, Found: %v",tc.output.emp,emp)
			}
		})
	}
}
//
func TestShouldDeleteEmployee(t *testing.T){
	var testCases = []struct{
		desc string
		input entity.Employee
		output Output
	}{
		{
			desc : "EmployeeStore Exists",
			input: entity.Employee{
				EmployeeId: 34,
			},
			output: Output{
				entity.Employee{
					EmployeeId: 34,
					FullName: "Raghav",
					Email: "raghav@ZopSmart.com",
					Phone: "8384852943",
					RoleId: 4,
				},
				errors.New("Success"),
			},
		},
		{
			desc: "EmployeeStore doesn't exist",
			input: entity.Employee{
				EmployeeId: 31,
			},
			output: Output{
				entity.Employee{
				},
				errors.New("Record does not exist"),
			},
		},
	}
	rows := sqlmock.NewRows([]string{"employeeId","name","email","phone","roleId"}).
		AddRow(34,"Raghav","raghav@ZopSmart.com","8384852943",4)
	db,mock,err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error occured : %v",err)
	}
	defer db.Close()
	empStore := New(db)
	mock.ExpectQuery("SELECT * FROM EmployeeService where employeeId = ?").
		WithArgs(34).
		WillReturnRows(rows)

	mock.ExpectPrepare("DELETE from EmployeeService where employeeid=?").
		ExpectExec().
		WithArgs(34).
		WillReturnResult(sqlmock.NewResult(1,1))

	mock.ExpectQuery("SELECT * FROM EmployeeService where employeeId = ?").
		WithArgs(31).
		WillReturnError(errors.New("Record does not exist"))

	mock.ExpectPrepare("DELETE from EmployeeService where employeeid=?").
		ExpectExec().
		WithArgs(31).
		WillReturnResult(sqlmock.NewResult(1,1))


	for _,tc := range testCases{
		t.Run(tc.desc,func(t *testing.T){
			emp,error := empStore.Delete(tc.input.EmployeeId)
			if tc.output.emp != emp && err != nil && error.Error() != tc.output.err.Error(){
				t.Errorf(err.Error()+"Expected: %v, Found: %v",tc.output.emp,emp)
			}
		})
	}
}
//
