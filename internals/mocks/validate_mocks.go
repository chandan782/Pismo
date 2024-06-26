// Code generated by MockGen. DO NOT EDIT.
// Source: ./validate/validate.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockValidate is a mock of Validate interface.
type MockValidate struct {
	ctrl     *gomock.Controller
	recorder *MockValidateMockRecorder
}

// MockValidateMockRecorder is the mock recorder for MockValidate.
type MockValidateMockRecorder struct {
	mock *MockValidate
}

// NewMockValidate creates a new mock instance.
func NewMockValidate(ctrl *gomock.Controller) *MockValidate {
	mock := &MockValidate{ctrl: ctrl}
	mock.recorder = &MockValidateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidate) EXPECT() *MockValidateMockRecorder {
	return m.recorder
}

// ValidateStructs mocks base method.
func (m *MockValidate) ValidateStructs(req interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateStructs", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateStructs indicates an expected call of ValidateStructs.
func (mr *MockValidateMockRecorder) ValidateStructs(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateStructs", reflect.TypeOf((*MockValidate)(nil).ValidateStructs), req)
}
