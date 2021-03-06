// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/cfgmgmt/request/nodes.proto

package request

import (
	fmt "fmt"
	query "github.com/chef/automate/api/external/common/query"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

type Nodes struct {
	// Filters to apply to the request for nodes list.
	Filter []string `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty"`
	// Pagination parameters to apply to the returned node list.
	Pagination *query.Pagination `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	// Sorting parameters to apply to the returned node list.
	Sorting *query.Sorting `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting,omitempty"`
	// Earliest most recent check-in node information to return.
	Start string `protobuf:"bytes,4,opt,name=start,proto3" json:"start,omitempty"`
	// Latest most recent check-in node information to return.
	End                  string   `protobuf:"bytes,5,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nodes) Reset()         { *m = Nodes{} }
func (m *Nodes) String() string { return proto.CompactTextString(m) }
func (*Nodes) ProtoMessage()    {}
func (*Nodes) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeeb0a9b2be1cdd5, []int{0}
}

func (m *Nodes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nodes.Unmarshal(m, b)
}
func (m *Nodes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nodes.Marshal(b, m, deterministic)
}
func (m *Nodes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nodes.Merge(m, src)
}
func (m *Nodes) XXX_Size() int {
	return xxx_messageInfo_Nodes.Size(m)
}
func (m *Nodes) XXX_DiscardUnknown() {
	xxx_messageInfo_Nodes.DiscardUnknown(m)
}

var xxx_messageInfo_Nodes proto.InternalMessageInfo

func (m *Nodes) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *Nodes) GetPagination() *query.Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func (m *Nodes) GetSorting() *query.Sorting {
	if m != nil {
		return m.Sorting
	}
	return nil
}

func (m *Nodes) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *Nodes) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

type Node struct {
	// Chef guid for the requested node.
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeeb0a9b2be1cdd5, []int{1}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type NodeRun struct {
	// Chef guid for the requested node.
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Run id for the node.
	RunId string `protobuf:"bytes,2,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	// End time on the node's run.
	EndTime              *timestamp.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *NodeRun) Reset()         { *m = NodeRun{} }
func (m *NodeRun) String() string { return proto.CompactTextString(m) }
func (*NodeRun) ProtoMessage()    {}
func (*NodeRun) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeeb0a9b2be1cdd5, []int{2}
}

func (m *NodeRun) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeRun.Unmarshal(m, b)
}
func (m *NodeRun) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeRun.Marshal(b, m, deterministic)
}
func (m *NodeRun) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeRun.Merge(m, src)
}
func (m *NodeRun) XXX_Size() int {
	return xxx_messageInfo_NodeRun.Size(m)
}
func (m *NodeRun) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeRun.DiscardUnknown(m)
}

var xxx_messageInfo_NodeRun proto.InternalMessageInfo

func (m *NodeRun) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *NodeRun) GetRunId() string {
	if m != nil {
		return m.RunId
	}
	return ""
}

func (m *NodeRun) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

type Runs struct {
	// Chef guid for the node.
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Filters to apply to the request for runs list.
	Filter []string `protobuf:"bytes,2,rep,name=filter,proto3" json:"filter,omitempty"`
	// Pagination parameters to apply to the returned runs list.
	Pagination *query.Pagination `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
	// Earliest (in history) run information to return for the runs list.
	Start string `protobuf:"bytes,4,opt,name=start,proto3" json:"start,omitempty"`
	// Latest (in history) run information to return for the runs list.
	End                  string   `protobuf:"bytes,5,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Runs) Reset()         { *m = Runs{} }
func (m *Runs) String() string { return proto.CompactTextString(m) }
func (*Runs) ProtoMessage()    {}
func (*Runs) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeeb0a9b2be1cdd5, []int{3}
}

