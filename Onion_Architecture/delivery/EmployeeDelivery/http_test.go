package EmployeeDelivery

import (
	"../../entity"
	"../../service/EmployeeService"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)


type mockDataStore struct{}

func (m mockDataStore) GetById(id int) (entity.Employee,error){
	if id == 1{
		return entity.Employee{
			EmployeeId: 1,
			FullName:   "Raghav Maheshwari",
			Email:      "raghav@ZopSmart.com",
			Phone:      "8384852943",
			RoleId:     0,
			FirstName:  "Raghav",
			LastName:   "Maheshwari",
		},nil
	} else{
		return entity.Employee{},errors.New("Id does not exist")
	}
}

func (m mockDataStore) Create(emp entity.Employee) (entity.Employee,error) {
	temp := entity.Employee{}
	if emp == temp {
		return entity.Employee{}, errors.New("Data inserted is empty")
	} else {
		return entity.Employee{
			EmployeeId: 1,
			FullName:   "Raghav Maheshwari",
			Email:      "raghav@ZopSmart.com",
			Phone:      "8384852943",
			RoleId:     0,
			FirstName:  "Raghav",
			LastName:   "Maheshwari",
		},nil
	}
}

func (m mockDataStore) Update(emp entity.Employee) (entity.Employee,error) {
	temp := entity.Employee{}
	if emp == temp {
		return entity.Employee{}, errors.New("The data updated is empty")
	} else {
		return emp,nil
	}
}

func (m mockDataStore) Delete(id int) (entity.Employee,error) {
	if id == 1 {
		return entity.Employee{
			EmployeeId: 1,
			FullName:   "Raghav Maheshwari",
			Email:      "raghav@ZopSmart.com",
			Phone:      "8384852943",
			RoleId:     0,
			FirstName:  "Raghav",
			LastName:   "Maheshwari",
		},nil
	} else {
		return entity.Employee{},errors.New("The Id does not exist")
	}
}

func TestEmployeeGet(t *testing.T){
	testCases := []struct{
		desc string
		id int
		output response
	}{
		{
			"Id exists",
			1,
			response{Emp: entity.Employee{
				EmployeeId: 1,
				FullName:   "Raghav Maheshwari",
				Email:      "raghav@ZopSmart.com",
				Phone:      "8384852943",
				RoleId:     0,
				FirstName:  "Raghav",
				LastName:   "Maheshwari",
			},Msg: "Success"},
		},
		{
			"Id does not exist",
			2,
			response{Emp: entity.Employee{},Msg: "Id does not exist"},
		},
	}

	for _,tc := range testCases {
		t.Run(tc.desc,func(t *testing.T){
			a := strconv.Itoa(tc.id)
			req := httptest.NewRequest("GET","/employee?id="+a,nil)
			res := httptest.NewRecorder()

			emp := New(mockDataStore{})
			emp.getById(res,req)
			var op response
			json.NewDecoder(res.Body).Decode(&op)
			if tc.output.Emp != op.Emp || tc.output.Msg != op.Msg {
				t.Fatalf("Test Failed: Expected: %v, Received: %v",tc.output,op)
			}
		})
	}
}

func TestEmployeeCreate(t *testing.T){
	testCases := []struct{
		desc string
		emp entity.Employee
		output response
	}{
		{
			"Data inserted is empty",
			entity.Employee{},
			response{Emp: entity.Employee{},Msg: "Data inserted is empty"},
		},
		{
			"Successful insertion",
			entity.Employee{
				EmployeeId: 1,
				FullName:   "Raghav Maheshwari",
				Email:      "raghav@ZopSmart.com",
				Phone:      "8384852943",
				RoleId:     0,
				FirstName:  "Raghav",
				LastName:   "Maheshwari",
			},
			response{Emp: entity.Employee{
				EmployeeId: 1,
				FullName:   "Raghav Maheshwari",
				Email:      "raghav@ZopSmart.com",
				Phone:      "8384852943",
				RoleId:     0,
				FirstName:  "Raghav",
				LastName:   "Maheshwari",
			},Msg: "Success"},
		},
	}

	for _,tc := range testCases {
		t.Run(tc.desc,func(t *testing.T){
			body, _ := json.Marshal(tc.emp)
			req := httptest.NewRequest("POST","/employee",bytes.NewBuffer(body))
			res := httptest.NewRecorder()

			emp := New(mockDataStore{})
			emp.create(res,req)
			var op response
			json.NewDecoder(res.Body).Decode(&op)
			if tc.output.Emp != op.Emp || tc.output.Msg != op.Msg {
				t.Fatalf("Test Failed: Expected: %v, Received: %v",tc.output,op)
			}
		})
	}
}

