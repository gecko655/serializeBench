// Code generated by protoc-gen-go. DO NOT EDIT.
// source: serializeBench/data/data.proto

package protoData

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

type ResponseTop struct {
	Ts                   int32    `protobuf:"varint,1,opt,name=Ts" json:"Ts,omitempty"`
	Pid                  int32    `protobuf:"varint,2,opt,name=Pid" json:"Pid,omitempty"`
	Rev                  int32    `protobuf:"varint,3,opt,name=Rev" json:"Rev,omitempty"`
	Login                *Login   `protobuf:"bytes,4,opt,name=Login" json:"Login,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseTop) Reset()         { *m = ResponseTop{} }
func (m *ResponseTop) String() string { return proto.CompactTextString(m) }
func (*ResponseTop) ProtoMessage()    {}
func (*ResponseTop) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_98d321a32b16285b, []int{0}
}
func (m *ResponseTop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseTop.Unmarshal(m, b)
}
func (m *ResponseTop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseTop.Marshal(b, m, deterministic)
}
func (dst *ResponseTop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseTop.Merge(dst, src)
}
func (m *ResponseTop) XXX_Size() int {
	return xxx_messageInfo_ResponseTop.Size(m)
}
func (m *ResponseTop) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseTop.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseTop proto.InternalMessageInfo

func (m *ResponseTop) GetTs() int32 {
	if m != nil {
		return m.Ts
	}
	return 0
}

func (m *ResponseTop) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *ResponseTop) GetRev() int32 {
	if m != nil {
		return m.Rev
	}
	return 0
}

func (m *ResponseTop) GetLogin() *Login {
	if m != nil {
		return m.Login
	}
	return nil
}

type Login struct {
	UserStatus           *UserStatus `protobuf:"bytes,1,opt,name=UserStatus" json:"UserStatus,omitempty"`
	UserCardList         []*UserCard `protobuf:"bytes,2,rep,name=UserCardList" json:"UserCardList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Login) Reset()         { *m = Login{} }
func (m *Login) String() string { return proto.CompactTextString(m) }
func (*Login) ProtoMessage()    {}
func (*Login) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_98d321a32b16285b, []int{1}
}
func (m *Login) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Login.Unmarshal(m, b)
}
func (m *Login) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Login.Marshal(b, m, deterministic)
}
func (dst *Login) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Login.Merge(dst, src)
}
func (m *Login) XXX_Size() int {
	return xxx_messageInfo_Login.Size(m)
}
func (m *Login) XXX_DiscardUnknown() {
	xxx_messageInfo_Login.DiscardUnknown(m)
}

var xxx_messageInfo_Login proto.InternalMessageInfo

func (m *Login) GetUserStatus() *UserStatus {
	if m != nil {
		return m.UserStatus
	}
	return nil
}

func (m *Login) GetUserCardList() []*UserCard {
	if m != nil {
		return m.UserCardList
	}
	return nil
}

