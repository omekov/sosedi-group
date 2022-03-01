// Code generated by MockGen. DO NOT EDIT.
// Source: internal/email/repository.go

// Package mock_email is a generated GoMock package.
package mock_email

import (
	gomock "github.com/golang/mock/gomock"
)

// MockEmail is a mock of Email interface.
type MockEmail struct {
	ctrl     *gomock.Controller
	recorder *MockEmailMockRecorder
}

// MockEmailMockRecorder is the mock recorder for MockEmail.
type MockEmailMockRecorder struct {
	mock *MockEmail
}

// NewMockEmail creates a new mock instance.
func NewMockEmail(ctrl *gomock.Controller) *MockEmail {
	mock := &MockEmail{ctrl: ctrl}
	mock.recorder = &MockEmailMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmail) EXPECT() *MockEmailMockRecorder {
	return m.recorder
}