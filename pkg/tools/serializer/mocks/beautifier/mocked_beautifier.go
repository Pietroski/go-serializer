// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/pietroski-software-company/tools/serializer/go-serializer/pkg/tools/serializer (interfaces: Beautifier)

// Package mocked_beautifier is a generated GoMock package.
package mocked_beautifier

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBeautifier is a mock of Beautifier interface.
type MockBeautifier struct {
	ctrl     *gomock.Controller
	recorder *MockBeautifierMockRecorder
}

// MockBeautifierMockRecorder is the mock recorder for MockBeautifier.
type MockBeautifierMockRecorder struct {
	mock *MockBeautifier
}

// NewMockBeautifier creates a new mock instance.
func NewMockBeautifier(ctrl *gomock.Controller) *MockBeautifier {
	mock := &MockBeautifier{ctrl: ctrl}
	mock.recorder = &MockBeautifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBeautifier) EXPECT() *MockBeautifierMockRecorder {
	return m.recorder
}

// Beautify mocks base method.
func (m *MockBeautifier) Beautify(arg0 interface{}, arg1, arg2 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Beautify", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Beautify indicates an expected call of Beautify.
func (mr *MockBeautifierMockRecorder) Beautify(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Beautify", reflect.TypeOf((*MockBeautifier)(nil).Beautify), arg0, arg1, arg2)
}

// DataRebind mocks base method.
func (m *MockBeautifier) DataRebind(arg0, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataRebind", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DataRebind indicates an expected call of DataRebind.
func (mr *MockBeautifierMockRecorder) DataRebind(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataRebind", reflect.TypeOf((*MockBeautifier)(nil).DataRebind), arg0, arg1)
}

// Deserialize mocks base method.
func (m *MockBeautifier) Deserialize(arg0 []byte, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deserialize", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deserialize indicates an expected call of Deserialize.
func (mr *MockBeautifierMockRecorder) Deserialize(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deserialize", reflect.TypeOf((*MockBeautifier)(nil).Deserialize), arg0, arg1)
}

// Serialize mocks base method.
func (m *MockBeautifier) Serialize(arg0 interface{}) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Serialize", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Serialize indicates an expected call of Serialize.
func (mr *MockBeautifierMockRecorder) Serialize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Serialize", reflect.TypeOf((*MockBeautifier)(nil).Serialize), arg0)
}
