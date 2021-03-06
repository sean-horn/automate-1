// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/authz/type.proto

package authz

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

type Type int32

const (
	Type_CHEF_MANAGED Type = 0
	Type_CUSTOM       Type = 1
)

var Type_name = map[int32]string{
	0: "CHEF_MANAGED",
	1: "CUSTOM",
}

var Type_value = map[string]int32{
	"CHEF_MANAGED": 0,
	"CUSTOM":       1,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_98248891a06d7196, []int{0}
}

func init() {
	proto.RegisterEnum("chef.automate.domain.authz.Type", Type_name, Type_value)
}

func init() {
	proto.RegisterFile("api/interservice/authz/type.proto", fileDescriptor_98248891a06d7196)
}

var fileDescriptor_98248891a06d7196 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0x2c, 0xc8, 0xd4,
	0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0x2a, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x4f, 0x2c, 0x2d,
	0xc9, 0xa8, 0xd2, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4a,
	0xce, 0x48, 0x4d, 0xd3, 0x4b, 0x2c, 0x2d, 0xc9, 0xcf, 0x4d, 0x2c, 0x49, 0xd5, 0x4b, 0xc9, 0xcf,
	0x4d, 0xcc, 0xcc, 0xd3, 0x03, 0x2b, 0xd3, 0x52, 0xe1, 0x62, 0x09, 0xa9, 0x2c, 0x48, 0x15, 0x12,
	0xe0, 0xe2, 0x71, 0xf6, 0x70, 0x75, 0x8b, 0xf7, 0x75, 0xf4, 0x73, 0x74, 0x77, 0x75, 0x11, 0x60,
	0x10, 0xe2, 0xe2, 0x62, 0x73, 0x0e, 0x0d, 0x0e, 0xf1, 0xf7, 0x15, 0x60, 0x74, 0x32, 0x8c, 0xd2,
	0x4f, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x07, 0x19, 0xa7, 0x0f, 0x33,
	0x4e, 0x1f, 0xbb, 0xfd, 0x49, 0x6c, 0x60, 0xbb, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x26,
	0xb2, 0x37, 0x3e, 0xa0, 0x00, 0x00, 0x00,
}
