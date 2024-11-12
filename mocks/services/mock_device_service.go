// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fromsi/jwt-oauth-sso/internal/services (interfaces: DeviceService)
//
// Generated by this command:
//
//	mockgen -destination=../../mocks/services/mock_device_service.go -package=services_mocks github.com/fromsi/jwt-oauth-sso/internal/services DeviceService
//

// Package services_mocks is a generated GoMock package.
package services_mocks

import (
	reflect "reflect"

	repositories "github.com/fromsi/jwt-oauth-sso/internal/repositories"
	gomock "go.uber.org/mock/gomock"
)

// MockDeviceService is a mock of DeviceService interface.
type MockDeviceService struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceServiceMockRecorder
	isgomock struct{}
}

// MockDeviceServiceMockRecorder is the mock recorder for MockDeviceService.
type MockDeviceServiceMockRecorder struct {
	mock *MockDeviceService
}

// NewMockDeviceService creates a new mock instance.
func NewMockDeviceService(ctrl *gomock.Controller) *MockDeviceService {
	mock := &MockDeviceService{ctrl: ctrl}
	mock.recorder = &MockDeviceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeviceService) EXPECT() *MockDeviceServiceMockRecorder {
	return m.recorder
}

// GenerateRefreshToken mocks base method.
func (m *MockDeviceService) GenerateRefreshToken() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRefreshToken")
	ret0, _ := ret[0].(string)
	return ret0
}

// GenerateRefreshToken indicates an expected call of GenerateRefreshToken.
func (mr *MockDeviceServiceMockRecorder) GenerateRefreshToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRefreshToken", reflect.TypeOf((*MockDeviceService)(nil).GenerateRefreshToken))
}

// GenerateUUID mocks base method.
func (m *MockDeviceService) GenerateUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GenerateUUID indicates an expected call of GenerateUUID.
func (mr *MockDeviceServiceMockRecorder) GenerateUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateUUID", reflect.TypeOf((*MockDeviceService)(nil).GenerateUUID))
}

// GetNewDeviceByUserUUIDAndIpAndUserAgent mocks base method.
func (m *MockDeviceService) GetNewDeviceByUserUUIDAndIpAndUserAgent(arg0, arg1, arg2 string) repositories.Device {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNewDeviceByUserUUIDAndIpAndUserAgent", arg0, arg1, arg2)
	ret0, _ := ret[0].(repositories.Device)
	return ret0
}

// GetNewDeviceByUserUUIDAndIpAndUserAgent indicates an expected call of GetNewDeviceByUserUUIDAndIpAndUserAgent.
func (mr *MockDeviceServiceMockRecorder) GetNewDeviceByUserUUIDAndIpAndUserAgent(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNewDeviceByUserUUIDAndIpAndUserAgent", reflect.TypeOf((*MockDeviceService)(nil).GetNewDeviceByUserUUIDAndIpAndUserAgent), arg0, arg1, arg2)
}

// GetNewRefreshDetailsByDevice mocks base method.
func (m *MockDeviceService) GetNewRefreshDetailsByDevice(arg0 repositories.Device) repositories.Device {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNewRefreshDetailsByDevice", arg0)
	ret0, _ := ret[0].(repositories.Device)
	return ret0
}

// GetNewRefreshDetailsByDevice indicates an expected call of GetNewRefreshDetailsByDevice.
func (mr *MockDeviceServiceMockRecorder) GetNewRefreshDetailsByDevice(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNewRefreshDetailsByDevice", reflect.TypeOf((*MockDeviceService)(nil).GetNewRefreshDetailsByDevice), arg0)
}

// GetOldDeviceByUserUUIDAndIpAndUserAgent mocks base method.
func (m *MockDeviceService) GetOldDeviceByUserUUIDAndIpAndUserAgent(arg0, arg1, arg2 string) repositories.Device {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOldDeviceByUserUUIDAndIpAndUserAgent", arg0, arg1, arg2)
	ret0, _ := ret[0].(repositories.Device)
	return ret0
}

// GetOldDeviceByUserUUIDAndIpAndUserAgent indicates an expected call of GetOldDeviceByUserUUIDAndIpAndUserAgent.
func (mr *MockDeviceServiceMockRecorder) GetOldDeviceByUserUUIDAndIpAndUserAgent(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOldDeviceByUserUUIDAndIpAndUserAgent", reflect.TypeOf((*MockDeviceService)(nil).GetOldDeviceByUserUUIDAndIpAndUserAgent), arg0, arg1, arg2)
}