func (m *Runs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Runs.Unmarshal(m, b)
}
func (m *Runs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Runs.Marshal(b, m, deterministic)
}
func (m *Runs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Runs.Merge(m, src)
}
func (m *Runs) XXX_Size() int {
	return xxx_messageInfo_Runs.Size(m)
}
func (m *Runs) XXX_DiscardUnknown() {
	xxx_messageInfo_Runs.DiscardUnknown(m)
}

var xxx_messageInfo_Runs proto.InternalMessageInfo

func (m *Runs) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *Runs) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *Runs) GetPagination() *query.Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func (m *Runs) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *Runs) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

type NodeExport struct {
	// Filters to apply to the request for nodes list.
	Filter []string `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty"`
	// Sorting parameters to apply to the returned node list.
	Sorting *query.Sorting `protobuf:"bytes,2,opt,name=sorting,proto3" json:"sorting,omitempty"`
	// File type, either JSON or CSV.
	OutputType           string   `protobuf:"bytes,3,opt,name=output_type,json=outputType,proto3" json:"output_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeExport) Reset()         { *m = NodeExport{} }
func (m *NodeExport) String() string { return proto.CompactTextString(m) }
func (*NodeExport) ProtoMessage()    {}
func (*NodeExport) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeeb0a9b2be1cdd5, []int{4}
}

func (m *NodeExport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExport.Unmarshal(m, b)
}
func (m *NodeExport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExport.Marshal(b, m, deterministic)
}
func (m *NodeExport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExport.Merge(m, src)
}
func (m *NodeExport) XXX_Size() int {
	return xxx_messageInfo_NodeExport.Size(m)
}
func (m *NodeExport) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExport.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExport proto.InternalMessageInfo

func (m *NodeExport) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *NodeExport) GetSorting() *query.Sorting {
	if m != nil {
		return m.Sorting
	}
	return nil
}

func (m *NodeExport) GetOutputType() string {
	if m != nil {
		return m.OutputType
	}
	return ""
}

type ReportExport struct {
	// Filters to apply to the request for node report list.
	Filter []string `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty"`
	// File type, either JSON or CSV.
	OutputType string `protobuf:"bytes,2,opt,name=output_type,json=outputType,proto3" json:"output_type,omitempty"`
	// The node ID of the reports
	NodeId string `protobuf:"bytes,3,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Earliest (in history) run information to return for the runs list.
	Start *timestamp.Timestamp `protobuf:"bytes,4,opt,name=start,proto3" json:"start,omitempty"`
	// Latest (in history) run information to return for the runs list.
	End                  *timestamp.Timestamp `protobuf:"bytes,5,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ReportExport) Reset()         { *m = ReportExport{} }
func (m *ReportExport) String() string { return proto.CompactTextString(m) }
func (*ReportExport) ProtoMessage()    {}
func (*ReportExport) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeeb0a9b2be1cdd5, []int{5}
}

func (m *ReportExport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportExport.Unmarshal(m, b)
}
func (m *ReportExport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportExport.Marshal(b, m, deterministic)
}
func (m *ReportExport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportExport.Merge(m, src)
}
func (m *ReportExport) XXX_Size() int {
	return xxx_messageInfo_ReportExport.Size(m)
}
func (m *ReportExport) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportExport.DiscardUnknown(m)
}

var xxx_messageInfo_ReportExport proto.InternalMessageInfo

func (m *ReportExport) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ReportExport) GetOutputType() string {
	if m != nil {
		return m.OutputType
	}
	return ""
}

func (m *ReportExport) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *ReportExport) GetStart() *timestamp.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *ReportExport) GetEnd() *timestamp.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

