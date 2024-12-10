// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fromsi/jwt-oauth-sso/internal/http/responses (interfaces: SuccessRegisterResponse)
//
// Generated by this command:
//
//	mockgen -destination=../../../mocks/http/responses/mock_success_register_response.go -package=responses_mocks github.com/fromsi/jwt-oauth-sso/internal/http/responses SuccessRegisterResponse
//

// Package responses_mocks is a generated GoMock package.
package responses_mocks

import (
	reflect "reflect"

	responses "github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	repositories "github.com/fromsi/jwt-oauth-sso/internal/repositories"
	gomock "go.uber.org/mock/gomock"
)

// MockSuccessRegisterResponse is a mock of SuccessRegisterResponse interface.
type MockSuccessRegisterResponse struct {
	ctrl     *gomock.Controller
	recorder *MockSuccessRegisterResponseMockRecorder
	isgomock struct{}
}

// MockSuccessRegisterResponseMockRecorder is the mock recorder for MockSuccessRegisterResponse.
type MockSuccessRegisterResponseMockRecorder struct {
	mock *MockSuccessRegisterResponse
}

// NewMockSuccessRegisterResponse creates a new mock instance.
func NewMockSuccessRegisterResponse(ctrl *gomock.Controller) *MockSuccessRegisterResponse {
	mock := &MockSuccessRegisterResponse{ctrl: ctrl}
	mock.recorder = &MockSuccessRegisterResponseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSuccessRegisterResponse) EXPECT() *MockSuccessRegisterResponseMockRecorder {
	return m.recorder
}

// Make mocks base method.
func (m *MockSuccessRegisterResponse) Make(arg0 repositories.Device) (responses.SuccessRegisterResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Make", arg0)
	ret0, _ := ret[0].(responses.SuccessRegisterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Make indicates an expected call of Make.
func (mr *MockSuccessRegisterResponseMockRecorder) Make(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Make", reflect.TypeOf((*MockSuccessRegisterResponse)(nil).Make), arg0)
}