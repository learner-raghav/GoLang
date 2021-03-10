package VariantDelivery

import (
	"Assignment/entity"
	"Assignment/mocks"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	"strconv"
	"testing"
)

type response struct{
	Variant entity.Variant
	Msg string
	err error
}

func TestGetVariant(t *testing.T){
	testCases := []struct {
		desc string
		id int
		result response
	}{
		{
			desc: "Id exists",
			id: 7,
			result: response{Variant: entity.Variant{Variant_id:7,Variant_name: "i10 Asta",Model_id: 2,Disp: 12.3,Peak_power: 110,Peak_torque: 30,Model_name: "I10",Brand_name: "Honda",Brand_Id: 1},Msg: "Success",err: nil},
		},
		{
			desc: "Id does not exist",
			id: 300,
			result: response{Variant: entity.Variant{}, Msg:"Record does not exist",err: errors.New("Record does not exist")},
		},
	}


	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _,tc := range testCases {
		a := strconv.Itoa(tc.id)
		req := httptest.NewRequest("GET","/variant?id="+a,nil)
		res := httptest.NewRecorder()
		m := mocks.NewMockVariantServicer(ctrl)
		m.EXPECT().GetById(tc.id).Return(tc.result.Variant,tc.result.err)
		emp := New(m)
		emp.Handle(res,req)
		var op response
		json.NewDecoder(res.Body).Decode(&op)
		fmt.Println(op)
		fmt.Println(tc.result.Variant.Variant_id,op.Variant.Variant_id)
		if tc.result.Variant != op.Variant || tc.result.Msg != op.Msg {
			t.Fatalf("Test Failed: Expected: %v, Received: %v",tc.result,op)
		}
	}
}

func TestCreateVariant(t *testing.T){
	testCases := []struct {
		desc string
		input entity.Variant
		result response
	}{
		{
			desc: "Created",
			input: entity.Variant{Variant_name: "i10 asta",Model_id: 2,Disp: 12.3,Peak_power: 120,Peak_torque: 30},
			result: response{Variant: entity.Variant{Variant_name: "i10 asta",Model_id: 2,Disp: 12.3,Peak_power: 120,Peak_torque: 30},err: nil,Msg: "Success"},
		},
		{
			desc: "Foreign key error",
			input: entity.Variant{Model_id: 200},
			result: response{Variant: entity.Variant{},err: errors.New("Foreign key does not exist"),Msg:"Foreign key does not exist" },
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _,tc := range testCases {
		body, _ := json.Marshal(tc.input)
		req := httptest.NewRequest("POST","/variant",bytes.NewBuffer(body))
		res := httptest.NewRecorder()
		m := mocks.NewMockVariantServicer(ctrl)
		m.EXPECT().Create(tc.input).Return(tc.result.Variant,tc.result.err)
		emp := New(m)
		emp.Handle(res,req)
		var op response
		json.NewDecoder(res.Body).Decode(&op)
		if tc.result.Variant != op.Variant || tc.result.Msg != op.Msg {
			t.Fatalf("Test Failed: Expected: %v, Received: %v ",tc.result,op)
		}
	}
}

func TestUpdateVariant(t *testing.T){
	testCases := []struct {
		desc string
		input entity.Variant
		result response
	}{
		{
			desc: "Update success",
			input: entity.Variant{Variant_id:5,Variant_name: "i10 Asta",Model_id: 2,Disp: 12.3,Peak_power: 120,Peak_torque: 30},
			result: response{Variant: entity.Variant{Variant_id:5,Variant_name: "i10 Asta",Model_id: 2,Disp: 12.3,Peak_power: 120,Peak_torque: 30},err: nil,Msg: "Success"},
		},
		{
			desc: "Id does nto exist",
			input: entity.Variant{Model_id: 400},
			result: response{Variant: entity.Variant{},err: errors.New("The id does not exist or no modifications were done"),Msg: "The id does not exist or no modifications were done"},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _,tc := range testCases {
		body,_ := json.Marshal(tc.input)
		a := strconv.Itoa(tc.input.Variant_id)
		req := httptest.NewRequest("PUT","/variant?id="+a,bytes.NewBuffer(body))
		res := httptest.NewRecorder()
		m := mocks.NewMockVariantServicer(ctrl)
		m.EXPECT().Update(tc.input).Return(tc.result.Variant,tc.result.err)
		emp := New(m)
		emp.Handle(res,req)
		var op response
		json.NewDecoder(res.Body).Decode(&op)
		fmt.Println(op)
		if tc.result.Variant != op.Variant || tc.result.Msg != op.Msg {
			t.Fatalf("Test Failed: Expected: %v, Received: %v",tc.result,op)
		}
	}
}

func TestDeleteVariant(t *testing.T) {
	testCases := []struct {
		desc string
		id int
		result response
	}{
		{
			desc: "Id exists",
			id: 5,
			result: response{Variant: entity.Variant{Variant_id:5,Variant_name: "i10 Asta",Model_id: 2,Disp: 12.3,Peak_power: 120,Peak_torque: 30},err: nil,Msg: "Success"},
		},
		{
			desc: "Id does not exist",
			id: 30,
			result: response{Variant: entity.Variant{}, err: errors.New("Record does not exist"),Msg: "Record does not exist"},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _,tc := range testCases {
		a := strconv.Itoa(tc.id)
		req := httptest.NewRequest("DELETE","/variant?id="+a,nil)
		res := httptest.NewRecorder()
		m := mocks.NewMockVariantServicer(ctrl)
		m.EXPECT().Delete(tc.id).Return(tc.result.Variant,tc.result.err)
		emp := New(m)
		emp.Handle(res,req)
		var op response
		json.NewDecoder(res.Body).Decode(&op)
		if tc.result.Variant != op.Variant || tc.result.Msg != op.Msg {
			t.Fatalf("Test Failed: Expected: %v, Received: %v",tc.result,op)
		}
	}
}