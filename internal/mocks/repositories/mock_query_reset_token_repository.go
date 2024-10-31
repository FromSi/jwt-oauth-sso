// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fromsi/jwt-oauth-sso/internal/repositories (interfaces: QueryResetTokenRepository)
//
// Generated by this command:
//
//	mockgen -destination=../mocks/repositories/mock_query_reset_token_repository.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories QueryResetTokenRepository
//

// Package repositories_mocks is a generated GoMock package.
package repositories_mocks

import (
	reflect "reflect"

	repositories "github.com/fromsi/jwt-oauth-sso/internal/repositories"
	gomock "go.uber.org/mock/gomock"
)

// MockQueryResetTokenRepository is a mock of QueryResetTokenRepository interface.
type MockQueryResetTokenRepository struct {
	ctrl     *gomock.Controller
	recorder *MockQueryResetTokenRepositoryMockRecorder
	isgomock struct{}
}

// MockQueryResetTokenRepositoryMockRecorder is the mock recorder for MockQueryResetTokenRepository.
type MockQueryResetTokenRepositoryMockRecorder struct {
	mock *MockQueryResetTokenRepository
}

// NewMockQueryResetTokenRepository creates a new mock instance.
func NewMockQueryResetTokenRepository(ctrl *gomock.Controller) *MockQueryResetTokenRepository {
	mock := &MockQueryResetTokenRepository{ctrl: ctrl}
	mock.recorder = &MockQueryResetTokenRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueryResetTokenRepository) EXPECT() *MockQueryResetTokenRepositoryMockRecorder {
	return m.recorder
}

// GetResetTokenByToken mocks base method.
func (m *MockQueryResetTokenRepository) GetResetTokenByToken(arg0 string) repositories.ResetToken {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResetTokenByToken", arg0)
	ret0, _ := ret[0].(repositories.ResetToken)
	return ret0
}

// GetResetTokenByToken indicates an expected call of GetResetTokenByToken.
func (mr *MockQueryResetTokenRepositoryMockRecorder) GetResetTokenByToken(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResetTokenByToken", reflect.TypeOf((*MockQueryResetTokenRepository)(nil).GetResetTokenByToken), arg0)
}

// HasToken mocks base method.
func (m *MockQueryResetTokenRepository) HasToken(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasToken", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasToken indicates an expected call of HasToken.
func (mr *MockQueryResetTokenRepositoryMockRecorder) HasToken(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasToken", reflect.TypeOf((*MockQueryResetTokenRepository)(nil).HasToken), arg0)
}
