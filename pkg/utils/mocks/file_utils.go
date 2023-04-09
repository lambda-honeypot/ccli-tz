// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/utils/file_utils.go

// Package mock_utils is a generated GoMock package.
package mock_utils

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFileUtilsInterface is a mock of FileUtilsInterface interface.
type MockFileUtilsInterface struct {
	ctrl     *gomock.Controller
	recorder *MockFileUtilsInterfaceMockRecorder
}

// MockFileUtilsInterfaceMockRecorder is the mock recorder for MockFileUtilsInterface.
type MockFileUtilsInterfaceMockRecorder struct {
	mock *MockFileUtilsInterface
}

// NewMockFileUtilsInterface creates a new mock instance.
func NewMockFileUtilsInterface(ctrl *gomock.Controller) *MockFileUtilsInterface {
	mock := &MockFileUtilsInterface{ctrl: ctrl}
	mock.recorder = &MockFileUtilsInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileUtilsInterface) EXPECT() *MockFileUtilsInterfaceMockRecorder {
	return m.recorder
}

// MkDir mocks base method.
func (m *MockFileUtilsInterface) MkDir(path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MkDir", path)
	ret0, _ := ret[0].(error)
	return ret0
}

// MkDir indicates an expected call of MkDir.
func (mr *MockFileUtilsInterfaceMockRecorder) MkDir(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MkDir", reflect.TypeOf((*MockFileUtilsInterface)(nil).MkDir), path)
}

// ReadFile mocks base method.
func (m *MockFileUtilsInterface) ReadFile(path string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFile", path)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFile indicates an expected call of ReadFile.
func (mr *MockFileUtilsInterfaceMockRecorder) ReadFile(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFile", reflect.TypeOf((*MockFileUtilsInterface)(nil).ReadFile), path)
}

// UserHomeDir mocks base method.
func (m *MockFileUtilsInterface) UserHomeDir() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserHomeDir")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserHomeDir indicates an expected call of UserHomeDir.
func (mr *MockFileUtilsInterfaceMockRecorder) UserHomeDir() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserHomeDir", reflect.TypeOf((*MockFileUtilsInterface)(nil).UserHomeDir))
}

// WriteFile mocks base method.
func (m *MockFileUtilsInterface) WriteFile(path string, contents []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteFile", path, contents)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteFile indicates an expected call of WriteFile.
func (mr *MockFileUtilsInterfaceMockRecorder) WriteFile(path, contents interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteFile", reflect.TypeOf((*MockFileUtilsInterface)(nil).WriteFile), path, contents)
}
