// Copyright 2020 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: api/v3/api_proto/permissions.proto

package monorail_v3

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Request message for the GetPermissionSet emthod.
// Next available tag: 2
type GetPermissionSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the resource permissions to retrieve.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetPermissionSetRequest) Reset() {
	*x = GetPermissionSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v3_api_proto_permissions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPermissionSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPermissionSetRequest) ProtoMessage() {}

func (x *GetPermissionSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v3_api_proto_permissions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPermissionSetRequest.ProtoReflect.Descriptor instead.
func (*GetPermissionSetRequest) Descriptor() ([]byte, []int) {
	return file_api_v3_api_proto_permissions_proto_rawDescGZIP(), []int{0}
}

func (x *GetPermissionSetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request message for the BatchGetPermissionSets method.
// Next available tag: 2
type BatchGetPermissionSetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource names of the resource permissions to retrieve.
	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *BatchGetPermissionSetsRequest) Reset() {
	*x = BatchGetPermissionSetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v3_api_proto_permissions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetPermissionSetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetPermissionSetsRequest) ProtoMessage() {}

func (x *BatchGetPermissionSetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v3_api_proto_permissions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetPermissionSetsRequest.ProtoReflect.Descriptor instead.
func (*BatchGetPermissionSetsRequest) Descriptor() ([]byte, []int) {
	return file_api_v3_api_proto_permissions_proto_rawDescGZIP(), []int{1}
}

func (x *BatchGetPermissionSetsRequest) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

// Response message for the BatchGetPermissionSets method.
// Next available tag: 2
type BatchGetPermissionSetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Permissions, one for each of the given resources.
	PermissionSets []*PermissionSet `protobuf:"bytes,1,rep,name=permission_sets,json=permissionSets,proto3" json:"permission_sets,omitempty"`
}

func (x *BatchGetPermissionSetsResponse) Reset() {
	*x = BatchGetPermissionSetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v3_api_proto_permissions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetPermissionSetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetPermissionSetsResponse) ProtoMessage() {}

func (x *BatchGetPermissionSetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v3_api_proto_permissions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetPermissionSetsResponse.ProtoReflect.Descriptor instead.
func (*BatchGetPermissionSetsResponse) Descriptor() ([]byte, []int) {
	return file_api_v3_api_proto_permissions_proto_rawDescGZIP(), []int{2}
}

func (x *BatchGetPermissionSetsResponse) GetPermissionSets() []*PermissionSet {
	if x != nil {
		return x.PermissionSets
	}
	return nil
}

var File_api_v3_api_proto_permissions_proto protoreflect.FileDescriptor

var file_api_v3_api_proto_permissions_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x61, 0x69, 0x6c, 0x2e, 0x76,
	0x33, 0x1a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x29, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a,
	0x0a, 0x1d, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x65, 0x0a, 0x1e, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0f,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x61, 0x69, 0x6c,
	0x2e, 0x76, 0x33, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x74, 0x52, 0x0e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74,
	0x73, 0x32, 0xda, 0x01, 0x0a, 0x0b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x56, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x74, 0x12, 0x24, 0x2e, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x61, 0x69, 0x6c,
	0x2e, 0x76, 0x33, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d, 0x6f,
	0x6e, 0x6f, 0x72, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x22, 0x00, 0x12, 0x73, 0x0a, 0x16, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x74, 0x73, 0x12, 0x2a, 0x2e, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x61, 0x69, 0x6c, 0x2e, 0x76,
	0x33, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x33, 0x2e, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v3_api_proto_permissions_proto_rawDescOnce sync.Once
	file_api_v3_api_proto_permissions_proto_rawDescData = file_api_v3_api_proto_permissions_proto_rawDesc
)

func file_api_v3_api_proto_permissions_proto_rawDescGZIP() []byte {
	file_api_v3_api_proto_permissions_proto_rawDescOnce.Do(func() {
		file_api_v3_api_proto_permissions_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v3_api_proto_permissions_proto_rawDescData)
	})
	return file_api_v3_api_proto_permissions_proto_rawDescData
}

