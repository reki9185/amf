// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/service/init.go
//
// Generated by this command:
//
//	mockgen -source=pkg/service/init.go -package=service
//

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	context "github.com/free5gc/amf/internal/context"
	consumer "github.com/free5gc/amf/internal/sbi/consumer"
	factory "github.com/free5gc/amf/pkg/factory"
	gomock "go.uber.org/mock/gomock"
)

// MockAmfAppInterface is a mock of AmfAppInterface interface.
type MockAmfAppInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAmfAppInterfaceMockRecorder
}

// MockAmfAppInterfaceMockRecorder is the mock recorder for MockAmfAppInterface.
type MockAmfAppInterfaceMockRecorder struct {
	mock *MockAmfAppInterface
}

// NewMockAmfAppInterface creates a new mock instance.
func NewMockAmfAppInterface(ctrl *gomock.Controller) *MockAmfAppInterface {
	mock := &MockAmfAppInterface{ctrl: ctrl}
	mock.recorder = &MockAmfAppInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAmfAppInterface) EXPECT() *MockAmfAppInterfaceMockRecorder {
	return m.recorder
}

// Config mocks base method.
func (m *MockAmfAppInterface) Config() *factory.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*factory.Config)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockAmfAppInterfaceMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockAmfAppInterface)(nil).Config))
}

// Consumer mocks base method.
func (m *MockAmfAppInterface) Consumer() *consumer.Consumer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consumer")
	ret0, _ := ret[0].(*consumer.Consumer)
	return ret0
}

// Consumer indicates an expected call of Consumer.
func (mr *MockAmfAppInterfaceMockRecorder) Consumer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consumer", reflect.TypeOf((*MockAmfAppInterface)(nil).Consumer))
}

// Context mocks base method.
func (m *MockAmfAppInterface) Context() *context.AMFContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(*context.AMFContext)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockAmfAppInterfaceMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockAmfAppInterface)(nil).Context))
}

// SetLogEnable mocks base method.
func (m *MockAmfAppInterface) SetLogEnable(enable bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLogEnable", enable)
}

// SetLogEnable indicates an expected call of SetLogEnable.
func (mr *MockAmfAppInterfaceMockRecorder) SetLogEnable(enable any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLogEnable", reflect.TypeOf((*MockAmfAppInterface)(nil).SetLogEnable), enable)
}

// SetLogLevel mocks base method.
func (m *MockAmfAppInterface) SetLogLevel(level string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLogLevel", level)
}

// SetLogLevel indicates an expected call of SetLogLevel.
func (mr *MockAmfAppInterfaceMockRecorder) SetLogLevel(level any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLogLevel", reflect.TypeOf((*MockAmfAppInterface)(nil).SetLogLevel), level)
}

// SetReportCaller mocks base method.
func (m *MockAmfAppInterface) SetReportCaller(reportCaller bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetReportCaller", reportCaller)
}

// SetReportCaller indicates an expected call of SetReportCaller.
func (mr *MockAmfAppInterfaceMockRecorder) SetReportCaller(reportCaller any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReportCaller", reflect.TypeOf((*MockAmfAppInterface)(nil).SetReportCaller), reportCaller)
}

// Start mocks base method.
func (m *MockAmfAppInterface) Start(tlsKeyLogPath string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start", tlsKeyLogPath)
}

// Start indicates an expected call of Start.
func (mr *MockAmfAppInterfaceMockRecorder) Start(tlsKeyLogPath any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockAmfAppInterface)(nil).Start), tlsKeyLogPath)
}

// Terminate mocks base method.
func (m *MockAmfAppInterface) Terminate() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Terminate")
}

// Terminate indicates an expected call of Terminate.
func (mr *MockAmfAppInterfaceMockRecorder) Terminate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Terminate", reflect.TypeOf((*MockAmfAppInterface)(nil).Terminate))
}
