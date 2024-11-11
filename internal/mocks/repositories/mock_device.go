// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fromsi/jwt-oauth-sso/internal/repositories (interfaces: Device)
//
// Generated by this command:
//
//	mockgen -destination=../mocks/repositories/mock_device.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories Device
//

// Package repositories_mocks is a generated GoMock package.
package repositories_mocks

import (
	reflect "reflect"

	configs "github.com/fromsi/jwt-oauth-sso/internal/configs"
	tokens "github.com/fromsi/jwt-oauth-sso/internal/tokens"
	gomock "go.uber.org/mock/gomock"
)

// MockDevice is a mock of Device interface.
type MockDevice struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceMockRecorder
	isgomock struct{}
}

// MockDeviceMockRecorder is the mock recorder for MockDevice.
type MockDeviceMockRecorder struct {
	mock *MockDevice
}

// NewMockDevice creates a new mock instance.
func NewMockDevice(ctrl *gomock.Controller) *MockDevice {
	mock := &MockDevice{ctrl: ctrl}
	mock.recorder = &MockDeviceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDevice) EXPECT() *MockDeviceMockRecorder {
	return m.recorder
}

// GenerateAccessToken mocks base method.
func (m *MockDevice) GenerateAccessToken(arg0 configs.TokenConfig) (*tokens.AccessToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAccessToken", arg0)
	ret0, _ := ret[0].(*tokens.AccessToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAccessToken indicates an expected call of GenerateAccessToken.
func (mr *MockDeviceMockRecorder) GenerateAccessToken(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAccessToken", reflect.TypeOf((*MockDevice)(nil).GenerateAccessToken), arg0)
}

// GetCreatedAt mocks base method.
func (m *MockDevice) GetCreatedAt() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreatedAt")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetCreatedAt indicates an expected call of GetCreatedAt.
func (mr *MockDeviceMockRecorder) GetCreatedAt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreatedAt", reflect.TypeOf((*MockDevice)(nil).GetCreatedAt))
}

// GetExpiresAt mocks base method.
func (m *MockDevice) GetExpiresAt() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExpiresAt")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetExpiresAt indicates an expected call of GetExpiresAt.
func (mr *MockDeviceMockRecorder) GetExpiresAt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExpiresAt", reflect.TypeOf((*MockDevice)(nil).GetExpiresAt))
}

// GetIp mocks base method.
func (m *MockDevice) GetIp() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIp")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetIp indicates an expected call of GetIp.
func (mr *MockDeviceMockRecorder) GetIp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIp", reflect.TypeOf((*MockDevice)(nil).GetIp))
}

// GetRefreshToken mocks base method.
func (m *MockDevice) GetRefreshToken() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRefreshToken")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetRefreshToken indicates an expected call of GetRefreshToken.
func (mr *MockDeviceMockRecorder) GetRefreshToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRefreshToken", reflect.TypeOf((*MockDevice)(nil).GetRefreshToken))
}

// GetUUID mocks base method.
func (m *MockDevice) GetUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetUUID indicates an expected call of GetUUID.
func (mr *MockDeviceMockRecorder) GetUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUUID", reflect.TypeOf((*MockDevice)(nil).GetUUID))
}

// GetUpdatedAt mocks base method.
func (m *MockDevice) GetUpdatedAt() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpdatedAt")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetUpdatedAt indicates an expected call of GetUpdatedAt.
func (mr *MockDeviceMockRecorder) GetUpdatedAt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpdatedAt", reflect.TypeOf((*MockDevice)(nil).GetUpdatedAt))
}

// GetUserAgent mocks base method.
func (m *MockDevice) GetUserAgent() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserAgent")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetUserAgent indicates an expected call of GetUserAgent.
func (mr *MockDeviceMockRecorder) GetUserAgent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAgent", reflect.TypeOf((*MockDevice)(nil).GetUserAgent))
}

// GetUserUUID mocks base method.
func (m *MockDevice) GetUserUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetUserUUID indicates an expected call of GetUserUUID.
func (mr *MockDeviceMockRecorder) GetUserUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserUUID", reflect.TypeOf((*MockDevice)(nil).GetUserUUID))
}

// SetCreatedAt mocks base method.
func (m *MockDevice) SetCreatedAt(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCreatedAt", arg0)
}

// SetCreatedAt indicates an expected call of SetCreatedAt.
func (mr *MockDeviceMockRecorder) SetCreatedAt(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCreatedAt", reflect.TypeOf((*MockDevice)(nil).SetCreatedAt), arg0)
}

// SetExpiresAt mocks base method.
func (m *MockDevice) SetExpiresAt(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetExpiresAt", arg0)
}

// SetExpiresAt indicates an expected call of SetExpiresAt.
func (mr *MockDeviceMockRecorder) SetExpiresAt(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetExpiresAt", reflect.TypeOf((*MockDevice)(nil).SetExpiresAt), arg0)
}

// SetIp mocks base method.
func (m *MockDevice) SetIp(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetIp", arg0)
}

// SetIp indicates an expected call of SetIp.
func (mr *MockDeviceMockRecorder) SetIp(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetIp", reflect.TypeOf((*MockDevice)(nil).SetIp), arg0)
}

// SetRefreshToken mocks base method.
func (m *MockDevice) SetRefreshToken(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRefreshToken", arg0)
}

// SetRefreshToken indicates an expected call of SetRefreshToken.
func (mr *MockDeviceMockRecorder) SetRefreshToken(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRefreshToken", reflect.TypeOf((*MockDevice)(nil).SetRefreshToken), arg0)
}

// SetUUID mocks base method.
func (m *MockDevice) SetUUID(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUUID", arg0)
}

// SetUUID indicates an expected call of SetUUID.
func (mr *MockDeviceMockRecorder) SetUUID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUUID", reflect.TypeOf((*MockDevice)(nil).SetUUID), arg0)
}

// SetUpdatedAt mocks base method.
func (m *MockDevice) SetUpdatedAt(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUpdatedAt", arg0)
}

// SetUpdatedAt indicates an expected call of SetUpdatedAt.
func (mr *MockDeviceMockRecorder) SetUpdatedAt(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUpdatedAt", reflect.TypeOf((*MockDevice)(nil).SetUpdatedAt), arg0)
}

// SetUserAgent mocks base method.
func (m *MockDevice) SetUserAgent(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUserAgent", arg0)
}

// SetUserAgent indicates an expected call of SetUserAgent.
func (mr *MockDeviceMockRecorder) SetUserAgent(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserAgent", reflect.TypeOf((*MockDevice)(nil).SetUserAgent), arg0)
}

// SetUserUUID mocks base method.
func (m *MockDevice) SetUserUUID(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUserUUID", arg0)
}

// SetUserUUID indicates an expected call of SetUserUUID.
func (mr *MockDeviceMockRecorder) SetUserUUID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserUUID", reflect.TypeOf((*MockDevice)(nil).SetUserUUID), arg0)
}
