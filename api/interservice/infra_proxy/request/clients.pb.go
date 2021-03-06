// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/infra_proxy/request/clients.proto

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

type Clients struct {
	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId             string   `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Clients) Reset()         { *m = Clients{} }
func (m *Clients) String() string { return proto.CompactTextString(m) }
func (*Clients) ProtoMessage()    {}
func (*Clients) Descriptor() ([]byte, []int) {
	return fileDescriptor_83a21eb78b74ff59, []int{0}
}

func (m *Clients) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Clients.Unmarshal(m, b)
}
func (m *Clients) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Clients.Marshal(b, m, deterministic)
}
func (m *Clients) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Clients.Merge(m, src)
}
func (m *Clients) XXX_Size() int {
	return xxx_messageInfo_Clients.Size(m)
}
func (m *Clients) XXX_DiscardUnknown() {
	xxx_messageInfo_Clients.DiscardUnknown(m)
}

var xxx_messageInfo_Clients proto.InternalMessageInfo

func (m *Clients) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *Clients) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

type Client struct {
	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// Client name.
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Client) Reset()         { *m = Client{} }
func (m *Client) String() string { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()    {}
func (*Client) Descriptor() ([]byte, []int) {
	return fileDescriptor_83a21eb78b74ff59, []int{1}
}

func (m *Client) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Client.Unmarshal(m, b)
}
func (m *Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Client.Marshal(b, m, deterministic)
}
func (m *Client) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Client.Merge(m, src)
}
func (m *Client) XXX_Size() int {
	return xxx_messageInfo_Client.Size(m)
}
func (m *Client) XXX_DiscardUnknown() {
	xxx_messageInfo_Client.DiscardUnknown(m)
}

var xxx_messageInfo_Client proto.InternalMessageInfo

func (m *Client) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *Client) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *Client) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Clients)(nil), "chef.automate.domain.infra_proxy.request.Clients")
	proto.RegisterType((*Client)(nil), "chef.automate.domain.infra_proxy.request.Client")
}

func init() {
	proto.RegisterFile("api/interservice/infra_proxy/request/clients.proto", fileDescriptor_83a21eb78b74ff59)
}

var fileDescriptor_83a21eb78b74ff59 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0xbf, 0x4b, 0xc5, 0x30,
	0x14, 0x85, 0xa9, 0x3f, 0xaa, 0xcd, 0x18, 0x10, 0x0a, 0x2e, 0xd2, 0xa9, 0x53, 0x02, 0x3a, 0x17,
	0x41, 0xa7, 0x6e, 0xe2, 0xe8, 0x52, 0xd2, 0xe4, 0xb6, 0xbd, 0x60, 0x72, 0xeb, 0x6d, 0x2a, 0xfa,
	0xdf, 0x4b, 0x53, 0x1f, 0xbc, 0xf1, 0xf1, 0xb6, 0xcb, 0x39, 0x7c, 0x5c, 0xce, 0x27, 0x1e, 0xcd,
	0x8c, 0x1a, 0x43, 0x04, 0x5e, 0x80, 0xbf, 0xd1, 0x82, 0xc6, 0x30, 0xb0, 0xe9, 0x66, 0xa6, 0x9f,
	0x5f, 0xcd, 0xf0, 0xb5, 0xc2, 0x12, 0xb5, 0xfd, 0x44, 0x08, 0x71, 0x51, 0x33, 0x53, 0x24, 0x59,
	0xdb, 0x09, 0x06, 0x65, 0xd6, 0x48, 0xde, 0x44, 0x50, 0x8e, 0xbc, 0xc1, 0xa0, 0x8e, 0x38, 0xf5,
	0xcf, 0x55, 0x8d, 0xb8, 0x79, 0xdd, 0x51, 0x79, 0x27, 0x72, 0xe2, 0xb1, 0x43, 0x57, 0x66, 0x0f,
	0x59, 0x5d, 0xbc, 0x5f, 0x13, 0x8f, 0xad, 0x93, 0xf7, 0xa2, 0xd8, 0x1e, 0x03, 0x6f, 0xcd, 0x45,
	0x6a, 0x6e, 0xf7, 0xa0, 0x75, 0xd5, 0x9b, 0xc8, 0x77, 0xfc, 0x1c, 0x5a, 0x4a, 0x71, 0x15, 0x8c,
	0x87, 0xf2, 0x32, 0xe5, 0xe9, 0x7e, 0x79, 0xfe, 0x68, 0x46, 0x8c, 0xd3, 0xda, 0x2b, 0x4b, 0x5e,
	0x6f, 0x3b, 0xf4, 0x61, 0x87, 0x3e, 0xc5, 0x44, 0x9f, 0x27, 0x05, 0x4f, 0x7f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xdd, 0x0a, 0xa3, 0xf6, 0x38, 0x01, 0x00, 0x00,
}