type MissingNodeDurationCounts struct {
	// A valid duration is any number zero or greater with one of these characters 'h', 'd', 'w', or 'M'.
	// 'h' is hours
	// 'd' is days
	// 'w' is weeks
	// 'M' is months
	// Will contain one or many.
	Durations            []string `protobuf:"bytes,1,rep,name=durations,proto3" json:"durations,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MissingNodeDurationCounts) Reset()         { *m = MissingNodeDurationCounts{} }
func (m *MissingNodeDurationCounts) String() string { return proto.CompactTextString(m) }
func (*MissingNodeDurationCounts) ProtoMessage()    {}
func (*MissingNodeDurationCounts) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeeb0a9b2be1cdd5, []int{6}
}

func (m *MissingNodeDurationCounts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MissingNodeDurationCounts.Unmarshal(m, b)
}
func (m *MissingNodeDurationCounts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MissingNodeDurationCounts.Marshal(b, m, deterministic)
}
func (m *MissingNodeDurationCounts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MissingNodeDurationCounts.Merge(m, src)
}
func (m *MissingNodeDurationCounts) XXX_Size() int {
	return xxx_messageInfo_MissingNodeDurationCounts.Size(m)
}
func (m *MissingNodeDurationCounts) XXX_DiscardUnknown() {
	xxx_messageInfo_MissingNodeDurationCounts.DiscardUnknown(m)
}

var xxx_messageInfo_MissingNodeDurationCounts proto.InternalMessageInfo

func (m *MissingNodeDurationCounts) GetDurations() []string {
	if m != nil {
		return m.Durations
	}
	return nil
}

type NodeMetadataCounts struct {
	// Types of node fields to collect value counts for
	Type []string `protobuf:"bytes,1,rep,name=type,proto3" json:"type,omitempty"`
	// Filters to apply to the counts returned
	Filter []string `protobuf:"bytes,2,rep,name=filter,proto3" json:"filter,omitempty"`
	// Earliest most recent check-in node information to return.
	Start string `protobuf:"bytes,3,opt,name=start,proto3" json:"start,omitempty"`
	// Latest most recent check-in node information to return.
	End                  string   `protobuf:"bytes,4,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeMetadataCounts) Reset()         { *m = NodeMetadataCounts{} }
func (m *NodeMetadataCounts) String() string { return proto.CompactTextString(m) }
func (*NodeMetadataCounts) ProtoMessage()    {}
func (*NodeMetadataCounts) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeeb0a9b2be1cdd5, []int{7}
}

func (m *NodeMetadataCounts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeMetadataCounts.Unmarshal(m, b)
}
func (m *NodeMetadataCounts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeMetadataCounts.Marshal(b, m, deterministic)
}
func (m *NodeMetadataCounts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeMetadataCounts.Merge(m, src)
}
func (m *NodeMetadataCounts) XXX_Size() int {
	return xxx_messageInfo_NodeMetadataCounts.Size(m)
}
func (m *NodeMetadataCounts) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeMetadataCounts.DiscardUnknown(m)
}

var xxx_messageInfo_NodeMetadataCounts proto.InternalMessageInfo

func (m *NodeMetadataCounts) GetType() []string {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *NodeMetadataCounts) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *NodeMetadataCounts) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *NodeMetadataCounts) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

