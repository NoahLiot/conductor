// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/search.proto

package search // import "github.com/netflix/conductor/client/gogrpc/conductor/grpc/search"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Start                int32    `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	Size                 int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Sort                 string   `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
	FreeText             string   `protobuf:"bytes,4,opt,name=free_text,json=freeText,proto3" json:"free_text,omitempty"`
	Query                string   `protobuf:"bytes,5,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_search_a50e745244da0fb1, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetStart() int32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Request) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Request) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

func (m *Request) GetFreeText() string {
	if m != nil {
		return m.FreeText
	}
	return ""
}

func (m *Request) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "conductor.grpc.search.Request")
}

func init() { proto.RegisterFile("grpc/search.proto", fileDescriptor_search_a50e745244da0fb1) }

var fileDescriptor_search_a50e745244da0fb1 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xa9, 0x6e, 0x75, 0x37, 0x37, 0x83, 0x42, 0xd0, 0xcb, 0xe2, 0x69, 0x4f, 0xc9, 0xc1,
	0x17, 0x90, 0x7d, 0x02, 0xa9, 0x9e, 0xbc, 0xc8, 0x36, 0x4e, 0xdb, 0x40, 0xdb, 0x69, 0x27, 0x13,
	0xa8, 0x7d, 0x7a, 0xe9, 0x54, 0xd4, 0xbd, 0xcd, 0x7c, 0xdf, 0x10, 0xfe, 0xfc, 0xea, 0xa6, 0xa6,
	0xc1, 0xbb, 0x08, 0x27, 0xf2, 0x8d, 0x1d, 0x08, 0x19, 0xf5, 0x9d, 0xc7, 0xfe, 0x33, 0x79, 0x46,
	0xb2, 0x8b, 0xb4, 0xab, 0x7c, 0x9c, 0xd5, 0x75, 0x01, 0x63, 0x82, 0xc8, 0xfa, 0x56, 0xe5, 0x91,
	0x4f, 0xc4, 0x26, 0xdb, 0x67, 0x87, 0xbc, 0x58, 0x17, 0xad, 0xd5, 0x26, 0x86, 0x19, 0xcc, 0x85,
	0x40, 0x99, 0x85, 0x21, 0xb1, 0xb9, 0xdc, 0x67, 0x87, 0x5d, 0x21, 0xb3, 0x7e, 0x50, 0xbb, 0x8a,
	0x00, 0x3e, 0x18, 0x26, 0x36, 0x1b, 0x11, 0xdb, 0x05, 0xbc, 0xc1, 0x24, 0x4f, 0x8f, 0x09, 0xe8,
	0xcb, 0xe4, 0x22, 0xd6, 0xe5, 0xd8, 0xa8, 0x7b, 0x8f, 0x9d, 0xed, 0x81, 0xab, 0x36, 0x4c, 0xf6,
	0x3c, 0xe0, 0x71, 0xfb, 0x2a, 0x09, 0x5f, 0xca, 0xf7, 0xe7, 0x3a, 0x70, 0x93, 0x4a, 0xeb, 0xb1,
	0x73, 0x3f, 0xc7, 0xee, 0xf7, 0xd8, 0xf9, 0x36, 0x40, 0xcf, 0xae, 0x46, 0xf9, 0xf3, 0x1f, 0xff,
	0x57, 0x41, 0x79, 0x25, 0x1d, 0x3c, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x4d, 0x39, 0xe7,
	0x18, 0x01, 0x00, 0x00,
}
