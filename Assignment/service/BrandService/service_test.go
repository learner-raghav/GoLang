package BrandService

import (
	"../../entity"
	"../../mocks"
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

type response struct{
	brand entity.Brand
	err error
}

func TestBrandService_GetById(t *testing.T) {
	testCases := []struct{
		desc string
		id int
		output response
	}{
		{
			desc: "Id exists",
			id: 1,
			output: response{brand: entity.Brand{
				Brand_id: 1,
				Brand_name: "Honda",
			},err: nil},
		},
		{
			desc: "Id does not exists",
			id: 100,
			output: response{brand: entity.Brand{},err: errors.New("Id does not exist")},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _,tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			m := mocks.NewMockBrandStoreHandler(ctrl)
			m.EXPECT().GetById(tc.id).Return(tc.output.brand,tc.output.err)
			brandService := New(m)
			brand,err := brandService.GetById(tc.id)
			if (err != nil && err.Error() != tc.output.err.Error()) || (tc.output.brand != brand) {
				t.Fatalf("Test Failed: Expected: %v, Received: %v %v",tc.output,brand,err)
			}
		})
	}
}

func TestBrandService_Update(t *testing.T) {

	testCases := []struct {
		desc string
		input entity.Brand
		output response
	}{
		{
			desc : "Id exists",
			input: entity.Brand{Brand_name: "M Suzuki",Brand_id: 2},
			output: response{brand: entity.Brand{Brand_name: "M Suzuki",Brand_id: 2},err: nil},
		},
		{
			desc: "Id does not exist",
			input: entity.Brand{Brand_id: 400},
			output: response{brand: entity.Brand{},err: errors.New("The id does not exist or no modifications were done")},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _,tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			m := mocks.NewMockBrandStoreHandler(ctrl)
			m.EXPECT().Update(tc.input).Return(tc.output.brand,tc.output.err)
			brandService := New(m)
			brand,err := brandService.Update(tc.input)
			if (err != nil && err.Error() != tc.output.err.Error()) || (tc.output.brand != brand) {
				t.Fatalf("Test Failed: Expected: %v, Received: %v %v",tc.output,brand,err)
			}
		})
	}
}

func TestBrandService_Create(t *testing.T) {
	testCases := []struct {
		desc string
		input entity.Brand
		output response
	}{
		{
			desc: "Contains data",
			input: entity.Brand{Brand_name: "Mercedes"},
			output: response{brand: entity.Brand{Brand_name: "Mercedes"},err: nil},
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _,tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			m := mocks.NewMockBrandStoreHandler(ctrl)
			m.EXPECT().Create(tc.input).Return(tc.output.brand,tc.output.err)
			brandService := New(m)
			brand,err := brandService.Create(tc.input)
			if (err != nil && err.Error() != tc.output.err.Error()) || (tc.output.brand != brand) {
				t.Fatalf("Test Failed: Expected: %v, Received: %v %v",tc.output,brand,err)
			}
		})
	}
}
func TestBrandService_Delete(t *testing.T) {
	testCases := []struct {
		desc string
		id int
		output response
	}{
		{
			desc: "Id exists",
			id: 6,
			output: response{brand: entity.Brand{Brand_name: "Audi",Brand_id: 6},err: nil},
		},
		{
			desc: "Id does not exist",
			id: 30,
			output: response{brand: entity.Brand{}, err: errors.New("BrandStore record does not exist")},
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _,tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			m := mocks.NewMockBrandStoreHandler(ctrl)
			m.EXPECT().Delete(tc.id).Return(tc.output.brand,tc.output.err)
			brandService := New(m)
			brand,err := brandService.Delete(tc.id)
			if (err != nil && err.Error() != tc.output.err.Error()) || (tc.output.brand != brand) {
				t.Fatalf("Test Failed: Expected: %v, Received: %v %v",tc.output,brand,err)
			}
		})
	}
}