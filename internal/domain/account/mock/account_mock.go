// Code generated by MockGen. DO NOT EDIT.
// Source: ./account_interface.go
//
// Generated by this command:
//
//	mockgen -destination=./mock/account_mock.go -package=mock -source=./account_interface.go
//

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	time "time"

	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockEvent is a mock of Event interface.
type MockEvent struct {
	ctrl     *gomock.Controller
	recorder *MockEventMockRecorder
	isgomock struct{}
}

// MockEventMockRecorder is the mock recorder for MockEvent.
type MockEventMockRecorder struct {
	mock *MockEvent
}

// NewMockEvent creates a new mock instance.
func NewMockEvent(ctrl *gomock.Controller) *MockEvent {
	mock := &MockEvent{ctrl: ctrl}
	mock.recorder = &MockEventMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEvent) EXPECT() *MockEventMockRecorder {
	return m.recorder
}

// GetCompletedAt mocks base method.
func (m *MockEvent) GetCompletedAt() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompletedAt")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetCompletedAt indicates an expected call of GetCompletedAt.
func (mr *MockEventMockRecorder) GetCompletedAt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompletedAt", reflect.TypeOf((*MockEvent)(nil).GetCompletedAt))
}

// GetContextID mocks base method.
func (m *MockEvent) GetContextID() uuid.UUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContextID")
	ret0, _ := ret[0].(uuid.UUID)
	return ret0
}

// GetContextID indicates an expected call of GetContextID.
func (mr *MockEventMockRecorder) GetContextID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContextID", reflect.TypeOf((*MockEvent)(nil).GetContextID))
}

// GetCreatedAt mocks base method.
func (m *MockEvent) GetCreatedAt() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreatedAt")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetCreatedAt indicates an expected call of GetCreatedAt.
func (mr *MockEventMockRecorder) GetCreatedAt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreatedAt", reflect.TypeOf((*MockEvent)(nil).GetCreatedAt))
}

// GetEventData mocks base method.
func (m *MockEvent) GetEventData() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventData")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetEventData indicates an expected call of GetEventData.
func (mr *MockEventMockRecorder) GetEventData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventData", reflect.TypeOf((*MockEvent)(nil).GetEventData))
}

// GetID mocks base method.
func (m *MockEvent) GetID() uuid.UUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(uuid.UUID)
	return ret0
}

// GetID indicates an expected call of GetID.
func (mr *MockEventMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockEvent)(nil).GetID))
}

// GetMaxRetry mocks base method.
func (m *MockEvent) GetMaxRetry() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxRetry")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetMaxRetry indicates an expected call of GetMaxRetry.
func (mr *MockEventMockRecorder) GetMaxRetry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxRetry", reflect.TypeOf((*MockEvent)(nil).GetMaxRetry))
}

// GetOrigin mocks base method.
func (m *MockEvent) GetOrigin() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrigin")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetOrigin indicates an expected call of GetOrigin.
func (mr *MockEventMockRecorder) GetOrigin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrigin", reflect.TypeOf((*MockEvent)(nil).GetOrigin))
}

// GetRetry mocks base method.
func (m *MockEvent) GetRetry() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRetry")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetRetry indicates an expected call of GetRetry.
func (mr *MockEventMockRecorder) GetRetry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRetry", reflect.TypeOf((*MockEvent)(nil).GetRetry))
}

// GetScheduledAt mocks base method.
func (m *MockEvent) GetScheduledAt() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScheduledAt")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetScheduledAt indicates an expected call of GetScheduledAt.
func (mr *MockEventMockRecorder) GetScheduledAt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScheduledAt", reflect.TypeOf((*MockEvent)(nil).GetScheduledAt))
}

// GetStartedAt mocks base method.
func (m *MockEvent) GetStartedAt() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStartedAt")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetStartedAt indicates an expected call of GetStartedAt.
func (mr *MockEventMockRecorder) GetStartedAt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStartedAt", reflect.TypeOf((*MockEvent)(nil).GetStartedAt))
}

// GetState mocks base method.
func (m *MockEvent) GetState() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetState")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetState indicates an expected call of GetState.
func (mr *MockEventMockRecorder) GetState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetState", reflect.TypeOf((*MockEvent)(nil).GetState))
}

// GetType mocks base method.
func (m *MockEvent) GetType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetType")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetType indicates an expected call of GetType.
func (mr *MockEventMockRecorder) GetType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetType", reflect.TypeOf((*MockEvent)(nil).GetType))
}

// GetTypeVersion mocks base method.
func (m *MockEvent) GetTypeVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTypeVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTypeVersion indicates an expected call of GetTypeVersion.
func (mr *MockEventMockRecorder) GetTypeVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTypeVersion", reflect.TypeOf((*MockEvent)(nil).GetTypeVersion))
}
