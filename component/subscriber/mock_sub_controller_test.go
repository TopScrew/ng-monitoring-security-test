// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pingcap/ng-monitoring/component/subscriber (interfaces: SubscribeController,Scraper)

// Package subscriber_test is a generated GoMock package.
package subscriber_test

import (
	"net/http"
	"sync"

	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	subscriber "github.com/pingcap/ng-monitoring/component/subscriber"
	topology "github.com/pingcap/ng-monitoring/component/topology"
	config "github.com/pingcap/ng-monitoring/config"
	pdvariable "github.com/pingcap/ng-monitoring/config/pdvariable"
)

// MockSubscribeController is a mock of SubscribeController interface.
type MockSubscribeController struct {
	ctrl     *gomock.Controller
	recorder *MockSubscribeControllerMockRecorder
}

// MockSubscribeControllerMockRecorder is the mock recorder for MockSubscribeController.
type MockSubscribeControllerMockRecorder struct {
	mock *MockSubscribeController
}

// NewMockSubscribeController creates a new mock instance.
func NewMockSubscribeController(ctrl *gomock.Controller) *MockSubscribeController {
	mock := &MockSubscribeController{ctrl: ctrl}
	mock.recorder = &MockSubscribeControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscribeController) EXPECT() *MockSubscribeControllerMockRecorder {
	return m.recorder
}

func (m *MockSubscribeController) NewHTTPClient() *http.Client {
	return nil
}

// IsEnabled mocks base method.
func (m *MockSubscribeController) IsEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEnabled indicates an expected call of IsEnabled.
func (mr *MockSubscribeControllerMockRecorder) IsEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEnabled", reflect.TypeOf((*MockSubscribeController)(nil).IsEnabled))
}

// Name mocks base method.
func (m *MockSubscribeController) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockSubscribeControllerMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockSubscribeController)(nil).Name))
}

// NewScraper mocks base method.
func (m *MockSubscribeController) NewScraper(arg0 context.Context, arg1 topology.Component, _ *sync.Map) subscriber.Scraper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewScraper", arg0, arg1)
	ret0, _ := ret[0].(subscriber.Scraper)
	return ret0
}

// NewScraper indicates an expected call of NewScraper.
func (mr *MockSubscribeControllerMockRecorder) NewScraper(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewScraper", reflect.TypeOf((*MockSubscribeController)(nil).NewScraper), arg0, arg1)
}

// UpdateConfig mocks base method.
func (m *MockSubscribeController) UpdateConfig(arg0 config.Config) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateConfig", arg0)
}

// UpdateConfig indicates an expected call of UpdateConfig.
func (mr *MockSubscribeControllerMockRecorder) UpdateConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConfig", reflect.TypeOf((*MockSubscribeController)(nil).UpdateConfig), arg0)
}

// UpdatePDVariable mocks base method.
func (m *MockSubscribeController) UpdatePDVariable(arg0 pdvariable.PDVariable) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePDVariable", arg0)
}

// UpdatePDVariable indicates an expected call of UpdatePDVariable.
func (mr *MockSubscribeControllerMockRecorder) UpdatePDVariable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePDVariable", reflect.TypeOf((*MockSubscribeController)(nil).UpdatePDVariable), arg0)
}

// UpdateTopology mocks base method.
func (m *MockSubscribeController) UpdateTopology(arg0 []topology.Component) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateTopology", arg0)
}

// UpdateTopology indicates an expected call of UpdateTopology.
func (mr *MockSubscribeControllerMockRecorder) UpdateTopology(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTopology", reflect.TypeOf((*MockSubscribeController)(nil).UpdateTopology), arg0)
}

// MockScraper is a mock of Scraper interface.
type MockScraper struct {
	ctrl     *gomock.Controller
	recorder *MockScraperMockRecorder
}

// MockScraperMockRecorder is the mock recorder for MockScraper.
type MockScraperMockRecorder struct {
	mock *MockScraper
}

// NewMockScraper creates a new mock instance.
func NewMockScraper(ctrl *gomock.Controller) *MockScraper {
	mock := &MockScraper{ctrl: ctrl}
	mock.recorder = &MockScraperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScraper) EXPECT() *MockScraperMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockScraper) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockScraperMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockScraper)(nil).Close))
}

// IsDown mocks base method.
func (m *MockScraper) IsDown() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDown")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDown indicates an expected call of IsDown.
func (mr *MockScraperMockRecorder) IsDown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDown", reflect.TypeOf((*MockScraper)(nil).IsDown))
}

// Run mocks base method.
func (m *MockScraper) Run() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Run")
}

// Run indicates an expected call of Run.
func (mr *MockScraperMockRecorder) Run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockScraper)(nil).Run))
}
