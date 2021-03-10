package mocks

import (
	entity "Assignment/entity"
	reflect "reflect"
	gomock "github.com/golang/mock/gomock"
)

// MockBrandStoreHandler is a mock of BrandStoreHandler interface.
type MockBrandStoreHandler struct {
	ctrl     *gomock.Controller
	recorder *MockBrandStoreHandlerMockRecorder
}

// MockBrandStoreHandlerMockRecorder is the mock recorder for MockBrandStoreHandler.
type MockBrandStoreHandlerMockRecorder struct {
	mock *MockBrandStoreHandler
}

// NewMockBrandStoreHandler creates a new mock instance.
func NewMockBrandStoreHandler(ctrl *gomock.Controller) *MockBrandStoreHandler {
	mock := &MockBrandStoreHandler{ctrl: ctrl}
	mock.recorder = &MockBrandStoreHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBrandStoreHandler) EXPECT() *MockBrandStoreHandlerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockBrandStoreHandler) Create(arg0 entity.Brand) (entity.Brand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(entity.Brand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockBrandStoreHandlerMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBrandStoreHandler)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockBrandStoreHandler) Delete(arg0 int) (entity.Brand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(entity.Brand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockBrandStoreHandlerMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBrandStoreHandler)(nil).Delete), arg0)
}

// GetById mocks base method.
func (m *MockBrandStoreHandler) GetById(arg0 int) (entity.Brand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0)
	ret0, _ := ret[0].(entity.Brand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockBrandStoreHandlerMockRecorder) GetById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockBrandStoreHandler)(nil).GetById), arg0)
}

// Update mocks base method.
func (m *MockBrandStoreHandler) Update(arg0 entity.Brand) (entity.Brand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(entity.Brand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockBrandStoreHandlerMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBrandStoreHandler)(nil).Update), arg0)
}
