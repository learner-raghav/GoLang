package mocks
import (
	entity "Assignment/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockVariantServicer is a mock of VariantServicer interface.
type MockVariantServicer struct {
	ctrl     *gomock.Controller
	recorder *MockVariantServicerMockRecorder
}

// MockVariantServicerMockRecorder is the mock recorder for MockVariantServicer.
type MockVariantServicerMockRecorder struct {
	mock *MockVariantServicer
}

// NewMockVariantServicer creates a new mock instance.
func NewMockVariantServicer(ctrl *gomock.Controller) *MockVariantServicer {
	mock := &MockVariantServicer{ctrl: ctrl}
	mock.recorder = &MockVariantServicerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVariantServicer) EXPECT() *MockVariantServicerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockVariantServicer) Create(arg0 entity.Variant) (entity.Variant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(entity.Variant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockVariantServicerMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVariantServicer)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockVariantServicer) Delete(arg0 int) (entity.Variant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(entity.Variant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockVariantServicerMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVariantServicer)(nil).Delete), arg0)
}

// GetById mocks base method.
func (m *MockVariantServicer) GetById(arg0 int) (entity.Variant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0)
	ret0, _ := ret[0].(entity.Variant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockVariantServicerMockRecorder) GetById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockVariantServicer)(nil).GetById), arg0)
}

// Update mocks base method.
func (m *MockVariantServicer) Update(arg0 entity.Variant) (entity.Variant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(entity.Variant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockVariantServicerMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockVariantServicer)(nil).Update), arg0)
}
