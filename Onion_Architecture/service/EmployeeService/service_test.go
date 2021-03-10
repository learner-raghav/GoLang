package EmployeeDelivery

import (
	"../../entity"
	"Onion_Architecture/service/EmployeeService"
	"Onion_Architecture/delivery/EmployeeDelivery"
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"../../mocks"
)
type response struct{
	Msg string
	Emp entity.Employee
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	for _,tc := range testCases {
		t.Run(tc.desc,func(t *testing.T){
			a := strconv.Itoa(tc.id)
			req := httptest.NewRequest("GET","/employee?id="+a,nil)
			res := httptest.NewRecorder()
			m := mock_store.NewMockEmployeeStoreHandler(ctrl)
			m.EXPECT().GetById(tc.id).Return(tc.output.Emp,tc.output.Msg)
			emp := New(m)
			emp.Handle(res,req)
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	for _,tc := range testCases {
		t.Run(tc.desc,func(t *testing.T){
			body, _ := json.Marshal(tc.emp)
			req := httptest.NewRequest("POST","/employee",bytes.NewBuffer(body))
			res := httptest.NewRecorder()
			m := mock_store.NewMockEmployeeStoreHandler(ctrl)
			m.EXPECT().Create(tc.emp).Return(tc.output.Emp,tc.output.Msg)
			emp := New(m)
			emp.Handle(res,req)
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
			req := httptest.NewRequest("PUT","/employee?id="+strconv.Itoa(tc.emp.EmployeeId),bytes.NewBuffer(body))
			res := httptest.NewRecorder()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_store.NewMockEmployeeStoreHandler(ctrl)
			m.EXPECT().Update(tc.emp).Return(tc.output.Emp,tc.output.Msg)
			emp := New(m)
			emp.Handle(res,req)
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
			req := httptest.NewRequest("DELETE","/employee?id="+a,nil)
			res := httptest.NewRecorder()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_store.NewMockEmployeeStoreHandler(ctrl)
			m.EXPECT().Delete(tc.id).Return(tc.output.Emp,tc.output.Msg)
			emp := New(m)
			emp.Handle(res,req)
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