// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fromsi/jwt-oauth-sso/internal/repositories (interfaces: ResetTokenRepository)
//
// Generated by this command:
//
//	mockgen -destination=../mocks/repositories/mock_reset_token_repository.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories ResetTokenRepository
//

// Package repositories_mocks is a generated GoMock package.
package repositories_mocks

import (
	reflect "reflect"

	repositories "github.com/fromsi/jwt-oauth-sso/internal/repositories"
	gomock "go.uber.org/mock/gomock"
)

// MockResetTokenRepository is a mock of ResetTokenRepository interface.
type MockResetTokenRepository struct {
	ctrl     *gomock.Controller
	recorder *MockResetTokenRepositoryMockRecorder
	isgomock struct{}
}

// MockResetTokenRepositoryMockRecorder is the mock recorder for MockResetTokenRepository.
type MockResetTokenRepositoryMockRecorder struct {
	mock *MockResetTokenRepository
}

// NewMockResetTokenRepository creates a new mock instance.
func NewMockResetTokenRepository(ctrl *gomock.Controller) *MockResetTokenRepository {
	mock := &MockResetTokenRepository{ctrl: ctrl}
	mock.recorder = &MockResetTokenRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResetTokenRepository) EXPECT() *MockResetTokenRepositoryMockRecorder {
	return m.recorder
}

// CreateResetToken mocks base method.
func (m *MockResetTokenRepository) CreateResetToken(arg0 repositories.ResetToken) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateResetToken", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateResetToken indicates an expected call of CreateResetToken.
func (mr *MockResetTokenRepositoryMockRecorder) CreateResetToken(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateResetToken", reflect.TypeOf((*MockResetTokenRepository)(nil).CreateResetToken), arg0)
}

// DeleteResetToken mocks base method.
func (m *MockResetTokenRepository) DeleteResetToken(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteResetToken", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteResetToken indicates an expected call of DeleteResetToken.
func (mr *MockResetTokenRepositoryMockRecorder) DeleteResetToken(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteResetToken", reflect.TypeOf((*MockResetTokenRepository)(nil).DeleteResetToken), arg0)
}

// GetActiveResetTokenByToken mocks base method.
func (m *MockResetTokenRepository) GetActiveResetTokenByToken(arg0 string) repositories.ResetToken {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveResetTokenByToken", arg0)
	ret0, _ := ret[0].(repositories.ResetToken)
	return ret0
}

// GetActiveResetTokenByToken indicates an expected call of GetActiveResetTokenByToken.
func (mr *MockResetTokenRepositoryMockRecorder) GetActiveResetTokenByToken(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveResetTokenByToken", reflect.TypeOf((*MockResetTokenRepository)(nil).GetActiveResetTokenByToken), arg0)
}

// HasToken mocks base method.
func (m *MockResetTokenRepository) HasToken(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasToken", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasToken indicates an expected call of HasToken.
func (mr *MockResetTokenRepositoryMockRecorder) HasToken(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasToken", reflect.TypeOf((*MockResetTokenRepository)(nil).HasToken), arg0)
}
