// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld/helloworld.proto

package helloworld

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

// The request message containing the user's name.
type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73149fedf49f4319, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_73149fedf49f4319, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
}

func init() { proto.RegisterFile("helloworld/helloworld.proto", fileDescriptor_73149fedf49f4319) }

var fileDescriptor_73149fedf49f4319 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x3f, 0x0b, 0xc2, 0x30,
	0x14, 0xc4, 0x2d, 0x88, 0xd5, 0x87, 0x53, 0x06, 0x29, 0xba, 0x48, 0x06, 0x71, 0x31, 0xa5, 0xea,
	0xec, 0xa0, 0x83, 0xba, 0x95, 0xba, 0xb9, 0xf5, 0xcf, 0xa3, 0x2d, 0xa4, 0x4d, 0x4d, 0x13, 0xa4,
	0xdf, 0x5e, 0x2c, 0xd6, 0x06, 0x74, 0xbb, 0xcb, 0xfd, 0x8e, 0x3c, 0x0e, 0x16, 0x19, 0x72, 0x2e,
	0x9e, 0x42, 0xf2, 0xc4, 0xed, 0x25, 0xab, 0xa4, 0x50, 0x82, 0x40, 0xff, 0x42, 0x29, 0x4c, 0x2f,
	0x6f, 0x17, 0xe0, 0x43, 0x63, 0xad, 0x08, 0x81, 0x61, 0x19, 0x16, 0xe8, 0x58, 0x4b, 0x6b, 0x3d,
	0x09, 0x5a, 0x4d, 0x57, 0x00, 0x1f, 0xa6, 0xe2, 0x0d, 0x71, 0xc0, 0x2e, 0xb0, 0xae, 0xc3, 0xb4,
	0x83, 0x3a, 0xbb, 0xbd, 0x82, 0x7d, 0x96, 0x88, 0x0a, 0x25, 0x39, 0xc0, 0xf8, 0x16, 0x36, 0x6d,
	0x8b, 0x38, 0xcc, 0xb8, 0xc0, 0xfc, 0x6c, 0x3e, 0xfb, 0x93, 0x54, 0xbc, 0xa1, 0x83, 0x63, 0x06,
	0x9e, 0x11, 0xa5, 0x81, 0x7f, 0xda, 0xe4, 0xa5, 0x92, 0x22, 0xd1, 0xb1, 0xca, 0x45, 0xc9, 0x38,
	0xc6, 0x4a, 0x4b, 0xdc, 0x6b, 0x96, 0xe6, 0x2a, 0xd3, 0x11, 0x8b, 0x45, 0xe1, 0x5b, 0x77, 0xaf,
	0x77, 0xee, 0x17, 0x71, 0x7f, 0xea, 0xc6, 0x24, 0xd1, 0xa8, 0xdd, 0x64, 0xf7, 0x0a, 0x00, 0x00,
	0xff, 0xff, 0x8f, 0xbe, 0xb6, 0xe4, 0x32, 0x01, 0x00, 0x00,
}