var file_api_v3_api_proto_permissions_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_v3_api_proto_permissions_proto_goTypes = []interface{}{
	(*GetPermissionSetRequest)(nil),        // 0: monorail.v3.GetPermissionSetRequest
	(*BatchGetPermissionSetsRequest)(nil),  // 1: monorail.v3.BatchGetPermissionSetsRequest
	(*BatchGetPermissionSetsResponse)(nil), // 2: monorail.v3.BatchGetPermissionSetsResponse
	(*PermissionSet)(nil),                  // 3: monorail.v3.PermissionSet
}
var file_api_v3_api_proto_permissions_proto_depIdxs = []int32{
	3, // 0: monorail.v3.BatchGetPermissionSetsResponse.permission_sets:type_name -> monorail.v3.PermissionSet
	0, // 1: monorail.v3.Permissions.GetPermissionSet:input_type -> monorail.v3.GetPermissionSetRequest
	1, // 2: monorail.v3.Permissions.BatchGetPermissionSets:input_type -> monorail.v3.BatchGetPermissionSetsRequest
	3, // 3: monorail.v3.Permissions.GetPermissionSet:output_type -> monorail.v3.PermissionSet
	2, // 4: monorail.v3.Permissions.BatchGetPermissionSets:output_type -> monorail.v3.BatchGetPermissionSetsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_v3_api_proto_permissions_proto_init() }
func file_api_v3_api_proto_permissions_proto_init() {
	if File_api_v3_api_proto_permissions_proto != nil {
		return
	}
	file_api_v3_api_proto_permission_objects_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_v3_api_proto_permissions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPermissionSetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v3_api_proto_permissions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetPermissionSetsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v3_api_proto_permissions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetPermissionSetsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v3_api_proto_permissions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v3_api_proto_permissions_proto_goTypes,
		DependencyIndexes: file_api_v3_api_proto_permissions_proto_depIdxs,
		MessageInfos:      file_api_v3_api_proto_permissions_proto_msgTypes,
	}.Build()
	File_api_v3_api_proto_permissions_proto = out.File
	file_api_v3_api_proto_permissions_proto_rawDesc = nil
	file_api_v3_api_proto_permissions_proto_goTypes = nil
	file_api_v3_api_proto_permissions_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PermissionsClient is the client API for Permissions service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PermissionsClient interface {
	// Returns the requester's permissions for the given resource.
	//
	// Raises:
	//  PERMISSION_DENIED if the given resource does not exist and/or the
	//      requester does not have permission to view the resource's name space.
	//  NOT_FOUND if the given resource does not exist.
	GetPermissionSet(ctx context.Context, in *GetPermissionSetRequest, opts ...grpc.CallOption) (*PermissionSet, error)
	// Returns the requester's permissions for all the given resources.
	//
	// Raises:
	//  PERMISSION_DENIED if any of the given resources do not exist and/or the
	//      requester does not have permission to view one of the resource's
	//      name space.
	// NOT_FOUND if one of the given resources do not exist.
	BatchGetPermissionSets(ctx context.Context, in *BatchGetPermissionSetsRequest, opts ...grpc.CallOption) (*BatchGetPermissionSetsResponse, error)
}
type permissionsPRPCClient struct {
	client *prpc.Client
}

func NewPermissionsPRPCClient(client *prpc.Client) PermissionsClient {
	return &permissionsPRPCClient{client}
}

func (c *permissionsPRPCClient) GetPermissionSet(ctx context.Context, in *GetPermissionSetRequest, opts ...grpc.CallOption) (*PermissionSet, error) {
	out := new(PermissionSet)
	err := c.client.Call(ctx, "monorail.v3.Permissions", "GetPermissionSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsPRPCClient) BatchGetPermissionSets(ctx context.Context, in *BatchGetPermissionSetsRequest, opts ...grpc.CallOption) (*BatchGetPermissionSetsResponse, error) {
	out := new(BatchGetPermissionSetsResponse)
	err := c.client.Call(ctx, "monorail.v3.Permissions", "BatchGetPermissionSets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type permissionsClient struct {
	cc grpc.ClientConnInterface
}

func NewPermissionsClient(cc grpc.ClientConnInterface) PermissionsClient {
	return &permissionsClient{cc}
}

func (c *permissionsClient) GetPermissionSet(ctx context.Context, in *GetPermissionSetRequest, opts ...grpc.CallOption) (*PermissionSet, error) {
	out := new(PermissionSet)
	err := c.cc.Invoke(ctx, "/monorail.v3.Permissions/GetPermissionSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsClient) BatchGetPermissionSets(ctx context.Context, in *BatchGetPermissionSetsRequest, opts ...grpc.CallOption) (*BatchGetPermissionSetsResponse, error) {
	out := new(BatchGetPermissionSetsResponse)
	err := c.cc.Invoke(ctx, "/monorail.v3.Permissions/BatchGetPermissionSets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionsServer is the server API for Permissions service.
type PermissionsServer interface {
	// Returns the requester's permissions for the given resource.
	//
	// Raises:
	//  PERMISSION_DENIED if the given resource does not exist and/or the
	//      requester does not have permission to view the resource's name space.
	//  NOT_FOUND if the given resource does not exist.
	GetPermissionSet(context.Context, *GetPermissionSetRequest) (*PermissionSet, error)
	// Returns the requester's permissions for all the given resources.
	//
	// Raises:
	//  PERMISSION_DENIED if any of the given resources do not exist and/or the
	//      requester does not have permission to view one of the resource's
	//      name space.
	// NOT_FOUND if one of the given resources do not exist.
	BatchGetPermissionSets(context.Context, *BatchGetPermissionSetsRequest) (*BatchGetPermissionSetsResponse, error)
}

// UnimplementedPermissionsServer can be embedded to have forward compatible implementations.
type UnimplementedPermissionsServer struct {
}

func (*UnimplementedPermissionsServer) GetPermissionSet(context.Context, *GetPermissionSetRequest) (*PermissionSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermissionSet not implemented")
}
func (*UnimplementedPermissionsServer) BatchGetPermissionSets(context.Context, *BatchGetPermissionSetsRequest) (*BatchGetPermissionSetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetPermissionSets not implemented")
}

func RegisterPermissionsServer(s prpc.Registrar, srv PermissionsServer) {
	s.RegisterService(&_Permissions_serviceDesc, srv)
}

func _Permissions_GetPermissionSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPermissionSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServer).GetPermissionSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monorail.v3.Permissions/GetPermissionSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServer).GetPermissionSet(ctx, req.(*GetPermissionSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Permissions_BatchGetPermissionSets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetPermissionSetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServer).BatchGetPermissionSets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monorail.v3.Permissions/BatchGetPermissionSets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServer).BatchGetPermissionSets(ctx, req.(*BatchGetPermissionSetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Permissions_serviceDesc = grpc.ServiceDesc{
	ServiceName: "monorail.v3.Permissions",
	HandlerType: (*PermissionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPermissionSet",
			Handler:    _Permissions_GetPermissionSet_Handler,
		},
		{
			MethodName: "BatchGetPermissionSets",
			Handler:    _Permissions_BatchGetPermissionSets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v3/api_proto/permissions.proto",
}
