// Code generated by protoc-gen-go. DO NOT EDIT.
// source: integers.proto

package integers

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

type SignedInteger struct {
	Integer              *int32   `protobuf:"varint,1,opt,name=integer" json:"integer,omitempty"`
	SInteger             *int32   `protobuf:"zigzag32,2,opt,name=s_integer,json=sInteger" json:"s_integer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignedInteger) Reset()         { *m = SignedInteger{} }
func (m *SignedInteger) String() string { return proto.CompactTextString(m) }
func (*SignedInteger) ProtoMessage()    {}
func (*SignedInteger) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7d2a4376ec32b36, []int{0}
}

func (m *SignedInteger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedInteger.Unmarshal(m, b)
}
func (m *SignedInteger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedInteger.Marshal(b, m, deterministic)
}
func (m *SignedInteger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedInteger.Merge(m, src)
}
func (m *SignedInteger) XXX_Size() int {
	return xxx_messageInfo_SignedInteger.Size(m)
}
func (m *SignedInteger) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedInteger.DiscardUnknown(m)
}

var xxx_messageInfo_SignedInteger proto.InternalMessageInfo

func (m *SignedInteger) GetInteger() int32 {
	if m != nil && m.Integer != nil {
		return *m.Integer
	}
	return 0
}

func (m *SignedInteger) GetSInteger() int32 {
	if m != nil && m.SInteger != nil {
		return *m.SInteger
	}
	return 0
}

func init() {
	proto.RegisterType((*SignedInteger)(nil), "integers.SignedInteger")
}

func init() { proto.RegisterFile("integers.proto", fileDescriptor_e7d2a4376ec32b36) }

var fileDescriptor_e7d2a4376ec32b36 = []byte{
	// 89 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0xcc, 0x2b, 0x49,
	0x4d, 0x4f, 0x2d, 0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0xdc,
	0xb8, 0x78, 0x83, 0x33, 0xd3, 0xf3, 0x52, 0x53, 0x3c, 0x21, 0x22, 0x42, 0x12, 0x5c, 0xec, 0x50,
	0x49, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x18, 0x57, 0x48, 0x9a, 0x8b, 0xb3, 0x38, 0x1e,
	0x26, 0xc7, 0xa4, 0xc0, 0xa8, 0x21, 0x18, 0xc4, 0x51, 0x0c, 0xd5, 0x06, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x94, 0xc2, 0x0b, 0x3d, 0x62, 0x00, 0x00, 0x00,
}
