// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infra/appengine/crosskylabadmin/api/fleet/v1/tracker.proto

package fleet

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// DutState specifies the valid values for DUT state.
type DutState int32

const (
	DutState_DutStateInvalid DutState = 0
	DutState_Ready           DutState = 1
	DutState_NeedsCleanup    DutState = 2
	DutState_NeedsRepair     DutState = 3
	DutState_NeedsReset      DutState = 4
	DutState_RepairFailed    DutState = 5
)

var DutState_name = map[int32]string{
	0: "DutStateInvalid",
	1: "Ready",
	2: "NeedsCleanup",
	3: "NeedsRepair",
	4: "NeedsReset",
	5: "RepairFailed",
}

var DutState_value = map[string]int32{
	"DutStateInvalid": 0,
	"Ready":           1,
	"NeedsCleanup":    2,
	"NeedsRepair":     3,
	"NeedsReset":      4,
	"RepairFailed":    5,
}

func (x DutState) String() string {
	return proto.EnumName(DutState_name, int32(x))
}

func (DutState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_474af594abe23e82, []int{0}
}

type Health int32

const (
	Health_HealthInvalid Health = 0
	// A Healthy bot may be used for external workload.
	Health_Healthy Health = 1
	// An Unhealthy bot is not usable for external workload.
	// Further classification of the problem is not available.
	Health_Unhealthy Health = 2
)

var Health_name = map[int32]string{
	0: "HealthInvalid",
	1: "Healthy",
	2: "Unhealthy",
}

var Health_value = map[string]int32{
	"HealthInvalid": 0,
	"Healthy":       1,
	"Unhealthy":     2,
}

func (x Health) String() string {
	return proto.EnumName(Health_name, int32(x))
}

func (Health) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_474af594abe23e82, []int{1}
}

