// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/controller/partitioner.go

// Package controller is a generated GoMock package.
package controller

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockWorkPartitioner is a mock of WorkPartitioner interface
type MockWorkPartitioner struct {
	ctrl     *gomock.Controller
	recorder *MockWorkPartitionerMockRecorder
}

// MockWorkPartitionerMockRecorder is the mock recorder for MockWorkPartitioner
type MockWorkPartitionerMockRecorder struct {
	mock *MockWorkPartitioner
}

// NewMockWorkPartitioner creates a new mock instance
func NewMockWorkPartitioner(ctrl *gomock.Controller) *MockWorkPartitioner {
	mock := &MockWorkPartitioner{ctrl: ctrl}
	mock.recorder = &MockWorkPartitionerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWorkPartitioner) EXPECT() *MockWorkPartitionerMockRecorder {
	return m.recorder
}

// Partition mocks base method
func (m *MockWorkPartitioner) Partition(id ID) (PartitionKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Partition", id)
	ret0, _ := ret[0].(PartitionKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Partition indicates an expected call of Partition
func (mr *MockWorkPartitionerMockRecorder) Partition(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Partition", reflect.TypeOf((*MockWorkPartitioner)(nil).Partition), id)
}
