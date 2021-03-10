package mocks

import (
	entity "Assignment/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBrandServicer is a mock of BrandServicer interface.
type MockBrandServicer struct {
	ctrl     *gomock.Controller
	recorder *MockBrandServicerMockRecorder
}

// MockBrandServicerMockRecorder is the mock recorder for MockBrandServicer.
type MockBrandServicerMockRecorder struct {
	mock *MockBrandServicer
}

// NewMockBrandServicer creates a new mock instance.
func NewMockBrandServicer(ctrl *gomock.Controller) *MockBrandServicer {
	mock := &MockBrandServicer{ctrl: ctrl}
	mock.recorder = &MockBrandServicerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBrandServicer) EXPECT() *MockBrandServicerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockBrandServicer) Create(arg0 entity.Brand) (entity.Brand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(entity.Brand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockBrandServicerMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBrandServicer)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockBrandServicer) Delete(arg0 int) (entity.Brand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(entity.Brand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockBrandServicerMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBrandServicer)(nil).Delete), arg0)
}

// GetById mocks base method.
func (m *MockBrandServicer) GetById(arg0 int) (entity.Brand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0)
	ret0, _ := ret[0].(entity.Brand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockBrandServicerMockRecorder) GetById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockBrandServicer)(nil).GetById), arg0)
}

// Update mocks base method.
func (m *MockBrandServicer) Update(arg0 entity.Brand) (entity.Brand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(entity.Brand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockBrandServicerMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBrandServicer)(nil).Update), arg0)
}

