// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fromsi/jwt-oauth-sso/internal/services (interfaces: MutableDeviceService)
//
// Generated by this command:
//
//	mockgen -destination=../mocks/services/mock_mutable_device_service.go -package=services_mocks github.com/fromsi/jwt-oauth-sso/internal/services MutableDeviceService
//

// Package services_mocks is a generated GoMock package.
package services_mocks

import (
	gomock "go.uber.org/mock/gomock"
)

// MockMutableDeviceService is a mock of MutableDeviceService interface.
type MockMutableDeviceService struct {
	ctrl     *gomock.Controller
	recorder *MockMutableDeviceServiceMockRecorder
	isgomock struct{}
}

// MockMutableDeviceServiceMockRecorder is the mock recorder for MockMutableDeviceService.
type MockMutableDeviceServiceMockRecorder struct {
	mock *MockMutableDeviceService
}

// NewMockMutableDeviceService creates a new mock instance.
func NewMockMutableDeviceService(ctrl *gomock.Controller) *MockMutableDeviceService {
	mock := &MockMutableDeviceService{ctrl: ctrl}
	mock.recorder = &MockMutableDeviceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMutableDeviceService) EXPECT() *MockMutableDeviceServiceMockRecorder {
	return m.recorder
}
