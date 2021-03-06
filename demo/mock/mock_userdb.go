// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\shilpa.sukumar\Documents\Golang\projects\src\github.com\shilpas131\OVN-Poc\demo\domain\userdb.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserDB is a mock of UserDB interface
type MockUserDB struct {
	ctrl     *gomock.Controller
	recorder *MockUserDBMockRecorder
}

// MockUserDBMockRecorder is the mock recorder for MockUserDB
type MockUserDBMockRecorder struct {
	mock *MockUserDB
}

// NewMockUserDB creates a new mock instance
func NewMockUserDB(ctrl *gomock.Controller) *MockUserDB {
	mock := &MockUserDB{ctrl: ctrl}
	mock.recorder = &MockUserDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserDB) EXPECT() *MockUserDBMockRecorder {
	return m.recorder
}

// UserExists mocks base method
func (m *MockUserDB) UserExists(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserExists", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// UserExists indicates an expected call of UserExists
func (mr *MockUserDBMockRecorder) UserExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserExists", reflect.TypeOf((*MockUserDB)(nil).UserExists), arg0)
}
