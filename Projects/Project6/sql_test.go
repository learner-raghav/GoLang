package main

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)
type Output struct{
	emp Employee
	err error
}
func TestShouldGetEmployee(t *testing.T){
	var testCases = []struct{
		desc string
		input Employee
		output Output
	}{
		{
			desc : "Employee Exists",
			input: Employee{
				EmployeeId: 34,
			},
			output: Output{
				Employee{
					EmployeeId: 34,
					Name: "Raghav",
					Email: "raghav@ZopSmart.com",
					Phone: "8384852943",
					RoleId: 4,
				},
				errors.New("Success"),
			},
		},
		{
			desc: "Employee doesn't exist",
			input: Employee{
				EmployeeId: 31,
			},
			output: Output{
				Employee{
				},
				errors.New("Employee record does not exist"),
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

	mock.ExpectQuery("SELECT * FROM employee where employeeId = ?").
		WithArgs(34).
		WillReturnRows(rows)

	for _,tc := range testCases {
		t.Run(tc.desc,func(t *testing.T){
			emp,error := GetById(DBObject{db: db},Employee{EmployeeId: tc.output.emp.EmployeeId})

			if tc.output.emp != emp && error.Error() != tc.output.err.Error(){
				t.Errorf(err.Error())
			}
			if err := mock.ExpectationsWereMet(); err != nil && error != nil{
				t.Errorf("All expectations were not fulfilled: %v",err)
			}
		})
	}
}

func TestShouldCreateEmployee(t *testing.T){
	var testCases = []struct{
		desc string
		input Employee
		output Output
	}{
		{
			desc : "Employee Exists",
			input: Employee{
				Name:       "Raghav",
				Email:      "raghav@ZopSmart.com",
				Phone:      "8384852943",
				RoleId:     4,
			},
			output: Output{
				Employee{
					Name:       "Raghav",
					Email:      "raghav@ZopSmart.com",
					Phone:      "8384852943",
					RoleId:     4,
				},
				errors.New("Success"),
			},
		},
		{
			desc: "No modifications done",
			input: Employee{
				EmployeeId: 2,
			},
			output: Output{
				Employee{
				},
				errors.New("Error occured"),
			},
		},
	}
	db,mock,err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf("Error occured : %v",err)
	}
	defer db.Close()

	for _,tc := range testCases {
		t.Run(tc.desc,func(t *testing.T){
			prepared := mock.ExpectPrepare("INSERT INTO employee(name,email,phone,roleId) values(?,?,?,?)").
				ExpectExec().
				WithArgs(tc.input.Name,tc.input.Email,tc.input.Phone,tc.input.RoleId)

			if tc.output.err.Error() != "Success"{
				prepared.WillReturnError(errors.New("Error occured"))
			} else{
				prepared.WillReturnResult(sqlmock.NewResult(1,1))
			}
			emp,error := CreateEmployee(DBObject{db: db},tc.input)
			if emp != tc.output.emp || error.Error() != tc.output.err.Error(){
				t.Errorf(err.Error())
			}
		})

	}
}

func TestShouldUpdateEmployee(t *testing.T){
	var testCases = []struct{
		desc string
		input Employee
		output Output
	}{
		{
			desc : "Employee Exists",
			input: Employee{
				EmployeeId: 1,
				Name: "Raghavs",
				Email: "raghavs@ZopSmart.com",
				Phone: "8384852943",
				RoleId: 4,
			},
			output: Output{
				Employee{
					EmployeeId: 1,
					Name: "Raghavs",
					Email: "raghavs@ZopSmart.com",
					Phone: "8384852943",
					RoleId: 4,
				},
				errors.New("Success"),
			},
		},
		{
			desc: "No modifications done",
			input: Employee{
				EmployeeId: 2,
			},
			output: Output{
				Employee{
				},
				errors.New("Error occured"),
			},
		},
	}
	db,mock,err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf("Error occured : %v",err)
	}
	defer db.Close()

	for _,tc := range testCases {
		t.Run(tc.desc,func(t *testing.T){
			prepared := mock.ExpectPrepare("UPDATE employee SET name=?,email=?,phone=?,roleId=? where employeeid = ?").
				ExpectExec().
				WithArgs(tc.input.Name,tc.input.Email,tc.input.Phone,tc.input.RoleId,tc.input.EmployeeId)

			if tc.output.err.Error() != "Success"{
				prepared.WillReturnError(errors.New("Error occured"))
			} else{
				prepared.WillReturnResult(sqlmock.NewResult(0,1))
			}
			emp,error := UpdateEmployee(DBObject{db: db},tc.input)

			if emp != tc.output.emp || error.Error() != tc.output.err.Error(){
				t.Errorf(err.Error())
			}
		})

	}

}

func TestShouldDeleteEmployee(t *testing.T){
	var testCases = []struct{
		desc string
		input Employee
		output Output
	}{
		{
			desc : "Employee Exists",
			input: Employee{
				EmployeeId: 34,
			},
			output: Output{
				Employee{
					EmployeeId: 34,
					Name: "Raghav",
					Email: "raghav@ZopSmart.com",
					Phone: "8384852943",
					RoleId: 4,
				},
				errors.New("Success"),
			},
		},
		{
			desc: "Employee doesn't exist",
			input: Employee{
				EmployeeId: 31,
			},
			output: Output{
				Employee{
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

	mock.ExpectQuery("SELECT * FROM employee where employeeId = ?").
		WithArgs(34).
		WillReturnRows(rows)

	mock.ExpectPrepare("DELETE from employee where employeeid=?").
		ExpectExec().
		WithArgs(34).
		WillReturnResult(sqlmock.NewResult(1,1))

	mock.ExpectQuery("SELECT * FROM employee where employeeId = ?").
		WithArgs(31).
		WillReturnError(errors.New("Record does not exist"))

	mock.ExpectPrepare("DELETE from employee where employeeid=?").
		ExpectExec().
		WithArgs(31).
		WillReturnResult(sqlmock.NewResult(1,1))


	for _,tc := range testCases{
		t.Run(tc.desc,func(t *testing.T){
			emp,error := DeleteEmployee(DBObject{db: db},Employee{EmployeeId: tc.input.EmployeeId})
			if emp != tc.output.emp || error.Error() != tc.output.err.Error() {
				t.Errorf(err.Error())
			}
		})
	}
}

