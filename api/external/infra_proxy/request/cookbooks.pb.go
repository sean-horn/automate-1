// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/infra_proxy/request/cookbooks.proto

package request

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

type Cookbooks struct {
	// ID of the Org.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// ID of the Server
	ServerId             string   `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cookbooks) Reset()         { *m = Cookbooks{} }
func (m *Cookbooks) String() string { return proto.CompactTextString(m) }
func (*Cookbooks) ProtoMessage()    {}
func (*Cookbooks) Descriptor() ([]byte, []int) {
	return fileDescriptor_334f5d297dde2e18, []int{0}
}

func (m *Cookbooks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cookbooks.Unmarshal(m, b)
}
func (m *Cookbooks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cookbooks.Marshal(b, m, deterministic)
}
func (m *Cookbooks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cookbooks.Merge(m, src)
}
func (m *Cookbooks) XXX_Size() int {
	return xxx_messageInfo_Cookbooks.Size(m)
}
func (m *Cookbooks) XXX_DiscardUnknown() {
	xxx_messageInfo_Cookbooks.DiscardUnknown(m)
}

var xxx_messageInfo_Cookbooks proto.InternalMessageInfo

func (m *Cookbooks) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *Cookbooks) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

type CookbookVersions struct {
	// ID of the Org.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// ID of the Server.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// Name of the cookbook.
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CookbookVersions) Reset()         { *m = CookbookVersions{} }
func (m *CookbookVersions) String() string { return proto.CompactTextString(m) }
func (*CookbookVersions) ProtoMessage()    {}
func (*CookbookVersions) Descriptor() ([]byte, []int) {
	return fileDescriptor_334f5d297dde2e18, []int{1}
}

func (m *CookbookVersions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CookbookVersions.Unmarshal(m, b)
}
func (m *CookbookVersions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CookbookVersions.Marshal(b, m, deterministic)
}
func (m *CookbookVersions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CookbookVersions.Merge(m, src)
}
func (m *CookbookVersions) XXX_Size() int {
	return xxx_messageInfo_CookbookVersions.Size(m)
}
func (m *CookbookVersions) XXX_DiscardUnknown() {
	xxx_messageInfo_CookbookVersions.DiscardUnknown(m)
}

var xxx_messageInfo_CookbookVersions proto.InternalMessageInfo

func (m *CookbookVersions) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *CookbookVersions) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *CookbookVersions) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Cookbook struct {
	// ID of the Org.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// ID of the Server.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// Name of the cookbook.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Version of the cookbook.
	Version              string   `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cookbook) Reset()         { *m = Cookbook{} }
func (m *Cookbook) String() string { return proto.CompactTextString(m) }
func (*Cookbook) ProtoMessage()    {}
func (*Cookbook) Descriptor() ([]byte, []int) {
	return fileDescriptor_334f5d297dde2e18, []int{2}
}

func (m *Cookbook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cookbook.Unmarshal(m, b)
}
func (m *Cookbook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cookbook.Marshal(b, m, deterministic)
}
func (m *Cookbook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cookbook.Merge(m, src)
}
func (m *Cookbook) XXX_Size() int {
	return xxx_messageInfo_Cookbook.Size(m)
}
func (m *Cookbook) XXX_DiscardUnknown() {
	xxx_messageInfo_Cookbook.DiscardUnknown(m)
}

var xxx_messageInfo_Cookbook proto.InternalMessageInfo

