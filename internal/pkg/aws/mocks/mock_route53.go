// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/aws/route53/route53.go

// Package mocks is a generated GoMock package.
package mocks

import (
	route53 "github.com/aws/aws-sdk-go/service/route53"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockLister is a mock of Lister interface
type MockLister struct {
	ctrl     *gomock.Controller
	recorder *MockListerMockRecorder
}

// MockListerMockRecorder is the mock recorder for MockLister
type MockListerMockRecorder struct {
	mock *MockLister
}

// NewMockLister creates a new mock instance
func NewMockLister(ctrl *gomock.Controller) *MockLister {
	mock := &MockLister{ctrl: ctrl}
	mock.recorder = &MockListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLister) EXPECT() *MockListerMockRecorder {
	return m.recorder
}

// ListHostedZonesByName mocks base method
func (m *MockLister) ListHostedZonesByName(in *route53.ListHostedZonesByNameInput) (*route53.ListHostedZonesByNameOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListHostedZonesByName", in)
	ret0, _ := ret[0].(*route53.ListHostedZonesByNameOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListHostedZonesByName indicates an expected call of ListHostedZonesByName
func (mr *MockListerMockRecorder) ListHostedZonesByName(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHostedZonesByName", reflect.TypeOf((*MockLister)(nil).ListHostedZonesByName), in)
}
