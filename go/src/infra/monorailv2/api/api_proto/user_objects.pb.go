// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/api_proto/user_objects.proto

package monorail

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// TODO(jojwang): monorail:1701, fill User with all info necessary for
// creating a user profile page.
// Next available tag: 7
type User struct {
	Email                string     `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	UserId               int64      `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	IsSiteAdmin          bool       `protobuf:"varint,3,opt,name=is_site_admin,json=isSiteAdmin,proto3" json:"is_site_admin,omitempty"`
	Availability         string     `protobuf:"bytes,4,opt,name=availability,proto3" json:"availability,omitempty"`
	LinkedParentRef      *UserRef   `protobuf:"bytes,5,opt,name=linked_parent_ref,json=linkedParentRef,proto3" json:"linked_parent_ref,omitempty"`
	LinkedChildRefs      []*UserRef `protobuf:"bytes,6,rep,name=linked_child_refs,json=linkedChildRefs,proto3" json:"linked_child_refs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_e651956a3fdc871c, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *User) GetIsSiteAdmin() bool {
	if m != nil {
		return m.IsSiteAdmin
	}
	return false
}

func (m *User) GetAvailability() string {
	if m != nil {
		return m.Availability
	}
	return ""
}

func (m *User) GetLinkedParentRef() *UserRef {
	if m != nil {
		return m.LinkedParentRef
	}
	return nil
}

func (m *User) GetLinkedChildRefs() []*UserRef {
	if m != nil {
		return m.LinkedChildRefs
	}
	return nil
}

// Next available tag: 3
type UserPrefValue struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserPrefValue) Reset()         { *m = UserPrefValue{} }
func (m *UserPrefValue) String() string { return proto.CompactTextString(m) }
func (*UserPrefValue) ProtoMessage()    {}
func (*UserPrefValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_e651956a3fdc871c, []int{1}
}

func (m *UserPrefValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserPrefValue.Unmarshal(m, b)
}
func (m *UserPrefValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserPrefValue.Marshal(b, m, deterministic)
}
func (m *UserPrefValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserPrefValue.Merge(m, src)
}
func (m *UserPrefValue) XXX_Size() int {
	return xxx_messageInfo_UserPrefValue.Size(m)
}
func (m *UserPrefValue) XXX_DiscardUnknown() {
	xxx_messageInfo_UserPrefValue.DiscardUnknown(m)
}

var xxx_messageInfo_UserPrefValue proto.InternalMessageInfo

func (m *UserPrefValue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserPrefValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Next available tag: 6
type UserProjects struct {
	UserRef              *UserRef `protobuf:"bytes,1,opt,name=user_ref,json=userRef,proto3" json:"user_ref,omitempty"`
	OwnerOf              []string `protobuf:"bytes,2,rep,name=owner_of,json=ownerOf,proto3" json:"owner_of,omitempty"`
	MemberOf             []string `protobuf:"bytes,3,rep,name=member_of,json=memberOf,proto3" json:"member_of,omitempty"`
	ContributorTo        []string `protobuf:"bytes,4,rep,name=contributor_to,json=contributorTo,proto3" json:"contributor_to,omitempty"`
	StarredProjects      []string `protobuf:"bytes,5,rep,name=starred_projects,json=starredProjects,proto3" json:"starred_projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserProjects) Reset()         { *m = UserProjects{} }
func (m *UserProjects) String() string { return proto.CompactTextString(m) }
func (*UserProjects) ProtoMessage()    {}
func (*UserProjects) Descriptor() ([]byte, []int) {
	return fileDescriptor_e651956a3fdc871c, []int{2}
}

func (m *UserProjects) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserProjects.Unmarshal(m, b)
}
func (m *UserProjects) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserProjects.Marshal(b, m, deterministic)
}
func (m *UserProjects) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserProjects.Merge(m, src)
}
func (m *UserProjects) XXX_Size() int {
	return xxx_messageInfo_UserProjects.Size(m)
}
func (m *UserProjects) XXX_DiscardUnknown() {
	xxx_messageInfo_UserProjects.DiscardUnknown(m)
}

var xxx_messageInfo_UserProjects proto.InternalMessageInfo

func (m *UserProjects) GetUserRef() *UserRef {
	if m != nil {
		return m.UserRef
	}
	return nil
}

func (m *UserProjects) GetOwnerOf() []string {
	if m != nil {
		return m.OwnerOf
	}
	return nil
}

func (m *UserProjects) GetMemberOf() []string {
	if m != nil {
		return m.MemberOf
	}
	return nil
}

func (m *UserProjects) GetContributorTo() []string {
	if m != nil {
		return m.ContributorTo
	}
	return nil
}

func (m *UserProjects) GetStarredProjects() []string {
	if m != nil {
		return m.StarredProjects
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "monorail.User")
	proto.RegisterType((*UserPrefValue)(nil), "monorail.UserPrefValue")
	proto.RegisterType((*UserProjects)(nil), "monorail.UserProjects")
}

func init() { proto.RegisterFile("api/api_proto/user_objects.proto", fileDescriptor_e651956a3fdc871c) }

