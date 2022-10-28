// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/dependency-watchdog/internal/prober (interfaces: ShootClientCreator)

// Package prober is a generated GoMock package.
package prober

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	kubernetes "k8s.io/client-go/kubernetes"
)

// MockShootClientCreator is a mock of ShootClientCreator interface.
type MockShootClientCreator struct {
	ctrl     *gomock.Controller
	recorder *MockShootClientCreatorMockRecorder
}

// MockShootClientCreatorMockRecorder is the mock recorder for MockShootClientCreator.
type MockShootClientCreatorMockRecorder struct {
	mock *MockShootClientCreator
}

// NewMockShootClientCreator creates a new mock instance.
func NewMockShootClientCreator(ctrl *gomock.Controller) *MockShootClientCreator {
	mock := &MockShootClientCreator{ctrl: ctrl}
	mock.recorder = &MockShootClientCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShootClientCreator) EXPECT() *MockShootClientCreatorMockRecorder {
	return m.recorder
}

// CreateClient mocks base method.
func (m *MockShootClientCreator) CreateClient(arg0 context.Context, arg1, arg2 string, arg3 time.Duration) (kubernetes.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateClient", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(kubernetes.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateClient indicates an expected call of CreateClient.
func (mr *MockShootClientCreatorMockRecorder) CreateClient(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateClient", reflect.TypeOf((*MockShootClientCreator)(nil).CreateClient), arg0, arg1, arg2, arg3)
}
