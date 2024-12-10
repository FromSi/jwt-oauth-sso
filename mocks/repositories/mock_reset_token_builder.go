// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fromsi/jwt-oauth-sso/internal/repositories (interfaces: ResetTokenBuilder)
//
// Generated by this command:
//
//	mockgen -destination=../../mocks/repositories/mock_reset_token_builder.go -package=repositories_mocks github.com/fromsi/jwt-oauth-sso/internal/repositories ResetTokenBuilder
//

// Package repositories_mocks is a generated GoMock package.
package repositories_mocks

import (
	reflect "reflect"

	repositories "github.com/fromsi/jwt-oauth-sso/internal/repositories"
	gomock "go.uber.org/mock/gomock"
)

// MockResetTokenBuilder is a mock of ResetTokenBuilder interface.
type MockResetTokenBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockResetTokenBuilderMockRecorder
	isgomock struct{}
}

// MockResetTokenBuilderMockRecorder is the mock recorder for MockResetTokenBuilder.
type MockResetTokenBuilderMockRecorder struct {
	mock *MockResetTokenBuilder
}

// NewMockResetTokenBuilder creates a new mock instance.
func NewMockResetTokenBuilder(ctrl *gomock.Controller) *MockResetTokenBuilder {
	mock := &MockResetTokenBuilder{ctrl: ctrl}
	mock.recorder = &MockResetTokenBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResetTokenBuilder) EXPECT() *MockResetTokenBuilderMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockResetTokenBuilder) Build() (repositories.ResetToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build")
	ret0, _ := ret[0].(repositories.ResetToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Build indicates an expected call of Build.
func (mr *MockResetTokenBuilderMockRecorder) Build() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockResetTokenBuilder)(nil).Build))
}

// BuildToGorm mocks base method.
func (m *MockResetTokenBuilder) BuildToGorm() (*repositories.GormResetToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildToGorm")
	ret0, _ := ret[0].(*repositories.GormResetToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildToGorm indicates an expected call of BuildToGorm.
func (mr *MockResetTokenBuilderMockRecorder) BuildToGorm() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildToGorm", reflect.TypeOf((*MockResetTokenBuilder)(nil).BuildToGorm))
}

// New mocks base method.
func (m *MockResetTokenBuilder) New() repositories.ResetTokenBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New")
	ret0, _ := ret[0].(repositories.ResetTokenBuilder)
	return ret0
}

// New indicates an expected call of New.
func (mr *MockResetTokenBuilderMockRecorder) New() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockResetTokenBuilder)(nil).New))
}

// NewFromResetToken mocks base method.
func (m *MockResetTokenBuilder) NewFromResetToken(arg0 repositories.ResetToken) repositories.ResetTokenBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewFromResetToken", arg0)
	ret0, _ := ret[0].(repositories.ResetTokenBuilder)
	return ret0
}

// NewFromResetToken indicates an expected call of NewFromResetToken.
func (mr *MockResetTokenBuilderMockRecorder) NewFromResetToken(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewFromResetToken", reflect.TypeOf((*MockResetTokenBuilder)(nil).NewFromResetToken), arg0)
}

// SetCreatedAt mocks base method.
func (m *MockResetTokenBuilder) SetCreatedAt(arg0 int) repositories.ResetTokenBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCreatedAt", arg0)
	ret0, _ := ret[0].(repositories.ResetTokenBuilder)
	return ret0
}

// SetCreatedAt indicates an expected call of SetCreatedAt.
func (mr *MockResetTokenBuilderMockRecorder) SetCreatedAt(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCreatedAt", reflect.TypeOf((*MockResetTokenBuilder)(nil).SetCreatedAt), arg0)
}

// SetExpiresAt mocks base method.
func (m *MockResetTokenBuilder) SetExpiresAt(arg0 int) repositories.ResetTokenBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetExpiresAt", arg0)
	ret0, _ := ret[0].(repositories.ResetTokenBuilder)
	return ret0
}

// SetExpiresAt indicates an expected call of SetExpiresAt.
func (mr *MockResetTokenBuilderMockRecorder) SetExpiresAt(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetExpiresAt", reflect.TypeOf((*MockResetTokenBuilder)(nil).SetExpiresAt), arg0)
}

// SetToken mocks base method.
func (m *MockResetTokenBuilder) SetToken(arg0 string) repositories.ResetTokenBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetToken", arg0)
	ret0, _ := ret[0].(repositories.ResetTokenBuilder)
	return ret0
}

// SetToken indicates an expected call of SetToken.
func (mr *MockResetTokenBuilderMockRecorder) SetToken(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetToken", reflect.TypeOf((*MockResetTokenBuilder)(nil).SetToken), arg0)
}

// SetUserUUID mocks base method.
func (m *MockResetTokenBuilder) SetUserUUID(arg0 string) repositories.ResetTokenBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUserUUID", arg0)
	ret0, _ := ret[0].(repositories.ResetTokenBuilder)
	return ret0
}

// SetUserUUID indicates an expected call of SetUserUUID.
func (mr *MockResetTokenBuilderMockRecorder) SetUserUUID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserUUID", reflect.TypeOf((*MockResetTokenBuilder)(nil).SetUserUUID), arg0)
}