var fileDescriptor_e651956a3fdc871c = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x95, 0x37, 0x69, 0x37, 0xf1, 0x6e, 0x59, 0xd6, 0x42, 0xc2, 0x2c, 0x97, 0x28, 0x12, 0x52,
	0x90, 0x50, 0x2b, 0xc1, 0x89, 0x03, 0x07, 0xc4, 0x89, 0x53, 0x2b, 0xf3, 0x71, 0x8d, 0x9c, 0x64,
	0x2c, 0x06, 0x62, 0x3b, 0xb2, 0xdd, 0x22, 0xfe, 0x20, 0x7f, 0x8a, 0x0b, 0xb2, 0xd3, 0x4a, 0xed,
	0x01, 0x6e, 0x99, 0xf7, 0x31, 0x7a, 0xf3, 0x62, 0x5a, 0xc9, 0x09, 0x37, 0x72, 0xc2, 0x76, 0x72,
	0x36, 0xd8, 0xcd, 0xde, 0x83, 0x6b, 0x6d, 0xf7, 0x1d, 0xfa, 0xe0, 0xd7, 0x09, 0x62, 0x85, 0xb6,
	0xc6, 0x3a, 0x89, 0xe3, 0xc3, 0xc3, 0xa5, 0xb6, 0xb7, 0x5a, 0x5b, 0x33, 0xab, 0xea, 0x3f, 0x84,
	0xe6, 0x5f, 0x3c, 0x38, 0xf6, 0x84, 0x2e, 0x40, 0x4b, 0x1c, 0x39, 0xa9, 0x48, 0x53, 0x8a, 0x79,
	0x60, 0x4f, 0xe9, 0x75, 0x5a, 0x8d, 0x03, 0xbf, 0xaa, 0x48, 0x93, 0x89, 0x65, 0x1c, 0x3f, 0x0e,
	0xac, 0xa6, 0x2b, 0xf4, 0xad, 0xc7, 0x00, 0xad, 0x1c, 0x34, 0x1a, 0x9e, 0x55, 0xa4, 0x29, 0xc4,
	0x0d, 0xfa, 0x4f, 0x18, 0xe0, 0x7d, 0x84, 0x58, 0x4d, 0x6f, 0xe5, 0x41, 0xe2, 0x28, 0x3b, 0x1c,
	0x31, 0xfc, 0xe2, 0x79, 0xda, 0x7c, 0x81, 0xb1, 0x77, 0xf4, 0x7e, 0x44, 0xf3, 0x03, 0x86, 0x76,
	0x92, 0x0e, 0x4c, 0x68, 0x1d, 0x28, 0xbe, 0xa8, 0x48, 0x73, 0xf3, 0xfa, 0x7e, 0x7d, 0xba, 0x60,
	0x1d, 0x13, 0x0a, 0x50, 0xe2, 0x6e, 0xd6, 0xee, 0x92, 0x54, 0x80, 0x3a, 0xb3, 0xf7, 0xdf, 0x70,
	0x1c, 0xa2, 0xdb, 0xf3, 0x65, 0x95, 0xfd, 0xd7, 0xfe, 0x21, 0x4a, 0x05, 0x28, 0x5f, 0xbf, 0xa5,
	0xab, 0xc8, 0xed, 0x1c, 0xa8, 0xaf, 0x72, 0xdc, 0x03, 0x63, 0x34, 0x37, 0x52, 0xc3, 0xb1, 0x84,
	0xf4, 0x1d, 0x9b, 0x39, 0x44, 0x32, 0x35, 0x50, 0x8a, 0x79, 0xa8, 0x7f, 0x13, 0x7a, 0x3b, 0x7b,
	0x6d, 0x6a, 0x9d, 0xbd, 0xa2, 0x45, 0xaa, 0x2a, 0x1e, 0x40, 0xfe, 0x75, 0x40, 0x6a, 0x33, 0x06,
	0x7f, 0x46, 0x0b, 0xfb, 0xd3, 0xc4, 0x9f, 0xa6, 0xf8, 0x55, 0x95, 0x35, 0xa5, 0xb8, 0x4e, 0xf3,
	0x56, 0xb1, 0xe7, 0xb4, 0xd4, 0xa0, 0xbb, 0x99, 0xcb, 0x12, 0x57, 0xcc, 0xc0, 0x56, 0xb1, 0x17,
	0xf4, 0x51, 0x6f, 0x4d, 0x70, 0xd8, 0xed, 0x83, 0x75, 0x6d, 0xb0, 0x3c, 0x4f, 0x8a, 0xd5, 0x19,
	0xfa, 0xd9, 0xb2, 0x97, 0xf4, 0xb1, 0x0f, 0xd2, 0xb9, 0xd8, 0xeb, 0x31, 0x20, 0x5f, 0x24, 0xe1,
	0xdd, 0x11, 0x3f, 0xe5, 0xee, 0x96, 0xe9, 0x21, 0xbc, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x63,
	0x06, 0x13, 0x49, 0x52, 0x02, 0x00, 0x00,
}
