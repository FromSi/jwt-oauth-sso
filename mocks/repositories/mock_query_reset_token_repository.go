// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fromsi/jwt-oauth-sso/internal/repositories (interfaces: QueryResetTokenRepository)
//
// Generated by this command:
//
//	mockgen -destination=../../mocks/repositories/mock_query_reset_token_repository.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories QueryResetTokenRepository
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

// GetActiveResetTokenByToken mocks base method.
func (m *MockQueryResetTokenRepository) GetActiveResetTokenByToken(arg0 string) repositories.ResetToken {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveResetTokenByToken", arg0)
	ret0, _ := ret[0].(repositories.ResetToken)
	return ret0
}

// GetActiveResetTokenByToken indicates an expected call of GetActiveResetTokenByToken.
func (mr *MockQueryResetTokenRepositoryMockRecorder) GetActiveResetTokenByToken(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveResetTokenByToken", reflect.TypeOf((*MockQueryResetTokenRepository)(nil).GetActiveResetTokenByToken), arg0)
}
