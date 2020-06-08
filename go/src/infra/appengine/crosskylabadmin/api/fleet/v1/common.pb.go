// Copyright 2018 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.6.1
// source: infra/appengine/crosskylabadmin/api/fleet/v1/common.proto

package fleet

import (
	proto "github.com/golang/protobuf/proto"
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

// ServerRole is adapted from chrome.chromeos_infra.skylab.proto.inventory
// (Server.Role enum). It is filtered down to only server roles that are
// applicable to Skylab.
//
// Commented-out fields are those that exist in the inventory proto schema, but
// do no apply to this API.
type ServerRole int32

const (
	ServerRole_ROLE_INVALID ServerRole = 0
	// ROLE_AFE = 1;
	ServerRole_ROLE_CRASH_SERVER ServerRole = 2
	// ROLE_DATABASE = 3;
	// ROLE_DATABASE_SLAVE = 4;
	ServerRole_ROLE_DEVSERVER ServerRole = 5
	// ROLE_DRONE = 6;
	// ROLE_GOLO_PROXY = 7;
	// ROLE_HOST_SCHEDULER = 8;
	// ROLE_SCHEDULER = 9;
	// ROLE_SENTINEL = 10;
	// ROLE_SHARD = 11;
	ServerRole_ROLE_SUITE_SCHEDULER    ServerRole = 12
	ServerRole_ROLE_SKYLAB_DRONE       ServerRole = 13
	ServerRole_ROLE_SKYLAB_SUITE_PROXY ServerRole = 14
	ServerRole_ROLE_RPMSERVER          ServerRole = 15
)

// Enum value maps for ServerRole.
var (
	ServerRole_name = map[int32]string{
		0:  "ROLE_INVALID",
		2:  "ROLE_CRASH_SERVER",
		5:  "ROLE_DEVSERVER",
		12: "ROLE_SUITE_SCHEDULER",
		13: "ROLE_SKYLAB_DRONE",
		14: "ROLE_SKYLAB_SUITE_PROXY",
		15: "ROLE_RPMSERVER",
	}
	ServerRole_value = map[string]int32{
		"ROLE_INVALID":            0,
		"ROLE_CRASH_SERVER":       2,
		"ROLE_DEVSERVER":          5,
		"ROLE_SUITE_SCHEDULER":    12,
		"ROLE_SKYLAB_DRONE":       13,
		"ROLE_SKYLAB_SUITE_PROXY": 14,
		"ROLE_RPMSERVER":          15,
	}
)

func (x ServerRole) Enum() *ServerRole {
	p := new(ServerRole)
	*p = x
	return p
}

func (x ServerRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServerRole) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_enumTypes[0].Descriptor()
}

func (ServerRole) Type() protoreflect.EnumType {
	return &file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_enumTypes[0]
}

func (x ServerRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServerRole.Descriptor instead.
func (ServerRole) EnumDescriptor() ([]byte, []int) {
	return file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_rawDescGZIP(), []int{0}
}

// BotSelector is used in various fleet RPCs to filter the Swarming bots that
// the RPC applies to.
//
// For example, it is used to select the bots to summarize by the Tracker RPCs,
// and to select the bots against which admin tasks are managed by the Tasker
// RPCs.
type BotSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// dut_id selects a bot by the dut_id dimension.
	DutId string `protobuf:"bytes,1,opt,name=dut_id,json=dutId,proto3" json:"dut_id,omitempty"`
	// dimensions select bots by Swarming dimensions.
	//
	// All fields in the dimension message must match for a bot to be selected.
	Dimensions *BotDimensions `protobuf:"bytes,2,opt,name=dimensions,proto3" json:"dimensions,omitempty"`
}

func (x *BotSelector) Reset() {
	*x = BotSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BotSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BotSelector) ProtoMessage() {}

func (x *BotSelector) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BotSelector.ProtoReflect.Descriptor instead.
func (*BotSelector) Descriptor() ([]byte, []int) {
	return file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *BotSelector) GetDutId() string {
	if x != nil {
		return x.DutId
	}
	return ""
}

func (x *BotSelector) GetDimensions() *BotDimensions {
	if x != nil {
		return x.Dimensions
	}
	return nil
}

