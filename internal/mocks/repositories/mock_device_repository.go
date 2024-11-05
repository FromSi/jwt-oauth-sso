// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fromsi/jwt-oauth-sso/internal/repositories (interfaces: DeviceRepository)
//
// Generated by this command:
//
//	mockgen -destination=../mocks/repositories/mock_device_repository.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories DeviceRepository
//

// Package repositories_mocks is a generated GoMock package.
package repositories_mocks

import (
	reflect "reflect"

	repositories "github.com/fromsi/jwt-oauth-sso/internal/repositories"
	gomock "go.uber.org/mock/gomock"
)

// MockDeviceRepository is a mock of DeviceRepository interface.
type MockDeviceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceRepositoryMockRecorder
	isgomock struct{}
}

// MockDeviceRepositoryMockRecorder is the mock recorder for MockDeviceRepository.
type MockDeviceRepositoryMockRecorder struct {
	mock *MockDeviceRepository
}

// NewMockDeviceRepository creates a new mock instance.
func NewMockDeviceRepository(ctrl *gomock.Controller) *MockDeviceRepository {
	mock := &MockDeviceRepository{ctrl: ctrl}
	mock.recorder = &MockDeviceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeviceRepository) EXPECT() *MockDeviceRepositoryMockRecorder {
	return m.recorder
}

// CreateDevice mocks base method.
func (m *MockDeviceRepository) CreateDevice(arg0 repositories.Device) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDevice", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDevice indicates an expected call of CreateDevice.
func (mr *MockDeviceRepositoryMockRecorder) CreateDevice(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDevice", reflect.TypeOf((*MockDeviceRepository)(nil).CreateDevice), arg0)
}

// DeleteAllDevicesByUserUUID mocks base method.
func (m *MockDeviceRepository) DeleteAllDevicesByUserUUID(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllDevicesByUserUUID", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllDevicesByUserUUID indicates an expected call of DeleteAllDevicesByUserUUID.
func (mr *MockDeviceRepositoryMockRecorder) DeleteAllDevicesByUserUUID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllDevicesByUserUUID", reflect.TypeOf((*MockDeviceRepository)(nil).DeleteAllDevicesByUserUUID), arg0)
}

// DeleteDeviceByUUID mocks base method.
func (m *MockDeviceRepository) DeleteDeviceByUUID(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDeviceByUUID", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDeviceByUUID indicates an expected call of DeleteDeviceByUUID.
func (mr *MockDeviceRepositoryMockRecorder) DeleteDeviceByUUID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDeviceByUUID", reflect.TypeOf((*MockDeviceRepository)(nil).DeleteDeviceByUUID), arg0)
}

// DeleteDeviceByUUIDAndUserUUID mocks base method.
func (m *MockDeviceRepository) DeleteDeviceByUUIDAndUserUUID(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDeviceByUUIDAndUserUUID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDeviceByUUIDAndUserUUID indicates an expected call of DeleteDeviceByUUIDAndUserUUID.
func (mr *MockDeviceRepositoryMockRecorder) DeleteDeviceByUUIDAndUserUUID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDeviceByUUIDAndUserUUID", reflect.TypeOf((*MockDeviceRepository)(nil).DeleteDeviceByUUIDAndUserUUID), arg0, arg1)
}

// GetDeviceByRefreshToken mocks base method.
func (m *MockDeviceRepository) GetDeviceByRefreshToken(arg0 string) repositories.Device {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceByRefreshToken", arg0)
	ret0, _ := ret[0].(repositories.Device)
	return ret0
}

// GetDeviceByRefreshToken indicates an expected call of GetDeviceByRefreshToken.
func (mr *MockDeviceRepositoryMockRecorder) GetDeviceByRefreshToken(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceByRefreshToken", reflect.TypeOf((*MockDeviceRepository)(nil).GetDeviceByRefreshToken), arg0)
}

// GetDeviceByUserUUIDAndIpAndUserAgent mocks base method.
func (m *MockDeviceRepository) GetDeviceByUserUUIDAndIpAndUserAgent(arg0, arg1, arg2 string) repositories.Device {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceByUserUUIDAndIpAndUserAgent", arg0, arg1, arg2)
	ret0, _ := ret[0].(repositories.Device)
	return ret0
}

// GetDeviceByUserUUIDAndIpAndUserAgent indicates an expected call of GetDeviceByUserUUIDAndIpAndUserAgent.
func (mr *MockDeviceRepositoryMockRecorder) GetDeviceByUserUUIDAndIpAndUserAgent(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceByUserUUIDAndIpAndUserAgent", reflect.TypeOf((*MockDeviceRepository)(nil).GetDeviceByUserUUIDAndIpAndUserAgent), arg0, arg1, arg2)
}

// GetDevicesByUserUUID mocks base method.
func (m *MockDeviceRepository) GetDevicesByUserUUID(arg0 string) []repositories.Device {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevicesByUserUUID", arg0)
	ret0, _ := ret[0].([]repositories.Device)
	return ret0
}

// GetDevicesByUserUUID indicates an expected call of GetDevicesByUserUUID.
func (mr *MockDeviceRepositoryMockRecorder) GetDevicesByUserUUID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevicesByUserUUID", reflect.TypeOf((*MockDeviceRepository)(nil).GetDevicesByUserUUID), arg0)
}

// UpdateDevice mocks base method.
func (m *MockDeviceRepository) UpdateDevice(arg0 repositories.Device) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDevice", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDevice indicates an expected call of UpdateDevice.
func (mr *MockDeviceRepositoryMockRecorder) UpdateDevice(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDevice", reflect.TypeOf((*MockDeviceRepository)(nil).UpdateDevice), arg0)
}
