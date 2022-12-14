// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	match "github.com/fbngrm/around-home/pkg/match"
	materials "github.com/fbngrm/around-home/pkg/materials"
	gomock "github.com/golang/mock/gomock"
)

// MockMatcher is a mock of Matcher interface.
type MockMatcher struct {
	ctrl     *gomock.Controller
	recorder *MockMatcherMockRecorder
}

// MockMatcherMockRecorder is the mock recorder for MockMatcher.
type MockMatcherMockRecorder struct {
	mock *MockMatcher
}

// NewMockMatcher creates a new mock instance.
func NewMockMatcher(ctrl *gomock.Controller) *MockMatcher {
	mock := &MockMatcher{ctrl: ctrl}
	mock.recorder = &MockMatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMatcher) EXPECT() *MockMatcherMockRecorder {
	return m.recorder
}

// GetMatches mocks base method.
func (m_2 *MockMatcher) GetMatches(ctx context.Context, m materials.Material, location string) (match.Matches, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "GetMatches", ctx, m, location)
	ret0, _ := ret[0].(match.Matches)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMatches indicates an expected call of GetMatches.
func (mr *MockMatcherMockRecorder) GetMatches(ctx, m, location interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMatches", reflect.TypeOf((*MockMatcher)(nil).GetMatches), ctx, m, location)
}
