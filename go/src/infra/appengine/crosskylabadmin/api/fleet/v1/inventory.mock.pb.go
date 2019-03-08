// Code generated by MockGen. DO NOT EDIT.
// Source: inventory.pb.go

// Package fleet is a generated GoMock package.
package fleet

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockInventoryClient is a mock of InventoryClient interface
type MockInventoryClient struct {
	ctrl     *gomock.Controller
	recorder *MockInventoryClientMockRecorder
}

// MockInventoryClientMockRecorder is the mock recorder for MockInventoryClient
type MockInventoryClientMockRecorder struct {
	mock *MockInventoryClient
}

// NewMockInventoryClient creates a new mock instance
func NewMockInventoryClient(ctrl *gomock.Controller) *MockInventoryClient {
	mock := &MockInventoryClient{ctrl: ctrl}
	mock.recorder = &MockInventoryClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInventoryClient) EXPECT() *MockInventoryClientMockRecorder {
	return m.recorder
}

// DeployDut mocks base method
func (m *MockInventoryClient) DeployDut(ctx context.Context, in *DeployDutRequest, opts ...grpc.CallOption) (*DeployDutResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeployDut", varargs...)
	ret0, _ := ret[0].(*DeployDutResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeployDut indicates an expected call of DeployDut
func (mr *MockInventoryClientMockRecorder) DeployDut(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployDut", reflect.TypeOf((*MockInventoryClient)(nil).DeployDut), varargs...)
}

// RedeployDut mocks base method
func (m *MockInventoryClient) RedeployDut(ctx context.Context, in *RedeployDutRequest, opts ...grpc.CallOption) (*RedeployDutResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RedeployDut", varargs...)
	ret0, _ := ret[0].(*RedeployDutResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RedeployDut indicates an expected call of RedeployDut
func (mr *MockInventoryClientMockRecorder) RedeployDut(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RedeployDut", reflect.TypeOf((*MockInventoryClient)(nil).RedeployDut), varargs...)
}

// GetDeploymentStatus mocks base method
func (m *MockInventoryClient) GetDeploymentStatus(ctx context.Context, in *GetDeploymentStatusRequest, opts ...grpc.CallOption) (*GetDeploymentStatusResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDeploymentStatus", varargs...)
	ret0, _ := ret[0].(*GetDeploymentStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeploymentStatus indicates an expected call of GetDeploymentStatus
func (mr *MockInventoryClientMockRecorder) GetDeploymentStatus(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeploymentStatus", reflect.TypeOf((*MockInventoryClient)(nil).GetDeploymentStatus), varargs...)
}

// DeleteDut mocks base method
func (m *MockInventoryClient) DeleteDut(ctx context.Context, in *DeleteDutRequest, opts ...grpc.CallOption) (*DeleteDutResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteDut", varargs...)
	ret0, _ := ret[0].(*DeleteDutResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteDut indicates an expected call of DeleteDut
func (mr *MockInventoryClientMockRecorder) DeleteDut(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDut", reflect.TypeOf((*MockInventoryClient)(nil).DeleteDut), varargs...)
}

// EnsurePoolHealthy mocks base method
func (m *MockInventoryClient) EnsurePoolHealthy(ctx context.Context, in *EnsurePoolHealthyRequest, opts ...grpc.CallOption) (*EnsurePoolHealthyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnsurePoolHealthy", varargs...)
	ret0, _ := ret[0].(*EnsurePoolHealthyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsurePoolHealthy indicates an expected call of EnsurePoolHealthy
func (mr *MockInventoryClientMockRecorder) EnsurePoolHealthy(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsurePoolHealthy", reflect.TypeOf((*MockInventoryClient)(nil).EnsurePoolHealthy), varargs...)
}

// EnsurePoolHealthyForAllModels mocks base method
func (m *MockInventoryClient) EnsurePoolHealthyForAllModels(ctx context.Context, in *EnsurePoolHealthyForAllModelsRequest, opts ...grpc.CallOption) (*EnsurePoolHealthyForAllModelsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnsurePoolHealthyForAllModels", varargs...)
	ret0, _ := ret[0].(*EnsurePoolHealthyForAllModelsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsurePoolHealthyForAllModels indicates an expected call of EnsurePoolHealthyForAllModels
func (mr *MockInventoryClientMockRecorder) EnsurePoolHealthyForAllModels(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsurePoolHealthyForAllModels", reflect.TypeOf((*MockInventoryClient)(nil).EnsurePoolHealthyForAllModels), varargs...)
}

// ResizePool mocks base method
func (m *MockInventoryClient) ResizePool(ctx context.Context, in *ResizePoolRequest, opts ...grpc.CallOption) (*ResizePoolResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ResizePool", varargs...)
	ret0, _ := ret[0].(*ResizePoolResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResizePool indicates an expected call of ResizePool
func (mr *MockInventoryClientMockRecorder) ResizePool(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResizePool", reflect.TypeOf((*MockInventoryClient)(nil).ResizePool), varargs...)
}

// RemoveDutsFromDrones mocks base method
func (m *MockInventoryClient) RemoveDutsFromDrones(ctx context.Context, in *RemoveDutsFromDronesRequest, opts ...grpc.CallOption) (*RemoveDutsFromDronesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveDutsFromDrones", varargs...)
	ret0, _ := ret[0].(*RemoveDutsFromDronesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveDutsFromDrones indicates an expected call of RemoveDutsFromDrones
func (mr *MockInventoryClientMockRecorder) RemoveDutsFromDrones(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDutsFromDrones", reflect.TypeOf((*MockInventoryClient)(nil).RemoveDutsFromDrones), varargs...)
}

// AssignDutsToDrones mocks base method
func (m *MockInventoryClient) AssignDutsToDrones(ctx context.Context, in *AssignDutsToDronesRequest, opts ...grpc.CallOption) (*AssignDutsToDronesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AssignDutsToDrones", varargs...)
	ret0, _ := ret[0].(*AssignDutsToDronesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignDutsToDrones indicates an expected call of AssignDutsToDrones
func (mr *MockInventoryClientMockRecorder) AssignDutsToDrones(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignDutsToDrones", reflect.TypeOf((*MockInventoryClient)(nil).AssignDutsToDrones), varargs...)
}

// ListServers mocks base method
func (m *MockInventoryClient) ListServers(ctx context.Context, in *ListServersRequest, opts ...grpc.CallOption) (*ListServersResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListServers", varargs...)
	ret0, _ := ret[0].(*ListServersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServers indicates an expected call of ListServers
func (mr *MockInventoryClientMockRecorder) ListServers(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServers", reflect.TypeOf((*MockInventoryClient)(nil).ListServers), varargs...)
}

// GetDutInfo mocks base method
func (m *MockInventoryClient) GetDutInfo(ctx context.Context, in *GetDutInfoRequest, opts ...grpc.CallOption) (*GetDutInfoResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDutInfo", varargs...)
	ret0, _ := ret[0].(*GetDutInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDutInfo indicates an expected call of GetDutInfo
func (mr *MockInventoryClientMockRecorder) GetDutInfo(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDutInfo", reflect.TypeOf((*MockInventoryClient)(nil).GetDutInfo), varargs...)
}

// GetDroneConfig mocks base method
func (m *MockInventoryClient) GetDroneConfig(ctx context.Context, in *GetDroneConfigRequest, opts ...grpc.CallOption) (*GetDroneConfigResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDroneConfig", varargs...)
	ret0, _ := ret[0].(*GetDroneConfigResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDroneConfig indicates an expected call of GetDroneConfig
func (mr *MockInventoryClientMockRecorder) GetDroneConfig(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDroneConfig", reflect.TypeOf((*MockInventoryClient)(nil).GetDroneConfig), varargs...)
}

// UpdateDutLabels mocks base method
func (m *MockInventoryClient) UpdateDutLabels(ctx context.Context, in *UpdateDutLabelsRequest, opts ...grpc.CallOption) (*UpdateDutLabelsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateDutLabels", varargs...)
	ret0, _ := ret[0].(*UpdateDutLabelsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDutLabels indicates an expected call of UpdateDutLabels
func (mr *MockInventoryClientMockRecorder) UpdateDutLabels(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDutLabels", reflect.TypeOf((*MockInventoryClient)(nil).UpdateDutLabels), varargs...)
}

// UpdateCachedInventory mocks base method
func (m *MockInventoryClient) UpdateCachedInventory(ctx context.Context, in *UpdateCachedInventoryRequest, opts ...grpc.CallOption) (*UpdateCachedInventoryResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateCachedInventory", varargs...)
	ret0, _ := ret[0].(*UpdateCachedInventoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCachedInventory indicates an expected call of UpdateCachedInventory
func (mr *MockInventoryClientMockRecorder) UpdateCachedInventory(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCachedInventory", reflect.TypeOf((*MockInventoryClient)(nil).UpdateCachedInventory), varargs...)
}

// MockInventoryServer is a mock of InventoryServer interface
type MockInventoryServer struct {
	ctrl     *gomock.Controller
	recorder *MockInventoryServerMockRecorder
}

// MockInventoryServerMockRecorder is the mock recorder for MockInventoryServer
type MockInventoryServerMockRecorder struct {
	mock *MockInventoryServer
}

// NewMockInventoryServer creates a new mock instance
func NewMockInventoryServer(ctrl *gomock.Controller) *MockInventoryServer {
	mock := &MockInventoryServer{ctrl: ctrl}
	mock.recorder = &MockInventoryServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInventoryServer) EXPECT() *MockInventoryServerMockRecorder {
	return m.recorder
}

// DeployDut mocks base method
func (m *MockInventoryServer) DeployDut(arg0 context.Context, arg1 *DeployDutRequest) (*DeployDutResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeployDut", arg0, arg1)
	ret0, _ := ret[0].(*DeployDutResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeployDut indicates an expected call of DeployDut
func (mr *MockInventoryServerMockRecorder) DeployDut(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployDut", reflect.TypeOf((*MockInventoryServer)(nil).DeployDut), arg0, arg1)
}

// RedeployDut mocks base method
func (m *MockInventoryServer) RedeployDut(arg0 context.Context, arg1 *RedeployDutRequest) (*RedeployDutResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RedeployDut", arg0, arg1)
	ret0, _ := ret[0].(*RedeployDutResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RedeployDut indicates an expected call of RedeployDut
func (mr *MockInventoryServerMockRecorder) RedeployDut(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RedeployDut", reflect.TypeOf((*MockInventoryServer)(nil).RedeployDut), arg0, arg1)
}

// GetDeploymentStatus mocks base method
func (m *MockInventoryServer) GetDeploymentStatus(arg0 context.Context, arg1 *GetDeploymentStatusRequest) (*GetDeploymentStatusResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeploymentStatus", arg0, arg1)
	ret0, _ := ret[0].(*GetDeploymentStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeploymentStatus indicates an expected call of GetDeploymentStatus
func (mr *MockInventoryServerMockRecorder) GetDeploymentStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeploymentStatus", reflect.TypeOf((*MockInventoryServer)(nil).GetDeploymentStatus), arg0, arg1)
}

// DeleteDut mocks base method
func (m *MockInventoryServer) DeleteDut(arg0 context.Context, arg1 *DeleteDutRequest) (*DeleteDutResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDut", arg0, arg1)
	ret0, _ := ret[0].(*DeleteDutResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteDut indicates an expected call of DeleteDut
func (mr *MockInventoryServerMockRecorder) DeleteDut(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDut", reflect.TypeOf((*MockInventoryServer)(nil).DeleteDut), arg0, arg1)
}

// EnsurePoolHealthy mocks base method
func (m *MockInventoryServer) EnsurePoolHealthy(arg0 context.Context, arg1 *EnsurePoolHealthyRequest) (*EnsurePoolHealthyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsurePoolHealthy", arg0, arg1)
	ret0, _ := ret[0].(*EnsurePoolHealthyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsurePoolHealthy indicates an expected call of EnsurePoolHealthy
func (mr *MockInventoryServerMockRecorder) EnsurePoolHealthy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsurePoolHealthy", reflect.TypeOf((*MockInventoryServer)(nil).EnsurePoolHealthy), arg0, arg1)
}

// EnsurePoolHealthyForAllModels mocks base method
func (m *MockInventoryServer) EnsurePoolHealthyForAllModels(arg0 context.Context, arg1 *EnsurePoolHealthyForAllModelsRequest) (*EnsurePoolHealthyForAllModelsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsurePoolHealthyForAllModels", arg0, arg1)
	ret0, _ := ret[0].(*EnsurePoolHealthyForAllModelsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsurePoolHealthyForAllModels indicates an expected call of EnsurePoolHealthyForAllModels
func (mr *MockInventoryServerMockRecorder) EnsurePoolHealthyForAllModels(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsurePoolHealthyForAllModels", reflect.TypeOf((*MockInventoryServer)(nil).EnsurePoolHealthyForAllModels), arg0, arg1)
}

// ResizePool mocks base method
func (m *MockInventoryServer) ResizePool(arg0 context.Context, arg1 *ResizePoolRequest) (*ResizePoolResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResizePool", arg0, arg1)
	ret0, _ := ret[0].(*ResizePoolResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResizePool indicates an expected call of ResizePool
func (mr *MockInventoryServerMockRecorder) ResizePool(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResizePool", reflect.TypeOf((*MockInventoryServer)(nil).ResizePool), arg0, arg1)
}

// RemoveDutsFromDrones mocks base method
func (m *MockInventoryServer) RemoveDutsFromDrones(arg0 context.Context, arg1 *RemoveDutsFromDronesRequest) (*RemoveDutsFromDronesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveDutsFromDrones", arg0, arg1)
	ret0, _ := ret[0].(*RemoveDutsFromDronesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveDutsFromDrones indicates an expected call of RemoveDutsFromDrones
func (mr *MockInventoryServerMockRecorder) RemoveDutsFromDrones(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDutsFromDrones", reflect.TypeOf((*MockInventoryServer)(nil).RemoveDutsFromDrones), arg0, arg1)
}

// AssignDutsToDrones mocks base method
func (m *MockInventoryServer) AssignDutsToDrones(arg0 context.Context, arg1 *AssignDutsToDronesRequest) (*AssignDutsToDronesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignDutsToDrones", arg0, arg1)
	ret0, _ := ret[0].(*AssignDutsToDronesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignDutsToDrones indicates an expected call of AssignDutsToDrones
func (mr *MockInventoryServerMockRecorder) AssignDutsToDrones(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignDutsToDrones", reflect.TypeOf((*MockInventoryServer)(nil).AssignDutsToDrones), arg0, arg1)
}

// ListServers mocks base method
func (m *MockInventoryServer) ListServers(arg0 context.Context, arg1 *ListServersRequest) (*ListServersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServers", arg0, arg1)
	ret0, _ := ret[0].(*ListServersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServers indicates an expected call of ListServers
func (mr *MockInventoryServerMockRecorder) ListServers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServers", reflect.TypeOf((*MockInventoryServer)(nil).ListServers), arg0, arg1)
}

// GetDutInfo mocks base method
func (m *MockInventoryServer) GetDutInfo(arg0 context.Context, arg1 *GetDutInfoRequest) (*GetDutInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDutInfo", arg0, arg1)
	ret0, _ := ret[0].(*GetDutInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDutInfo indicates an expected call of GetDutInfo
func (mr *MockInventoryServerMockRecorder) GetDutInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDutInfo", reflect.TypeOf((*MockInventoryServer)(nil).GetDutInfo), arg0, arg1)
}

// GetDroneConfig mocks base method
func (m *MockInventoryServer) GetDroneConfig(arg0 context.Context, arg1 *GetDroneConfigRequest) (*GetDroneConfigResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDroneConfig", arg0, arg1)
	ret0, _ := ret[0].(*GetDroneConfigResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDroneConfig indicates an expected call of GetDroneConfig
func (mr *MockInventoryServerMockRecorder) GetDroneConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDroneConfig", reflect.TypeOf((*MockInventoryServer)(nil).GetDroneConfig), arg0, arg1)
}

// UpdateDutLabels mocks base method
func (m *MockInventoryServer) UpdateDutLabels(arg0 context.Context, arg1 *UpdateDutLabelsRequest) (*UpdateDutLabelsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDutLabels", arg0, arg1)
	ret0, _ := ret[0].(*UpdateDutLabelsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDutLabels indicates an expected call of UpdateDutLabels
func (mr *MockInventoryServerMockRecorder) UpdateDutLabels(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDutLabels", reflect.TypeOf((*MockInventoryServer)(nil).UpdateDutLabels), arg0, arg1)
}

// UpdateCachedInventory mocks base method
func (m *MockInventoryServer) UpdateCachedInventory(arg0 context.Context, arg1 *UpdateCachedInventoryRequest) (*UpdateCachedInventoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCachedInventory", arg0, arg1)
	ret0, _ := ret[0].(*UpdateCachedInventoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCachedInventory indicates an expected call of UpdateCachedInventory
func (mr *MockInventoryServerMockRecorder) UpdateCachedInventory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCachedInventory", reflect.TypeOf((*MockInventoryServer)(nil).UpdateCachedInventory), arg0, arg1)
}
