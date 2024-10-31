// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fromsi/jwt-oauth-sso/internal/services (interfaces: MutableUserService)
//
// Generated by this command:
//
//	mockgen -destination=../mocks/services/mock_mutable_user_service.go -package=services_mocks github.com/fromsi/jwt-oauth-sso/internal/services MutableUserService
//

// Package services_mocks is a generated GoMock package.
package services_mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockMutableUserService is a mock of MutableUserService interface.
type MockMutableUserService struct {
	ctrl     *gomock.Controller
	recorder *MockMutableUserServiceMockRecorder
	isgomock struct{}
}

// MockMutableUserServiceMockRecorder is the mock recorder for MockMutableUserService.
type MockMutableUserServiceMockRecorder struct {
	mock *MockMutableUserService
}

// NewMockMutableUserService creates a new mock instance.
func NewMockMutableUserService(ctrl *gomock.Controller) *MockMutableUserService {
	mock := &MockMutableUserService{ctrl: ctrl}
	mock.recorder = &MockMutableUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMutableUserService) EXPECT() *MockMutableUserServiceMockRecorder {
	return m.recorder
}

// CreateUserByUUIDAndEmailAndPassword mocks base method.
func (m *MockMutableUserService) CreateUserByUUIDAndEmailAndPassword(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserByUUIDAndEmailAndPassword", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUserByUUIDAndEmailAndPassword indicates an expected call of CreateUserByUUIDAndEmailAndPassword.
func (mr *MockMutableUserServiceMockRecorder) CreateUserByUUIDAndEmailAndPassword(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserByUUIDAndEmailAndPassword", reflect.TypeOf((*MockMutableUserService)(nil).CreateUserByUUIDAndEmailAndPassword), arg0, arg1, arg2)
}