// Code generated by MockGen. DO NOT EDIT.
// Source: controller.go

// Package cluster is a generated GoMock package.
package cluster

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/vanus-labs/vanus/proto/pkg/controller"
)

// MockCluster is a mock of Cluster interface.
type MockCluster struct {
	ctrl     *gomock.Controller
	recorder *MockClusterMockRecorder
}

// MockClusterMockRecorder is the mock recorder for MockCluster.
type MockClusterMockRecorder struct {
	mock *MockCluster
}

// NewMockCluster creates a new mock instance.
func NewMockCluster(ctrl *gomock.Controller) *MockCluster {
	mock := &MockCluster{ctrl: ctrl}
	mock.recorder = &MockClusterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCluster) EXPECT() *MockClusterMockRecorder {
	return m.recorder
}

// EventbusService mocks base method.
func (m *MockCluster) EventbusService() EventbusService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventbusService")
	ret0, _ := ret[0].(EventbusService)
	return ret0
}

// EventbusService indicates an expected call of EventbusService.
func (mr *MockClusterMockRecorder) EventbusService() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventbusService", reflect.TypeOf((*MockCluster)(nil).EventbusService))
}

// EventlogService mocks base method.
func (m *MockCluster) EventlogService() EventlogService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventlogService")
	ret0, _ := ret[0].(EventlogService)
	return ret0
}

// EventlogService indicates an expected call of EventlogService.
func (mr *MockClusterMockRecorder) EventlogService() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventlogService", reflect.TypeOf((*MockCluster)(nil).EventlogService))
}

// IDService mocks base method.
func (m *MockCluster) IDService() IDService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IDService")
	ret0, _ := ret[0].(IDService)
	return ret0
}

// IDService indicates an expected call of IDService.
func (mr *MockClusterMockRecorder) IDService() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IDService", reflect.TypeOf((*MockCluster)(nil).IDService))
}

// IsReady mocks base method.
func (m *MockCluster) IsReady(createEventbus bool) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsReady", createEventbus)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsReady indicates an expected call of IsReady.
func (mr *MockClusterMockRecorder) IsReady(createEventbus interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsReady", reflect.TypeOf((*MockCluster)(nil).IsReady), createEventbus)
}

// SegmentService mocks base method.
func (m *MockCluster) SegmentService() SegmentService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SegmentService")
	ret0, _ := ret[0].(SegmentService)
	return ret0
}

// SegmentService indicates an expected call of SegmentService.
func (mr *MockClusterMockRecorder) SegmentService() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SegmentService", reflect.TypeOf((*MockCluster)(nil).SegmentService))
}

// Status mocks base method.
func (m *MockCluster) Status() Topology {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(Topology)
	return ret0
}

// Status indicates an expected call of Status.
func (mr *MockClusterMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockCluster)(nil).Status))
}

// TriggerService mocks base method.
func (m *MockCluster) TriggerService() TriggerService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TriggerService")
	ret0, _ := ret[0].(TriggerService)
	return ret0
}

// TriggerService indicates an expected call of TriggerService.
func (mr *MockClusterMockRecorder) TriggerService() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TriggerService", reflect.TypeOf((*MockCluster)(nil).TriggerService))
}

// WaitForControllerReady mocks base method.
func (m *MockCluster) WaitForControllerReady(createEventbus bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForControllerReady", createEventbus)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForControllerReady indicates an expected call of WaitForControllerReady.
func (mr *MockClusterMockRecorder) WaitForControllerReady(createEventbus interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForControllerReady", reflect.TypeOf((*MockCluster)(nil).WaitForControllerReady), createEventbus)
}

// MockEventbusService is a mock of EventbusService interface.
type MockEventbusService struct {
	ctrl     *gomock.Controller
	recorder *MockEventbusServiceMockRecorder
}

// MockEventbusServiceMockRecorder is the mock recorder for MockEventbusService.
type MockEventbusServiceMockRecorder struct {
	mock *MockEventbusService
}

// NewMockEventbusService creates a new mock instance.
func NewMockEventbusService(ctrl *gomock.Controller) *MockEventbusService {
	mock := &MockEventbusService{ctrl: ctrl}
	mock.recorder = &MockEventbusServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventbusService) EXPECT() *MockEventbusServiceMockRecorder {
	return m.recorder
}

// CreateSystemEventbusIfNotExist mocks base method.
func (m *MockEventbusService) CreateSystemEventbusIfNotExist(ctx context.Context, name, desc string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSystemEventbusIfNotExist", ctx, name, desc)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSystemEventbusIfNotExist indicates an expected call of CreateSystemEventbusIfNotExist.
func (mr *MockEventbusServiceMockRecorder) CreateSystemEventbusIfNotExist(ctx, name, desc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSystemEventbusIfNotExist", reflect.TypeOf((*MockEventbusService)(nil).CreateSystemEventbusIfNotExist), ctx, name, desc)
}

// Delete mocks base method.
func (m *MockEventbusService) Delete(ctx context.Context, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockEventbusServiceMockRecorder) Delete(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockEventbusService)(nil).Delete), ctx, name)
}

// IsExist mocks base method.
func (m *MockEventbusService) IsExist(ctx context.Context, name string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsExist", ctx, name)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsExist indicates an expected call of IsExist.
func (mr *MockEventbusServiceMockRecorder) IsExist(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExist", reflect.TypeOf((*MockEventbusService)(nil).IsExist), ctx, name)
}

// RawClient mocks base method.
func (m *MockEventbusService) RawClient() controller.EventbusControllerClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RawClient")
	ret0, _ := ret[0].(controller.EventbusControllerClient)
	return ret0
}

