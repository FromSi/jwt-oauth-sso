// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fromsi/jwt-oauth-sso/internal/repositories (interfaces: QueryUserRepository)
//
// Generated by this command:
//
//	mockgen -destination=../mocks/repositories/mock_query_user_repository.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories QueryUserRepository
//

// Package repositories_mocks is a generated GoMock package.
package repositories_mocks

import (
	reflect "reflect"

	repositories "github.com/fromsi/jwt-oauth-sso/internal/repositories"
	gomock "go.uber.org/mock/gomock"
)

// MockQueryUserRepository is a mock of QueryUserRepository interface.
type MockQueryUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockQueryUserRepositoryMockRecorder
	isgomock struct{}
}

// MockQueryUserRepositoryMockRecorder is the mock recorder for MockQueryUserRepository.
type MockQueryUserRepositoryMockRecorder struct {
	mock *MockQueryUserRepository
}

// NewMockQueryUserRepository creates a new mock instance.
func NewMockQueryUserRepository(ctrl *gomock.Controller) *MockQueryUserRepository {
	mock := &MockQueryUserRepository{ctrl: ctrl}
	mock.recorder = &MockQueryUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueryUserRepository) EXPECT() *MockQueryUserRepositoryMockRecorder {
	return m.recorder
}

// GetUserByEmail mocks base method.
func (m *MockQueryUserRepository) GetUserByEmail(arg0 string) repositories.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0)
	ret0, _ := ret[0].(repositories.User)
	return ret0
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockQueryUserRepositoryMockRecorder) GetUserByEmail(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockQueryUserRepository)(nil).GetUserByEmail), arg0)
}

// HasUserByEmail mocks base method.
func (m *MockQueryUserRepository) HasUserByEmail(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasUserByEmail", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasUserByEmail indicates an expected call of HasUserByEmail.
func (mr *MockQueryUserRepositoryMockRecorder) HasUserByEmail(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasUserByEmail", reflect.TypeOf((*MockQueryUserRepository)(nil).HasUserByEmail), arg0)
}

// HasUserByEmailAndPassword mocks base method.
func (m *MockQueryUserRepository) HasUserByEmailAndPassword(arg0, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasUserByEmailAndPassword", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasUserByEmailAndPassword indicates an expected call of HasUserByEmailAndPassword.
func (mr *MockQueryUserRepositoryMockRecorder) HasUserByEmailAndPassword(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasUserByEmailAndPassword", reflect.TypeOf((*MockQueryUserRepository)(nil).HasUserByEmailAndPassword), arg0, arg1)
}

// HasUserByUUID mocks base method.
func (m *MockQueryUserRepository) HasUserByUUID(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasUserByUUID", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasUserByUUID indicates an expected call of HasUserByUUID.
func (mr *MockQueryUserRepositoryMockRecorder) HasUserByUUID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasUserByUUID", reflect.TypeOf((*MockQueryUserRepository)(nil).HasUserByUUID), arg0)
}

// HasUserByUUIDAndPassword mocks base method.
func (m *MockQueryUserRepository) HasUserByUUIDAndPassword(arg0, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasUserByUUIDAndPassword", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasUserByUUIDAndPassword indicates an expected call of HasUserByUUIDAndPassword.
func (mr *MockQueryUserRepositoryMockRecorder) HasUserByUUIDAndPassword(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasUserByUUIDAndPassword", reflect.TypeOf((*MockQueryUserRepository)(nil).HasUserByUUIDAndPassword), arg0, arg1)
}
