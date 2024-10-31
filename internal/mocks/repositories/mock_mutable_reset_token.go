// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fromsi/jwt-oauth-sso/internal/repositories (interfaces: MutableResetToken)
//
// Generated by this command:
//
//	mockgen -destination=../mocks/repositories/mock_mutable_reset_token.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories MutableResetToken
//

// Package repositories_mocks is a generated GoMock package.
package repositories_mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockMutableResetToken is a mock of MutableResetToken interface.
type MockMutableResetToken struct {
	ctrl     *gomock.Controller
	recorder *MockMutableResetTokenMockRecorder
	isgomock struct{}
}

// MockMutableResetTokenMockRecorder is the mock recorder for MockMutableResetToken.
type MockMutableResetTokenMockRecorder struct {
	mock *MockMutableResetToken
}

// NewMockMutableResetToken creates a new mock instance.
func NewMockMutableResetToken(ctrl *gomock.Controller) *MockMutableResetToken {
	mock := &MockMutableResetToken{ctrl: ctrl}
	mock.recorder = &MockMutableResetTokenMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMutableResetToken) EXPECT() *MockMutableResetTokenMockRecorder {
	return m.recorder
}

// SetCreatedAt mocks base method.
func (m *MockMutableResetToken) SetCreatedAt(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCreatedAt", arg0)
}

// SetCreatedAt indicates an expected call of SetCreatedAt.
func (mr *MockMutableResetTokenMockRecorder) SetCreatedAt(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCreatedAt", reflect.TypeOf((*MockMutableResetToken)(nil).SetCreatedAt), arg0)
}

// SetExpiredAt mocks base method.
func (m *MockMutableResetToken) SetExpiredAt(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetExpiredAt", arg0)
}

// SetExpiredAt indicates an expected call of SetExpiredAt.
func (mr *MockMutableResetTokenMockRecorder) SetExpiredAt(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetExpiredAt", reflect.TypeOf((*MockMutableResetToken)(nil).SetExpiredAt), arg0)
}

// SetToken mocks base method.
func (m *MockMutableResetToken) SetToken(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetToken", arg0)
}

// SetToken indicates an expected call of SetToken.
func (mr *MockMutableResetTokenMockRecorder) SetToken(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetToken", reflect.TypeOf((*MockMutableResetToken)(nil).SetToken), arg0)
}

// SetUserUUID mocks base method.
func (m *MockMutableResetToken) SetUserUUID(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUserUUID", arg0)
}

// SetUserUUID indicates an expected call of SetUserUUID.
func (mr *MockMutableResetTokenMockRecorder) SetUserUUID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserUUID", reflect.TypeOf((*MockMutableResetToken)(nil).SetUserUUID), arg0)
}