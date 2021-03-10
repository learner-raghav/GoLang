package ModelStore

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
	ModelStorer := New(db)
	//testModelStorer_Get(t,ModelStorer)
	//testModelStorer_Create(t,ModelStorer)
	//testModelStorer_Update(t,ModelStorer)
	//testModelStorer_Delete(t,ModelStorer)
	testModelStorer_Read(t,ModelStorer)
}

func testModelStorer_Read(t *testing.T, store ModelStore) {
	type response struct {
		err error
		models []entity.Model
	}
	
	testCases := []struct{
		desc string
		filter map[string]string
		modelStore ModelStore
		result response
	}{
		{
			desc:       "filter 1",
			filter:     map[string]string{"model_name": "I10"},
			modelStore: store,
			result:     response{err: nil, models: []entity.Model{{13, "I10", 9, "Honda"}}},
		},
		{
			desc:       "filter 2",
			filter:     map[string]string{"brand_id": "10"},
			modelStore: store,
			result:     response{err: nil, models: []entity.Model{{14, "Swift", 10, "Suzuki"}, {15, "Baleno", 10, "Suzuki"}, {16, "Creta", 10, "Suzuki"}},},
		},
	}
	for i,tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			resp,err := tc.modelStore.Read(tc.filter)
			if !reflect.DeepEqual(resp,tc.result.models) || err != tc.result.err {
				t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp,tc.result.models)
				t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, err,tc.result.err)
			}
		})
	}
}


func testModelStorer_Get(t *testing.T,store ModelStore) {
	type response struct{
		err error
		brand entity.Model
	}
	testCases := []struct {
		id int
		ModelStore ModelStore
		result response
	}{
		{
			id: 1,
			ModelStore: store,
			result: response{brand: entity.Model{Model_id: 1,Model_name: "Swift",Brand_id: 2,Brand_name: "M Suzuki"},err: nil},
		},
		{
			id: 3,
			ModelStore: store,
			result: response{brand: entity.Model{}, err: errors.New("Record does not exist")},
		},
		{
			id: 100,
			ModelStore: ModelStore{db: nil},
			result: response{brand: entity.Model{},err: errors.New("DB not configured properly")},
		},
	}
	for i,tc := range testCases {
		resp, err := tc.ModelStore.GetById(tc.id)
		if (err != nil && err.Error() != tc.result.err.Error()) || (resp != tc.result.brand) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp,tc.result.brand)
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, err,tc.result.err)
		}
	}
}

func testModelStorer_Create(t *testing.T,store ModelStore) {
	type response struct{
		err error
		brand entity.Model
	}
	testCases := []struct {
		input entity.Model
		ModelStore ModelStore
		result response
	}{
		{
			input: entity.Model{Model_name: "Swift",Brand_id: 2},
			ModelStore: store,
			result: response{brand: entity.Model{Model_name: "Swift",Brand_id: 2},err: nil},
		},
		{
			input: entity.Model{},
			ModelStore: ModelStore{db: nil},
			result: response{brand: entity.Model{},err: errors.New("DB not configured properly")},
		},
		{
			input: entity.Model{Brand_id: 200},
			ModelStore: store,
			result: response{brand: entity.Model{},err: errors.New("Foreign key does not exist")},
		},
	}
	for i,tc := range testCases {
		resp, err := tc.ModelStore.Create(tc.input)
		if (err != nil && err.Error() != tc.result.err.Error()) || (resp != tc.result.brand) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp,tc.result.brand)
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, err,tc.result.err)
		}
	}
}

func testModelStorer_Update(t *testing.T,store ModelStore) {
	type response struct{
		err error
		brand entity.Model
	}
	testCases := []struct {
		input entity.Model
		ModelStore ModelStore
		result response
	}{
		{
			input: entity.Model{Model_name: "Swift",Brand_id: 2,Model_id: 1},
			ModelStore: store,
			result: response{brand: entity.Model{Model_name: "Swift",Brand_id: 2,Model_id: 1},err: nil},
		},
		{
			input: entity.Model{Model_id: 400},
			ModelStore: store,
			result: response{brand: entity.Model{},err: errors.New("The id does not exist or no modifications were done")},
		},
		{
			input: entity.Model{},
			ModelStore: ModelStore{db: nil},
			result: response{brand: entity.Model{},err: errors.New("DB not configured properly")},
		},
	}
	for i,tc := range testCases {
		resp, err := tc.ModelStore.Update(tc.input)
		if (err != nil && err.Error() != tc.result.err.Error()) || (resp != tc.result.brand) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp,tc.result.brand)
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, err,tc.result.err)
		}
	}
}

func testModelStorer_Delete(t *testing.T,store ModelStore) {
	type response struct{
		err error
		brand entity.Model
	}
	testCases := []struct {
		id int
		ModelStore ModelStore
		result response
	}{
		{
			id: 10,
			ModelStore: store,
			result: response{brand: entity.Model{Model_id: 10,Model_name: "A3",Brand_id: 1,Brand_name: "Honda"},err: nil},
		},
		{
			id: 30,
			ModelStore: store,
			result: response{brand: entity.Model{}, err: errors.New("Record does not exist")},
		},
		{
			id: 30,
			ModelStore: ModelStore{db: nil},
			result: response{brand: entity.Model{}, err: errors.New("DB not configured properly")},
		},
	}
	for i,tc := range testCases {
		resp, err := tc.ModelStore.Delete(tc.id)
		if (err != nil && err.Error() != tc.result.err.Error()) || (resp != tc.result.brand) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp,tc.result.brand)
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, err,tc.result.err)
		}
	}
}
