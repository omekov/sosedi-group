// Code generated by MockGen. DO NOT EDIT.
// Source: internal/email/service.go

// Package mock_email is a generated GoMock package.
package mock_email

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEmailService is a mock of EmailService interface.
type MockEmailService struct {
	ctrl     *gomock.Controller
	recorder *MockEmailServiceMockRecorder
}

// MockEmailServiceMockRecorder is the mock recorder for MockEmailService.
type MockEmailServiceMockRecorder struct {
	mock *MockEmailService
}

// NewMockEmailService creates a new mock instance.
func NewMockEmailService(ctrl *gomock.Controller) *MockEmailService {
	mock := &MockEmailService{ctrl: ctrl}
	mock.recorder = &MockEmailServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmailService) EXPECT() *MockEmailServiceMockRecorder {
	return m.recorder
}

// FindEmailFromText mocks base method.
func (m *MockEmailService) FindEmailFromText(text string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEmailFromText", text)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEmailFromText indicates an expected call of FindEmailFromText.
func (mr *MockEmailServiceMockRecorder) FindEmailFromText(text interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEmailFromText", reflect.TypeOf((*MockEmailService)(nil).FindEmailFromText), text)
}

// FindIINFromText mocks base method.
func (m *MockEmailService) FindIINFromText(text string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindIINFromText", text)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindIINFromText indicates an expected call of FindIINFromText.
func (mr *MockEmailServiceMockRecorder) FindIINFromText(text interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindIINFromText", reflect.TypeOf((*MockEmailService)(nil).FindIINFromText), text)
}
