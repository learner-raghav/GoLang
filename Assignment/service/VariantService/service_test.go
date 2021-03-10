package VariantService

import (
	"../../entity"
	"../../mocks"
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)
type response struct{
	variant entity.Variant
	err error
}

func TestVariantService_Create(t *testing.T) {
	testCases := []struct {
		desc string
		input entity.Variant
		result response
	}{
		{
			desc: "Created",
			input: entity.Variant{Variant_name: "i10 asta",Model_id: 2,Disp: 12.3,Peak_power: 120,Peak_torque: 30},
			result: response{variant: entity.Variant{Variant_name: "i10 asta",Model_id: 2,Disp: 12.3,Peak_power: 120,Peak_torque: 30},err: nil},
		},
		{
			desc: "Foreign key error",
			input: entity.Variant{Model_id: 200},
			result: response{variant: entity.Variant{},err: errors.New("Foreign key does not exist")},
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _,tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			m := mocks.NewMockVariantStoreHandler(ctrl)
			m.EXPECT().Create(tc.input).Return(tc.result.variant,tc.result.err)
			variantService := New(m)
			variant,err := variantService.Create(tc.input)
			if (err != nil && err.Error() != tc.result.err.Error()) || (tc.result.variant != variant) {
				t.Fatalf("Test Failed: Expected: %v, Received: %v %v",tc.result,variant,err)
			}
		})
	}
}

func TestVariantService_GetById(t *testing.T) {

	testCases := []struct {
		desc string
		id int
		result response
	}{
		{
			desc: "Id exists",
			id: 7,
			result: response{variant: entity.Variant{Variant_id:7,Variant_name: "i10 Asta",Model_id: 2,Disp: 12.3,Peak_power: 110,Peak_torque: 30,Model_name: "I10",Brand_name: "Honda",Brand_Id: 1},err: nil},
		},
		{
			desc: "Id does not exist",
			id: 300,
			result: response{variant: entity.Variant{}, err: errors.New("Record does not exist")},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _,tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			m := mocks.NewMockVariantStoreHandler(ctrl)
			m.EXPECT().GetById(tc.id).Return(tc.result.variant,tc.result.err)
			variantService := New(m)
			variant,err := variantService.GetById(tc.id)
			if (err != nil && err.Error() != tc.result.err.Error()) || (tc.result.variant != variant) {
				t.Fatalf("Test Failed: Expected: %v, Received: %v %v",tc.result,variant,err)
			}
		})
	}
}

func TestVariantService_Update(t *testing.T) {
	testCases := []struct {
		desc string
		input entity.Variant
		result response
	}{
		{
			desc: "Update success",
			input: entity.Variant{Variant_id:5,Variant_name: "i10 Asta",Model_id: 2,Disp: 12.3,Peak_power: 120,Peak_torque: 30},
			result: response{variant: entity.Variant{Variant_id:5,Variant_name: "i10 Asta",Model_id: 2,Disp: 12.3,Peak_power: 120,Peak_torque: 30},err: nil},
		},
		{
			desc: "Id does nto exist",
			input: entity.Variant{Model_id: 400},
			result: response{variant: entity.Variant{},err: errors.New("The id does not exist or no modifications were done")},
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _,tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			m := mocks.NewMockVariantStoreHandler(ctrl)
			m.EXPECT().Update(tc.input).Return(tc.result.variant,tc.result.err)
			variantService := New(m)
			variant,err := variantService.Update(tc.input)
			if (err != nil && err.Error() != tc.result.err.Error()) || (tc.result.variant != variant) {
				t.Fatalf("Test Failed: Expected: %v, Received: %v %v",tc.result,variant,err)
			}
		})
	}
}

func TestVariantService_Delete(t *testing.T) {
	testCases := []struct {
		desc string
		id int
		result response
	}{
		{
			desc: "Id exists",
			id: 5,
			result: response{variant: entity.Variant{Variant_id:5,Variant_name: "i10 Asta",Model_id: 2,Disp: 12.3,Peak_power: 120,Peak_torque: 30},err: nil},
		},
		{
			desc: "Id does not exist",
			id: 30,
			result: response{variant: entity.Variant{}, err: errors.New("Record does not exist")},
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _,tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			m := mocks.NewMockVariantStoreHandler(ctrl)
			m.EXPECT().Delete(tc.id).Return(tc.result.variant,tc.result.err)
			variantService := New(m)
			variant,err := variantService.Delete(tc.id)
			if (err != nil && err.Error() != tc.result.err.Error()) || (tc.result.variant != variant) {
				t.Fatalf("Test Failed: Expected: %v, Received: %v %v",tc.result,variant,err)
			}
		})
	}
}
