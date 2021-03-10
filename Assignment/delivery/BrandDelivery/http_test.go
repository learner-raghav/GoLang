package BrandDelivery

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
	Brand entity.Brand
	Msg string
	err error
}

func TestGetBrand(t *testing.T){
	testCases := []struct {
		desc string
		id int
		result response
	}{
		{
			desc: "Id exists",
			id: 1,
			result: response{Brand: entity.Brand{
				Brand_id: 1,
				Brand_name: "Honda",
			},err: nil,Msg: "Success"},
		},
		{
			desc: "Id does not exists",
			id: 100,
			result: response{Brand: entity.Brand{},err: errors.New("Id does not exist"),Msg: "Id does not exist"},
		},
	}


	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _,tc := range testCases {
		a := strconv.Itoa(tc.id)
		req := httptest.NewRequest("GET","/brand?id="+a,nil)
		res := httptest.NewRecorder()
		m := mocks.NewMockBrandServicer(ctrl)
		m.EXPECT().GetById(tc.id).Return(tc.result.Brand,tc.result.err)
		emp := New(m)
		emp.Handle(res,req)
		var op response
		json.NewDecoder(res.Body).Decode(&op)
		if tc.result.Brand != op.Brand || tc.result.Msg != op.Msg {
			t.Fatalf("Test Failed: Expected: %v, Received: %v",tc.result,op)
		}
	}
}

func TestCreateBrand(t *testing.T){
	testCases := []struct {
		desc string
		input entity.Brand
		result response
	}{
		{
			desc: "Contains data",
			input: entity.Brand{Brand_name: "Mercedes"},
			result: response{Brand: entity.Brand{Brand_name: "Mercedes"},err: nil,Msg: "Success"},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _,tc := range testCases {
		body, _ := json.Marshal(tc.input)
		req := httptest.NewRequest("POST","/brand",bytes.NewBuffer(body))
		res := httptest.NewRecorder()
		m := mocks.NewMockBrandServicer(ctrl)
		m.EXPECT().Create(tc.input).Return(tc.result.Brand,tc.result.err)
		emp := New(m)
		emp.Handle(res,req)
		var op response
		json.NewDecoder(res.Body).Decode(&op)
		if tc.result.Brand != op.Brand || tc.result.Msg != op.Msg {
			t.Fatalf("Test Failed: Expected: %v, Received: %v ",tc.result,op)
		}
	}
}

func TestUpdateBrand(t *testing.T){
	testCases := []struct {
		desc string
		input entity.Brand
		result response
	}{
		{
			desc : "Id exists",
			input: entity.Brand{Brand_name: "M Suzuki",Brand_id: 2},
			result: response{Brand: entity.Brand{Brand_name: "M Suzuki",Brand_id: 2},err: nil,Msg: "Success"},
		},
		{
			desc: "Id does not exist",
			input: entity.Brand{Brand_id: 400},
			result: response{Brand: entity.Brand{},err: errors.New("The id does not exist or no modifications were done"),Msg: "The id does not exist or no modifications were done"},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _,tc := range testCases {
		body,_ := json.Marshal(tc.input)
		a := strconv.Itoa(tc.input.Brand_id)
		req := httptest.NewRequest("PUT","/brand?id="+a,bytes.NewBuffer(body))
		res := httptest.NewRecorder()
		m := mocks.NewMockBrandServicer(ctrl)
		m.EXPECT().Update(tc.input).Return(tc.result.Brand,tc.result.err)
		emp := New(m)
		emp.Handle(res,req)
		var op response
		json.NewDecoder(res.Body).Decode(&op)
		if tc.result.Brand != op.Brand || tc.result.Msg != op.Msg {
			t.Fatalf("Test Failed: Expected: %v, Received: %v",tc.result,op)
		}
	}
}

func TestDeleteBrand(t *testing.T) {
	testCases := []struct {
		desc string
		id int
		result response
	}{
		{
			desc: "Id exists",
			id: 6,
			result: response{Brand: entity.Brand{Brand_name: "Audi",Brand_id: 6},err: nil,Msg: "Success"},
		},
		{
			desc: "Id does not exist",
			id: 30,
			result: response{Brand: entity.Brand{}, err: errors.New("BrandStore record does not exist"),Msg: "BrandStore record does not exist"},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _,tc := range testCases {
		a := strconv.Itoa(tc.id)
		req := httptest.NewRequest("DELETE","/brand?id="+a,nil)
		res := httptest.NewRecorder()
		m := mocks.NewMockBrandServicer(ctrl)
		m.EXPECT().Delete(tc.id).Return(tc.result.Brand,tc.result.err)
		emp := New(m)
		emp.Handle(res,req)
		var op response
		json.NewDecoder(res.Body).Decode(&op)
		if tc.result.Brand != op.Brand || tc.result.Msg != op.Msg {
			t.Fatalf("Test Failed: Expected: %v, Received: %v",tc.result,op)
		}
	}
}