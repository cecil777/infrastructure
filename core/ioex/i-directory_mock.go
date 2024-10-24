// Code generated by MockGen. DO NOT EDIT.
// Source: ioex\i-directory.go

// Package ioex is a generated GoMock package.
package ioex

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIDirectory is a mock of IDirectory interface.
type MockIDirectory struct {
	ctrl     *gomock.Controller
	recorder *MockIDirectoryMockRecorder
}

// MockIDirectoryMockRecorder is the mock recorder for MockIDirectory.
type MockIDirectoryMockRecorder struct {
	mock *MockIDirectory
}

// NewMockIDirectory creates a new mock instance.
func NewMockIDirectory(ctrl *gomock.Controller) *MockIDirectory {
	mock := &MockIDirectory{ctrl: ctrl}
	mock.recorder = &MockIDirectoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDirectory) EXPECT() *MockIDirectoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIDirectory) Create() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create")
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIDirectoryMockRecorder) Create() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIDirectory)(nil).Create))
}

// FindDirectories mocks base method.
func (m *MockIDirectory) FindDirectories() []IDirectory {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindDirectories")
	ret0, _ := ret[0].([]IDirectory)
	return ret0
}

// FindDirectories indicates an expected call of FindDirectories.
func (mr *MockIDirectoryMockRecorder) FindDirectories() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDirectories", reflect.TypeOf((*MockIDirectory)(nil).FindDirectories))
}

// FindFiles mocks base method.
func (m *MockIDirectory) FindFiles() []IFile {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindFiles")
	ret0, _ := ret[0].([]IFile)
	return ret0
}

// FindFiles indicates an expected call of FindFiles.
func (mr *MockIDirectoryMockRecorder) FindFiles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindFiles", reflect.TypeOf((*MockIDirectory)(nil).FindFiles))
}

// GetName mocks base method.
func (m *MockIDirectory) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockIDirectoryMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockIDirectory)(nil).GetName))
}

// GetParent mocks base method.
func (m *MockIDirectory) GetParent() IDirectory {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParent")
	ret0, _ := ret[0].(IDirectory)
	return ret0
}

// GetParent indicates an expected call of GetParent.
func (mr *MockIDirectoryMockRecorder) GetParent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParent", reflect.TypeOf((*MockIDirectory)(nil).GetParent))
}

// GetPath mocks base method.
func (m *MockIDirectory) GetPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPath indicates an expected call of GetPath.
func (mr *MockIDirectoryMockRecorder) GetPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPath", reflect.TypeOf((*MockIDirectory)(nil).GetPath))
}

// IsExist mocks base method.
func (m *MockIDirectory) IsExist() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsExist")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsExist indicates an expected call of IsExist.
func (mr *MockIDirectoryMockRecorder) IsExist() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExist", reflect.TypeOf((*MockIDirectory)(nil).IsExist))
}

// Move mocks base method.
func (m *MockIDirectory) Move(pathArgs ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range pathArgs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Move", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Move indicates an expected call of Move.
func (mr *MockIDirectoryMockRecorder) Move(pathArgs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Move", reflect.TypeOf((*MockIDirectory)(nil).Move), pathArgs...)
}

// Remove mocks base method.
func (m *MockIDirectory) Remove() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove")
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockIDirectoryMockRecorder) Remove() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockIDirectory)(nil).Remove))
}
