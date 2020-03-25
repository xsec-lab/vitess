// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/xsec-lab/go/vt/discovery (interfaces: HealthCheck)

// Package txthrottler is a generated GoMock package.
package txthrottler

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	discovery "github.com/xsec-lab/go/vt/discovery"
	topodata "github.com/xsec-lab/go/vt/proto/topodata"
	queryservice "github.com/xsec-lab/go/vt/vttablet/queryservice"
)

// MockHealthCheck is a mock of HealthCheck interface
type MockHealthCheck struct {
	ctrl     *gomock.Controller
	recorder *MockHealthCheckMockRecorder
}

// MockHealthCheckMockRecorder is the mock recorder for MockHealthCheck
type MockHealthCheckMockRecorder struct {
	mock *MockHealthCheck
}

// NewMockHealthCheck creates a new mock instance
func NewMockHealthCheck(ctrl *gomock.Controller) *MockHealthCheck {
	mock := &MockHealthCheck{ctrl: ctrl}
	mock.recorder = &MockHealthCheckMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHealthCheck) EXPECT() *MockHealthCheckMockRecorder {
	return m.recorder
}

// AddTablet mocks base method
func (m *MockHealthCheck) AddTablet(arg0 *topodata.Tablet, arg1 string) {
	m.ctrl.Call(m, "AddTablet", arg0, arg1)
}

// AddTablet indicates an expected call of AddTablet
func (mr *MockHealthCheckMockRecorder) AddTablet(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTablet", reflect.TypeOf((*MockHealthCheck)(nil).AddTablet), arg0, arg1)
}

// CacheStatus mocks base method
func (m *MockHealthCheck) CacheStatus() discovery.TabletsCacheStatusList {
	ret := m.ctrl.Call(m, "CacheStatus")
	ret0, _ := ret[0].(discovery.TabletsCacheStatusList)
	return ret0
}

// CacheStatus indicates an expected call of CacheStatus
func (mr *MockHealthCheckMockRecorder) CacheStatus() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CacheStatus", reflect.TypeOf((*MockHealthCheck)(nil).CacheStatus))
}

// Close mocks base method
func (m *MockHealthCheck) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockHealthCheckMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockHealthCheck)(nil).Close))
}

// GetConnection mocks base method
func (m *MockHealthCheck) GetConnection(arg0 string) queryservice.QueryService {
	ret := m.ctrl.Call(m, "GetConnection", arg0)
	ret0, _ := ret[0].(queryservice.QueryService)
	return ret0
}

// GetConnection indicates an expected call of GetConnection
func (mr *MockHealthCheckMockRecorder) GetConnection(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConnection", reflect.TypeOf((*MockHealthCheck)(nil).GetConnection), arg0)
}

// RegisterStats mocks base method
func (m *MockHealthCheck) RegisterStats() {
	m.ctrl.Call(m, "RegisterStats")
}

// RegisterStats indicates an expected call of RegisterStats
func (mr *MockHealthCheckMockRecorder) RegisterStats() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterStats", reflect.TypeOf((*MockHealthCheck)(nil).RegisterStats))
}

// RemoveTablet mocks base method
func (m *MockHealthCheck) RemoveTablet(arg0 *topodata.Tablet) {
	m.ctrl.Call(m, "RemoveTablet", arg0)
}

// RemoveTablet indicates an expected call of RemoveTablet
func (mr *MockHealthCheckMockRecorder) RemoveTablet(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTablet", reflect.TypeOf((*MockHealthCheck)(nil).RemoveTablet), arg0)
}

// ReplaceTablet mocks base method
func (m *MockHealthCheck) ReplaceTablet(arg0, arg1 *topodata.Tablet, arg2 string) {
	m.ctrl.Call(m, "ReplaceTablet", arg0, arg1, arg2)
}

// ReplaceTablet indicates an expected call of ReplaceTablet
func (mr *MockHealthCheckMockRecorder) ReplaceTablet(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplaceTablet", reflect.TypeOf((*MockHealthCheck)(nil).ReplaceTablet), arg0, arg1, arg2)
}

// SetListener mocks base method
func (m *MockHealthCheck) SetListener(arg0 discovery.HealthCheckStatsListener, arg1 bool) {
	m.ctrl.Call(m, "SetListener", arg0, arg1)
}

// SetListener indicates an expected call of SetListener
func (mr *MockHealthCheckMockRecorder) SetListener(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetListener", reflect.TypeOf((*MockHealthCheck)(nil).SetListener), arg0, arg1)
}

// WaitForInitialStatsUpdates mocks base method
func (m *MockHealthCheck) WaitForInitialStatsUpdates() {
	m.ctrl.Call(m, "WaitForInitialStatsUpdates")
}

// WaitForInitialStatsUpdates indicates an expected call of WaitForInitialStatsUpdates
func (mr *MockHealthCheckMockRecorder) WaitForInitialStatsUpdates() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForInitialStatsUpdates", reflect.TypeOf((*MockHealthCheck)(nil).WaitForInitialStatsUpdates))
}
