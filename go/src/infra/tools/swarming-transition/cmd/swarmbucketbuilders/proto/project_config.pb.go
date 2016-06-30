// Code generated by protoc-gen-go.
// source: project_config.proto
// DO NOT EDIT!

/*
Package buildbucket is a generated protocol buffer package.

It is generated from these files:
	project_config.proto

It has these top-level messages:
	Acl
	Swarming
	Bucket
	BuildbucketCfg
*/
package buildbucket

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Acl_Role int32

const (
	// Can do read-only operations, such as search for builds.
	Acl_READER Acl_Role = 0
	// Same as READER + can schedule and cancel builds.
	Acl_SCHEDULER Acl_Role = 1
	// Can do all write operations.
	Acl_WRITER Acl_Role = 2
)

var Acl_Role_name = map[int32]string{
	0: "READER",
	1: "SCHEDULER",
	2: "WRITER",
}
var Acl_Role_value = map[string]int32{
	"READER":    0,
	"SCHEDULER": 1,
	"WRITER":    2,
}

func (x Acl_Role) Enum() *Acl_Role {
	p := new(Acl_Role)
	*p = x
	return p
}
func (x Acl_Role) String() string {
	return proto.EnumName(Acl_Role_name, int32(x))
}
func (x *Acl_Role) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Acl_Role_value, data, "Acl_Role")
	if err != nil {
		return err
	}
	*x = Acl_Role(value)
	return nil
}
func (Acl_Role) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// A single access control rule.
type Acl struct {
	// Role denotes a list of actions that an identity can perform.
	Role *Acl_Role `protobuf:"varint,1,opt,name=role,enum=buildbucket.Acl_Role" json:"role,omitempty"`
	// Name of the group defined in the auth service.
	Group *string `protobuf:"bytes,2,opt,name=group" json:"group,omitempty"`
	// An email address or a full identity string "kind:name". See auth service
	// on kinds of identities. Anonymous users are "anonymous:anonymous".
	// Either identity or group must be present, not both.
	Identity         *string `protobuf:"bytes,3,opt,name=identity" json:"identity,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Acl) Reset()                    { *m = Acl{} }
func (m *Acl) String() string            { return proto.CompactTextString(m) }
func (*Acl) ProtoMessage()               {}
func (*Acl) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Acl) GetRole() Acl_Role {
	if m != nil && m.Role != nil {
		return *m.Role
	}
	return Acl_READER
}

func (m *Acl) GetGroup() string {
	if m != nil && m.Group != nil {
		return *m.Group
	}
	return ""
}

func (m *Acl) GetIdentity() string {
	if m != nil && m.Identity != nil {
		return *m.Identity
	}
	return ""
}

// Configuration of buildbucket-swarming integration for one bucket.
type Swarming struct {
	// Hostname of the swarming instance, e.g. "chromium-swarm.appspot.com".
	Hostname *string `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	// Used to generate a URL for Build, may contain parameters
	// {swarming_hostname}, {task_id}, {bucket} and {builder}. Defaults to:
	// https://{swarming_hostname}/user/task/{task_id}
	UrlFormat *string `protobuf:"bytes,2,opt,name=url_format,json=urlFormat" json:"url_format,omitempty"`
	// Will be put to all swarming tasks.
	CommonSwarmingTags []string `protobuf:"bytes,3,rep,name=common_swarming_tags,json=commonSwarmingTags" json:"common_swarming_tags,omitempty"`
	// Colon-delimited key-value pair of common task dimensions.
	CommonDimensions []string `protobuf:"bytes,4,rep,name=common_dimensions,json=commonDimensions" json:"common_dimensions,omitempty"`
	// Configuration for each builder.
	// Swarming tasks are created only for builds for builders that are not
	// explicitly specified.
	Builders []*Swarming_Builder `protobuf:"bytes,5,rep,name=builders" json:"builders,omitempty"`
	// Defines a default recipe for all builders in this bucket.
	// A builder may override it or parts of it with something else.
	CommonRecipe     *Swarming_Recipe `protobuf:"bytes,6,opt,name=common_recipe,json=commonRecipe" json:"common_recipe,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *Swarming) Reset()                    { *m = Swarming{} }
func (m *Swarming) String() string            { return proto.CompactTextString(m) }
func (*Swarming) ProtoMessage()               {}
func (*Swarming) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Swarming) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func (m *Swarming) GetUrlFormat() string {
	if m != nil && m.UrlFormat != nil {
		return *m.UrlFormat
	}
	return ""
}

func (m *Swarming) GetCommonSwarmingTags() []string {
	if m != nil {
		return m.CommonSwarmingTags
	}
	return nil
}

func (m *Swarming) GetCommonDimensions() []string {
	if m != nil {
		return m.CommonDimensions
	}
	return nil
}

func (m *Swarming) GetBuilders() []*Swarming_Builder {
	if m != nil {
		return m.Builders
	}
	return nil
}

func (m *Swarming) GetCommonRecipe() *Swarming_Recipe {
	if m != nil {
		return m.CommonRecipe
	}
	return nil
}

type Swarming_Recipe struct {
	// Repository URL of the recipe package.
	Repository *string `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	// Name of the recipe to run.
	Name *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// A list of colon-separated build property names and values.
	// A value has the following format:
	// -
	// A property can be overriden by "properties" build parameter.
	//
	// Use this field for string properties and use properties_j for other
	// types.
	Properties []string `protobuf:"bytes,3,rep,name=properties" json:"properties,omitempty"`
	// Same as properties, but the value must valid JSON. For example
	//   properties_j: "a:1"
	// means property a is a number 1, not string "1".
	//
	// Fields properties and properties_j can be used together, but cannot both
	// specify values for same property.
	PropertiesJ      []string `protobuf:"bytes,4,rep,name=properties_j,json=propertiesJ" json:"properties_j,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Swarming_Recipe) Reset()                    { *m = Swarming_Recipe{} }
func (m *Swarming_Recipe) String() string            { return proto.CompactTextString(m) }
func (*Swarming_Recipe) ProtoMessage()               {}
func (*Swarming_Recipe) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Swarming_Recipe) GetRepository() string {
	if m != nil && m.Repository != nil {
		return *m.Repository
	}
	return ""
}

func (m *Swarming_Recipe) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Swarming_Recipe) GetProperties() []string {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *Swarming_Recipe) GetPropertiesJ() []string {
	if m != nil {
		return m.PropertiesJ
	}
	return nil
}

type Swarming_Builder struct {
	// Name of the builder. Will be propagated to "builder" build tag and
	// "buildername" recipe property.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Builder category. Will be used for visual grouping, for example in Code Review.
	Category *string `protobuf:"bytes,6,opt,name=category" json:"category,omitempty"`
	// Will be become to swarming task tags.
	// Each tag will end up in "swarming_tag" buildbucket tag, for example
	// "swarming_tag:builder:release"
	SwarmingTags []string `protobuf:"bytes,2,rep,name=swarming_tags,json=swarmingTags" json:"swarming_tags,omitempty"`
	// Colon-delimited key-value pair of task dimensions.
	Dimensions []string `protobuf:"bytes,3,rep,name=dimensions" json:"dimensions,omitempty"`
	// Specifies that a recipe to run.
	Recipe *Swarming_Recipe `protobuf:"bytes,4,opt,name=recipe" json:"recipe,omitempty"`
	// Swarming task priority.
	Priority *int32 `protobuf:"varint,5,opt,name=priority" json:"priority,omitempty"`
	// Maximum build execution time. Not to be confused with pending time.
	// If not set, defaults to server defaults.
	ExecutionTimeoutSecs *int32 `protobuf:"varint,7,opt,name=execution_timeout_secs,json=executionTimeoutSecs" json:"execution_timeout_secs,omitempty"`
	XXX_unrecognized     []byte `json:"-"`
}

func (m *Swarming_Builder) Reset()                    { *m = Swarming_Builder{} }
func (m *Swarming_Builder) String() string            { return proto.CompactTextString(m) }
func (*Swarming_Builder) ProtoMessage()               {}
func (*Swarming_Builder) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 1} }

func (m *Swarming_Builder) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Swarming_Builder) GetCategory() string {
	if m != nil && m.Category != nil {
		return *m.Category
	}
	return ""
}

func (m *Swarming_Builder) GetSwarmingTags() []string {
	if m != nil {
		return m.SwarmingTags
	}
	return nil
}

func (m *Swarming_Builder) GetDimensions() []string {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *Swarming_Builder) GetRecipe() *Swarming_Recipe {
	if m != nil {
		return m.Recipe
	}
	return nil
}

func (m *Swarming_Builder) GetPriority() int32 {
	if m != nil && m.Priority != nil {
		return *m.Priority
	}
	return 0
}

func (m *Swarming_Builder) GetExecutionTimeoutSecs() int32 {
	if m != nil && m.ExecutionTimeoutSecs != nil {
		return *m.ExecutionTimeoutSecs
	}
	return 0
}

// Defines one bucket in buildbucket.cfg
type Bucket struct {
	// Name of the bucket. Names are unique within one instance of buildbucket.
	// If another project already uses this name, a config will be rejected.
	// Name reservation is first-come first-serve.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// List of access control rules for the bucket.
	Acls []*Acl `protobuf:"bytes,2,rep,name=acls" json:"acls,omitempty"`
	// Buildbucket-swarming integration.
	Swarming         *Swarming `protobuf:"bytes,3,opt,name=swarming" json:"swarming,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Bucket) Reset()                    { *m = Bucket{} }
func (m *Bucket) String() string            { return proto.CompactTextString(m) }
func (*Bucket) ProtoMessage()               {}
func (*Bucket) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Bucket) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Bucket) GetAcls() []*Acl {
	if m != nil {
		return m.Acls
	}
	return nil
}

