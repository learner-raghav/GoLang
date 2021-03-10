package mocks
import (
	entity "Assignment/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockVariantStoreHandler is a mock of VariantStoreHandler interface.
type MockVariantStoreHandler struct {
	ctrl     *gomock.Controller
	recorder *MockVariantStoreHandlerMockRecorder
}

// MockVariantStoreHandlerMockRecorder is the mock recorder for MockVariantStoreHandler.
type MockVariantStoreHandlerMockRecorder struct {
	mock *MockVariantStoreHandler
}

// NewMockVariantStoreHandler creates a new mock instance.
func NewMockVariantStoreHandler(ctrl *gomock.Controller) *MockVariantStoreHandler {
	mock := &MockVariantStoreHandler{ctrl: ctrl}
	mock.recorder = &MockVariantStoreHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVariantStoreHandler) EXPECT() *MockVariantStoreHandlerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockVariantStoreHandler) Create(arg0 entity.Variant) (entity.Variant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(entity.Variant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockVariantStoreHandlerMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVariantStoreHandler)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockVariantStoreHandler) Delete(arg0 int) (entity.Variant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(entity.Variant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockVariantStoreHandlerMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVariantStoreHandler)(nil).Delete), arg0)
}

// GetById mocks base method.
func (m *MockVariantStoreHandler) GetById(arg0 int) (entity.Variant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0)
	ret0, _ := ret[0].(entity.Variant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockVariantStoreHandlerMockRecorder) GetById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockVariantStoreHandler)(nil).GetById), arg0)
}

// Update mocks base method.
func (m *MockVariantStoreHandler) Update(arg0 entity.Variant) (entity.Variant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(entity.Variant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockVariantStoreHandlerMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockVariantStoreHandler)(nil).Update), arg0)
}