type UserCard struct {
	CardId               int32    `protobuf:"varint,1,opt,name=CardId" json:"CardId,omitempty"`
	Level                int32    `protobuf:"varint,2,opt,name=Level" json:"Level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserCard) Reset()         { *m = UserCard{} }
func (m *UserCard) String() string { return proto.CompactTextString(m) }
func (*UserCard) ProtoMessage()    {}
func (*UserCard) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_98d321a32b16285b, []int{2}
}
func (m *UserCard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCard.Unmarshal(m, b)
}
func (m *UserCard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCard.Marshal(b, m, deterministic)
}
func (dst *UserCard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCard.Merge(dst, src)
}
func (m *UserCard) XXX_Size() int {
	return xxx_messageInfo_UserCard.Size(m)
}
func (m *UserCard) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCard.DiscardUnknown(m)
}

var xxx_messageInfo_UserCard proto.InternalMessageInfo

func (m *UserCard) GetCardId() int32 {
	if m != nil {
		return m.CardId
	}
	return 0
}

func (m *UserCard) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

type UserStatus struct {
	UserId               int32    `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=UserName" json:"UserName,omitempty"`
	Exp                  int32    `protobuf:"varint,3,opt,name=Exp" json:"Exp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserStatus) Reset()         { *m = UserStatus{} }
func (m *UserStatus) String() string { return proto.CompactTextString(m) }
func (*UserStatus) ProtoMessage()    {}
func (*UserStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_data_98d321a32b16285b, []int{3}
}
func (m *UserStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserStatus.Unmarshal(m, b)
}
func (m *UserStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserStatus.Marshal(b, m, deterministic)
}
func (dst *UserStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserStatus.Merge(dst, src)
}
func (m *UserStatus) XXX_Size() int {
	return xxx_messageInfo_UserStatus.Size(m)
}
func (m *UserStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_UserStatus.DiscardUnknown(m)
}

var xxx_messageInfo_UserStatus proto.InternalMessageInfo

func (m *UserStatus) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserStatus) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserStatus) GetExp() int32 {
	if m != nil {
		return m.Exp
	}
	return 0
}

func init() {
	proto.RegisterType((*ResponseTop)(nil), "protoData.ResponseTop")
	proto.RegisterType((*Login)(nil), "protoData.Login")
	proto.RegisterType((*UserCard)(nil), "protoData.UserCard")
	proto.RegisterType((*UserStatus)(nil), "protoData.UserStatus")
}

func init() {
	proto.RegisterFile("serializeBench/data/data.proto", fileDescriptor_data_98d321a32b16285b)
}

var fileDescriptor_data_98d321a32b16285b = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0x69, 0x6b, 0x17, 0x77, 0x2a, 0xb2, 0xc4, 0x3f, 0x04, 0x0f, 0x52, 0x7a, 0x90, 0x9e,
	0x2a, 0x54, 0x44, 0xcf, 0xea, 0x1e, 0x84, 0x22, 0x12, 0xeb, 0x03, 0x44, 0x3b, 0x68, 0x60, 0x6d,
	0x42, 0x13, 0xeb, 0xe2, 0xd3, 0x4b, 0x92, 0xee, 0xb6, 0xee, 0xa5, 0xd3, 0xdf, 0x37, 0xf3, 0x4d,
	0x66, 0x06, 0xce, 0x35, 0x76, 0x82, 0xaf, 0xc4, 0x2f, 0xde, 0x61, 0xfb, 0xfe, 0x79, 0xd9, 0x70,
	0xc3, 0xdd, 0xa7, 0x50, 0x9d, 0x34, 0x92, 0xcc, 0x5d, 0x78, 0xe0, 0x86, 0x67, 0x02, 0x12, 0x86,
	0x5a, 0xc9, 0x56, 0x63, 0x2d, 0x15, 0x39, 0x84, 0xb0, 0xd6, 0x34, 0x48, 0x83, 0x3c, 0x66, 0x61,
	0xad, 0xc9, 0x02, 0xa2, 0x67, 0xd1, 0xd0, 0xd0, 0x09, 0xf6, 0xd7, 0x2a, 0x0c, 0x7b, 0x1a, 0x79,
	0x85, 0x61, 0x4f, 0x2e, 0x20, 0xae, 0xe4, 0x87, 0x68, 0xe9, 0x5e, 0x1a, 0xe4, 0x49, 0xb9, 0x28,
	0xb6, 0xdd, 0x0b, 0xa7, 0x33, 0x9f, 0xce, 0x7e, 0x86, 0x3a, 0x72, 0x0d, 0xf0, 0xaa, 0xb1, 0x7b,
	0x31, 0xdc, 0x7c, 0xfb, 0xc7, 0x92, 0xf2, 0x64, 0xe2, 0x1a, 0x93, 0x6c, 0x52, 0x48, 0x6e, 0xe0,
	0xc0, 0xd2, 0x3d, 0xef, 0x9a, 0x4a, 0x68, 0x43, 0xc3, 0x34, 0xca, 0x93, 0xf2, 0x68, 0xc7, 0x68,
	0xd3, 0xec, 0x5f, 0x61, 0x76, 0x0b, 0xfb, 0x1b, 0x26, 0xa7, 0x30, 0xb3, 0xf1, 0xb1, 0x19, 0x96,
	0x1c, 0x88, 0x1c, 0x43, 0x5c, 0x61, 0x8f, 0xab, 0x61, 0x55, 0x0f, 0x19, 0x9b, 0x4e, 0x6a, 0xbd,
	0x96, 0x46, 0xaf, 0x27, 0x72, 0xe6, 0xfb, 0x3f, 0xf1, 0x2f, 0x74, 0xf6, 0x39, 0xdb, 0xb2, 0x3d,
	0xd7, 0x72, 0xad, 0x36, 0xe7, 0x5a, 0xae, 0xd5, 0xdb, 0xcc, 0xcd, 0x7b, 0xf5, 0x17, 0x00, 0x00,
	0xff, 0xff, 0x28, 0xdc, 0xc1, 0x57, 0xa5, 0x01, 0x00, 0x00,
}