// RawClient indicates an expected call of RawClient.
func (mr *MockEventbusServiceMockRecorder) RawClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RawClient", reflect.TypeOf((*MockEventbusService)(nil).RawClient))
}

// MockEventlogService is a mock of EventlogService interface.
type MockEventlogService struct {
	ctrl     *gomock.Controller
	recorder *MockEventlogServiceMockRecorder
}

// MockEventlogServiceMockRecorder is the mock recorder for MockEventlogService.
type MockEventlogServiceMockRecorder struct {
	mock *MockEventlogService
}

// NewMockEventlogService creates a new mock instance.
func NewMockEventlogService(ctrl *gomock.Controller) *MockEventlogService {
	mock := &MockEventlogService{ctrl: ctrl}
	mock.recorder = &MockEventlogServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventlogService) EXPECT() *MockEventlogServiceMockRecorder {
	return m.recorder
}

// RawClient mocks base method.
func (m *MockEventlogService) RawClient() controller.EventlogControllerClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RawClient")
	ret0, _ := ret[0].(controller.EventlogControllerClient)
	return ret0
}

// RawClient indicates an expected call of RawClient.
func (mr *MockEventlogServiceMockRecorder) RawClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RawClient", reflect.TypeOf((*MockEventlogService)(nil).RawClient))
}

// MockTriggerService is a mock of TriggerService interface.
type MockTriggerService struct {
	ctrl     *gomock.Controller
	recorder *MockTriggerServiceMockRecorder
}

// MockTriggerServiceMockRecorder is the mock recorder for MockTriggerService.
type MockTriggerServiceMockRecorder struct {
	mock *MockTriggerService
}

// NewMockTriggerService creates a new mock instance.
func NewMockTriggerService(ctrl *gomock.Controller) *MockTriggerService {
	mock := &MockTriggerService{ctrl: ctrl}
	mock.recorder = &MockTriggerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTriggerService) EXPECT() *MockTriggerServiceMockRecorder {
	return m.recorder
}

// RawClient mocks base method.
func (m *MockTriggerService) RawClient() controller.TriggerControllerClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RawClient")
	ret0, _ := ret[0].(controller.TriggerControllerClient)
	return ret0
}

// RawClient indicates an expected call of RawClient.
func (mr *MockTriggerServiceMockRecorder) RawClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RawClient", reflect.TypeOf((*MockTriggerService)(nil).RawClient))
}

// RegisterHeartbeat mocks base method.
func (m *MockTriggerService) RegisterHeartbeat(ctx context.Context, interval time.Duration, reqFunc func() interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterHeartbeat", ctx, interval, reqFunc)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterHeartbeat indicates an expected call of RegisterHeartbeat.
func (mr *MockTriggerServiceMockRecorder) RegisterHeartbeat(ctx, interval, reqFunc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterHeartbeat", reflect.TypeOf((*MockTriggerService)(nil).RegisterHeartbeat), ctx, interval, reqFunc)
}

// MockIDService is a mock of IDService interface.
type MockIDService struct {
	ctrl     *gomock.Controller
	recorder *MockIDServiceMockRecorder
}

// MockIDServiceMockRecorder is the mock recorder for MockIDService.
type MockIDServiceMockRecorder struct {
	mock *MockIDService
}

// NewMockIDService creates a new mock instance.
func NewMockIDService(ctrl *gomock.Controller) *MockIDService {
	mock := &MockIDService{ctrl: ctrl}
	mock.recorder = &MockIDServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDService) EXPECT() *MockIDServiceMockRecorder {
	return m.recorder
}

// RawClient mocks base method.
func (m *MockIDService) RawClient() controller.SnowflakeControllerClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RawClient")
	ret0, _ := ret[0].(controller.SnowflakeControllerClient)
	return ret0
}

// RawClient indicates an expected call of RawClient.
func (mr *MockIDServiceMockRecorder) RawClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RawClient", reflect.TypeOf((*MockIDService)(nil).RawClient))
}

// MockSegmentService is a mock of SegmentService interface.
type MockSegmentService struct {
	ctrl     *gomock.Controller
	recorder *MockSegmentServiceMockRecorder
}

// MockSegmentServiceMockRecorder is the mock recorder for MockSegmentService.
type MockSegmentServiceMockRecorder struct {
	mock *MockSegmentService
}

// NewMockSegmentService creates a new mock instance.
func NewMockSegmentService(ctrl *gomock.Controller) *MockSegmentService {
	mock := &MockSegmentService{ctrl: ctrl}
	mock.recorder = &MockSegmentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSegmentService) EXPECT() *MockSegmentServiceMockRecorder {
	return m.recorder
}

// RawClient mocks base method.
func (m *MockSegmentService) RawClient() controller.SegmentControllerClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RawClient")
	ret0, _ := ret[0].(controller.SegmentControllerClient)
	return ret0
}

// RawClient indicates an expected call of RawClient.
func (mr *MockSegmentServiceMockRecorder) RawClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RawClient", reflect.TypeOf((*MockSegmentService)(nil).RawClient))
}

// RegisterHeartbeat mocks base method.
func (m *MockSegmentService) RegisterHeartbeat(ctx context.Context, interval time.Duration, reqFunc func() interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterHeartbeat", ctx, interval, reqFunc)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterHeartbeat indicates an expected call of RegisterHeartbeat.
func (mr *MockSegmentServiceMockRecorder) RegisterHeartbeat(ctx, interval, reqFunc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterHeartbeat", reflect.TypeOf((*MockSegmentService)(nil).RegisterHeartbeat), ctx, interval, reqFunc)
}
