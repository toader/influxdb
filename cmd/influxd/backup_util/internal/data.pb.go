// Code generated by protoc-gen-gogo.
// source: internal/data.proto
// DO NOT EDIT!

/*
Package backup_util is a generated protocol buffer package.

It is generated from these files:
	internal/data.proto

It has these top-level messages:
	PortableData
*/
package backup_util

import (
	fmt "fmt"

	proto "github.com/gogo/protobuf/proto"

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type PortableData struct {
	Data             []byte  `protobuf:"bytes,1,req,name=Data" json:"Data,omitempty"`
	MaxNodeID        *uint64 `protobuf:"varint,2,req,name=MaxNodeID" json:"MaxNodeID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PortableData) Reset()                    { *m = PortableData{} }
func (m *PortableData) String() string            { return proto.CompactTextString(m) }
func (*PortableData) ProtoMessage()               {}
func (*PortableData) Descriptor() ([]byte, []int) { return fileDescriptorData, []int{0} }

func (m *PortableData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *PortableData) GetMaxNodeID() uint64 {
	if m != nil && m.MaxNodeID != nil {
		return *m.MaxNodeID
	}
	return 0
}

func init() {
	proto.RegisterType((*PortableData)(nil), "backup_util.PortableData")
}

func init() { proto.RegisterFile("internal/data.proto", fileDescriptorData) }

var fileDescriptorData = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x4f, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x4e, 0x4a, 0x4c, 0xce, 0x2e, 0x2d, 0x88, 0x2f, 0x2d, 0xc9, 0xcc, 0x51, 0x72, 0xe0, 0xe2,
	0x09, 0xc8, 0x2f, 0x2a, 0x49, 0x4c, 0xca, 0x49, 0x75, 0x49, 0x2c, 0x49, 0x14, 0x12, 0xe2, 0x62,
	0x01, 0xd1, 0x12, 0x8c, 0x0a, 0x4c, 0x1a, 0x3c, 0x41, 0x60, 0xb6, 0x90, 0x0c, 0x17, 0xa7, 0x6f,
	0x62, 0x85, 0x5f, 0x7e, 0x4a, 0xaa, 0xa7, 0x8b, 0x04, 0x93, 0x02, 0x93, 0x06, 0x4b, 0x10, 0x42,
	0x00, 0x10, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x54, 0xdc, 0x48, 0x64, 0x00, 0x00, 0x00,
}
