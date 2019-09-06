// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: unixfs.proto

package unixfs_pb

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

type Data_DataType int32

const (
	Data_Raw       Data_DataType = 0
	Data_Directory Data_DataType = 1
	Data_File      Data_DataType = 2
	Data_Metadata  Data_DataType = 3
	Data_Symlink   Data_DataType = 4
	Data_HAMTShard Data_DataType = 5
)

var Data_DataType_name = map[int32]string{
	0: "Raw",
	1: "Directory",
	2: "File",
	3: "Metadata",
	4: "Symlink",
	5: "HAMTShard",
}

var Data_DataType_value = map[string]int32{
	"Raw":       0,
	"Directory": 1,
	"File":      2,
	"Metadata":  3,
	"Symlink":   4,
	"HAMTShard": 5,
}

func (x Data_DataType) Enum() *Data_DataType {
	p := new(Data_DataType)
	*p = x
	return p
}

func (x Data_DataType) String() string {
	return proto.EnumName(Data_DataType_name, int32(x))
}

func (x *Data_DataType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Data_DataType_value, data, "Data_DataType")
	if err != nil {
		return err
	}
	*x = Data_DataType(value)
	return nil
}

func (Data_DataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e2fd76cc44dfc7c3, []int{0, 0}
}

type Data struct {
	Type                 *Data_DataType `protobuf:"varint,1,req,name=Type,enum=unixfs.pb.Data_DataType" json:"Type,omitempty"`
	Data                 []byte         `protobuf:"bytes,2,opt,name=Data" json:"Data,omitempty"`
	Filesize             *uint64        `protobuf:"varint,3,opt,name=filesize" json:"filesize,omitempty"`
	Blocksizes           []uint64       `protobuf:"varint,4,rep,name=blocksizes" json:"blocksizes,omitempty"`
	HashType             *uint64        `protobuf:"varint,5,opt,name=hashType" json:"hashType,omitempty"`
	Fanout               *uint64        `protobuf:"varint,6,opt,name=fanout" json:"fanout,omitempty"`
	TokenMeta            *bool          `protobuf:"varint,7,opt,name=tokenMeta" json:"tokenMeta,omitempty"`
	MetaOffset           *uint64        `protobuf:"varint,9,opt,name=metaOffset" json:"metaOffset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2fd76cc44dfc7c3, []int{0}
}
func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetType() Data_DataType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Data_Raw
}

func (m *Data) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Data) GetFilesize() uint64 {
	if m != nil && m.Filesize != nil {
		return *m.Filesize
	}
	return 0
}

func (m *Data) GetBlocksizes() []uint64 {
	if m != nil {
		return m.Blocksizes
	}
	return nil
}

func (m *Data) GetHashType() uint64 {
	if m != nil && m.HashType != nil {
		return *m.HashType
	}
	return 0
}

func (m *Data) GetFanout() uint64 {
	if m != nil && m.Fanout != nil {
		return *m.Fanout
	}
	return 0
}

func (m *Data) GetTokenMeta() bool {
	if m != nil && m.TokenMeta != nil {
		return *m.TokenMeta
	}
	return false
}

func (m *Data) GetMetaOffset() uint64 {
	if m != nil && m.MetaOffset != nil {
		return *m.MetaOffset
	}
	return 0
}

type Metadata struct {
	MimeType             *string  `protobuf:"bytes,1,opt,name=MimeType" json:"MimeType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2fd76cc44dfc7c3, []int{1}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetMimeType() string {
	if m != nil && m.MimeType != nil {
		return *m.MimeType
	}
	return ""
}

func init() {
	proto.RegisterEnum("unixfs.pb.Data_DataType", Data_DataType_name, Data_DataType_value)
	proto.RegisterType((*Data)(nil), "unixfs.pb.Data")
	proto.RegisterType((*Metadata)(nil), "unixfs.pb.Metadata")
}

