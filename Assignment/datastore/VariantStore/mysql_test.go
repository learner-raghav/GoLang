package VariantStore

import (
	"../../driver"
	"../../entity"
	"database/sql"
	"errors"
	"reflect"
	"testing"
)

func initializeMySQL(t *testing.T) *sql.DB {
	conf := entity.MySQLConfig{
		DbName: "testDB",
		DbUser: "raghav",
		DbPass: "raghav@123M",
	}

	var err error
	db, err := driver.ConnectToDB(conf)
	if err != nil {
		t.Errorf("could not connect to sql, err:%v", err)
	}

	return db
}

func TestDatastore(t *testing.T) {
	db := initializeMySQL(t)
	VariantStorer := New(db)
	//testVariantStorer_Get(t,VariantStorer)
	//testVariantStorer_Create(t,VariantStorer)
	//testVariantStorer_Update(t,VariantStorer)
	//testVariantStorer_Delete(t,VariantStorer)
	testModelStorer_Read(t,VariantStorer)
}


func testModelStorer_Read(t *testing.T, store VariantStore) {
	type response struct {
		err error
		variants []entity.Variant
	}

	testCases := []struct{
		desc string
		filter map[string]string
		varStore VariantStore
		result response
	}{
		{
			desc:     "filter 1",
			filter:   map[string]string{"model_id": "13"},
			varStore: store,
			result:   response{err: nil, variants: []entity.Variant{{9, 13, "i10 sportz", 23.2, 110, 120, "I10", "Honda", 9}, {10, 13, "i10 Magna", 23.2, 110, 120, "I10", "Honda", 9}, {11, 13, "i10 Asta", 23.2, 110, 120, "I10", "Honda", 9}},},
		},
		{
			desc:       "filter 2",
			filter:     map[string]string{"model_id": "14","peak_torque":"130"},
			varStore: store,
			result:     response{err: nil, variants: []entity.Variant{{12, 14, "Swift Dzire", 23.2, 110 ,130, "Swift","Suzuki", 10}, {13, 14, "Swift Xs", 23.2, 110, 130, "Swift", "Suzuki", 10}, {14, 14, "Swift sportz" ,23.2, 110, 130, "Swift", "Suzuki", 10}}},
		},
	}
	for i,tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			resp,err := tc.varStore.Read(tc.filter)
			if !reflect.DeepEqual(resp,tc.result.variants) || err != tc.result.err {
				t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp,tc.result.variants)
				t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, err,tc.result.err)
			}
		})
	}
}


func testVariantStorer_Get(t *testing.T,store VariantStore) {
	type response struct{
		err error
		brand entity.Variant
	}
	testCases := []struct {
		id int
		VariantStore VariantStore
		result response
	}{
		{
			id: 7,
			VariantStore: store,
			result: response{brand: entity.Variant{Variant_id:7,Variant_name: "i10 Asta",Model_id: 2,Disp: 12.3,Peak_power: 110,Peak_torque: 30,Model_name: "I10",Brand_name: "Honda",Brand_Id: 1},err: nil},
		},
		{
			id: 300,
			VariantStore: store,
			result: response{brand: entity.Variant{}, err: errors.New("Record does not exist")},
		},
		{
			id: 100,
			VariantStore: VariantStore{db: nil},
			result: response{brand: entity.Variant{},err: errors.New("DB not configured properly")},
		},
	}
	for i,tc := range testCases {
		resp, err := tc.VariantStore.GetById(tc.id)
		if (err != nil && err.Error() != tc.result.err.Error()) || (resp != tc.result.brand) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp,tc.result.brand)
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, err,tc.result.err)
		}
	}
}

func testVariantStorer_Create(t *testing.T,store VariantStore) {
	type response struct{
		err error
		brand entity.Variant
	}
	testCases := []struct {
		input entity.Variant
		VariantStore VariantStore
		result response
	}{
		{
			input: entity.Variant{Variant_name: "i10 asta",Model_id: 2,Disp: 12.3,Peak_power: 120,Peak_torque: 30},
			VariantStore: store,
			result: response{brand: entity.Variant{Variant_name: "i10 asta",Model_id: 2,Disp: 12.3,Peak_power: 120,Peak_torque: 30},err: nil},
		},
		{
			input: entity.Variant{},
			VariantStore: VariantStore{db: nil},
			result: response{brand: entity.Variant{},err: errors.New("DB not configured properly")},
		},
		{
			input: entity.Variant{Model_id: 200},
			VariantStore: store,
			result: response{brand: entity.Variant{},err: errors.New("Foreign key does not exist")},
		},
	}
	for i,tc := range testCases {
		resp, err := tc.VariantStore.Create(tc.input)
		if (err != nil && err.Error() != tc.result.err.Error()) || (resp != tc.result.brand) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp,tc.result.brand)
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, err,tc.result.err)
		}
	}
}

func testVariantStorer_Update(t *testing.T,store VariantStore) {
	type response struct{
		err error
		brand entity.Variant
	}
	testCases := []struct {
		input entity.Variant
		VariantStore VariantStore
		result response
	}{
		{
			input: entity.Variant{Variant_id:5,Variant_name: "i10 Asta",Model_id: 2,Disp: 12.3,Peak_power: 120,Peak_torque: 30},
			VariantStore: store,
			result: response{brand: entity.Variant{Variant_id:5,Variant_name: "i10 Asta",Model_id: 2,Disp: 12.3,Peak_power: 120,Peak_torque: 30},err: nil},
		},
		{
			input: entity.Variant{Model_id: 400},
			VariantStore: store,
			result: response{brand: entity.Variant{},err: errors.New("The id does not exist or no modifications were done")},
		},
		{
			input: entity.Variant{},
			VariantStore: VariantStore{db: nil},
			result: response{brand: entity.Variant{},err: errors.New("DB not configured properly")},
		},
	}
	for i,tc := range testCases {
		resp, err := tc.VariantStore.Update(tc.input)
		if (err != nil && err.Error() != tc.result.err.Error()) || (resp != tc.result.brand) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp,tc.result.brand)
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, err,tc.result.err)
		}
	}
}

func testVariantStorer_Delete(t *testing.T,store VariantStore) {
	type response struct{
		err error
		brand entity.Variant
	}
	testCases := []struct {
		id int
		VariantStore VariantStore
		result response
	}{
		{
			id: 5,
			VariantStore: store,
			result: response{brand: entity.Variant{Variant_id:5,Variant_name: "i10 Asta",Model_id: 2,Disp: 12.3,Peak_power: 120,Peak_torque: 30},err: nil},
		},
		{
			id: 30,
			VariantStore: store,
			result: response{brand: entity.Variant{}, err: errors.New("Record does not exist")},
		},
		{
			id: 30,
			VariantStore: VariantStore{db: nil},
			result: response{brand: entity.Variant{}, err: errors.New("DB not configured properly")},
		},
	}
	for i,tc := range testCases {
		resp, err := tc.VariantStore.Delete(tc.id)
		if (err != nil && err.Error() != tc.result.err.Error()) || (resp != tc.result.brand) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp,tc.result.brand)
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, err,tc.result.err)
		}
	}
}
