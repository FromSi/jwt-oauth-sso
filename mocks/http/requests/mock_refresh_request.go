// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fromsi/jwt-oauth-sso/internal/http/requests (interfaces: RefreshRequest)
//
// Generated by this command:
//
//	mockgen -destination=../../../mocks/http/requests/mock_refresh_request.go -package=requests_mocks github.com/fromsi/jwt-oauth-sso/internal/http/requests RefreshRequest
//

// Package requests_mocks is a generated GoMock package.
package requests_mocks

import (
	reflect "reflect"

	requests "github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	gin "github.com/gin-gonic/gin"
	gomock "go.uber.org/mock/gomock"
)

// MockRefreshRequest is a mock of RefreshRequest interface.
type MockRefreshRequest struct {
	ctrl     *gomock.Controller
	recorder *MockRefreshRequestMockRecorder
	isgomock struct{}
}

// MockRefreshRequestMockRecorder is the mock recorder for MockRefreshRequest.
type MockRefreshRequestMockRecorder struct {
	mock *MockRefreshRequest
}

// NewMockRefreshRequest creates a new mock instance.
func NewMockRefreshRequest(ctrl *gomock.Controller) *MockRefreshRequest {
	mock := &MockRefreshRequest{ctrl: ctrl}
	mock.recorder = &MockRefreshRequestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRefreshRequest) EXPECT() *MockRefreshRequestMockRecorder {
	return m.recorder
}

// GetBody mocks base method.
func (m *MockRefreshRequest) GetBody() requests.RefreshRequestBody {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBody")
	ret0, _ := ret[0].(requests.RefreshRequestBody)
	return ret0
}

// GetBody indicates an expected call of GetBody.
func (mr *MockRefreshRequestMockRecorder) GetBody() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBody", reflect.TypeOf((*MockRefreshRequest)(nil).GetBody))
}

// Make mocks base method.
func (m *MockRefreshRequest) Make(arg0 *gin.Context) (requests.RefreshRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Make", arg0)
	ret0, _ := ret[0].(requests.RefreshRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Make indicates an expected call of Make.
func (mr *MockRefreshRequestMockRecorder) Make(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Make", reflect.TypeOf((*MockRefreshRequest)(nil).Make), arg0)
}
