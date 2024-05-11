// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hr3lxphr6j/bililive-go/src/live (interfaces: Live)

// Package mock is a generated GoMock package.
package mock

import (
	url "net/url"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	live "github.com/kira1928/xlive/src/live"
)

// MockLive is a mock of Live interface.
type MockLive struct {
	ctrl     *gomock.Controller
	recorder *MockLiveMockRecorder
}

// MockLiveMockRecorder is the mock recorder for MockLive.
type MockLiveMockRecorder struct {
	mock *MockLive
}

// NewMockLive creates a new mock instance.
func NewMockLive(ctrl *gomock.Controller) *MockLive {
	mock := &MockLive{ctrl: ctrl}
	mock.recorder = &MockLiveMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLive) EXPECT() *MockLiveMockRecorder {
	return m.recorder
}

// GetHeadersForDownloader mocks base method.
func (m *MockLive) GetHeadersForDownloader() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeadersForDownloader")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetHeadersForDownloader indicates an expected call of GetHeadersForDownloader.
func (mr *MockLiveMockRecorder) GetHeadersForDownloader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeadersForDownloader", reflect.TypeOf((*MockLive)(nil).GetHeadersForDownloader))
}

// GetInfo mocks base method.
func (m *MockLive) GetInfo() (*live.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfo")
	ret0, _ := ret[0].(*live.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfo indicates an expected call of GetInfo.
func (mr *MockLiveMockRecorder) GetInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfo", reflect.TypeOf((*MockLive)(nil).GetInfo))
}

// GetLastStartTime mocks base method.
func (m *MockLive) GetLastStartTime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastStartTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetLastStartTime indicates an expected call of GetLastStartTime.
func (mr *MockLiveMockRecorder) GetLastStartTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastStartTime", reflect.TypeOf((*MockLive)(nil).GetLastStartTime))
}

// GetLiveId mocks base method.
func (m *MockLive) GetLiveId() live.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLiveId")
	ret0, _ := ret[0].(live.ID)
	return ret0
}

// GetLiveId indicates an expected call of GetLiveId.
func (mr *MockLiveMockRecorder) GetLiveId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLiveId", reflect.TypeOf((*MockLive)(nil).GetLiveId))
}

// GetPlatformCNName mocks base method.
func (m *MockLive) GetPlatformCNName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlatformCNName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPlatformCNName indicates an expected call of GetPlatformCNName.
func (mr *MockLiveMockRecorder) GetPlatformCNName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlatformCNName", reflect.TypeOf((*MockLive)(nil).GetPlatformCNName))
}

// GetRawUrl mocks base method.
func (m *MockLive) GetRawUrl() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRawUrl")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetRawUrl indicates an expected call of GetRawUrl.
func (mr *MockLiveMockRecorder) GetRawUrl() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawUrl", reflect.TypeOf((*MockLive)(nil).GetRawUrl))
}

// GetStreamUrls mocks base method.
func (m *MockLive) GetStreamUrls() ([]*url.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStreamUrls")
	ret0, _ := ret[0].([]*url.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStreamUrls indicates an expected call of GetStreamUrls.
func (mr *MockLiveMockRecorder) GetStreamUrls() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStreamUrls", reflect.TypeOf((*MockLive)(nil).GetStreamUrls))
}

// SetLastStartTime mocks base method.
func (m *MockLive) SetLastStartTime(arg0 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLastStartTime", arg0)
}

// SetLastStartTime indicates an expected call of SetLastStartTime.
func (mr *MockLiveMockRecorder) SetLastStartTime(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLastStartTime", reflect.TypeOf((*MockLive)(nil).SetLastStartTime), arg0)
}

// SetLiveIdByString mocks base method.
func (m *MockLive) SetLiveIdByString(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLiveIdByString", arg0)
}

// SetLiveIdByString indicates an expected call of SetLiveIdByString.
func (mr *MockLiveMockRecorder) SetLiveIdByString(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLiveIdByString", reflect.TypeOf((*MockLive)(nil).SetLiveIdByString), arg0)
}