func TestEmployeeUpdate(t *testing.T){
	testCases := []struct{
		desc string
		emp entity.Employee
		output response
	}{
		{
			"Data updated is empty",
			entity.Employee{},
			response{Emp: entity.Employee{},Msg: "The data updated is empty"},
		},
		{
			"Successful updation",
			entity.Employee{
				EmployeeId: 1,
				FullName:   "Raghav Maheshwari",
				Email:      "raghav@ZopSmart.com",
				Phone:      "8384852943",
				RoleId:     0,
				FirstName:  "Raghav",
				LastName:   "Maheshwari",
			},
			response{Emp: entity.Employee{
				EmployeeId: 1,
				FullName:   "Raghav Maheshwari",
				Email:      "raghav@ZopSmart.com",
				Phone:      "8384852943",
				RoleId:     0,
				FirstName:  "Raghav",
				LastName:   "Maheshwari",
			},Msg: "Employee Updated Successfully!!"},
		},
	}

	for _,tc := range testCases {
		t.Run(tc.desc,func(t *testing.T){
			body, _ := json.Marshal(tc.emp)
			req := httptest.NewRequest("POST","/employee?id="+strconv.Itoa(tc.emp.EmployeeId),bytes.NewBuffer(body))
			res := httptest.NewRecorder()

			emp := New(mockDataStore{})
			emp.update(res,req)
			var op response
			json.NewDecoder(res.Body).Decode(&op)

			if tc.output.Emp != op.Emp || tc.output.Msg != op.Msg {
				t.Fatalf("Test Failed: Expected: %v, Received: %v",tc.output,op)
			}
		})
	}
}

func TestEmployeeDelete(t *testing.T){
	testCases := []struct{
		desc string
		id int
		output response
	}{
		{
			"Id exists",
			1,
			response{Emp: entity.Employee{
				EmployeeId: 1,
				FullName:   "Raghav Maheshwari",
				Email:      "raghav@ZopSmart.com",
				Phone:      "8384852943",
				RoleId:     0,
				FirstName:  "Raghav",
				LastName:   "Maheshwari",
			},Msg: "Employee Deleted Successfully"},
		},
		{
			"Id does not exist",
			2,
			response{Emp: entity.Employee{},Msg: "The Id does not exist"},
		},
	}

	for _,tc := range testCases {
		t.Run(tc.desc,func(t *testing.T){
			a := strconv.Itoa(tc.id)
			req := httptest.NewRequest("GET","/employee?id="+a,nil)
			res := httptest.NewRecorder()
			emp := New(mockDataStore{})
			emp.delete(res,req)
			var op response
			json.NewDecoder(res.Body).Decode(&op)
			if tc.output.Emp != op.Emp || tc.output.Msg != op.Msg {
				t.Fatalf("Test Failed: Expected: %v, Received: %v",tc.output,op)
			}
		})
	}
}

func TestEmployeeHandler_Handle(t *testing.T) {
	testCases := []struct{
		method string
		expectedStatusCode int
		description string
	}{
		{

			"GET",
			http.StatusOK,
			"Testing GET",
		},
		{
			method: "POST",
			expectedStatusCode: http.StatusOK,
			description: "Testing POST",
		},
		{
			method: "DELETE",
			expectedStatusCode: http.StatusOK,
			description: "Testing DELETE",
		},
		{
			method: "PUT",
			expectedStatusCode: http.StatusOK,
			description: "Testing PUT",
		},
		{
			method: "PATCH",
			expectedStatusCode: http.StatusMethodNotAllowed,
			description: "Testing PATCH",
		},
	}

	for _,tc := range testCases {
		t.Run(tc.description,func(t *testing.T){
			req := httptest.NewRequest(tc.method,"http://localhost:8000/employee",nil)
			w := httptest.NewRecorder()
			e := New(EmployeeService.New(mockDataStore{}))
			e.Handle(w,req)
			if w.Code != tc.expectedStatusCode {
				t.Errorf("Expected %v\tGot %v", tc.expectedStatusCode, w.Code)
			}

		})
	}
}