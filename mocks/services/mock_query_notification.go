// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fromsi/jwt-oauth-sso/internal/services (interfaces: QueryNotificationService)
//
// Generated by this command:
//
//	mockgen -destination=../../mocks/services/mock_query_notification.go -package=services_mocks github.com/fromsi/jwt-oauth-sso/internal/services QueryNotificationService
//

// Package services_mocks is a generated GoMock package.
package services_mocks

import (
	gomock "go.uber.org/mock/gomock"
)

// MockQueryNotificationService is a mock of QueryNotificationService interface.
type MockQueryNotificationService struct {
	ctrl     *gomock.Controller
	recorder *MockQueryNotificationServiceMockRecorder
	isgomock struct{}
}

// MockQueryNotificationServiceMockRecorder is the mock recorder for MockQueryNotificationService.
type MockQueryNotificationServiceMockRecorder struct {
	mock *MockQueryNotificationService
}

// NewMockQueryNotificationService creates a new mock instance.
func NewMockQueryNotificationService(ctrl *gomock.Controller) *MockQueryNotificationService {
	mock := &MockQueryNotificationService{ctrl: ctrl}
	mock.recorder = &MockQueryNotificationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueryNotificationService) EXPECT() *MockQueryNotificationServiceMockRecorder {
	return m.recorder
}