func (m *Bucket) GetSwarming() *Swarming {
	if m != nil {
		return m.Swarming
	}
	return nil
}

// Schema of buildbucket.cfg file, a project config.
type BuildbucketCfg struct {
	// All buckets defined for this project.
	Buckets          []*Bucket `protobuf:"bytes,1,rep,name=buckets" json:"buckets,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *BuildbucketCfg) Reset()                    { *m = BuildbucketCfg{} }
func (m *BuildbucketCfg) String() string            { return proto.CompactTextString(m) }
func (*BuildbucketCfg) ProtoMessage()               {}
func (*BuildbucketCfg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BuildbucketCfg) GetBuckets() []*Bucket {
	if m != nil {
		return m.Buckets
	}
	return nil
}

func init() {
	proto.RegisterType((*Acl)(nil), "buildbucket.Acl")
	proto.RegisterType((*Swarming)(nil), "buildbucket.Swarming")
	proto.RegisterType((*Swarming_Recipe)(nil), "buildbucket.Swarming.Recipe")
	proto.RegisterType((*Swarming_Builder)(nil), "buildbucket.Swarming.Builder")
	proto.RegisterType((*Bucket)(nil), "buildbucket.Bucket")
	proto.RegisterType((*BuildbucketCfg)(nil), "buildbucket.BuildbucketCfg")
	proto.RegisterEnum("buildbucket.Acl_Role", Acl_Role_name, Acl_Role_value)
}

var fileDescriptor0 = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0xa5, 0xdd, 0xf4, 0x6b, 0xd2, 0xae, 0x8a, 0x29, 0x28, 0xaa, 0x58, 0x04, 0x85, 0x03, 0x08,
	0x6d, 0x05, 0x15, 0x17, 0x4e, 0xa8, 0xdd, 0x16, 0x01, 0xe2, 0xe4, 0x16, 0x71, 0x8c, 0xb2, 0xa9,
	0x1b, 0xbc, 0x34, 0x71, 0x64, 0x3b, 0x82, 0x3d, 0x71, 0xe3, 0xce, 0xff, 0xe0, 0x47, 0x32, 0xb1,
	0x93, 0x34, 0x5d, 0xed, 0x81, 0x5b, 0xde, 0xbc, 0x37, 0x33, 0xcf, 0xcf, 0x31, 0x8c, 0x52, 0x29,
	0xae, 0x58, 0xa8, 0xfd, 0x50, 0x24, 0x3b, 0x1e, 0x4d, 0x11, 0x6a, 0x41, 0xdc, 0xcb, 0x8c, 0xef,
	0xb7, 0x97, 0x59, 0xf8, 0x9d, 0xe9, 0xc9, 0x9f, 0x06, 0x9c, 0xcc, 0xc3, 0x3d, 0x79, 0x01, 0x8e,
	0x14, 0x7b, 0xe6, 0x35, 0x1e, 0x37, 0x9e, 0x9f, 0xce, 0xee, 0x4f, 0x6b, 0x9a, 0x29, 0xf2, 0x53,
	0x8a, 0x24, 0x35, 0x12, 0x32, 0x82, 0x56, 0x24, 0x45, 0x96, 0x7a, 0x4d, 0xd4, 0xf6, 0xa8, 0x05,
	0x64, 0x0c, 0x5d, 0xbe, 0x65, 0x89, 0xe6, 0xfa, 0xda, 0x3b, 0x31, 0x44, 0x85, 0x27, 0xe7, 0xe0,
	0xe4, 0xfd, 0x04, 0xa0, 0x4d, 0x57, 0xf3, 0xe5, 0x8a, 0x0e, 0xef, 0x90, 0x01, 0xf4, 0xd6, 0x17,
	0x1f, 0x56, 0xcb, 0x2f, 0x9f, 0x11, 0x36, 0x72, 0xea, 0x2b, 0xfd, 0xb8, 0xc1, 0xef, 0xe6, 0xe4,
	0x6f, 0x0b, 0xba, 0xeb, 0x1f, 0x81, 0x8c, 0x79, 0x12, 0xe5, 0x73, 0xbf, 0x09, 0xa5, 0x93, 0x20,
	0xb6, 0xe6, 0x70, 0x6e, 0x89, 0xc9, 0x19, 0x40, 0x26, 0xf7, 0xfe, 0x4e, 0xc8, 0x38, 0xd0, 0x85,
	0x9d, 0x1e, 0x56, 0xde, 0x9b, 0x02, 0x79, 0x05, 0xa3, 0x50, 0xc4, 0xb1, 0x48, 0x7c, 0x55, 0x4c,
	0xf3, 0x75, 0x10, 0x29, 0xb4, 0x77, 0x82, 0x42, 0x62, 0xb9, 0x72, 0xd1, 0x06, 0x19, 0xf2, 0x12,
	0xee, 0x16, 0x1d, 0x5b, 0x1e, 0xb3, 0x44, 0x71, 0x91, 0x28, 0xcf, 0x31, 0xf2, 0xa1, 0x25, 0x96,
	0x55, 0x9d, 0xbc, 0x85, 0xae, 0x49, 0x89, 0x49, 0xe5, 0xb5, 0x50, 0xe3, 0xce, 0xce, 0x8e, 0x62,
	0x2b, 0x27, 0x4f, 0x17, 0x56, 0x45, 0x2b, 0x39, 0x99, 0xc3, 0xa0, 0xd8, 0x23, 0x59, 0xc8, 0x53,
	0xe6, 0xb5, 0xd1, 0xbb, 0x3b, 0x7b, 0x78, 0x7b, 0x3f, 0x35, 0x1a, 0xda, 0xb7, 0x2d, 0x16, 0x8d,
	0x7f, 0x61, 0x96, 0xe6, 0x8b, 0x3c, 0x02, 0x90, 0x2c, 0x15, 0x8a, 0x6b, 0x21, 0xaf, 0x8b, 0x8c,
	0x6a, 0x15, 0x42, 0xc0, 0x31, 0xe9, 0xd9, 0x7c, 0xcc, 0x77, 0xde, 0x83, 0x3f, 0x43, 0xca, 0xa4,
	0xe6, 0xac, 0x0c, 0xa4, 0x56, 0x21, 0x4f, 0xa0, 0x7f, 0x40, 0xfe, 0x55, 0x91, 0x81, 0x7b, 0xa8,
	0x7d, 0x1a, 0xff, 0x6e, 0x42, 0xa7, 0x38, 0x59, 0xb5, 0xa2, 0x51, 0x5b, 0x81, 0x17, 0x17, 0x06,
	0x9a, 0x45, 0xb9, 0xa9, 0xb6, 0xbd, 0xb8, 0x12, 0x93, 0xa7, 0x30, 0x38, 0xbe, 0x92, 0xa6, 0x99,
	0xdf, 0x57, 0xf5, 0xcb, 0x40, 0x8f, 0xb5, 0x5b, 0x28, 0x3c, 0x1e, 0x2a, 0xe4, 0x0d, 0xb4, 0x8b,
	0xf4, 0x9c, 0xff, 0x48, 0xaf, 0xd0, 0xe6, 0xb6, 0x52, 0xc9, 0x85, 0xcc, 0xff, 0xd3, 0x16, 0xf6,
	0xb5, 0x68, 0x85, 0x71, 0xe2, 0x03, 0xf6, 0x93, 0x85, 0x99, 0xc6, 0xf9, 0xbe, 0xc6, 0x4d, 0x22,
	0xd3, 0xbe, 0x62, 0xa1, 0xf2, 0x3a, 0x46, 0x39, 0xaa, 0xd8, 0x8d, 0x25, 0xd7, 0xc8, 0x4d, 0x32,
	0x68, 0x2f, 0xcc, 0xce, 0x5b, 0x63, 0x78, 0x06, 0x4e, 0x10, 0xee, 0xed, 0x09, 0xdd, 0xd9, 0xf0,
	0xe6, 0xc3, 0xa2, 0x86, 0x25, 0xaf, 0xa1, 0x5b, 0x9e, 0xdd, 0xbc, 0x1e, 0xf7, 0xc6, 0x13, 0x2c,
	0x4f, 0x43, 0x2b, 0xd9, 0xe4, 0x1d, 0x9c, 0x2e, 0x0e, 0x8a, 0x8b, 0x5d, 0x44, 0xce, 0xa1, 0x63,
	0x81, 0x42, 0x07, 0xf9, 0xb6, 0x7b, 0x47, 0x33, 0xac, 0x49, 0x5a, 0x6a, 0xfe, 0x05, 0x00, 0x00,
	0xff, 0xff, 0x39, 0x3d, 0xb9, 0x48, 0x1e, 0x04, 0x00, 0x00,
}
