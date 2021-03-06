// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package protobuf

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

type ByteArray struct {
	Data                 []byte   `protobuf:"bytes,1,req,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ByteArray) Reset()         { *m = ByteArray{} }
func (m *ByteArray) String() string { return proto.CompactTextString(m) }
func (*ByteArray) ProtoMessage()    {}
func (*ByteArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

func (m *ByteArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ByteArray.Unmarshal(m, b)
}
func (m *ByteArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ByteArray.Marshal(b, m, deterministic)
}
func (m *ByteArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ByteArray.Merge(m, src)
}
func (m *ByteArray) XXX_Size() int {
	return xxx_messageInfo_ByteArray.Size(m)
}
func (m *ByteArray) XXX_DiscardUnknown() {
	xxx_messageInfo_ByteArray.DiscardUnknown(m)
}

var xxx_messageInfo_ByteArray proto.InternalMessageInfo

func (m *ByteArray) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ProtoTransactionCoreInfo struct {
	Version              *uint32    `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Nonce                *uint64    `protobuf:"varint,2,opt,name=nonce" json:"nonce,omitempty"`
	Toaddr               []byte     `protobuf:"bytes,3,opt,name=toaddr" json:"toaddr,omitempty"`
	Senderpubkey         *ByteArray `protobuf:"bytes,4,opt,name=senderpubkey" json:"senderpubkey,omitempty"`
	Amount               *ByteArray `protobuf:"bytes,5,opt,name=amount" json:"amount,omitempty"`
	Gasprice             *ByteArray `protobuf:"bytes,6,opt,name=gasprice" json:"gasprice,omitempty"`
	Gaslimit             *uint64    `protobuf:"varint,7,opt,name=gaslimit" json:"gaslimit,omitempty"`
	Code                 []byte     `protobuf:"bytes,8,opt,name=code" json:"code,omitempty"`
	Data                 []byte     `protobuf:"bytes,9,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ProtoTransactionCoreInfo) Reset()         { *m = ProtoTransactionCoreInfo{} }
func (m *ProtoTransactionCoreInfo) String() string { return proto.CompactTextString(m) }
func (*ProtoTransactionCoreInfo) ProtoMessage()    {}
func (*ProtoTransactionCoreInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{1}
}

func (m *ProtoTransactionCoreInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoTransactionCoreInfo.Unmarshal(m, b)
}
func (m *ProtoTransactionCoreInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoTransactionCoreInfo.Marshal(b, m, deterministic)
}
func (m *ProtoTransactionCoreInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoTransactionCoreInfo.Merge(m, src)
}
func (m *ProtoTransactionCoreInfo) XXX_Size() int {
	return xxx_messageInfo_ProtoTransactionCoreInfo.Size(m)
}
func (m *ProtoTransactionCoreInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoTransactionCoreInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoTransactionCoreInfo proto.InternalMessageInfo

func (m *ProtoTransactionCoreInfo) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *ProtoTransactionCoreInfo) GetNonce() uint64 {
	if m != nil && m.Nonce != nil {
		return *m.Nonce
	}
	return 0
}

func (m *ProtoTransactionCoreInfo) GetToaddr() []byte {
	if m != nil {
		return m.Toaddr
	}
	return nil
}

func (m *ProtoTransactionCoreInfo) GetSenderpubkey() *ByteArray {
	if m != nil {
		return m.Senderpubkey
	}
	return nil
}

func (m *ProtoTransactionCoreInfo) GetAmount() *ByteArray {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *ProtoTransactionCoreInfo) GetGasprice() *ByteArray {
	if m != nil {
		return m.Gasprice
	}
	return nil
}

func (m *ProtoTransactionCoreInfo) GetGaslimit() uint64 {
	if m != nil && m.Gaslimit != nil {
		return *m.Gaslimit
	}
	return 0
}

func (m *ProtoTransactionCoreInfo) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *ProtoTransactionCoreInfo) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ProtoTransaction struct {
	Tranid               []byte                    `protobuf:"bytes,1,opt,name=tranid" json:"tranid,omitempty"`
	Info                 *ProtoTransactionCoreInfo `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
	Signature            *ByteArray                `protobuf:"bytes,3,opt,name=signature" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ProtoTransaction) Reset()         { *m = ProtoTransaction{} }
func (m *ProtoTransaction) String() string { return proto.CompactTextString(m) }
func (*ProtoTransaction) ProtoMessage()    {}
func (*ProtoTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{2}
}

func (m *ProtoTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoTransaction.Unmarshal(m, b)
}
func (m *ProtoTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoTransaction.Marshal(b, m, deterministic)
}
func (m *ProtoTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoTransaction.Merge(m, src)
}
func (m *ProtoTransaction) XXX_Size() int {
	return xxx_messageInfo_ProtoTransaction.Size(m)
}
func (m *ProtoTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoTransaction proto.InternalMessageInfo

func (m *ProtoTransaction) GetTranid() []byte {
	if m != nil {
		return m.Tranid
	}
	return nil
}

func (m *ProtoTransaction) GetInfo() *ProtoTransactionCoreInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ProtoTransaction) GetSignature() *ByteArray {
	if m != nil {
		return m.Signature
	}
	return nil
}

type ProtoTransactionReceipt struct {
	Receipt              []byte   `protobuf:"bytes,1,opt,name=receipt" json:"receipt,omitempty"`
	Cumgas               *uint64  `protobuf:"varint,2,opt,name=cumgas" json:"cumgas,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtoTransactionReceipt) Reset()         { *m = ProtoTransactionReceipt{} }
func (m *ProtoTransactionReceipt) String() string { return proto.CompactTextString(m) }
func (*ProtoTransactionReceipt) ProtoMessage()    {}
func (*ProtoTransactionReceipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{3}
}

func (m *ProtoTransactionReceipt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoTransactionReceipt.Unmarshal(m, b)
}
func (m *ProtoTransactionReceipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoTransactionReceipt.Marshal(b, m, deterministic)
}
func (m *ProtoTransactionReceipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoTransactionReceipt.Merge(m, src)
}
func (m *ProtoTransactionReceipt) XXX_Size() int {
	return xxx_messageInfo_ProtoTransactionReceipt.Size(m)
}
func (m *ProtoTransactionReceipt) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoTransactionReceipt.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoTransactionReceipt proto.InternalMessageInfo

func (m *ProtoTransactionReceipt) GetReceipt() []byte {
	if m != nil {
		return m.Receipt
	}
	return nil
}

func (m *ProtoTransactionReceipt) GetCumgas() uint64 {
	if m != nil && m.Cumgas != nil {
		return *m.Cumgas
	}
	return 0
}

type ProtoTransactionWithReceipt struct {
	Transaction          *ProtoTransaction        `protobuf:"bytes,1,opt,name=transaction" json:"transaction,omitempty"`
	Receipt              *ProtoTransactionReceipt `protobuf:"bytes,2,opt,name=receipt" json:"receipt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ProtoTransactionWithReceipt) Reset()         { *m = ProtoTransactionWithReceipt{} }
func (m *ProtoTransactionWithReceipt) String() string { return proto.CompactTextString(m) }
func (*ProtoTransactionWithReceipt) ProtoMessage()    {}
func (*ProtoTransactionWithReceipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{4}
}

func (m *ProtoTransactionWithReceipt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoTransactionWithReceipt.Unmarshal(m, b)
}
func (m *ProtoTransactionWithReceipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoTransactionWithReceipt.Marshal(b, m, deterministic)
}
func (m *ProtoTransactionWithReceipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoTransactionWithReceipt.Merge(m, src)
}
func (m *ProtoTransactionWithReceipt) XXX_Size() int {
	return xxx_messageInfo_ProtoTransactionWithReceipt.Size(m)
}
func (m *ProtoTransactionWithReceipt) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoTransactionWithReceipt.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoTransactionWithReceipt proto.InternalMessageInfo

func (m *ProtoTransactionWithReceipt) GetTransaction() *ProtoTransaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *ProtoTransactionWithReceipt) GetReceipt() *ProtoTransactionReceipt {
	if m != nil {
		return m.Receipt
	}
	return nil
}

func init() {
	proto.RegisterType((*ByteArray)(nil), "ZilliqaMessage.ByteArray")
	proto.RegisterType((*ProtoTransactionCoreInfo)(nil), "ZilliqaMessage.ProtoTransactionCoreInfo")
	proto.RegisterType((*ProtoTransaction)(nil), "ZilliqaMessage.ProtoTransaction")
	proto.RegisterType((*ProtoTransactionReceipt)(nil), "ZilliqaMessage.ProtoTransactionReceipt")
	proto.RegisterType((*ProtoTransactionWithReceipt)(nil), "ZilliqaMessage.ProtoTransactionWithReceipt")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5) }

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xdf, 0x6a, 0xdb, 0x30,
	0x14, 0xc6, 0x51, 0xe6, 0xfc, 0xf1, 0x71, 0x16, 0x86, 0x18, 0x9b, 0xb6, 0x5d, 0xcc, 0xf8, 0x66,
	0xbe, 0x0a, 0x2c, 0x30, 0x76, 0xb3, 0x5d, 0x24, 0xbb, 0x1a, 0x63, 0x50, 0x44, 0xa1, 0x90, 0x3b,
	0xd5, 0x56, 0x5c, 0x51, 0x5b, 0x72, 0x25, 0xb9, 0x90, 0xbe, 0x4d, 0xa1, 0xef, 0xd3, 0x57, 0x2a,
	0x96, 0x1d, 0xa7, 0x31, 0xb4, 0xbe, 0x3b, 0x9f, 0x74, 0xbe, 0xc3, 0x77, 0x7e, 0x07, 0x16, 0x05,
	0x37, 0x86, 0x65, 0xdc, 0x2c, 0x4b, 0xad, 0xac, 0xc2, 0x8b, 0xad, 0xc8, 0x73, 0x71, 0xc3, 0xfe,
	0x37, 0xcf, 0xd1, 0x57, 0xf0, 0x37, 0x7b, 0xcb, 0xd7, 0x5a, 0xb3, 0x3d, 0xc6, 0xe0, 0xa5, 0xcc,
	0x32, 0x82, 0xc2, 0x51, 0x3c, 0xa7, 0xae, 0x8e, 0x1e, 0x47, 0x40, 0xce, 0x6a, 0xeb, 0xb9, 0x66,
	0xd2, 0xb0, 0xc4, 0x0a, 0x25, 0xff, 0x28, 0xcd, 0xff, 0xca, 0x9d, 0xc2, 0x04, 0xa6, 0xb7, 0x5c,
	0x1b, 0xa1, 0x24, 0x41, 0x21, 0x8a, 0xdf, 0xd2, 0x83, 0xc4, 0xef, 0x61, 0x2c, 0x95, 0x4c, 0x38,
	0x19, 0x85, 0x28, 0xf6, 0x68, 0x23, 0xf0, 0x07, 0x98, 0x58, 0xc5, 0xd2, 0x54, 0x93, 0x37, 0x21,
	0x8a, 0xe7, 0xb4, 0x55, 0xf8, 0x37, 0xcc, 0x0d, 0x97, 0x29, 0xd7, 0x65, 0x75, 0x79, 0xcd, 0xf7,
	0xc4, 0x0b, 0x51, 0x1c, 0xac, 0x3e, 0x2d, 0x4f, 0xc3, 0x2e, 0xbb, 0xa4, 0xf4, 0xa4, 0x1d, 0x7f,
	0x87, 0x09, 0x2b, 0x54, 0x25, 0x2d, 0x19, 0x0f, 0x19, 0xdb, 0x46, 0xfc, 0x03, 0x66, 0x19, 0x33,
	0xa5, 0x16, 0x09, 0x27, 0x93, 0x21, 0x53, 0xd7, 0x8a, 0x3f, 0x3b, 0x5b, 0x2e, 0x0a, 0x61, 0xc9,
	0xd4, 0x6d, 0xd6, 0xe9, 0x9a, 0x5e, 0xa2, 0x52, 0x4e, 0x66, 0x6e, 0x35, 0x57, 0x77, 0x44, 0xfd,
	0xe6, 0xcd, 0x11, 0xbd, 0x47, 0xf0, 0xae, 0x4f, 0xd4, 0x91, 0xd1, 0x4c, 0x8a, 0xd4, 0x81, 0xac,
	0xc9, 0x38, 0x85, 0x7f, 0x81, 0x27, 0xe4, 0x4e, 0x39, 0x8c, 0xc1, 0x2a, 0xee, 0x67, 0x7c, 0xe9,
	0x32, 0xd4, 0xb9, 0xf0, 0x4f, 0xf0, 0x8d, 0xc8, 0x24, 0xb3, 0x95, 0xe6, 0x0e, 0xf9, 0xab, 0x6b,
	0x1e, 0x7b, 0xa3, 0x7f, 0xf0, 0xb1, 0x3f, 0x9a, 0xf2, 0x84, 0x8b, 0xd2, 0xd6, 0x37, 0xd7, 0x4d,
	0xd9, 0x46, 0x3d, 0xc8, 0x7a, 0x87, 0xa4, 0x2a, 0x32, 0x66, 0xda, 0xa3, 0xb7, 0x2a, 0x7a, 0x40,
	0xf0, 0xa5, 0x3f, 0xed, 0x42, 0xd8, 0xab, 0xc3, 0xc4, 0x0d, 0x04, 0xf6, 0xf8, 0xe3, 0xa6, 0x06,
	0xab, 0x70, 0x68, 0x55, 0xfa, 0xdc, 0x84, 0xd7, 0xc7, 0x54, 0x0d, 0xaa, 0x6f, 0x83, 0xfe, 0xa6,
	0xbd, 0x8b, 0xbf, 0xf1, 0xb7, 0xd3, 0x3b, 0x91, 0xe7, 0x9c, 0xc9, 0xa7, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x85, 0x7f, 0x42, 0xb7, 0x36, 0x03, 0x00, 0x00,
}
