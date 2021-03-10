package ModelDelivery

import (
	"Assignment/entity"
	"Assignment/mocks"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	"strconv"
	"testing"
)

type response struct{
	Model entity.Model
	Msg string
	err error
}

func TestGetModel(t *testing.T){
	testCases := []struct {
		desc string
		id int
		result response
	}{
		{
			desc: "Id exists",
			id: 1,
			result: response{Model: entity.Model{Model_id: 1,Model_name: "Swift",Brand_id: 2,Brand_name: "M Suzuki"},err: nil,Msg: "Success"},
		},
		{
			desc: "Id does not exist",
			id: 3,
			result: response{Model: entity.Model{}, err: errors.New("Record does not exist"),Msg: "Record does not exist"},
		},
	}


	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _,tc := range testCases {
		a := strconv.Itoa(tc.id)
		req := httptest.NewRequest("GET","/model?id="+a,nil)
		res := httptest.NewRecorder()
		m := mocks.NewMockModelServicer(ctrl)
		m.EXPECT().GetById(tc.id).Return(tc.result.Model,tc.result.err)
		emp := New(m)
		emp.Handle(res,req)
		var op response
		json.NewDecoder(res.Body).Decode(&op)
		if tc.result.Model != op.Model || tc.result.Msg != op.Msg {
			t.Fatalf("Test Failed: Expected: %v, Received: %v",tc.result,op)
		}
	}
}

func TestCreateModel(t *testing.T){
	testCases := []struct {
		desc string
		input entity.Model
		result response
	}{
		{
			desc: "Create success",
			input: entity.Model{Model_name: "Swift",Brand_id: 2},
			result: response{Model: entity.Model{Model_name: "Swift",Brand_id: 2},err: nil,Msg: "Success"},
		},
		{
			desc: "Create error",
			input: entity.Model{Brand_id: 200},
			result: response{Model: entity.Model{},err: errors.New("Foreign key does not exist"),Msg: "Foreign key does not exist"},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _,tc := range testCases {
		body, _ := json.Marshal(tc.input)
		req := httptest.NewRequest("POST","/model",bytes.NewBuffer(body))
		res := httptest.NewRecorder()
		m := mocks.NewMockModelServicer(ctrl)
		m.EXPECT().Create(tc.input).Return(tc.result.Model,tc.result.err)
		emp := New(m)
		emp.Handle(res,req)
		var op response
		json.NewDecoder(res.Body).Decode(&op)
		if tc.result.Model != op.Model || tc.result.Msg != op.Msg {
			t.Fatalf("Test Failed: Expected: %v, Received: %v ",tc.result,op)
		}
	}
}

func TestUpdateModel(t *testing.T){
	testCases := []struct {
		desc string
		input entity.Model
		result response
	}{
		{
			desc: "Update success",
			input: entity.Model{Model_name: "Swift",Brand_id: 2,Model_id: 1},
			result: response{Model: entity.Model{Model_name: "Swift",Brand_id: 2,Model_id: 1},err: nil,Msg: "Success"},
		},
		{
			desc: "Id does not exist",
			input: entity.Model{Model_id: 400},
			result: response{Model: entity.Model{},err: errors.New("The id does not exist or no modifications were done"),Msg:"The id does not exist or no modifications were done" },
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _,tc := range testCases {
		body,_ := json.Marshal(tc.input)
		a := strconv.Itoa(tc.input.Model_id)
		req := httptest.NewRequest("PUT","/model?id="+a,bytes.NewBuffer(body))
		res := httptest.NewRecorder()
		m := mocks.NewMockModelServicer(ctrl)
		m.EXPECT().Update(tc.input).Return(tc.result.Model,tc.result.err)
		emp := New(m)
		emp.Handle(res,req)
		var op response
		json.NewDecoder(res.Body).Decode(&op)
		if tc.result.Model != op.Model || tc.result.Msg != op.Msg {
			t.Fatalf("Test Failed: Expected: %v, Received: %v",tc.result,op)
		}
	}
}

func TestDeleteModel(t *testing.T) {
	testCases := []struct {
		desc string
		id int
		result response
	}{
		{
			desc: "Delete success",
			id: 10,
			result: response{Model: entity.Model{Model_id: 10,Model_name: "A3",Brand_id: 1,Brand_name: "Honda"},err: nil,Msg: "Success"},
		},
		{
			desc: "Id does not exist",
			id: 100,
			result: response{Model: entity.Model{},err: errors.New("Id does not exist"),Msg: "Id does not exist"},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _,tc := range testCases {
		a := strconv.Itoa(tc.id)
		req := httptest.NewRequest("DELETE","/model?id="+a,nil)
		res := httptest.NewRecorder()
		m := mocks.NewMockModelServicer(ctrl)
		m.EXPECT().Delete(tc.id).Return(tc.result.Model,tc.result.err)
		emp := New(m)
		emp.Handle(res,req)
		var op response
		json.NewDecoder(res.Body).Decode(&op)
		if tc.result.Model != op.Model|| tc.result.Msg != op.Msg {
			t.Fatalf("Test Failed: Expected: %v, Received: %v",tc.result,op)
		}
	}
}