func init() { proto.RegisterFile("unixfs.proto", fileDescriptor_e2fd76cc44dfc7c3) }

var fileDescriptor_e2fd76cc44dfc7c3 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xcb, 0x6a, 0xeb, 0x30,
	0x10, 0x86, 0x8f, 0x6c, 0x25, 0xb1, 0xe7, 0xa4, 0x45, 0xcc, 0xa2, 0x88, 0x52, 0x8a, 0xc9, 0xa2,
	0x68, 0x51, 0xb2, 0xe8, 0x1b, 0x14, 0x42, 0xe9, 0x26, 0x14, 0x94, 0xd0, 0xbd, 0x92, 0xc8, 0x58,
	0xf8, 0xa2, 0x60, 0x2b, 0xb4, 0xee, 0x43, 0xf6, 0x99, 0x8a, 0xec, 0xd8, 0xc9, 0x46, 0xf0, 0xfd,
	0x97, 0x91, 0x46, 0x30, 0x3f, 0x55, 0xe6, 0x3b, 0x6d, 0x96, 0xc7, 0xda, 0x3a, 0x8b, 0xf1, 0x40,
	0xbb, 0xc5, 0x6f, 0x00, 0x74, 0xa5, 0x9c, 0xc2, 0x67, 0xa0, 0xdb, 0xf6, 0xa8, 0x39, 0x49, 0x02,
	0x71, 0xfb, 0xc2, 0x97, 0x63, 0x64, 0xe9, 0xed, 0xee, 0xf0, 0xbe, 0xec, 0x52, 0x88, 0x7d, 0x8b,
	0x07, 0x09, 0x11, 0x73, 0xd9, 0x4f, 0xb8, 0x87, 0x28, 0x35, 0x85, 0x6e, 0xcc, 0x8f, 0xe6, 0x61,
	0x42, 0x04, 0x95, 0x23, 0xe3, 0x23, 0xc0, 0xae, 0xb0, 0xfb, 0xdc, 0x43, 0xc3, 0x69, 0x12, 0x0a,
	0x2a, 0xaf, 0x14, 0xdf, 0xcd, 0x54, 0x93, 0x75, 0x2f, 0x98, 0xf4, 0xdd, 0x81, 0xf1, 0x0e, 0xa6,
	0xa9, 0xaa, 0xec, 0xc9, 0xf1, 0x69, 0xe7, 0x9c, 0x09, 0x1f, 0x20, 0x76, 0x36, 0xd7, 0xd5, 0x5a,
	0x3b, 0xc5, 0x67, 0x09, 0x11, 0x91, 0xbc, 0x08, 0xfe, 0xc6, 0x52, 0x3b, 0xf5, 0x91, 0xa6, 0x8d,
	0x76, 0x3c, 0xee, 0x9a, 0x57, 0xca, 0xe2, 0x13, 0xa2, 0x61, 0x27, 0x9c, 0x41, 0x28, 0xd5, 0x17,
	0xfb, 0x87, 0x37, 0x10, 0xaf, 0x4c, 0xad, 0xf7, 0xce, 0xd6, 0x2d, 0x23, 0x18, 0x01, 0x7d, 0x33,
	0x85, 0x66, 0x01, 0xce, 0x21, 0xf2, 0x53, 0x0f, 0xca, 0x29, 0x16, 0xe2, 0x7f, 0x98, 0x6d, 0xda,
	0xb2, 0x30, 0x55, 0xce, 0xa8, 0xef, 0xbc, 0xbf, 0xae, 0xb7, 0x9b, 0x4c, 0xd5, 0x07, 0x36, 0x59,
	0x3c, 0x5d, 0x92, 0x7e, 0xab, 0xb5, 0x29, 0xf5, 0xf9, 0x5f, 0x89, 0x88, 0xe5, 0xc8, 0x7f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xf5, 0xff, 0xd0, 0xf4, 0x92, 0x01, 0x00, 0x00,
}
