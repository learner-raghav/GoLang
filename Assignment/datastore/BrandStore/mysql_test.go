package BrandStore

import (
	"../../driver"
	"../../entity"
	"database/sql"
	"errors"
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
	brandStorer := New(db)
	testBrandStorer_Get(t,brandStorer)
	testBrandStorer_Create(t,brandStorer)
	testBrandStorer_Update(t,brandStorer)
	testBrandStorer_Delete(t,brandStorer)
}


func testBrandStorer_Get(t *testing.T,store BrandStore) {
	type response struct{
		err error
		brand entity.Brand
	}
	testCases := []struct {
		id int
		brandStore BrandStore
		result response
	}{
		{
			id: 1,
			brandStore: store,
			result: response{brand: entity.Brand{Brand_name: "Honda",Brand_id: 1},err: nil},
		},
		{
			id: 3,
			brandStore: store,
			result: response{brand: entity.Brand{}, err: errors.New("Record does not exist")},
		},
		{
			id: 100,
			brandStore: BrandStore{db: nil},
			result: response{brand: entity.Brand{},err: errors.New("DB not configured properly")},
		},
	}
	for i,tc := range testCases {
		resp, err := tc.brandStore.GetById(tc.id)
		if (err != nil && err.Error() != tc.result.err.Error()) || (resp != tc.result.brand) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp,tc.result.brand)
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, err,tc.result.err)
		}
	}
}

func testBrandStorer_Create(t *testing.T,store BrandStore) {
	type response struct{
		err error
		brand entity.Brand
	}
	testCases := []struct {
		input entity.Brand
		brandStore BrandStore
		result response
	}{
		{
			input: entity.Brand{Brand_name: "Mercedes"},
			brandStore: store,
			result: response{brand: entity.Brand{Brand_name: "Mercedes"},err: nil},
		},
		{
			input: entity.Brand{},
			brandStore: BrandStore{db: nil},
			result: response{brand: entity.Brand{},err: errors.New("DB not configured properly")},
		},
	}
	for i,tc := range testCases {
		resp, err := tc.brandStore.Create(tc.input)
		if (err != nil && err.Error() != tc.result.err.Error()) || (resp != tc.result.brand) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp,tc.result.brand)
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, err,tc.result.err)
		}
	}
}

func testBrandStorer_Update(t *testing.T,store BrandStore) {
	type response struct{
		err error
		brand entity.Brand
	}
	testCases := []struct {
		input entity.Brand
		brandStore BrandStore
		result response
	}{
		{
			input: entity.Brand{Brand_name: "M Suzuki",Brand_id: 2},
			brandStore: store,
			result: response{brand: entity.Brand{Brand_name: "M Suzuki",Brand_id: 2},err: nil},
		},
		{
			input: entity.Brand{Brand_id: 400},
			brandStore: store,
			result: response{brand: entity.Brand{},err: errors.New("The id does not exist or no modifications were done")},
		},
		{
			input: entity.Brand{},
			brandStore: BrandStore{db: nil},
			result: response{brand: entity.Brand{},err: errors.New("DB not configured properly")},
		},
	}
	for i,tc := range testCases {
		resp, err := tc.brandStore.Update(tc.input)
		if (err != nil && err.Error() != tc.result.err.Error()) || (resp != tc.result.brand) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp,tc.result.brand)
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, err,tc.result.err)
		}
	}
}

func testBrandStorer_Delete(t *testing.T,store BrandStore) {
	type response struct{
		err error
		brand entity.Brand
	}
	testCases := []struct {
		id int
		brandStore BrandStore
		result response
	}{
		{
			id: 6,
			brandStore: store,
			result: response{brand: entity.Brand{Brand_name: "Audi",Brand_id: 6},err: nil},
		},
		{
			id: 30,
			brandStore: store,
			result: response{brand: entity.Brand{}, err: errors.New("BrandStore record does not exist")},
		},
		{
			id: 30,
			brandStore: BrandStore{db: nil},
			result: response{brand: entity.Brand{}, err: errors.New("DB not configured properly")},
		},
	}
	for i,tc := range testCases {
		resp, err := tc.brandStore.Delete(tc.id)
		if (err != nil && err.Error() != tc.result.err.Error()) || (resp != tc.result.brand) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp,tc.result.brand)
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, err,tc.result.err)
		}
	}
}
