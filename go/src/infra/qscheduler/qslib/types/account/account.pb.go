// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infra/qscheduler/qslib/types/account/account.proto

package account

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	vector "infra/qscheduler/qslib/types/vector"
	math "math"
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

// Config represents per-quota-account configuration information, such
// as the recharge parameters. This does not represent anything about the
// current state of an account.
type Config struct {
	// ChargeRate is the rates (per second) at which per-priority accounts grow.
	//
	// Conceptually this is the time-averaged number of workers that this account
	// may use, at each priority level.
	ChargeRate *vector.Vector `protobuf:"bytes,1,opt,name=charge_rate,json=chargeRate,proto3" json:"charge_rate,omitempty"`
	// MaxChargeSeconds is the maximum amount of time over which this account can
	// accumulate quota before hitting its cap.
	//
	// Conceptually this sets the time window over which the time averaged
	// utilization by this account is measured. Very bursty clients will need to
	// use a wider window, whereas very consistent clients will use a narrow one.
	MaxChargeSeconds float64 `protobuf:"fixed64,2,opt,name=max_charge_seconds,json=maxChargeSeconds,proto3" json:"max_charge_seconds,omitempty"`
	// MaxFanout is the maximum number of concurrent paid jobs that this account
	// will pay for (0 = no limit).
	//
	// Additional jobs beyond this may run if there is idle capacity, but they
	// will run in the FreeBucket priority level.
	MaxFanout            int32    `protobuf:"varint,3,opt,name=max_fanout,json=maxFanout,proto3" json:"max_fanout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a30a156a0712643, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetChargeRate() *vector.Vector {
	if m != nil {
		return m.ChargeRate
	}
	return nil
}

func (m *Config) GetMaxChargeSeconds() float64 {
	if m != nil {
		return m.MaxChargeSeconds
	}
	return 0
}

func (m *Config) GetMaxFanout() int32 {
	if m != nil {
		return m.MaxFanout
	}
	return 0
}

func init() {
	proto.RegisterType((*Config)(nil), "account.Config")
}

func init() {
	proto.RegisterFile("infra/qscheduler/qslib/types/account/account.proto", fileDescriptor_2a30a156a0712643)
}

var fileDescriptor_2a30a156a0712643 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8e, 0xcd, 0x6a, 0x84, 0x30,
	0x14, 0x46, 0x49, 0x4b, 0x2d, 0x8d, 0x50, 0x4a, 0x56, 0x52, 0x28, 0x48, 0x57, 0x2e, 0x8a, 0x29,
	0xf6, 0x11, 0x84, 0x3e, 0x40, 0x06, 0x66, 0x2b, 0xd7, 0x78, 0xfd, 0x01, 0x4d, 0x34, 0x89, 0x83,
	0xf3, 0x02, 0xf3, 0xdc, 0xc3, 0xc4, 0xcc, 0x76, 0x56, 0x87, 0x7c, 0x39, 0x07, 0x2e, 0x2d, 0x06,
	0xd5, 0x1a, 0xe0, 0x8b, 0x95, 0x3d, 0x36, 0xeb, 0x88, 0x86, 0x2f, 0x76, 0x1c, 0x6a, 0xee, 0xce,
	0x33, 0x5a, 0x0e, 0x52, 0xea, 0x55, 0xb9, 0x3b, 0xf3, 0xd9, 0x68, 0xa7, 0xd9, 0x6b, 0x78, 0x7e,
	0xfe, 0x3e, 0x8c, 0x4f, 0x28, 0x9d, 0x36, 0x01, 0x7b, 0xfa, 0x7d, 0x21, 0x34, 0x2a, 0xb5, 0x6a,
	0x87, 0x8e, 0x71, 0x1a, 0xcb, 0x1e, 0x4c, 0x87, 0x95, 0x01, 0x87, 0x09, 0x49, 0x49, 0x16, 0x17,
	0xef, 0x79, 0xd0, 0x8f, 0x1e, 0x82, 0xee, 0x8a, 0x00, 0x87, 0xec, 0x87, 0xb2, 0x09, 0xb6, 0x2a,
	0x44, 0x16, 0xa5, 0x56, 0x8d, 0x4d, 0x9e, 0x52, 0x92, 0x11, 0xf1, 0x31, 0xc1, 0x56, 0xfa, 0x8f,
	0xc3, 0xbe, 0xb3, 0x2f, 0x4a, 0x6f, 0x76, 0x0b, 0x4a, 0xaf, 0x2e, 0x79, 0x4e, 0x49, 0xf6, 0x22,
	0xde, 0x26, 0xd8, 0xfe, 0xfd, 0x50, 0x47, 0xfe, 0x9e, 0xbf, 0x6b, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x6c, 0x71, 0xd8, 0xa1, 0x00, 0x01, 0x00, 0x00,
}
