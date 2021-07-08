// Code generated by protoc-gen-go. DO NOT EDIT.
// source: publish.proto

package protofiles

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

type Publish struct {
	Type                 uint32   `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Sender               string   `protobuf:"bytes,2,opt,name=Sender,proto3" json:"Sender,omitempty"`
	Message              []byte   `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Publish) Reset()         { *m = Publish{} }
func (m *Publish) String() string { return proto.CompactTextString(m) }
func (*Publish) ProtoMessage()    {}
func (*Publish) Descriptor() ([]byte, []int) {
	return fileDescriptor_34180b7635741fb2, []int{0}
}

func (m *Publish) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Publish.Unmarshal(m, b)
}
func (m *Publish) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Publish.Marshal(b, m, deterministic)
}
func (m *Publish) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Publish.Merge(m, src)
}
func (m *Publish) XXX_Size() int {
	return xxx_messageInfo_Publish.Size(m)
}
func (m *Publish) XXX_DiscardUnknown() {
	xxx_messageInfo_Publish.DiscardUnknown(m)
}

var xxx_messageInfo_Publish proto.InternalMessageInfo

func (m *Publish) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Publish) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Publish) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func init() {
	proto.RegisterType((*Publish)(nil), "protofiles.Publish")
}

func init() { proto.RegisterFile("publish.proto", fileDescriptor_34180b7635741fb2) }

var fileDescriptor_34180b7635741fb2 = []byte{
	// 115 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x28, 0x4d, 0xca,
	0xc9, 0x2c, 0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x02, 0x53, 0x69, 0x99, 0x39,
	0xa9, 0xc5, 0x4a, 0xfe, 0x5c, 0xec, 0x01, 0x10, 0x49, 0x21, 0x21, 0x2e, 0x96, 0x90, 0xca, 0x82,
	0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xde, 0x20, 0x30, 0x5b, 0x48, 0x8c, 0x8b, 0x2d, 0x38, 0x35,
	0x2f, 0x25, 0xb5, 0x48, 0x82, 0x49, 0x81, 0x51, 0x83, 0x33, 0x08, 0xca, 0x13, 0x92, 0xe0, 0x62,
	0xf7, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x95, 0x60, 0x56, 0x60, 0xd4, 0xe0, 0x09, 0x82, 0x71,
	0x93, 0xd8, 0xc0, 0x86, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xb3, 0x81, 0x9b, 0x74,
	0x00, 0x00, 0x00,
}
