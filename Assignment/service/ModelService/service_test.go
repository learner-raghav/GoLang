package ModelService

import (
	"../../entity"
	"../../mocks"
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

type response struct{
	brand entity.Model
	err error
}

func TestVariantService_GetById(t *testing.T) {
	testCases := []struct {

		desc string
		id int
		result response
	}{
		{
			desc: "Id exists",
			id: 1,
			result: response{brand: entity.Model{Model_id: 1,Model_name: "Swift",Brand_id: 2,Brand_name: "M Suzuki"},err: nil},
		},
		{
			desc: "Id does not exist",
			id: 3,
			result: response{brand: entity.Model{}, err: errors.New("Record does not exist")},
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _,tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			m := mocks.NewMockModelStoreHandler(ctrl)
			m.EXPECT().GetById(tc.id).Return(tc.result.brand,tc.result.err)
			modelService := New(m)
			brand,err := modelService.GetById(tc.id)
			if (err != nil && err.Error() != tc.result.err.Error()) || (tc.result.brand != brand) {
				t.Fatalf("Test Failed: Expected: %v, Received: %v %v",tc.result,brand,err)
			}
		})
	}
}

func TestVariantService_Update(t *testing.T) {
	testCases := []struct {
		desc string
		input entity.Model
		result response
	}{
		{
			desc: "Update success",
			input: entity.Model{Model_name: "Swift",Brand_id: 2,Model_id: 1},
			result: response{brand: entity.Model{Model_name: "Swift",Brand_id: 2,Model_id: 1},err: nil},
		},
		{
			desc: "Id does not exist",
			input: entity.Model{Model_id: 400},
			result: response{brand: entity.Model{},err: errors.New("The id does not exist or no modifications were done")},
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _,tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			m := mocks.NewMockModelStoreHandler(ctrl)
			m.EXPECT().GetById(tc.input).Return(tc.result.brand,tc.result.err)
			modelService := New(m)
			brand,err := modelService.Update(tc.input)
			if (err != nil && err.Error() != tc.result.err.Error()) || (tc.result.brand != brand) {
				t.Fatalf("Test Failed: Expected: %v, Received: %v %v",tc.result,brand,err)
			}
		})
	}
}

func TestVariantService_Create(t *testing.T) {

	testCases := []struct {
		desc string
		input entity.Model
		result response
	}{
		{
			desc: "Create success",
			input: entity.Model{Model_name: "Swift",Brand_id: 2},
			result: response{brand: entity.Model{Model_name: "Swift",Brand_id: 2},err: nil},
		},
		{
			desc: "Create error",
			input: entity.Model{Brand_id: 200},
			result: response{brand: entity.Model{},err: errors.New("Foreign key does not exist")},
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _,tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			m := mocks.NewMockModelStoreHandler(ctrl)
			m.EXPECT().Create(tc.input).Return(tc.result.brand,tc.result.err)
			modelService := New(m)
			brand,err := modelService.Create(tc.input)
			if (err != nil && err.Error() != tc.result.err.Error()) || (tc.result.brand != brand) {
				t.Fatalf("Test Failed: Expected: %v, Received: %v %v",tc.result,brand,err)
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
			desc: "Delete success",
			id: 10,
			result: response{brand: entity.Model{Model_id: 10,Model_name: "A3",Brand_id: 1,Brand_name: "Honda"},err: nil},
		},
		{
			desc: "Id does not exist",
			id: 100,
			result: response{brand: entity.Model{},err: errors.New("Id does not exist")},
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _,tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			m := mocks.NewMockModelStoreHandler(ctrl)
			m.EXPECT().GetById(tc.id).Return(tc.result.brand,tc.result.err)
			modelService := New(m)
			brand,err := modelService.Delete(tc.id)
			if (err != nil && err.Error() != tc.result.err.Error()) || (tc.result.brand != brand) {
				t.Fatalf("Test Failed: Expected: %v, Received: %v %v",tc.result,brand,err)
			}
		})
	}
}