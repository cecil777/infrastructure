// Code generated by MockGen. DO NOT EDIT.
// Source: .\i-unit-of-work.go

// Package db is a generated GoMock package.
package db

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIUnitOfWork is a mock of IUnitOfWork interface.
type MockIUnitOfWork struct {
	ctrl     *gomock.Controller
	recorder *MockIUnitOfWorkMockRecorder
}

// MockIUnitOfWorkMockRecorder is the mock recorder for MockIUnitOfWork.
type MockIUnitOfWorkMockRecorder struct {
	mock *MockIUnitOfWork
}

// NewMockIUnitOfWork creates a new mock instance.
func NewMockIUnitOfWork(ctrl *gomock.Controller) *MockIUnitOfWork {
	mock := &MockIUnitOfWork{ctrl: ctrl}
	mock.recorder = &MockIUnitOfWorkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUnitOfWork) EXPECT() *MockIUnitOfWorkMockRecorder {
	return m.recorder
}

// Commit mocks base method.
func (m *MockIUnitOfWork) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockIUnitOfWorkMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockIUnitOfWork)(nil).Commit))
}