// BotDimensions is a subset of Swarming bot dimensions.
type BotDimensions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pools   []string `protobuf:"bytes,1,rep,name=pools,proto3" json:"pools,omitempty"`
	Model   string   `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	DutName string   `protobuf:"bytes,3,opt,name=dut_name,json=dutName,proto3" json:"dut_name,omitempty"`
}

func (x *BotDimensions) Reset() {
	*x = BotDimensions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BotDimensions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BotDimensions) ProtoMessage() {}

func (x *BotDimensions) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BotDimensions.ProtoReflect.Descriptor instead.
func (*BotDimensions) Descriptor() ([]byte, []int) {
	return file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *BotDimensions) GetPools() []string {
	if x != nil {
		return x.Pools
	}
	return nil
}

func (x *BotDimensions) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *BotDimensions) GetDutName() string {
	if x != nil {
		return x.DutName
	}
	return ""
}

// Server is adapted from chrome.chromeos_infra.skylab.proto.inventory
// (Server message type). It is filtered down to only aspects applicable
// to Skylab, and of interest to clients of the crosskylabadmin APIs.
type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname string       `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Roles    []ServerRole `protobuf:"varint,2,rep,packed,name=roles,proto3,enum=crosskylabadmin.fleet.ServerRole" json:"roles,omitempty"`
	DutUids  []string     `protobuf:"bytes,3,rep,name=dut_uids,json=dutUids,proto3" json:"dut_uids,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_rawDescGZIP(), []int{2}
}

func (x *Server) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Server) GetRoles() []ServerRole {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *Server) GetDutUids() []string {
	if x != nil {
		return x.DutUids
	}
	return nil
}

var File_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto protoreflect.FileDescriptor

var file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_rawDesc = []byte{
	0x0a, 0x39, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x6b, 0x79, 0x6c, 0x61, 0x62, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x72, 0x6f,
	0x73, 0x73, 0x6b, 0x79, 0x6c, 0x61, 0x62, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x66, 0x6c, 0x65,
	0x65, 0x74, 0x22, 0x6a, 0x0a, 0x0b, 0x42, 0x6f, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x12, 0x15, 0x0a, 0x06, 0x64, 0x75, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x64, 0x75, 0x74, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x0a, 0x64, 0x69, 0x6d, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63,
	0x72, 0x6f, 0x73, 0x73, 0x6b, 0x79, 0x6c, 0x61, 0x62, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x2e, 0x42, 0x6f, 0x74, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x0a, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x56,
	0x0a, 0x0d, 0x42, 0x6f, 0x74, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x64,
	0x75, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64,
	0x75, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x78, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x05,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x63, 0x72,
	0x6f, 0x73, 0x73, 0x6b, 0x79, 0x6c, 0x61, 0x62, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x66, 0x6c,
	0x65, 0x65, 0x74, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x75, 0x74, 0x5f, 0x75, 0x69, 0x64,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64, 0x75, 0x74, 0x55, 0x69, 0x64, 0x73,
	0x2a, 0xab, 0x01, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x10, 0x0a, 0x0c, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10,
	0x00, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x43, 0x52, 0x41, 0x53, 0x48, 0x5f,
	0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x4f, 0x4c, 0x45,
	0x5f, 0x44, 0x45, 0x56, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x05, 0x12, 0x18, 0x0a, 0x14,
	0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x53, 0x55, 0x49, 0x54, 0x45, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x44,
	0x55, 0x4c, 0x45, 0x52, 0x10, 0x0c, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x53,
	0x4b, 0x59, 0x4c, 0x41, 0x42, 0x5f, 0x44, 0x52, 0x4f, 0x4e, 0x45, 0x10, 0x0d, 0x12, 0x1b, 0x0a,
	0x17, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x53, 0x4b, 0x59, 0x4c, 0x41, 0x42, 0x5f, 0x53, 0x55, 0x49,
	0x54, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x58, 0x59, 0x10, 0x0e, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x4f,
	0x4c, 0x45, 0x5f, 0x52, 0x50, 0x4d, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x0f, 0x42, 0x07,
	0x5a, 0x05, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_rawDescOnce sync.Once
	file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_rawDescData = file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_rawDesc
)

func file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_rawDescGZIP() []byte {
	file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_rawDescOnce.Do(func() {
		file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_rawDescData)
	})
	return file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_rawDescData
}

var file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_goTypes = []interface{}{
	(ServerRole)(0),       // 0: crosskylabadmin.fleet.ServerRole
	(*BotSelector)(nil),   // 1: crosskylabadmin.fleet.BotSelector
	(*BotDimensions)(nil), // 2: crosskylabadmin.fleet.BotDimensions
	(*Server)(nil),        // 3: crosskylabadmin.fleet.Server
}
var file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_depIdxs = []int32{
	2, // 0: crosskylabadmin.fleet.BotSelector.dimensions:type_name -> crosskylabadmin.fleet.BotDimensions
	0, // 1: crosskylabadmin.fleet.Server.roles:type_name -> crosskylabadmin.fleet.ServerRole
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_init() }
func file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_init() {
	if File_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BotSelector); i {
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
		file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BotDimensions); i {
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
		file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
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
			RawDescriptor: file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_goTypes,
		DependencyIndexes: file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_depIdxs,
		EnumInfos:         file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_enumTypes,
		MessageInfos:      file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_msgTypes,
	}.Build()
	File_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto = out.File
	file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_rawDesc = nil
	file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_goTypes = nil
	file_infra_appengine_crosskylabadmin_api_fleet_v1_common_proto_depIdxs = nil
}
