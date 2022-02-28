// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLib3 is a mock of Lib3 interface.
type MockLib3 struct {
	ctrl     *gomock.Controller
	recorder *MockLib3MockRecorder
}

// MockLib3MockRecorder is the mock recorder for MockLib3.
type MockLib3MockRecorder struct {
	mock *MockLib3
}

// NewMockLib3 creates a new mock instance.
func NewMockLib3(ctrl *gomock.Controller) *MockLib3 {
	mock := &MockLib3{ctrl: ctrl}
	mock.recorder = &MockLib3MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLib3) EXPECT() *MockLib3MockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockLib3) Do() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do")
	ret0, _ := ret[0].(error)
	return ret0
}

// Do indicates an expected call of Do.
func (mr *MockLib3MockRecorder) Do() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockLib3)(nil).Do))
}
