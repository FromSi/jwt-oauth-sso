// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fromsi/jwt-oauth-sso/internal/repositories (interfaces: QueryDeviceRepository)
//
// Generated by this command:
//
//	mockgen -destination=../mocks/repositories/mock_query_device_repository.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories QueryDeviceRepository
//

// Package repositories_mocks is a generated GoMock package.
package repositories_mocks

import (
	reflect "reflect"

	repositories "github.com/fromsi/jwt-oauth-sso/internal/repositories"
	gomock "go.uber.org/mock/gomock"
)

// MockQueryDeviceRepository is a mock of QueryDeviceRepository interface.
type MockQueryDeviceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockQueryDeviceRepositoryMockRecorder
	isgomock struct{}
}

// MockQueryDeviceRepositoryMockRecorder is the mock recorder for MockQueryDeviceRepository.
type MockQueryDeviceRepositoryMockRecorder struct {
	mock *MockQueryDeviceRepository
}

// NewMockQueryDeviceRepository creates a new mock instance.
func NewMockQueryDeviceRepository(ctrl *gomock.Controller) *MockQueryDeviceRepository {
	mock := &MockQueryDeviceRepository{ctrl: ctrl}
	mock.recorder = &MockQueryDeviceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueryDeviceRepository) EXPECT() *MockQueryDeviceRepositoryMockRecorder {
	return m.recorder
}

// GetDeviceByUserUUIDAndIpAndUserAgent mocks base method.
func (m *MockQueryDeviceRepository) GetDeviceByUserUUIDAndIpAndUserAgent(arg0, arg1, arg2 string) repositories.Device {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceByUserUUIDAndIpAndUserAgent", arg0, arg1, arg2)
	ret0, _ := ret[0].(repositories.Device)
	return ret0
}

// GetDeviceByUserUUIDAndIpAndUserAgent indicates an expected call of GetDeviceByUserUUIDAndIpAndUserAgent.
func (mr *MockQueryDeviceRepositoryMockRecorder) GetDeviceByUserUUIDAndIpAndUserAgent(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceByUserUUIDAndIpAndUserAgent", reflect.TypeOf((*MockQueryDeviceRepository)(nil).GetDeviceByUserUUIDAndIpAndUserAgent), arg0, arg1, arg2)
}

// GetDevicesByUserUUID mocks base method.
func (m *MockQueryDeviceRepository) GetDevicesByUserUUID(arg0 string) []repositories.Device {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevicesByUserUUID", arg0)
	ret0, _ := ret[0].([]repositories.Device)
	return ret0
}

// GetDevicesByUserUUID indicates an expected call of GetDevicesByUserUUID.
func (mr *MockQueryDeviceRepositoryMockRecorder) GetDevicesByUserUUID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevicesByUserUUID", reflect.TypeOf((*MockQueryDeviceRepository)(nil).GetDevicesByUserUUID), arg0)
}