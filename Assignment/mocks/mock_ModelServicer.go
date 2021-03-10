package mocks

import (
	entity "Assignment/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockModelServicer is a mock of ModelServicer interface.
type MockModelServicer struct {
	ctrl     *gomock.Controller
	recorder *MockModelServicerMockRecorder
}

// MockModelServicerMockRecorder is the mock recorder for MockModelServicer.
type MockModelServicerMockRecorder struct {
	mock *MockModelServicer
}

// NewMockModelServicer creates a new mock instance.
func NewMockModelServicer(ctrl *gomock.Controller) *MockModelServicer {
	mock := &MockModelServicer{ctrl: ctrl}
	mock.recorder = &MockModelServicerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelServicer) EXPECT() *MockModelServicerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockModelServicer) Create(arg0 entity.Model) (entity.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(entity.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockModelServicerMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockModelServicer)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockModelServicer) Delete(arg0 int) (entity.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(entity.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockModelServicerMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockModelServicer)(nil).Delete), arg0)
}

// GetById mocks base method.
func (m *MockModelServicer) GetById(arg0 int) (entity.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0)
	ret0, _ := ret[0].(entity.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockModelServicerMockRecorder) GetById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockModelServicer)(nil).GetById), arg0)
}

// Update mocks base method.
func (m *MockModelServicer) Update(arg0 entity.Model) (entity.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(entity.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockModelServicerMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockModelServicer)(nil).Update), arg0)
}
