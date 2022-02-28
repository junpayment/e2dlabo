// Code generated by MockGen. DO NOT EDIT.
// Source: db.go

// Package mocks is a generated GoMock package.
package mocks

import (
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSqlDB is a mock of SqlDB interface.
type MockSqlDB struct {
	ctrl     *gomock.Controller
	recorder *MockSqlDBMockRecorder
}

// MockSqlDBMockRecorder is the mock recorder for MockSqlDB.
type MockSqlDBMockRecorder struct {
	mock *MockSqlDB
}

// NewMockSqlDB creates a new mock instance.
func NewMockSqlDB(ctrl *gomock.Controller) *MockSqlDB {
	mock := &MockSqlDB{ctrl: ctrl}
	mock.recorder = &MockSqlDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSqlDB) EXPECT() *MockSqlDBMockRecorder {
	return m.recorder
}

// Exec mocks base method.
func (m *MockSqlDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockSqlDBMockRecorder) Exec(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockSqlDB)(nil).Exec), varargs...)
}