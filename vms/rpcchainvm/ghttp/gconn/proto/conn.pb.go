// Code generated by protoc-gen-go. DO NOT EDIT.
// source: conn.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ReadRequest struct {
	Length               int32    `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f401a58c1fc7ceef, []int{0}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

type ReadResponse struct {
	Read                 []byte   `protobuf:"bytes,1,opt,name=read,proto3" json:"read,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Errored              bool     `protobuf:"varint,3,opt,name=errored,proto3" json:"errored,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f401a58c1fc7ceef, []int{1}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetRead() []byte {
	if m != nil {
		return m.Read
	}
	return nil
}

func (m *ReadResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ReadResponse) GetErrored() bool {
	if m != nil {
		return m.Errored
	}
	return false
}

type WriteRequest struct {
	Payload              []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteRequest) Reset()         { *m = WriteRequest{} }
func (m *WriteRequest) String() string { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()    {}
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f401a58c1fc7ceef, []int{2}
}

func (m *WriteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteRequest.Unmarshal(m, b)
}
func (m *WriteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteRequest.Marshal(b, m, deterministic)
}
func (m *WriteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRequest.Merge(m, src)
}
func (m *WriteRequest) XXX_Size() int {
	return xxx_messageInfo_WriteRequest.Size(m)
}
func (m *WriteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRequest proto.InternalMessageInfo

func (m *WriteRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type WriteResponse struct {
	Length               int32    `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Errored              bool     `protobuf:"varint,3,opt,name=errored,proto3" json:"errored,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteResponse) Reset()         { *m = WriteResponse{} }
func (m *WriteResponse) String() string { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()    {}
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f401a58c1fc7ceef, []int{3}
}

func (m *WriteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteResponse.Unmarshal(m, b)
}
func (m *WriteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteResponse.Marshal(b, m, deterministic)
}
func (m *WriteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteResponse.Merge(m, src)
}
func (m *WriteResponse) XXX_Size() int {
	return xxx_messageInfo_WriteResponse.Size(m)
}
func (m *WriteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WriteResponse proto.InternalMessageInfo

func (m *WriteResponse) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *WriteResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *WriteResponse) GetErrored() bool {
	if m != nil {
		return m.Errored
	}
	return false
}

type CloseRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseRequest) Reset()         { *m = CloseRequest{} }
func (m *CloseRequest) String() string { return proto.CompactTextString(m) }
func (*CloseRequest) ProtoMessage()    {}
func (*CloseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f401a58c1fc7ceef, []int{4}
}

func (m *CloseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseRequest.Unmarshal(m, b)
}
func (m *CloseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseRequest.Marshal(b, m, deterministic)
}
func (m *CloseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseRequest.Merge(m, src)
}
func (m *CloseRequest) XXX_Size() int {
	return xxx_messageInfo_CloseRequest.Size(m)
}
func (m *CloseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CloseRequest proto.InternalMessageInfo

type CloseResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseResponse) Reset()         { *m = CloseResponse{} }
func (m *CloseResponse) String() string { return proto.CompactTextString(m) }
func (*CloseResponse) ProtoMessage()    {}
func (*CloseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f401a58c1fc7ceef, []int{5}
}

func (m *CloseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseResponse.Unmarshal(m, b)
}
func (m *CloseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseResponse.Marshal(b, m, deterministic)
}
func (m *CloseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseResponse.Merge(m, src)
}
func (m *CloseResponse) XXX_Size() int {
	return xxx_messageInfo_CloseResponse.Size(m)
}
func (m *CloseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CloseResponse proto.InternalMessageInfo

type SetDeadlineRequest struct {
	Time                 []byte   `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetDeadlineRequest) Reset()         { *m = SetDeadlineRequest{} }
func (m *SetDeadlineRequest) String() string { return proto.CompactTextString(m) }
func (*SetDeadlineRequest) ProtoMessage()    {}
func (*SetDeadlineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f401a58c1fc7ceef, []int{6}
}

func (m *SetDeadlineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDeadlineRequest.Unmarshal(m, b)
}
func (m *SetDeadlineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDeadlineRequest.Marshal(b, m, deterministic)
}
func (m *SetDeadlineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDeadlineRequest.Merge(m, src)
}
func (m *SetDeadlineRequest) XXX_Size() int {
	return xxx_messageInfo_SetDeadlineRequest.Size(m)
}
func (m *SetDeadlineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDeadlineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetDeadlineRequest proto.InternalMessageInfo

func (m *SetDeadlineRequest) GetTime() []byte {
	if m != nil {
		return m.Time
	}
	return nil
}

type SetDeadlineResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetDeadlineResponse) Reset()         { *m = SetDeadlineResponse{} }
func (m *SetDeadlineResponse) String() string { return proto.CompactTextString(m) }
func (*SetDeadlineResponse) ProtoMessage()    {}
func (*SetDeadlineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f401a58c1fc7ceef, []int{7}
}

func (m *SetDeadlineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDeadlineResponse.Unmarshal(m, b)
}
func (m *SetDeadlineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDeadlineResponse.Marshal(b, m, deterministic)
}
func (m *SetDeadlineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDeadlineResponse.Merge(m, src)
}
func (m *SetDeadlineResponse) XXX_Size() int {
	return xxx_messageInfo_SetDeadlineResponse.Size(m)
}
func (m *SetDeadlineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDeadlineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetDeadlineResponse proto.InternalMessageInfo

type SetReadDeadlineRequest struct {
	Time                 []byte   `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetReadDeadlineRequest) Reset()         { *m = SetReadDeadlineRequest{} }
func (m *SetReadDeadlineRequest) String() string { return proto.CompactTextString(m) }
func (*SetReadDeadlineRequest) ProtoMessage()    {}
func (*SetReadDeadlineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f401a58c1fc7ceef, []int{8}
}

func (m *SetReadDeadlineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetReadDeadlineRequest.Unmarshal(m, b)
}
func (m *SetReadDeadlineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetReadDeadlineRequest.Marshal(b, m, deterministic)
}
func (m *SetReadDeadlineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetReadDeadlineRequest.Merge(m, src)
}
func (m *SetReadDeadlineRequest) XXX_Size() int {
	return xxx_messageInfo_SetReadDeadlineRequest.Size(m)
}
func (m *SetReadDeadlineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetReadDeadlineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetReadDeadlineRequest proto.InternalMessageInfo

func (m *SetReadDeadlineRequest) GetTime() []byte {
	if m != nil {
		return m.Time
	}
	return nil
}

type SetReadDeadlineResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetReadDeadlineResponse) Reset()         { *m = SetReadDeadlineResponse{} }
func (m *SetReadDeadlineResponse) String() string { return proto.CompactTextString(m) }
func (*SetReadDeadlineResponse) ProtoMessage()    {}
func (*SetReadDeadlineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f401a58c1fc7ceef, []int{9}
}

func (m *SetReadDeadlineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetReadDeadlineResponse.Unmarshal(m, b)
}
func (m *SetReadDeadlineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetReadDeadlineResponse.Marshal(b, m, deterministic)
}
func (m *SetReadDeadlineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetReadDeadlineResponse.Merge(m, src)
}
func (m *SetReadDeadlineResponse) XXX_Size() int {
	return xxx_messageInfo_SetReadDeadlineResponse.Size(m)
}
func (m *SetReadDeadlineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetReadDeadlineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetReadDeadlineResponse proto.InternalMessageInfo

type SetWriteDeadlineRequest struct {
	Time                 []byte   `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetWriteDeadlineRequest) Reset()         { *m = SetWriteDeadlineRequest{} }
func (m *SetWriteDeadlineRequest) String() string { return proto.CompactTextString(m) }
func (*SetWriteDeadlineRequest) ProtoMessage()    {}
func (*SetWriteDeadlineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f401a58c1fc7ceef, []int{10}
}

func (m *SetWriteDeadlineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetWriteDeadlineRequest.Unmarshal(m, b)
}
func (m *SetWriteDeadlineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetWriteDeadlineRequest.Marshal(b, m, deterministic)
}
func (m *SetWriteDeadlineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetWriteDeadlineRequest.Merge(m, src)
}
func (m *SetWriteDeadlineRequest) XXX_Size() int {
	return xxx_messageInfo_SetWriteDeadlineRequest.Size(m)
}
func (m *SetWriteDeadlineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetWriteDeadlineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetWriteDeadlineRequest proto.InternalMessageInfo

func (m *SetWriteDeadlineRequest) GetTime() []byte {
	if m != nil {
		return m.Time
	}
	return nil
}

type SetWriteDeadlineResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetWriteDeadlineResponse) Reset()         { *m = SetWriteDeadlineResponse{} }
func (m *SetWriteDeadlineResponse) String() string { return proto.CompactTextString(m) }
func (*SetWriteDeadlineResponse) ProtoMessage()    {}
func (*SetWriteDeadlineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f401a58c1fc7ceef, []int{11}
}

func (m *SetWriteDeadlineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetWriteDeadlineResponse.Unmarshal(m, b)
}
func (m *SetWriteDeadlineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetWriteDeadlineResponse.Marshal(b, m, deterministic)
}
func (m *SetWriteDeadlineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetWriteDeadlineResponse.Merge(m, src)
}
func (m *SetWriteDeadlineResponse) XXX_Size() int {
	return xxx_messageInfo_SetWriteDeadlineResponse.Size(m)
}
func (m *SetWriteDeadlineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetWriteDeadlineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetWriteDeadlineResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ReadRequest)(nil), "proto.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "proto.ReadResponse")
	proto.RegisterType((*WriteRequest)(nil), "proto.WriteRequest")
	proto.RegisterType((*WriteResponse)(nil), "proto.WriteResponse")
	proto.RegisterType((*CloseRequest)(nil), "proto.CloseRequest")
	proto.RegisterType((*CloseResponse)(nil), "proto.CloseResponse")
	proto.RegisterType((*SetDeadlineRequest)(nil), "proto.SetDeadlineRequest")
	proto.RegisterType((*SetDeadlineResponse)(nil), "proto.SetDeadlineResponse")
	proto.RegisterType((*SetReadDeadlineRequest)(nil), "proto.SetReadDeadlineRequest")
	proto.RegisterType((*SetReadDeadlineResponse)(nil), "proto.SetReadDeadlineResponse")
	proto.RegisterType((*SetWriteDeadlineRequest)(nil), "proto.SetWriteDeadlineRequest")
	proto.RegisterType((*SetWriteDeadlineResponse)(nil), "proto.SetWriteDeadlineResponse")
}

func init() { proto.RegisterFile("conn.proto", fileDescriptor_f401a58c1fc7ceef) }

var fileDescriptor_f401a58c1fc7ceef = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xd1, 0x4e, 0xea, 0x40,
	0x10, 0x4d, 0x2f, 0x2d, 0x5c, 0x87, 0x22, 0x66, 0x40, 0x2c, 0x9b, 0xa8, 0x64, 0x13, 0x93, 0x3e,
	0x28, 0x26, 0xf8, 0x09, 0xf0, 0x01, 0x66, 0x79, 0xe0, 0xb9, 0xda, 0x89, 0x92, 0xd4, 0x5d, 0x6c,
	0xd7, 0x07, 0xff, 0xc1, 0x8f, 0x36, 0x6c, 0xb7, 0xa5, 0x05, 0x9a, 0xe0, 0x53, 0xf7, 0xec, 0x39,
	0x73, 0x66, 0x7a, 0x66, 0x01, 0x5e, 0x95, 0x94, 0xd3, 0x4d, 0xaa, 0xb4, 0x42, 0xcf, 0x7c, 0xf8,
	0x1d, 0x74, 0x05, 0x45, 0xb1, 0xa0, 0xcf, 0x2f, 0xca, 0x34, 0x8e, 0xa0, 0x9d, 0x90, 0x7c, 0xd3,
	0xef, 0x81, 0x33, 0x71, 0x42, 0x4f, 0x58, 0xc4, 0x05, 0xf8, 0xb9, 0x2c, 0xdb, 0x28, 0x99, 0x11,
	0x22, 0xb8, 0x29, 0x45, 0xb1, 0x51, 0xf9, 0xc2, 0x9c, 0x71, 0x08, 0x1e, 0xa5, 0xa9, 0x4a, 0x83,
	0x7f, 0x13, 0x27, 0x3c, 0x13, 0x39, 0xc0, 0x00, 0x3a, 0xe6, 0x40, 0x71, 0xd0, 0x9a, 0x38, 0xe1,
	0x7f, 0x51, 0x40, 0x1e, 0x82, 0xbf, 0x4a, 0xd7, 0x9a, 0x8a, 0xde, 0x01, 0x74, 0x36, 0xd1, 0x77,
	0xa2, 0x4a, 0xdb, 0x02, 0xf2, 0x15, 0xf4, 0xac, 0xd2, 0xb6, 0x6f, 0x18, 0xf3, 0xcf, 0x23, 0x9c,
	0x83, 0x3f, 0x4f, 0x54, 0x56, 0x8c, 0xc0, 0xfb, 0xd0, 0xb3, 0x38, 0x6f, 0xc4, 0x43, 0xc0, 0x25,
	0xe9, 0x05, 0x45, 0x71, 0xb2, 0x96, 0xe5, 0xa4, 0x08, 0xae, 0x5e, 0x7f, 0x50, 0xf1, 0xf7, 0xdb,
	0x33, 0xbf, 0x84, 0x41, 0x4d, 0x69, 0x0d, 0xee, 0x61, 0xb4, 0x24, 0xbd, 0xcd, 0xee, 0x14, 0x93,
	0x31, 0x5c, 0x1d, 0xa8, 0xad, 0xd1, 0x83, 0xa1, 0x4c, 0x0c, 0xa7, 0x38, 0x31, 0x08, 0x0e, 0xe5,
	0xb9, 0xd5, 0xec, 0xa7, 0x05, 0xee, 0x5c, 0x49, 0x89, 0x8f, 0xe0, 0x6e, 0x7b, 0x21, 0xe6, 0x6f,
	0x62, 0x5a, 0x79, 0x09, 0x6c, 0x50, 0xbb, 0xb3, 0xb9, 0xcf, 0xc0, 0x33, 0x96, 0x58, 0xb0, 0xd5,
	0x05, 0xb2, 0x61, 0xfd, 0x72, 0x57, 0x63, 0x32, 0x2d, 0x6b, 0xaa, 0x89, 0x97, 0x35, 0xb5, 0xd8,
	0x71, 0x01, 0xdd, 0x4a, 0x98, 0x38, 0xb6, 0xa2, 0xc3, 0x55, 0x30, 0x76, 0x8c, 0xb2, 0x2e, 0xcf,
	0xd0, 0xdf, 0x4b, 0x13, 0xaf, 0x77, 0xf2, 0x23, 0x3b, 0x61, 0x37, 0x4d, 0xb4, 0x75, 0x5c, 0xc2,
	0xc5, 0x7e, 0xaa, 0x58, 0xa9, 0x39, 0xb6, 0x1d, 0x76, 0xdb, 0xc8, 0xe7, 0xa6, 0x2f, 0x6d, 0xc3,
	0x3f, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x73, 0x4d, 0x43, 0x9e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConnClient is the client API for Conn service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConnClient interface {
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error)
	Close(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*CloseResponse, error)
	SetDeadline(ctx context.Context, in *SetDeadlineRequest, opts ...grpc.CallOption) (*SetDeadlineResponse, error)
	SetReadDeadline(ctx context.Context, in *SetReadDeadlineRequest, opts ...grpc.CallOption) (*SetReadDeadlineResponse, error)
	SetWriteDeadline(ctx context.Context, in *SetWriteDeadlineRequest, opts ...grpc.CallOption) (*SetWriteDeadlineResponse, error)
}

type connClient struct {
	cc grpc.ClientConnInterface
}

func NewConnClient(cc grpc.ClientConnInterface) ConnClient {
	return &connClient{cc}
}

func (c *connClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/proto.Conn/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := c.cc.Invoke(ctx, "/proto.Conn/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connClient) Close(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*CloseResponse, error) {
	out := new(CloseResponse)
	err := c.cc.Invoke(ctx, "/proto.Conn/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connClient) SetDeadline(ctx context.Context, in *SetDeadlineRequest, opts ...grpc.CallOption) (*SetDeadlineResponse, error) {
	out := new(SetDeadlineResponse)
	err := c.cc.Invoke(ctx, "/proto.Conn/SetDeadline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connClient) SetReadDeadline(ctx context.Context, in *SetReadDeadlineRequest, opts ...grpc.CallOption) (*SetReadDeadlineResponse, error) {
	out := new(SetReadDeadlineResponse)
	err := c.cc.Invoke(ctx, "/proto.Conn/SetReadDeadline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connClient) SetWriteDeadline(ctx context.Context, in *SetWriteDeadlineRequest, opts ...grpc.CallOption) (*SetWriteDeadlineResponse, error) {
	out := new(SetWriteDeadlineResponse)
	err := c.cc.Invoke(ctx, "/proto.Conn/SetWriteDeadline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnServer is the server API for Conn service.
type ConnServer interface {
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	Write(context.Context, *WriteRequest) (*WriteResponse, error)
	Close(context.Context, *CloseRequest) (*CloseResponse, error)
	SetDeadline(context.Context, *SetDeadlineRequest) (*SetDeadlineResponse, error)
	SetReadDeadline(context.Context, *SetReadDeadlineRequest) (*SetReadDeadlineResponse, error)
	SetWriteDeadline(context.Context, *SetWriteDeadlineRequest) (*SetWriteDeadlineResponse, error)
}

// UnimplementedConnServer can be embedded to have forward compatible implementations.
type UnimplementedConnServer struct {
}

func (*UnimplementedConnServer) Read(ctx context.Context, req *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedConnServer) Write(ctx context.Context, req *WriteRequest) (*WriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (*UnimplementedConnServer) Close(ctx context.Context, req *CloseRequest) (*CloseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}
func (*UnimplementedConnServer) SetDeadline(ctx context.Context, req *SetDeadlineRequest) (*SetDeadlineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDeadline not implemented")
}
func (*UnimplementedConnServer) SetReadDeadline(ctx context.Context, req *SetReadDeadlineRequest) (*SetReadDeadlineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetReadDeadline not implemented")
}
func (*UnimplementedConnServer) SetWriteDeadline(ctx context.Context, req *SetWriteDeadlineRequest) (*SetWriteDeadlineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetWriteDeadline not implemented")
}

func RegisterConnServer(s *grpc.Server, srv ConnServer) {
	s.RegisterService(&_Conn_serviceDesc, srv)
}

func _Conn_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Conn/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conn_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Conn/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnServer).Write(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conn_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Conn/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnServer).Close(ctx, req.(*CloseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conn_SetDeadline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDeadlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnServer).SetDeadline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Conn/SetDeadline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnServer).SetDeadline(ctx, req.(*SetDeadlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conn_SetReadDeadline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetReadDeadlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnServer).SetReadDeadline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Conn/SetReadDeadline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnServer).SetReadDeadline(ctx, req.(*SetReadDeadlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conn_SetWriteDeadline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetWriteDeadlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnServer).SetWriteDeadline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Conn/SetWriteDeadline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnServer).SetWriteDeadline(ctx, req.(*SetWriteDeadlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Conn_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Conn",
	HandlerType: (*ConnServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _Conn_Read_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _Conn_Write_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _Conn_Close_Handler,
		},
		{
			MethodName: "SetDeadline",
			Handler:    _Conn_SetDeadline_Handler,
		},
		{
			MethodName: "SetReadDeadline",
			Handler:    _Conn_SetReadDeadline_Handler,
		},
		{
			MethodName: "SetWriteDeadline",
			Handler:    _Conn_SetWriteDeadline_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "conn.proto",
}