// RefreshBotsRequest can be used to restrict the Swarming bots to refresh via
// the Tracker.RefreshBots rpc.
type RefreshBotsRequest struct {
	// selectors whitelists the bots to refresh. This includes new bots
	// discovered from Swarming matching the selectors.
	// Bots selected via repeated selectors are unioned together.
	//
	// If no selectors are provided, all bots are selected.
	Selectors            []*BotSelector `protobuf:"bytes,2,rep,name=selectors,proto3" json:"selectors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RefreshBotsRequest) Reset()         { *m = RefreshBotsRequest{} }
func (m *RefreshBotsRequest) String() string { return proto.CompactTextString(m) }
func (*RefreshBotsRequest) ProtoMessage()    {}
func (*RefreshBotsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_474af594abe23e82, []int{0}
}

func (m *RefreshBotsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshBotsRequest.Unmarshal(m, b)
}
func (m *RefreshBotsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshBotsRequest.Marshal(b, m, deterministic)
}
func (m *RefreshBotsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshBotsRequest.Merge(m, src)
}
func (m *RefreshBotsRequest) XXX_Size() int {
	return xxx_messageInfo_RefreshBotsRequest.Size(m)
}
func (m *RefreshBotsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshBotsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshBotsRequest proto.InternalMessageInfo

func (m *RefreshBotsRequest) GetSelectors() []*BotSelector {
	if m != nil {
		return m.Selectors
	}
	return nil
}

// RefreshBotsResponse contains information about the Swarming bots actually
// refreshed in response to a Tracker.RefreshBots rpc.
type RefreshBotsResponse struct {
	// dut_ids lists the dut_id of of the bots refreshed.
	DutIds               []string `protobuf:"bytes,1,rep,name=dut_ids,json=dutIds,proto3" json:"dut_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshBotsResponse) Reset()         { *m = RefreshBotsResponse{} }
func (m *RefreshBotsResponse) String() string { return proto.CompactTextString(m) }
func (*RefreshBotsResponse) ProtoMessage()    {}
func (*RefreshBotsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_474af594abe23e82, []int{1}
}

func (m *RefreshBotsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshBotsResponse.Unmarshal(m, b)
}
func (m *RefreshBotsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshBotsResponse.Marshal(b, m, deterministic)
}
func (m *RefreshBotsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshBotsResponse.Merge(m, src)
}
func (m *RefreshBotsResponse) XXX_Size() int {
	return xxx_messageInfo_RefreshBotsResponse.Size(m)
}
func (m *RefreshBotsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshBotsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshBotsResponse proto.InternalMessageInfo

func (m *RefreshBotsResponse) GetDutIds() []string {
	if m != nil {
		return m.DutIds
	}
	return nil
}

// SummarizeBotsRequest can be used to restrict the Swarming bots to summarize
// via the Tracker.SummarizeBots rpc.
type SummarizeBotsRequest struct {
	// selectors whitelists the bots to refresh, from the already known bots to
	// Tracker. Bots selected via repeated selectors are unioned together.
	//
	// If no selectors are provided, all bots are selected.
	Selectors            []*BotSelector `protobuf:"bytes,1,rep,name=selectors,proto3" json:"selectors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SummarizeBotsRequest) Reset()         { *m = SummarizeBotsRequest{} }
func (m *SummarizeBotsRequest) String() string { return proto.CompactTextString(m) }
func (*SummarizeBotsRequest) ProtoMessage()    {}
func (*SummarizeBotsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_474af594abe23e82, []int{2}
}

func (m *SummarizeBotsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SummarizeBotsRequest.Unmarshal(m, b)
}
func (m *SummarizeBotsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SummarizeBotsRequest.Marshal(b, m, deterministic)
}
func (m *SummarizeBotsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SummarizeBotsRequest.Merge(m, src)
}
func (m *SummarizeBotsRequest) XXX_Size() int {
	return xxx_messageInfo_SummarizeBotsRequest.Size(m)
}
func (m *SummarizeBotsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SummarizeBotsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SummarizeBotsRequest proto.InternalMessageInfo

func (m *SummarizeBotsRequest) GetSelectors() []*BotSelector {
	if m != nil {
		return m.Selectors
	}
	return nil
}

// SummarizeBotsResponse contains summary information about Swarming bots
// returned by the Tracker.SummarizeBots rpc.
type SummarizeBotsResponse struct {
	Bots                 []*BotSummary `protobuf:"bytes,1,rep,name=bots,proto3" json:"bots,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SummarizeBotsResponse) Reset()         { *m = SummarizeBotsResponse{} }
func (m *SummarizeBotsResponse) String() string { return proto.CompactTextString(m) }
func (*SummarizeBotsResponse) ProtoMessage()    {}
func (*SummarizeBotsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_474af594abe23e82, []int{3}
}

func (m *SummarizeBotsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SummarizeBotsResponse.Unmarshal(m, b)
}
func (m *SummarizeBotsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SummarizeBotsResponse.Marshal(b, m, deterministic)
}
func (m *SummarizeBotsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SummarizeBotsResponse.Merge(m, src)
}
func (m *SummarizeBotsResponse) XXX_Size() int {
	return xxx_messageInfo_SummarizeBotsResponse.Size(m)
}
func (m *SummarizeBotsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SummarizeBotsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SummarizeBotsResponse proto.InternalMessageInfo

func (m *SummarizeBotsResponse) GetBots() []*BotSummary {
	if m != nil {
		return m.Bots
	}
	return nil
}

// BotSummary contains the summary information tracked by Tracker for a single
// Skylab Swarming bot.
type BotSummary struct {
	// dut_id contains the dut_id dimension for the bot.
	DutId string `protobuf:"bytes,1,opt,name=dut_id,json=dutId,proto3" json:"dut_id,omitempty"`
	// dut_state contains the current Autotest state of the dut corresponding to
	// this bot.
	DutState DutState `protobuf:"varint,2,opt,name=dut_state,json=dutState,proto3,enum=crosskylabadmin.fleet.DutState" json:"dut_state,omitempty"`
	// idle_duration contains the time since this bot last ran a task.
	//
	// A bot is considered idle for the time that it wasn't running any task.
	// Killed tasks are counted as legitimate tasks (i.e., time spent running a
	// task that is then killed does not count as idle time)
	IdleDuration *duration.Duration `protobuf:"bytes,3,opt,name=idle_duration,json=idleDuration,proto3" json:"idle_duration,omitempty"`
	// Subset of Swarming dimensions for the current bot.
	Dimensions *BotDimensions `protobuf:"bytes,4,opt,name=dimensions,proto3" json:"dimensions,omitempty"`
	// health is the history aware health of the bot.
	//
	// A healthy bot is safe to use for external workload. For unhealthy bots,
	// this field summarizes the reason for the unhealthy state of the bot.
	Health Health `protobuf:"varint,5,opt,name=health,proto3,enum=crosskylabadmin.fleet.Health" json:"health,omitempty"`
	// diagnosis contains the tasks that explain how the DUT got into
	// its present state.
	Diagnosis            []*Task  `protobuf:"bytes,6,rep,name=diagnosis,proto3" json:"diagnosis,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BotSummary) Reset()         { *m = BotSummary{} }
func (m *BotSummary) String() string { return proto.CompactTextString(m) }
func (*BotSummary) ProtoMessage()    {}
func (*BotSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_474af594abe23e82, []int{4}
}

func (m *BotSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BotSummary.Unmarshal(m, b)
}
func (m *BotSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BotSummary.Marshal(b, m, deterministic)
}
func (m *BotSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BotSummary.Merge(m, src)
}
func (m *BotSummary) XXX_Size() int {
	return xxx_messageInfo_BotSummary.Size(m)
}
func (m *BotSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_BotSummary.DiscardUnknown(m)
}

var xxx_messageInfo_BotSummary proto.InternalMessageInfo

func (m *BotSummary) GetDutId() string {
	if m != nil {
		return m.DutId
	}
	return ""
}

func (m *BotSummary) GetDutState() DutState {
	if m != nil {
		return m.DutState
	}
	return DutState_DutStateInvalid
}

func (m *BotSummary) GetIdleDuration() *duration.Duration {
	if m != nil {
		return m.IdleDuration
	}
	return nil
}

func (m *BotSummary) GetDimensions() *BotDimensions {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *BotSummary) GetHealth() Health {
	if m != nil {
		return m.Health
	}
	return Health_HealthInvalid
}

func (m *BotSummary) GetDiagnosis() []*Task {
	if m != nil {
		return m.Diagnosis
	}
	return nil
}

// Task contains information about a Swarming task.
type Task struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	StateBefore          DutState             `protobuf:"varint,3,opt,name=state_before,json=stateBefore,proto3,enum=crosskylabadmin.fleet.DutState" json:"state_before,omitempty"`
	StateAfter           DutState             `protobuf:"varint,4,opt,name=state_after,json=stateAfter,proto3,enum=crosskylabadmin.fleet.DutState" json:"state_after,omitempty"`
	StartedTs            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=started_ts,json=startedTs,proto3" json:"started_ts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_474af594abe23e82, []int{5}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Task) GetStateBefore() DutState {
	if m != nil {
		return m.StateBefore
	}
	return DutState_DutStateInvalid
}

func (m *Task) GetStateAfter() DutState {
	if m != nil {
		return m.StateAfter
	}
	return DutState_DutStateInvalid
}

func (m *Task) GetStartedTs() *timestamp.Timestamp {
	if m != nil {
		return m.StartedTs
	}
	return nil
}

func init() {
	proto.RegisterEnum("crosskylabadmin.fleet.DutState", DutState_name, DutState_value)
	proto.RegisterEnum("crosskylabadmin.fleet.Health", Health_name, Health_value)
	proto.RegisterType((*RefreshBotsRequest)(nil), "crosskylabadmin.fleet.RefreshBotsRequest")
	proto.RegisterType((*RefreshBotsResponse)(nil), "crosskylabadmin.fleet.RefreshBotsResponse")
	proto.RegisterType((*SummarizeBotsRequest)(nil), "crosskylabadmin.fleet.SummarizeBotsRequest")
	proto.RegisterType((*SummarizeBotsResponse)(nil), "crosskylabadmin.fleet.SummarizeBotsResponse")
	proto.RegisterType((*BotSummary)(nil), "crosskylabadmin.fleet.BotSummary")
	proto.RegisterType((*Task)(nil), "crosskylabadmin.fleet.Task")
}

func init() {
	proto.RegisterFile("infra/appengine/crosskylabadmin/api/fleet/v1/tracker.proto", fileDescriptor_474af594abe23e82)
}

var fileDescriptor_474af594abe23e82 = []byte{
	// 656 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xdf, 0x6f, 0xd3, 0x3a,
	0x14, 0xc7, 0x6f, 0xfa, 0x73, 0x39, 0x5d, 0xb7, 0x5e, 0xef, 0x4e, 0x37, 0xb7, 0x57, 0xb0, 0x52,
	0xf1, 0x50, 0x06, 0x4a, 0x44, 0xd1, 0x84, 0x86, 0x10, 0x1a, 0xa5, 0x42, 0xec, 0x65, 0x0f, 0x5e,
	0x41, 0x88, 0x97, 0xca, 0xad, 0x4f, 0x5b, 0xb3, 0x24, 0x0e, 0xb1, 0x33, 0x69, 0xfc, 0xa7, 0x88,
	0x3f, 0x81, 0x7f, 0x02, 0xc5, 0x49, 0xd6, 0xfd, 0x2a, 0x2a, 0xbc, 0xd9, 0xc7, 0x9f, 0xef, 0xf1,
	0xf1, 0xd7, 0x3e, 0x86, 0x17, 0x22, 0x9c, 0xc5, 0xcc, 0x63, 0x51, 0x84, 0xe1, 0x5c, 0x84, 0xe8,
	0x4d, 0x63, 0xa9, 0xd4, 0xd9, 0x85, 0xcf, 0x26, 0x8c, 0x07, 0x22, 0xf4, 0x58, 0x24, 0xbc, 0x99,
	0x8f, 0xa8, 0xbd, 0xf3, 0xa7, 0x9e, 0x8e, 0xd9, 0xf4, 0x0c, 0x63, 0x37, 0x8a, 0xa5, 0x96, 0x64,
	0xf7, 0x06, 0xeb, 0x1a, 0xae, 0x7d, 0x7f, 0x2e, 0xe5, 0xdc, 0x47, 0xcf, 0x40, 0x93, 0x64, 0xe6,
	0xf1, 0x24, 0x66, 0x5a, 0xc8, 0x30, 0x93, 0xb5, 0xf7, 0x6e, 0xae, 0x6b, 0x11, 0xa0, 0xd2, 0x2c,
	0x88, 0x72, 0xe0, 0xf0, 0xb7, 0x6a, 0x9a, 0xca, 0x20, 0x28, 0x72, 0x77, 0x3f, 0x00, 0xa1, 0x38,
	0x8b, 0x51, 0x2d, 0x06, 0x52, 0x2b, 0x8a, 0x5f, 0x12, 0x54, 0x9a, 0x1c, 0x81, 0xad, 0xd0, 0xc7,
	0xa9, 0x96, 0xb1, 0x72, 0x4a, 0x9d, 0x72, 0xaf, 0xd1, 0xef, 0xba, 0x77, 0x16, 0xef, 0x0e, 0xa4,
	0x3e, 0xcd, 0x51, 0xba, 0x14, 0x75, 0x5d, 0xd8, 0xb9, 0x96, 0x57, 0x45, 0x32, 0x54, 0x48, 0xfe,
	0x85, 0x3a, 0x4f, 0xf4, 0x58, 0x70, 0xe5, 0x58, 0x9d, 0x72, 0xcf, 0xa6, 0x35, 0x9e, 0xe8, 0x63,
	0xae, 0xba, 0x1f, 0xe1, 0x9f, 0xd3, 0x24, 0x08, 0x58, 0x2c, 0xbe, 0xe2, 0xca, 0x4a, 0xac, 0x3f,
	0xa9, 0xe4, 0x04, 0x76, 0x6f, 0x64, 0xce, 0x6b, 0x39, 0x80, 0xca, 0x44, 0xea, 0x22, 0xeb, 0x83,
	0x5f, 0x64, 0x35, 0xf2, 0x0b, 0x6a, 0xf0, 0xee, 0xb7, 0x12, 0xc0, 0x32, 0x48, 0x76, 0xa1, 0x96,
	0x9d, 0xc8, 0xb1, 0x3a, 0x56, 0xcf, 0xa6, 0x55, 0x73, 0x20, 0xf2, 0x12, 0xec, 0x34, 0xac, 0x34,
	0xd3, 0xe8, 0x94, 0x3a, 0x56, 0x6f, 0xab, 0xbf, 0xb7, 0x62, 0x87, 0x61, 0xa2, 0x4f, 0x53, 0x8c,
	0x6e, 0xf0, 0x7c, 0x44, 0x5e, 0x41, 0x53, 0x70, 0x1f, 0xc7, 0xc5, 0x43, 0x70, 0xca, 0x1d, 0xab,
	0xd7, 0xe8, 0xff, 0xe7, 0x66, 0x2f, 0xc1, 0x2d, 0x5e, 0x82, 0x3b, 0xcc, 0x01, 0xba, 0x99, 0xf2,
	0xc5, 0x8c, 0x0c, 0x01, 0xb8, 0x08, 0x30, 0x54, 0x42, 0x86, 0xca, 0xa9, 0x18, 0xf1, 0xc3, 0xd5,
	0x07, 0x1c, 0x5e, 0xb2, 0xf4, 0x8a, 0x8e, 0x1c, 0x40, 0x6d, 0x81, 0xcc, 0xd7, 0x0b, 0xa7, 0x6a,
	0x0e, 0x70, 0x6f, 0x45, 0x86, 0x77, 0x06, 0xa2, 0x39, 0x4c, 0x0e, 0xc1, 0xe6, 0x82, 0xcd, 0x43,
	0xa9, 0x84, 0x72, 0x6a, 0xc6, 0xdc, 0xff, 0x57, 0x28, 0x47, 0x4c, 0x9d, 0xd1, 0x25, 0xdd, 0xfd,
	0x61, 0x41, 0x25, 0x8d, 0x91, 0x2d, 0x28, 0x5d, 0x3a, 0x5a, 0x12, 0x9c, 0x10, 0xa8, 0x84, 0x2c,
	0xc8, 0x9c, 0xb4, 0xa9, 0x19, 0x93, 0x01, 0x6c, 0x1a, 0x7b, 0xc7, 0x13, 0x9c, 0xc9, 0x18, 0x8d,
	0x47, 0x6b, 0xb8, 0xdc, 0x30, 0xa2, 0x81, 0xd1, 0x90, 0x23, 0xc8, 0xa6, 0x63, 0x36, 0xd3, 0x18,
	0x1b, 0xa7, 0xd6, 0x48, 0x01, 0x46, 0xf3, 0x3a, 0x95, 0x90, 0x43, 0x48, 0x67, 0xb1, 0x46, 0x3e,
	0xd6, 0xca, 0x18, 0xd5, 0xe8, 0xb7, 0x6f, 0xdd, 0xd3, 0xa8, 0xe8, 0x58, 0x6a, 0xe7, 0xf4, 0x48,
	0xed, 0x4b, 0xd8, 0x28, 0x52, 0x92, 0x1d, 0xd8, 0x2e, 0xc6, 0xc7, 0xe1, 0x39, 0xf3, 0x05, 0x6f,
	0xfd, 0x45, 0x6c, 0xa8, 0x52, 0x64, 0xfc, 0xa2, 0x65, 0x91, 0x16, 0x6c, 0x9e, 0x20, 0x72, 0xf5,
	0xc6, 0x47, 0x16, 0x26, 0x51, 0xab, 0x44, 0xb6, 0xa1, 0x61, 0x22, 0x14, 0x23, 0x26, 0xe2, 0x56,
	0x99, 0x6c, 0x01, 0xe4, 0x01, 0x85, 0xba, 0x55, 0x49, 0x25, 0xd9, 0xda, 0x5b, 0x26, 0x7c, 0xe4,
	0xad, 0xea, 0xfe, 0x73, 0xa8, 0x65, 0x77, 0x45, 0xfe, 0x86, 0x66, 0x36, 0x5a, 0x6e, 0xd6, 0x80,
	0x7a, 0x16, 0x4a, 0xb7, 0x6b, 0x82, 0xfd, 0x3e, 0x5c, 0xe4, 0xd3, 0x52, 0xff, 0xbb, 0x05, 0xf5,
	0x51, 0xf6, 0x95, 0x11, 0x0e, 0x8d, 0x2b, 0x9d, 0x4d, 0x1e, 0xad, 0x30, 0xeb, 0xf6, 0xaf, 0xd2,
	0xde, 0x5f, 0x07, 0xcd, 0x9b, 0xf3, 0x33, 0x34, 0xaf, 0x75, 0x2d, 0x79, 0xbc, 0x42, 0x7c, 0xd7,
	0xaf, 0xd1, 0x7e, 0xb2, 0x1e, 0x9c, 0xed, 0x35, 0xa8, 0x7f, 0xaa, 0x9a, 0xe5, 0x49, 0xcd, 0xdc,
	0xd7, 0xb3, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x58, 0xf5, 0xf4, 0xe4, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TrackerClient is the client API for Tracker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TrackerClient interface {
	// RefreshBots instructs the Tracker service to update Swarming bot
	// information from the Swarming server hosting ChromeOS Skylab bots.
	//
	// RefreshBots stops at the first error encountered and returns the error. A
	// failed RefreshBots call may have refreshed some of the bots requested.
	// It is safe to call RefreshBots to continue from a partially failed call.
	RefreshBots(ctx context.Context, in *RefreshBotsRequest, opts ...grpc.CallOption) (*RefreshBotsResponse, error)
	// SummarizeBots returns summary information about Swarming bots.
	// This includes ChromeOS Skylab specific dimensions/state information as
	// well as a summary of the recenty history of administrative tasks.
	//
	// SummarizeBots stops at the first error encountered and returns the error.
	SummarizeBots(ctx context.Context, in *SummarizeBotsRequest, opts ...grpc.CallOption) (*SummarizeBotsResponse, error)
}
type trackerPRPCClient struct {
	client *prpc.Client
}

func NewTrackerPRPCClient(client *prpc.Client) TrackerClient {
	return &trackerPRPCClient{client}
}

func (c *trackerPRPCClient) RefreshBots(ctx context.Context, in *RefreshBotsRequest, opts ...grpc.CallOption) (*RefreshBotsResponse, error) {
	out := new(RefreshBotsResponse)
	err := c.client.Call(ctx, "crosskylabadmin.fleet.Tracker", "RefreshBots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackerPRPCClient) SummarizeBots(ctx context.Context, in *SummarizeBotsRequest, opts ...grpc.CallOption) (*SummarizeBotsResponse, error) {
	out := new(SummarizeBotsResponse)
	err := c.client.Call(ctx, "crosskylabadmin.fleet.Tracker", "SummarizeBots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type trackerClient struct {
	cc *grpc.ClientConn
}

func NewTrackerClient(cc *grpc.ClientConn) TrackerClient {
	return &trackerClient{cc}
}

func (c *trackerClient) RefreshBots(ctx context.Context, in *RefreshBotsRequest, opts ...grpc.CallOption) (*RefreshBotsResponse, error) {
	out := new(RefreshBotsResponse)
	err := c.cc.Invoke(ctx, "/crosskylabadmin.fleet.Tracker/RefreshBots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackerClient) SummarizeBots(ctx context.Context, in *SummarizeBotsRequest, opts ...grpc.CallOption) (*SummarizeBotsResponse, error) {
	out := new(SummarizeBotsResponse)
	err := c.cc.Invoke(ctx, "/crosskylabadmin.fleet.Tracker/SummarizeBots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrackerServer is the server API for Tracker service.
type TrackerServer interface {
	// RefreshBots instructs the Tracker service to update Swarming bot
	// information from the Swarming server hosting ChromeOS Skylab bots.
	//
	// RefreshBots stops at the first error encountered and returns the error. A
	// failed RefreshBots call may have refreshed some of the bots requested.
	// It is safe to call RefreshBots to continue from a partially failed call.
	RefreshBots(context.Context, *RefreshBotsRequest) (*RefreshBotsResponse, error)
	// SummarizeBots returns summary information about Swarming bots.
	// This includes ChromeOS Skylab specific dimensions/state information as
	// well as a summary of the recenty history of administrative tasks.
	//
	// SummarizeBots stops at the first error encountered and returns the error.
	SummarizeBots(context.Context, *SummarizeBotsRequest) (*SummarizeBotsResponse, error)
}

func RegisterTrackerServer(s prpc.Registrar, srv TrackerServer) {
	s.RegisterService(&_Tracker_serviceDesc, srv)
}

func _Tracker_RefreshBots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshBotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackerServer).RefreshBots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crosskylabadmin.fleet.Tracker/RefreshBots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackerServer).RefreshBots(ctx, req.(*RefreshBotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tracker_SummarizeBots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SummarizeBotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackerServer).SummarizeBots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crosskylabadmin.fleet.Tracker/SummarizeBots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackerServer).SummarizeBots(ctx, req.(*SummarizeBotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tracker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "crosskylabadmin.fleet.Tracker",
	HandlerType: (*TrackerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RefreshBots",
			Handler:    _Tracker_RefreshBots_Handler,
		},
		{
			MethodName: "SummarizeBots",
			Handler:    _Tracker_SummarizeBots_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "infra/appengine/crosskylabadmin/api/fleet/v1/tracker.proto",
}
