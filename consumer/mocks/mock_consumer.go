// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Lubwama-Emmanuel/Kafka-and-CLIs/consumer (interfaces: Provider)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	config "github.com/Lubwama-Emmanuel/Kafka-and-CLIs/config"
	models "github.com/Lubwama-Emmanuel/Kafka-and-CLIs/models"
	gomock "github.com/golang/mock/gomock"
)

// MockProvider is a mock of Provider interface.
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider.
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance.
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockProvider) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockProviderMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockProvider)(nil).Close))
}

// ReadMessage mocks base method.
func (m *MockProvider) ReadMessage(arg0 time.Duration) (models.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadMessage", arg0)
	ret0, _ := ret[0].(models.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadMessage indicates an expected call of ReadMessage.
func (mr *MockProviderMockRecorder) ReadMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMessage", reflect.TypeOf((*MockProvider)(nil).ReadMessage), arg0)
}

// SetUp mocks base method.
func (m *MockProvider) SetUp(arg0 config.ProviderConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUp", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUp indicates an expected call of SetUp.
func (mr *MockProviderMockRecorder) SetUp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUp", reflect.TypeOf((*MockProvider)(nil).SetUp), arg0)
}

// Subscribe mocks base method.
func (m *MockProvider) Subscribe(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockProviderMockRecorder) Subscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockProvider)(nil).Subscribe), arg0)
}