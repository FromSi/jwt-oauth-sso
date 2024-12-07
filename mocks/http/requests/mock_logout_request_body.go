// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fromsi/jwt-oauth-sso/internal/http/requests (interfaces: LogoutRequestBody)
//
// Generated by this command:
//
//	mockgen -destination=../../../mocks/http/requests/mock_logout_request_body.go -package=requests_mocks github.com/fromsi/jwt-oauth-sso/internal/http/requests LogoutRequestBody
//

// Package requests_mocks is a generated GoMock package.
package requests_mocks

import (
	reflect "reflect"

	requests "github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	gin "github.com/gin-gonic/gin"
	gomock "go.uber.org/mock/gomock"
)

// MockLogoutRequestBody is a mock of LogoutRequestBody interface.
type MockLogoutRequestBody struct {
	ctrl     *gomock.Controller
	recorder *MockLogoutRequestBodyMockRecorder
	isgomock struct{}
}

// MockLogoutRequestBodyMockRecorder is the mock recorder for MockLogoutRequestBody.
type MockLogoutRequestBodyMockRecorder struct {
	mock *MockLogoutRequestBody
}

// NewMockLogoutRequestBody creates a new mock instance.
func NewMockLogoutRequestBody(ctrl *gomock.Controller) *MockLogoutRequestBody {
	mock := &MockLogoutRequestBody{ctrl: ctrl}
	mock.recorder = &MockLogoutRequestBodyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogoutRequestBody) EXPECT() *MockLogoutRequestBodyMockRecorder {
	return m.recorder
}

// Make mocks base method.
func (m *MockLogoutRequestBody) Make(arg0 *gin.Context) requests.LogoutRequestBody {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Make", arg0)
	ret0, _ := ret[0].(requests.LogoutRequestBody)
	return ret0
}

// Make indicates an expected call of Make.
func (mr *MockLogoutRequestBodyMockRecorder) Make(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Make", reflect.TypeOf((*MockLogoutRequestBody)(nil).Make), arg0)
}
