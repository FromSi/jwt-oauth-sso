// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fromsi/jwt-oauth-sso/internal/repositories (interfaces: MutableDevice)
//
// Generated by this command:
//
//	mockgen -destination=../mocks/repositories/mock_mutable_device.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories MutableDevice
//

// Package repositories_mocks is a generated GoMock package.
package repositories_mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockMutableDevice is a mock of MutableDevice interface.
type MockMutableDevice struct {
	ctrl     *gomock.Controller
	recorder *MockMutableDeviceMockRecorder
	isgomock struct{}
}

// MockMutableDeviceMockRecorder is the mock recorder for MockMutableDevice.
type MockMutableDeviceMockRecorder struct {
	mock *MockMutableDevice
}

// NewMockMutableDevice creates a new mock instance.
func NewMockMutableDevice(ctrl *gomock.Controller) *MockMutableDevice {
	mock := &MockMutableDevice{ctrl: ctrl}
	mock.recorder = &MockMutableDeviceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMutableDevice) EXPECT() *MockMutableDeviceMockRecorder {
	return m.recorder
}

// SetCreatedAt mocks base method.
func (m *MockMutableDevice) SetCreatedAt(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCreatedAt", arg0)
}

// SetCreatedAt indicates an expected call of SetCreatedAt.
func (mr *MockMutableDeviceMockRecorder) SetCreatedAt(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCreatedAt", reflect.TypeOf((*MockMutableDevice)(nil).SetCreatedAt), arg0)
}

// SetExpiredAt mocks base method.
func (m *MockMutableDevice) SetExpiredAt(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetExpiredAt", arg0)
}

// SetExpiredAt indicates an expected call of SetExpiredAt.
func (mr *MockMutableDeviceMockRecorder) SetExpiredAt(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetExpiredAt", reflect.TypeOf((*MockMutableDevice)(nil).SetExpiredAt), arg0)
}

// SetIp mocks base method.
func (m *MockMutableDevice) SetIp(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetIp", arg0)
}

// SetIp indicates an expected call of SetIp.
func (mr *MockMutableDeviceMockRecorder) SetIp(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetIp", reflect.TypeOf((*MockMutableDevice)(nil).SetIp), arg0)
}

// SetRefreshToken mocks base method.
func (m *MockMutableDevice) SetRefreshToken(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRefreshToken", arg0)
}

// SetRefreshToken indicates an expected call of SetRefreshToken.
func (mr *MockMutableDeviceMockRecorder) SetRefreshToken(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRefreshToken", reflect.TypeOf((*MockMutableDevice)(nil).SetRefreshToken), arg0)
}

// SetUUID mocks base method.
func (m *MockMutableDevice) SetUUID(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUUID", arg0)
}

// SetUUID indicates an expected call of SetUUID.
func (mr *MockMutableDeviceMockRecorder) SetUUID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUUID", reflect.TypeOf((*MockMutableDevice)(nil).SetUUID), arg0)
}

// SetUpdatedAt mocks base method.
func (m *MockMutableDevice) SetUpdatedAt(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUpdatedAt", arg0)
}

// SetUpdatedAt indicates an expected call of SetUpdatedAt.
func (mr *MockMutableDeviceMockRecorder) SetUpdatedAt(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUpdatedAt", reflect.TypeOf((*MockMutableDevice)(nil).SetUpdatedAt), arg0)
}

// SetUserAgent mocks base method.
func (m *MockMutableDevice) SetUserAgent(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUserAgent", arg0)
}

// SetUserAgent indicates an expected call of SetUserAgent.
func (mr *MockMutableDeviceMockRecorder) SetUserAgent(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserAgent", reflect.TypeOf((*MockMutableDevice)(nil).SetUserAgent), arg0)
}

// SetUserUUID mocks base method.
func (m *MockMutableDevice) SetUserUUID(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUserUUID", arg0)
}

// SetUserUUID indicates an expected call of SetUserUUID.
func (mr *MockMutableDeviceMockRecorder) SetUserUUID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserUUID", reflect.TypeOf((*MockMutableDevice)(nil).SetUserUUID), arg0)
}