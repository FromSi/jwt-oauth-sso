// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fromsi/jwt-oauth-sso/internal/services (interfaces: NotificationService)
//
// Generated by this command:
//
//	mockgen -destination=../mocks/services/mock_notification_service.go -package=services_mocks github.com/fromsi/jwt-oauth-sso/internal/services NotificationService
//

// Package services_mocks is a generated GoMock package.
package services_mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockNotificationService is a mock of NotificationService interface.
type MockNotificationService struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationServiceMockRecorder
	isgomock struct{}
}

// MockNotificationServiceMockRecorder is the mock recorder for MockNotificationService.
type MockNotificationServiceMockRecorder struct {
	mock *MockNotificationService
}

// NewMockNotificationService creates a new mock instance.
func NewMockNotificationService(ctrl *gomock.Controller) *MockNotificationService {
	mock := &MockNotificationService{ctrl: ctrl}
	mock.recorder = &MockNotificationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotificationService) EXPECT() *MockNotificationServiceMockRecorder {
	return m.recorder
}

// SendText mocks base method.
func (m *MockNotificationService) SendText(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendText", arg0)
}

// SendText indicates an expected call of SendText.
func (mr *MockNotificationServiceMockRecorder) SendText(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendText", reflect.TypeOf((*MockNotificationService)(nil).SendText), arg0)
}