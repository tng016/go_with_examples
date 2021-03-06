// Code generated by protoc-gen-go. DO NOT EDIT.
// source: embeddedmessages.proto

package embeddedmessages

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

type TopMessage struct {
	M                    *EmbeddedMessage `protobuf:"bytes,1,opt,name=m" json:"m,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TopMessage) Reset()         { *m = TopMessage{} }
func (m *TopMessage) String() string { return proto.CompactTextString(m) }
func (*TopMessage) ProtoMessage()    {}
func (*TopMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb284566b89e2b4e, []int{0}
}

func (m *TopMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopMessage.Unmarshal(m, b)
}
func (m *TopMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopMessage.Marshal(b, m, deterministic)
}
func (m *TopMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopMessage.Merge(m, src)
}
func (m *TopMessage) XXX_Size() int {
	return xxx_messageInfo_TopMessage.Size(m)
}
func (m *TopMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TopMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TopMessage proto.InternalMessageInfo

func (m *TopMessage) GetM() *EmbeddedMessage {
	if m != nil {
		return m.M
	}
	return nil
}

type EmbeddedMessage struct {
	I                    *uint64  `protobuf:"varint,1,opt,name=i" json:"i,omitempty"`
	S                    *string  `protobuf:"bytes,2,opt,name=s" json:"s,omitempty"`
	B                    *bool    `protobuf:"varint,3,opt,name=b" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmbeddedMessage) Reset()         { *m = EmbeddedMessage{} }
func (m *EmbeddedMessage) String() string { return proto.CompactTextString(m) }
func (*EmbeddedMessage) ProtoMessage()    {}
func (*EmbeddedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb284566b89e2b4e, []int{1}
}

func (m *EmbeddedMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmbeddedMessage.Unmarshal(m, b)
}
func (m *EmbeddedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmbeddedMessage.Marshal(b, m, deterministic)
}
func (m *EmbeddedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmbeddedMessage.Merge(m, src)
}
func (m *EmbeddedMessage) XXX_Size() int {
	return xxx_messageInfo_EmbeddedMessage.Size(m)
}
func (m *EmbeddedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EmbeddedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EmbeddedMessage proto.InternalMessageInfo

func (m *EmbeddedMessage) GetI() uint64 {
	if m != nil && m.I != nil {
		return *m.I
	}
	return 0
}

func (m *EmbeddedMessage) GetS() string {
	if m != nil && m.S != nil {
		return *m.S
	}
	return ""
}

func (m *EmbeddedMessage) GetB() bool {
	if m != nil && m.B != nil {
		return *m.B
	}
	return false
}

func init() {
	proto.RegisterType((*TopMessage)(nil), "embeddedmessages.TopMessage")
	proto.RegisterType((*EmbeddedMessage)(nil), "embeddedmessages.EmbeddedMessage")
}

func init() { proto.RegisterFile("embeddedmessages.proto", fileDescriptor_fb284566b89e2b4e) }

var fileDescriptor_fb284566b89e2b4e = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xcd, 0x4d, 0x4a,
	0x4d, 0x49, 0x49, 0x4d, 0xc9, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0x40, 0x17, 0x57, 0xb2, 0xe5, 0xe2, 0x0a, 0xc9, 0x2f, 0xf0, 0x85, 0x70,
	0x85, 0xf4, 0xb9, 0x18, 0x73, 0x25, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x14, 0xf5, 0x30, 0xcc,
	0x70, 0x85, 0x0a, 0x40, 0x55, 0x07, 0x31, 0xe6, 0x2a, 0x59, 0x73, 0xf1, 0xa3, 0x89, 0x0a, 0xf1,
	0x70, 0x31, 0x66, 0x82, 0xcd, 0x60, 0x09, 0x62, 0xcc, 0x04, 0xf1, 0x8a, 0x25, 0x98, 0x14, 0x18,
	0x35, 0x38, 0x83, 0x18, 0x8b, 0x41, 0xbc, 0x24, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0xc6,
	0x24, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16, 0x68, 0x0e, 0x87, 0xa6, 0x00, 0x00, 0x00,
}
