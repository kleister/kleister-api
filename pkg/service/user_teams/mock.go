// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/service/user_teams/service.go

// Package userteams is a generated GoMock package.
package userteams

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/kleister/kleister-api/pkg/model"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Attach mocks base method.
func (m *MockService) Attach(arg0 context.Context, arg1 model.UserTeamParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Attach", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Attach indicates an expected call of Attach.
func (mr *MockServiceMockRecorder) Attach(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attach", reflect.TypeOf((*MockService)(nil).Attach), arg0, arg1)
}

// Drop mocks base method.
func (m *MockService) Drop(arg0 context.Context, arg1 model.UserTeamParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Drop", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Drop indicates an expected call of Drop.
func (mr *MockServiceMockRecorder) Drop(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Drop", reflect.TypeOf((*MockService)(nil).Drop), arg0, arg1)
}

// List mocks base method.
func (m *MockService) List(arg0 context.Context, arg1 model.UserTeamParams) ([]*model.UserTeam, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*model.UserTeam)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockService)(nil).List), arg0, arg1)
}

// Permit mocks base method.
func (m *MockService) Permit(arg0 context.Context, arg1 model.UserTeamParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Permit", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Permit indicates an expected call of Permit.
func (mr *MockServiceMockRecorder) Permit(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Permit", reflect.TypeOf((*MockService)(nil).Permit), arg0, arg1)
}

// WithPrincipal mocks base method.
func (m *MockService) WithPrincipal(arg0 *model.User) Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithPrincipal", arg0)
	ret0, _ := ret[0].(Service)
	return ret0
}

// WithPrincipal indicates an expected call of WithPrincipal.
func (mr *MockServiceMockRecorder) WithPrincipal(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithPrincipal", reflect.TypeOf((*MockService)(nil).WithPrincipal), arg0)
}
