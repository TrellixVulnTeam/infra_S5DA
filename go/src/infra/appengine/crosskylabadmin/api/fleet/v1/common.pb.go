// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infra/appengine/crosskylabadmin/api/fleet/v1/common.proto

package fleet

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

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

var ServerRole_name = map[int32]string{
	0:  "ROLE_INVALID",
	2:  "ROLE_CRASH_SERVER",
	5:  "ROLE_DEVSERVER",
	12: "ROLE_SUITE_SCHEDULER",
	13: "ROLE_SKYLAB_DRONE",
	14: "ROLE_SKYLAB_SUITE_PROXY",
	15: "ROLE_RPMSERVER",
}

var ServerRole_value = map[string]int32{
	"ROLE_INVALID":            0,
	"ROLE_CRASH_SERVER":       2,
	"ROLE_DEVSERVER":          5,
	"ROLE_SUITE_SCHEDULER":    12,
	"ROLE_SKYLAB_DRONE":       13,
	"ROLE_SKYLAB_SUITE_PROXY": 14,
	"ROLE_RPMSERVER":          15,
}

func (x ServerRole) String() string {
	return proto.EnumName(ServerRole_name, int32(x))
}

func (ServerRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e1fa3f8860ba148c, []int{0}
}

// BotSelector is used in various fleet RPCs to filter the Swarming bots that
// the RPC applies to.
//
// For example, it is used to select the bots to summarize by the Tracker RPCs,
// and to select the bots against which admin tasks are managed by the Tasker
// RPCs.
type BotSelector struct {
	// dut_id selects a bot by the dut_id dimension.
	DutId string `protobuf:"bytes,1,opt,name=dut_id,json=dutId,proto3" json:"dut_id,omitempty"`
	// dimensions select bots by Swarming dimensions.
	//
	// All fields in the dimension message must match for a bot to be selected.
	Dimensions           *BotDimensions `protobuf:"bytes,2,opt,name=dimensions,proto3" json:"dimensions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BotSelector) Reset()         { *m = BotSelector{} }
func (m *BotSelector) String() string { return proto.CompactTextString(m) }
func (*BotSelector) ProtoMessage()    {}
func (*BotSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1fa3f8860ba148c, []int{0}
}

func (m *BotSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BotSelector.Unmarshal(m, b)
}
func (m *BotSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BotSelector.Marshal(b, m, deterministic)
}
func (m *BotSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BotSelector.Merge(m, src)
}
func (m *BotSelector) XXX_Size() int {
	return xxx_messageInfo_BotSelector.Size(m)
}
func (m *BotSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_BotSelector.DiscardUnknown(m)
}

var xxx_messageInfo_BotSelector proto.InternalMessageInfo

func (m *BotSelector) GetDutId() string {
	if m != nil {
		return m.DutId
	}
	return ""
}

func (m *BotSelector) GetDimensions() *BotDimensions {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

// BotDimensions is a subset of Swarming bot dimensions.
type BotDimensions struct {
	Pools                []string `protobuf:"bytes,1,rep,name=pools,proto3" json:"pools,omitempty"`
	Model                string   `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	DutName              string   `protobuf:"bytes,3,opt,name=dut_name,json=dutName,proto3" json:"dut_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BotDimensions) Reset()         { *m = BotDimensions{} }
func (m *BotDimensions) String() string { return proto.CompactTextString(m) }
func (*BotDimensions) ProtoMessage()    {}
func (*BotDimensions) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1fa3f8860ba148c, []int{1}
}

func (m *BotDimensions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BotDimensions.Unmarshal(m, b)
}
func (m *BotDimensions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BotDimensions.Marshal(b, m, deterministic)
}
func (m *BotDimensions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BotDimensions.Merge(m, src)
}
func (m *BotDimensions) XXX_Size() int {
	return xxx_messageInfo_BotDimensions.Size(m)
}
func (m *BotDimensions) XXX_DiscardUnknown() {
	xxx_messageInfo_BotDimensions.DiscardUnknown(m)
}

var xxx_messageInfo_BotDimensions proto.InternalMessageInfo

func (m *BotDimensions) GetPools() []string {
	if m != nil {
		return m.Pools
	}
	return nil
}

func (m *BotDimensions) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *BotDimensions) GetDutName() string {
	if m != nil {
		return m.DutName
	}
	return ""
}

// Server is adapted from chrome.chromeos_infra.skylab.proto.inventory
// (Server message type). It is filtered down to only aspects applicable
// to Skylab, and of interest to clients of the crosskylabadmin APIs.
type Server struct {
	Hostname             string       `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Roles                []ServerRole `protobuf:"varint,2,rep,packed,name=roles,proto3,enum=crosskylabadmin.fleet.ServerRole" json:"roles,omitempty"`
	DutUids              []string     `protobuf:"bytes,3,rep,name=dut_uids,json=dutUids,proto3" json:"dut_uids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Server) Reset()         { *m = Server{} }
func (m *Server) String() string { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()    {}
func (*Server) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1fa3f8860ba148c, []int{2}
}

func (m *Server) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Server.Unmarshal(m, b)
}
func (m *Server) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Server.Marshal(b, m, deterministic)
}
func (m *Server) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Server.Merge(m, src)
}
func (m *Server) XXX_Size() int {
	return xxx_messageInfo_Server.Size(m)
}
func (m *Server) XXX_DiscardUnknown() {
	xxx_messageInfo_Server.DiscardUnknown(m)
}

var xxx_messageInfo_Server proto.InternalMessageInfo

func (m *Server) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Server) GetRoles() []ServerRole {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *Server) GetDutUids() []string {
	if m != nil {
		return m.DutUids
	}
	return nil
}

func init() {
	proto.RegisterEnum("crosskylabadmin.fleet.ServerRole", ServerRole_name, ServerRole_value)
	proto.RegisterType((*BotSelector)(nil), "crosskylabadmin.fleet.BotSelector")
	proto.RegisterType((*BotDimensions)(nil), "crosskylabadmin.fleet.BotDimensions")
	proto.RegisterType((*Server)(nil), "crosskylabadmin.fleet.Server")
}

func init() {
	proto.RegisterFile("infra/appengine/crosskylabadmin/api/fleet/v1/common.proto", fileDescriptor_e1fa3f8860ba148c)
}

var fileDescriptor_e1fa3f8860ba148c = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0xc9, 0xa2, 0x74, 0xdb, 0xdb, 0x56, 0x82, 0xb5, 0x8a, 0x00, 0x97, 0x52, 0x71, 0xa8,
	0x38, 0x24, 0x62, 0x1c, 0x10, 0xc7, 0x66, 0xb1, 0xb4, 0x88, 0xd0, 0x4e, 0x0e, 0x8d, 0x18, 0x97,
	0x28, 0xab, 0x3d, 0x30, 0x24, 0x76, 0x14, 0x3b, 0x15, 0xfc, 0x4d, 0xfc, 0x93, 0x28, 0x4e, 0x09,
	0x05, 0xb1, 0xe3, 0xfb, 0xfc, 0xbd, 0xef, 0xf7, 0x3d, 0xc9, 0xf0, 0x96, 0x8b, 0xbb, 0xa6, 0x08,
	0x8a, 0xba, 0x66, 0xe2, 0x33, 0x17, 0x2c, 0xd8, 0x34, 0x52, 0xa9, 0x6f, 0x3f, 0xca, 0xe2, 0xb6,
	0xa0, 0x15, 0x17, 0x41, 0x51, 0xf3, 0xe0, 0xae, 0x64, 0x4c, 0x07, 0xdb, 0x57, 0xc1, 0x46, 0x56,
	0x95, 0x14, 0x7e, 0xdd, 0x48, 0x2d, 0xd1, 0xe4, 0x1f, 0xab, 0x6f, 0x6c, 0xb3, 0xaf, 0x70, 0x12,
	0x4a, 0x9d, 0xb2, 0x92, 0x6d, 0xb4, 0x6c, 0xd0, 0x04, 0x46, 0xb4, 0xd5, 0x39, 0xa7, 0x9e, 0x35,
	0xb5, 0xe6, 0xc7, 0xc4, 0xa1, 0xad, 0x8e, 0x29, 0x8a, 0x00, 0x28, 0xaf, 0x98, 0x50, 0x5c, 0x0a,
	0xe5, 0x1d, 0x4c, 0xad, 0xf9, 0xc9, 0xc5, 0x0b, 0xff, 0xbf, 0x89, 0x7e, 0x28, 0x75, 0x34, 0x78,
	0xc9, 0xde, 0xde, 0x2c, 0x83, 0xb3, 0xbf, 0x1e, 0xd1, 0x39, 0x38, 0xb5, 0x94, 0xa5, 0xf2, 0xac,
	0xa9, 0xdd, 0xc1, 0xcc, 0xd0, 0xa9, 0x95, 0xa4, 0xac, 0x34, 0x9c, 0x63, 0xd2, 0x0f, 0xe8, 0x09,
	0x1c, 0x75, 0xcd, 0x44, 0x51, 0x31, 0xcf, 0x36, 0x0f, 0x87, 0xb4, 0xd5, 0xcb, 0xa2, 0x62, 0xb3,
	0xef, 0x30, 0x4a, 0x59, 0xb3, 0x65, 0x0d, 0x7a, 0x0a, 0x47, 0x5f, 0xa4, 0xd2, 0xc6, 0xd4, 0x1f,
	0x30, 0xcc, 0xe8, 0x0d, 0x38, 0x8d, 0x2c, 0x59, 0x57, 0xdf, 0x9e, 0x8f, 0x2f, 0x9e, 0xdf, 0x53,
	0xbf, 0x4f, 0x22, 0xb2, 0x64, 0xa4, 0xf7, 0xff, 0x26, 0xb7, 0x9c, 0x2a, 0xcf, 0x36, 0x45, 0x3b,
	0xf2, 0x9a, 0x53, 0xf5, 0xf2, 0xa7, 0x05, 0xf0, 0x67, 0x01, 0xb9, 0x70, 0x4a, 0x56, 0x09, 0xce,
	0xe3, 0x65, 0xb6, 0x48, 0xe2, 0xc8, 0x7d, 0x80, 0x26, 0xf0, 0xc8, 0x28, 0x97, 0x64, 0x91, 0x5e,
	0xe5, 0x29, 0x26, 0x19, 0x26, 0xee, 0x01, 0x42, 0x30, 0x36, 0x72, 0x84, 0xb3, 0x9d, 0xe6, 0x20,
	0x0f, 0xce, 0x8d, 0x96, 0xae, 0xe3, 0x0f, 0x38, 0x4f, 0x2f, 0xaf, 0x70, 0xb4, 0x4e, 0x30, 0x71,
	0x4f, 0x87, 0x90, 0xf4, 0xdd, 0x4d, 0xb2, 0x08, 0xf3, 0x88, 0xac, 0x96, 0xd8, 0x3d, 0x43, 0xcf,
	0xe0, 0xf1, 0xbe, 0xdc, 0xef, 0x5d, 0x93, 0xd5, 0xc7, 0x1b, 0x77, 0x3c, 0x10, 0xc8, 0xf5, 0xfb,
	0x1d, 0xe1, 0x61, 0x78, 0xf8, 0xc9, 0x31, 0x37, 0xde, 0x8e, 0xcc, 0x97, 0x78, 0xfd, 0x2b, 0x00,
	0x00, 0xff, 0xff, 0x26, 0x9d, 0xfb, 0x38, 0x4f, 0x02, 0x00, 0x00,
}
