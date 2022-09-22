// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/leader/leader.go

// Package mock_leader is a generated GoMock package.
package mock_leader

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCommandRunner is a mock of CommandRunner interface.
type MockCommandRunner struct {
	ctrl     *gomock.Controller
	recorder *MockCommandRunnerMockRecorder
}

// MockCommandRunnerMockRecorder is the mock recorder for MockCommandRunner.
type MockCommandRunnerMockRecorder struct {
	mock *MockCommandRunner
}

// NewMockCommandRunner creates a new mock instance.
func NewMockCommandRunner(ctrl *gomock.Controller) *MockCommandRunner {
	mock := &MockCommandRunner{ctrl: ctrl}
	mock.recorder = &MockCommandRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommandRunner) EXPECT() *MockCommandRunnerMockRecorder {
	return m.recorder
}

// GetSchedule mocks base method.
func (m *MockCommandRunner) GetSchedule(period, shelleyGenesisFile, poolId, vrfKeysFile, testnetMagic string, dryRun bool) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSchedule", period, shelleyGenesisFile, poolId, vrfKeysFile, testnetMagic, dryRun)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchedule indicates an expected call of GetSchedule.
func (mr *MockCommandRunnerMockRecorder) GetSchedule(period, shelleyGenesisFile, poolId, vrfKeysFile, testnetMagic, dryRun interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchedule", reflect.TypeOf((*MockCommandRunner)(nil).GetSchedule), period, shelleyGenesisFile, poolId, vrfKeysFile, testnetMagic, dryRun)
}