func (m *Cookbook) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *Cookbook) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *Cookbook) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Cookbook) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type CookbookFileContent struct {
	// ID of the org.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// ID of the server.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// Name of the cookbook.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Version of the cookbook.
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	// Cookbook data file URL.
	Url                  string   `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CookbookFileContent) Reset()         { *m = CookbookFileContent{} }
func (m *CookbookFileContent) String() string { return proto.CompactTextString(m) }
func (*CookbookFileContent) ProtoMessage()    {}
func (*CookbookFileContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_334f5d297dde2e18, []int{3}
}

func (m *CookbookFileContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CookbookFileContent.Unmarshal(m, b)
}
func (m *CookbookFileContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CookbookFileContent.Marshal(b, m, deterministic)
}
func (m *CookbookFileContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CookbookFileContent.Merge(m, src)
}
func (m *CookbookFileContent) XXX_Size() int {
	return xxx_messageInfo_CookbookFileContent.Size(m)
}
func (m *CookbookFileContent) XXX_DiscardUnknown() {
	xxx_messageInfo_CookbookFileContent.DiscardUnknown(m)
}

var xxx_messageInfo_CookbookFileContent proto.InternalMessageInfo

func (m *CookbookFileContent) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *CookbookFileContent) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *CookbookFileContent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CookbookFileContent) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *CookbookFileContent) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*Cookbooks)(nil), "chef.automate.api.infra_proxy.request.Cookbooks")
	proto.RegisterType((*CookbookVersions)(nil), "chef.automate.api.infra_proxy.request.CookbookVersions")
	proto.RegisterType((*Cookbook)(nil), "chef.automate.api.infra_proxy.request.Cookbook")
	proto.RegisterType((*CookbookFileContent)(nil), "chef.automate.api.infra_proxy.request.CookbookFileContent")
}

func init() {
	proto.RegisterFile("api/external/infra_proxy/request/cookbooks.proto", fileDescriptor_334f5d297dde2e18)
}

var fileDescriptor_334f5d297dde2e18 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x87, 0x89, 0xfd, 0x63, 0x33, 0xa7, 0xb2, 0x22, 0x2c, 0x78, 0x91, 0x80, 0xe0, 0x69, 0x57,
	0xf0, 0x24, 0x1e, 0x04, 0x0b, 0x42, 0xaf, 0x1e, 0x3c, 0xf4, 0x52, 0x36, 0xc9, 0x34, 0x5d, 0x9a,
	0xec, 0xc4, 0xc9, 0xa6, 0xd4, 0x47, 0xf0, 0xad, 0x25, 0x5b, 0x17, 0xbc, 0x09, 0x62, 0x6f, 0xb3,
	0xdf, 0xf0, 0xfb, 0x66, 0x61, 0x06, 0xee, 0x4c, 0x6b, 0x35, 0x1e, 0x3c, 0xb2, 0x33, 0xb5, 0xb6,
	0x6e, 0xc3, 0x66, 0xdd, 0x32, 0x1d, 0x3e, 0x34, 0xe3, 0x7b, 0x8f, 0x9d, 0xd7, 0x05, 0xd1, 0x2e,
	0x27, 0xda, 0x75, 0xaa, 0x65, 0xf2, 0x24, 0x6e, 0x8a, 0x2d, 0x6e, 0x94, 0xe9, 0x3d, 0x35, 0xc6,
	0xa3, 0x32, 0xad, 0x55, 0x3f, 0x62, 0xea, 0x3b, 0x96, 0x3d, 0x41, 0xba, 0x88, 0x49, 0x71, 0x09,
	0x53, 0xe2, 0x6a, 0x6d, 0x4b, 0x99, 0x5c, 0x27, 0xb7, 0xe9, 0xeb, 0x84, 0xb8, 0x5a, 0x96, 0xe2,
	0x0a, 0xd2, 0x0e, 0x79, 0x8f, 0x3c, 0x74, 0xce, 0x42, 0x67, 0x76, 0x04, 0xcb, 0x32, 0x5b, 0xc1,
	0x3c, 0x0a, 0xde, 0x90, 0x3b, 0x4b, 0xee, 0x4f, 0x1e, 0x21, 0x60, 0xec, 0x4c, 0x83, 0x72, 0x14,
	0x78, 0xa8, 0xb3, 0x1a, 0x66, 0xd1, 0xfd, 0x5f, 0x4e, 0x21, 0xe1, 0x7c, 0x7f, 0xfc, 0xa7, 0x1c,
	0x07, 0x1c, 0x9f, 0xd9, 0x67, 0x02, 0x17, 0x71, 0xdc, 0x8b, 0xad, 0x71, 0x41, 0xce, 0xa3, 0xf3,
	0xa7, 0x9f, 0x2c, 0xe6, 0x30, 0xea, 0xb9, 0x96, 0x93, 0x40, 0x87, 0xf2, 0xf9, 0x71, 0xf5, 0x50,
	0x59, 0xbf, 0xed, 0x73, 0x55, 0x50, 0xa3, 0x87, 0x55, 0xea, 0xb8, 0x4a, 0xfd, 0xdb, 0x29, 0xe4,
	0xd3, 0x70, 0x01, 0xf7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbd, 0xdf, 0xa6, 0x7a, 0x35, 0x02,
	0x00, 0x00,
}