type NodeRunsDailyStatusTimeSeries struct {
	// Node ID of the runs
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Number of past days
	DaysAgo              int32    `protobuf:"varint,2,opt,name=days_ago,json=daysAgo,proto3" json:"days_ago,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeRunsDailyStatusTimeSeries) Reset()         { *m = NodeRunsDailyStatusTimeSeries{} }
func (m *NodeRunsDailyStatusTimeSeries) String() string { return proto.CompactTextString(m) }
func (*NodeRunsDailyStatusTimeSeries) ProtoMessage()    {}
func (*NodeRunsDailyStatusTimeSeries) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeeb0a9b2be1cdd5, []int{8}
}

func (m *NodeRunsDailyStatusTimeSeries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeRunsDailyStatusTimeSeries.Unmarshal(m, b)
}
func (m *NodeRunsDailyStatusTimeSeries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeRunsDailyStatusTimeSeries.Marshal(b, m, deterministic)
}
func (m *NodeRunsDailyStatusTimeSeries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeRunsDailyStatusTimeSeries.Merge(m, src)
}
func (m *NodeRunsDailyStatusTimeSeries) XXX_Size() int {
	return xxx_messageInfo_NodeRunsDailyStatusTimeSeries.Size(m)
}
func (m *NodeRunsDailyStatusTimeSeries) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeRunsDailyStatusTimeSeries.DiscardUnknown(m)
}

var xxx_messageInfo_NodeRunsDailyStatusTimeSeries proto.InternalMessageInfo

func (m *NodeRunsDailyStatusTimeSeries) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *NodeRunsDailyStatusTimeSeries) GetDaysAgo() int32 {
	if m != nil {
		return m.DaysAgo
	}
	return 0
}

func init() {
	proto.RegisterType((*Nodes)(nil), "chef.automate.api.cfgmgmt.request.Nodes")
	proto.RegisterType((*Node)(nil), "chef.automate.api.cfgmgmt.request.Node")
	proto.RegisterType((*NodeRun)(nil), "chef.automate.api.cfgmgmt.request.NodeRun")
	proto.RegisterType((*Runs)(nil), "chef.automate.api.cfgmgmt.request.Runs")
	proto.RegisterType((*NodeExport)(nil), "chef.automate.api.cfgmgmt.request.NodeExport")
	proto.RegisterType((*ReportExport)(nil), "chef.automate.api.cfgmgmt.request.ReportExport")
	proto.RegisterType((*MissingNodeDurationCounts)(nil), "chef.automate.api.cfgmgmt.request.MissingNodeDurationCounts")
	proto.RegisterType((*NodeMetadataCounts)(nil), "chef.automate.api.cfgmgmt.request.NodeMetadataCounts")
	proto.RegisterType((*NodeRunsDailyStatusTimeSeries)(nil), "chef.automate.api.cfgmgmt.request.NodeRunsDailyStatusTimeSeries")
}

func init() {
	proto.RegisterFile("api/external/cfgmgmt/request/nodes.proto", fileDescriptor_eeeb0a9b2be1cdd5)
}

var fileDescriptor_eeeb0a9b2be1cdd5 = []byte{
	// 690 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcb, 0x4e, 0xdb, 0x40,
	0x14, 0x95, 0xf3, 0x24, 0x03, 0x8b, 0xca, 0xea, 0x23, 0x44, 0xad, 0x48, 0xef, 0xa6, 0x69, 0x6a,
	0xec, 0x12, 0x84, 0x2a, 0xa5, 0x52, 0xa5, 0x14, 0xba, 0xa0, 0x52, 0x5a, 0xe4, 0xb0, 0x6a, 0x85,
	0xa2, 0x09, 0x9e, 0x18, 0x4b, 0xf1, 0x8c, 0x99, 0x87, 0xc0, 0x42, 0xec, 0xbb, 0xe6, 0x37, 0xfa,
	0x07, 0x5d, 0xf5, 0x1b, 0xf8, 0xa2, 0x6a, 0xc6, 0x0e, 0x09, 0x88, 0x84, 0xaa, 0xdd, 0x24, 0x73,
	0xaf, 0xee, 0xe3, 0xcc, 0x39, 0x67, 0x8c, 0x5a, 0x38, 0x89, 0x3c, 0x72, 0x2e, 0x09, 0xa7, 0x78,
	0xe2, 0x1d, 0x8f, 0xc3, 0x38, 0x8c, 0xa5, 0xc7, 0xc9, 0xa9, 0x22, 0x42, 0x7a, 0x94, 0x05, 0x44,
	0xb8, 0x09, 0x67, 0x92, 0xd9, 0x2f, 0x8f, 0x4f, 0xc8, 0xd8, 0xc5, 0x4a, 0xb2, 0x18, 0x4b, 0xe2,
	0xe2, 0x24, 0x72, 0xf3, 0x72, 0x37, 0x2f, 0x6f, 0x6c, 0x84, 0x8c, 0x85, 0x13, 0xe2, 0x99, 0x86,
	0x91, 0x1a, 0x7b, 0x32, 0x8a, 0x89, 0x90, 0x38, 0x4e, 0xb2, 0x19, 0x8d, 0xf6, 0xed, 0x6d, 0x2c,
	0x8e, 0x19, 0xf5, 0x4e, 0x15, 0xe1, 0xa9, 0x97, 0x60, 0x8e, 0x63, 0x22, 0x09, 0xcf, 0xf7, 0x35,
	0x1c, 0xf3, 0x77, 0xbc, 0x19, 0x12, 0xba, 0x29, 0xce, 0x70, 0x18, 0x12, 0xee, 0xb1, 0x44, 0x46,
	0x8c, 0x0a, 0x0f, 0x53, 0xca, 0x24, 0x36, 0xe7, 0xac, 0x1a, 0x7e, 0x15, 0x50, 0xf9, 0x8b, 0x46,
	0x6b, 0x3f, 0x45, 0x95, 0x71, 0x34, 0x91, 0x84, 0xd7, 0xad, 0x66, 0xb1, 0x55, 0xf3, 0xf3, 0xc8,
	0xfe, 0x8c, 0x50, 0x82, 0xc3, 0x88, 0x9a, 0xb6, 0x7a, 0xa1, 0x69, 0xb5, 0x56, 0x3b, 0x6d, 0xf7,
	0x9e, 0x4b, 0x19, 0x54, 0xae, 0x41, 0xe5, 0x1e, 0xdc, 0x74, 0xf8, 0x73, 0xdd, 0x76, 0x0f, 0x55,
	0x05, 0xe3, 0x32, 0xa2, 0x61, 0xbd, 0x68, 0x06, 0xbd, 0x7a, 0x68, 0xd0, 0x20, 0x2b, 0xf7, 0xa7,
	0x7d, 0xf6, 0x63, 0x54, 0x16, 0x12, 0x73, 0x59, 0x2f, 0x35, 0xad, 0x56, 0xcd, 0xcf, 0x02, 0xfb,
	0x11, 0x2a, 0x12, 0x1a, 0xd4, 0xcb, 0x26, 0xa7, 0x8f, 0xdd, 0xd1, 0x55, 0x6f, 0x88, 0x6a, 0xd7,
	0x56, 0x7e, 0x8d, 0x8e, 0x6f, 0x1f, 0x5c, 0x40, 0x76, 0x86, 0x6e, 0xf3, 0x3b, 0x50, 0x1c, 0x93,
	0x6e, 0x9c, 0x0e, 0xbe, 0xb6, 0xc1, 0x81, 0x64, 0x82, 0xe5, 0x98, 0xf1, 0xb8, 0xab, 0x46, 0x8a,
	0xb6, 0xe1, 0xc8, 0x69, 0xc2, 0x0c, 0x35, 0x74, 0x9b, 0xb0, 0x05, 0x4e, 0x13, 0x72, 0x08, 0x3a,
	0xee, 0x0d, 0x76, 0xe1, 0x12, 0x36, 0x50, 0x49, 0x73, 0x67, 0x3f, 0x43, 0x55, 0xad, 0xf8, 0x30,
	0x0a, 0xea, 0x96, 0x41, 0x50, 0xd1, 0xe1, 0x7e, 0x00, 0x09, 0xaa, 0xea, 0x02, 0x5f, 0xd1, 0x85,
	0x35, 0xf6, 0x13, 0x54, 0xe1, 0x8a, 0xea, 0x7c, 0x21, 0xbb, 0x11, 0x57, 0x74, 0x3f, 0xb0, 0x77,
	0xd0, 0x0a, 0xa1, 0xc1, 0x50, 0x3b, 0x21, 0xe7, 0xaa, 0xe1, 0x66, 0x36, 0x71, 0xa7, 0x36, 0x71,
	0x0f, 0xa7, 0x36, 0xf1, 0xab, 0x84, 0x06, 0x3a, 0x82, 0x9f, 0x16, 0x2a, 0xf9, 0x8a, 0x8a, 0xc5,
	0xfb, 0x66, 0x3a, 0x17, 0x96, 0xe8, 0x5c, 0xfc, 0x2f, 0x9d, 0xff, 0x52, 0x24, 0xf8, 0x61, 0x21,
	0xa4, 0x09, 0xfa, 0x74, 0x9e, 0x30, 0x2e, 0x17, 0x5a, 0x70, 0xce, 0x36, 0x85, 0x7f, 0xb4, 0xcd,
	0x06, 0x5a, 0x65, 0x4a, 0x26, 0x4a, 0x0e, 0x65, 0x9a, 0x64, 0x8c, 0xd6, 0x7c, 0x94, 0xa5, 0x0e,
	0xd3, 0x84, 0xc0, 0x6f, 0x0b, 0xad, 0xf9, 0x44, 0xc3, 0x78, 0x00, 0xcc, 0x9d, 0x49, 0x85, 0xbb,
	0x93, 0xe6, 0x99, 0x2f, 0xde, 0x62, 0xfe, 0xed, 0x3c, 0x2b, 0xcb, 0xf5, 0xcc, 0x19, 0x73, 0x66,
	0x8c, 0x2d, 0xaf, 0x37, 0x6c, 0xa6, 0x68, 0xbd, 0x1f, 0x09, 0x11, 0xd1, 0x50, 0x73, 0xba, 0xa7,
	0xb8, 0x11, 0x63, 0x97, 0x29, 0x2a, 0x85, 0xfd, 0x1c, 0xd5, 0x82, 0x3c, 0x23, 0xf2, 0x1b, 0xcd,
	0x12, 0xdd, 0x0f, 0x57, 0xbd, 0xf7, 0x68, 0xed, 0xda, 0x9a, 0xa5, 0x3a, 0x6f, 0xec, 0xd7, 0x17,
	0x70, 0x13, 0x9a, 0x37, 0xb3, 0x1d, 0x80, 0x03, 0x5b, 0x67, 0xe0, 0x40, 0x47, 0xff, 0x6c, 0xf5,
	0xc1, 0x81, 0xed, 0x3e, 0x1c, 0x5d, 0xc2, 0x09, 0xb2, 0xf5, 0xce, 0x3e, 0x91, 0x38, 0xc0, 0x12,
	0xe7, 0x3b, 0x6d, 0x54, 0x32, 0x1c, 0x65, 0xeb, 0xcc, 0x79, 0xa1, 0xfd, 0x6e, 0x2c, 0x53, 0xbc,
	0xc7, 0x32, 0xa5, 0x99, 0x65, 0x06, 0xe8, 0x45, 0xfe, 0xa4, 0xc4, 0x1e, 0x8e, 0x26, 0xe9, 0x40,
	0x62, 0xa9, 0x84, 0x66, 0x62, 0x40, 0x78, 0x44, 0x96, 0x18, 0x7f, 0x1d, 0xad, 0x04, 0x38, 0x15,
	0x43, 0x1c, 0x32, 0xa3, 0x5a, 0xd9, 0xaf, 0xea, 0xb8, 0x17, 0xb2, 0x8f, 0xef, 0xbe, 0xed, 0x84,
	0x91, 0x3c, 0x51, 0x23, 0xed, 0x22, 0x4f, 0x7b, 0xcb, 0x9b, 0x7a, 0xcb, 0x5b, 0xf6, 0xa1, 0x1f,
	0x55, 0x8c, 0x16, 0xdb, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x65, 0x8f, 0x94, 0x11, 0x0f, 0x06,
	0x00, 0x00,
}
