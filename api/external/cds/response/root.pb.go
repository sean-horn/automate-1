// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/cds/response/root.proto

package response

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

type ContentItems struct {
	Items                []*ContentItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ContentItems) Reset()         { *m = ContentItems{} }
func (m *ContentItems) String() string { return proto.CompactTextString(m) }
func (*ContentItems) ProtoMessage()    {}
func (*ContentItems) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b612a741fd15c61, []int{0}
}

func (m *ContentItems) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentItems.Unmarshal(m, b)
}
func (m *ContentItems) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentItems.Marshal(b, m, deterministic)
}
func (m *ContentItems) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentItems.Merge(m, src)
}
func (m *ContentItems) XXX_Size() int {
	return xxx_messageInfo_ContentItems.Size(m)
}
func (m *ContentItems) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentItems.DiscardUnknown(m)
}

var xxx_messageInfo_ContentItems proto.InternalMessageInfo

func (m *ContentItems) GetItems() []*ContentItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type ContentItem struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Version              string   `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Platform             string   `protobuf:"bytes,5,opt,name=platform,proto3" json:"platform,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContentItem) Reset()         { *m = ContentItem{} }
func (m *ContentItem) String() string { return proto.CompactTextString(m) }
func (*ContentItem) ProtoMessage()    {}
func (*ContentItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b612a741fd15c61, []int{1}
}

func (m *ContentItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentItem.Unmarshal(m, b)
}
func (m *ContentItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentItem.Marshal(b, m, deterministic)
}
func (m *ContentItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentItem.Merge(m, src)
}
func (m *ContentItem) XXX_Size() int {
	return xxx_messageInfo_ContentItem.Size(m)
}
func (m *ContentItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentItem.DiscardUnknown(m)
}

var xxx_messageInfo_ContentItem proto.InternalMessageInfo

func (m *ContentItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContentItem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ContentItem) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ContentItem) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ContentItem) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func init() {
	proto.RegisterType((*ContentItems)(nil), "chef.automate.api.cds.response.ContentItems")
	proto.RegisterType((*ContentItem)(nil), "chef.automate.api.cds.response.ContentItem")
}

func init() {
	proto.RegisterFile("api/external/cds/response/root.proto", fileDescriptor_4b612a741fd15c61)
}

var fileDescriptor_4b612a741fd15c61 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xbd, 0x4e, 0x03, 0x31,
	0x10, 0x84, 0x75, 0x24, 0xe1, 0xc7, 0x47, 0xe5, 0xca, 0xa2, 0x40, 0xa7, 0x88, 0x22, 0x12, 0xd2,
	0x5a, 0x0a, 0xbc, 0x00, 0x50, 0x51, 0x92, 0x92, 0xce, 0xf1, 0x6d, 0x88, 0xa5, 0xd8, 0x6b, 0xd9,
	0x1b, 0x04, 0x0f, 0xc1, 0x3b, 0x23, 0xfb, 0x74, 0xe8, 0x1a, 0xd2, 0xcd, 0xce, 0xec, 0x57, 0xcc,
	0x88, 0x3b, 0x13, 0x9d, 0xc6, 0x2f, 0xc6, 0x14, 0xcc, 0x41, 0xdb, 0x3e, 0xeb, 0x84, 0x39, 0x52,
	0xc8, 0xa8, 0x13, 0x11, 0x43, 0x4c, 0xc4, 0x24, 0x6f, 0xed, 0x1e, 0x77, 0x60, 0x8e, 0x4c, 0xde,
	0x30, 0x82, 0x89, 0x0e, 0x6c, 0x9f, 0x61, 0x7c, 0x5d, 0xbe, 0x89, 0xeb, 0x17, 0x0a, 0x8c, 0x81,
	0x5f, 0x19, 0x7d, 0x96, 0x4f, 0x62, 0xe1, 0x8a, 0x50, 0x4d, 0x37, 0x5b, 0xb5, 0xeb, 0x7b, 0x38,
	0xcd, 0xc3, 0x04, 0xde, 0x0c, 0xe4, 0xf2, 0xa7, 0x11, 0xed, 0xc4, 0x96, 0x52, 0xcc, 0x83, 0xf1,
	0xa8, 0x9a, 0xae, 0x59, 0x5d, 0x6d, 0xaa, 0x96, 0x9d, 0x68, 0x7b, 0xcc, 0x36, 0xb9, 0xc8, 0x8e,
	0x82, 0x3a, 0xab, 0xd1, 0xd4, 0x2a, 0x14, 0x7f, 0x47, 0x54, 0xb3, 0x81, 0x2a, 0x5a, 0x2a, 0x71,
	0xf1, 0x89, 0x29, 0x17, 0x62, 0x5e, 0xed, 0xf1, 0x94, 0x37, 0xe2, 0x32, 0x1e, 0x0c, 0xef, 0x28,
	0x79, 0xb5, 0xa8, 0xd1, 0xdf, 0xfd, 0xfc, 0xf8, 0xbe, 0xfe, 0x70, 0xbc, 0x3f, 0x6e, 0xc1, 0x92,
	0xd7, 0xa5, 0x8f, 0x1e, 0xfb, 0xe8, 0x7f, 0x37, 0xdc, 0x9e, 0xd7, 0xfd, 0x1e, 0x7e, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xe4, 0x4c, 0x7a, 0x01, 0x67, 0x01, 0x00, 0x00,
}
