// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: datarepair.proto

package pb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// InjuredSegment is the queue item used for the data repair queue
type InjuredSegment struct {
	Path                 []byte    `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	LostPieces           []int32   `protobuf:"varint,2,rep,packed,name=lost_pieces,json=lostPieces,proto3" json:"lost_pieces,omitempty"`
	InsertedTime         time.Time `protobuf:"bytes,3,opt,name=inserted_time,json=insertedTime,proto3,stdtime" json:"inserted_time"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *InjuredSegment) Reset()         { *m = InjuredSegment{} }
func (m *InjuredSegment) String() string { return proto.CompactTextString(m) }
func (*InjuredSegment) ProtoMessage()    {}
func (*InjuredSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1b08e6fe9398aa6, []int{0}
}
func (m *InjuredSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InjuredSegment.Unmarshal(m, b)
}
func (m *InjuredSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InjuredSegment.Marshal(b, m, deterministic)
}
func (m *InjuredSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InjuredSegment.Merge(m, src)
}
func (m *InjuredSegment) XXX_Size() int {
	return xxx_messageInfo_InjuredSegment.Size(m)
}
func (m *InjuredSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_InjuredSegment.DiscardUnknown(m)
}

var xxx_messageInfo_InjuredSegment proto.InternalMessageInfo

func (m *InjuredSegment) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *InjuredSegment) GetLostPieces() []int32 {
	if m != nil {
		return m.LostPieces
	}
	return nil
}

func (m *InjuredSegment) GetInsertedTime() time.Time {
	if m != nil {
		return m.InsertedTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*InjuredSegment)(nil), "repair.InjuredSegment")
}

func init() { proto.RegisterFile("datarepair.proto", fileDescriptor_b1b08e6fe9398aa6) }

var fileDescriptor_b1b08e6fe9398aa6 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x49, 0x2c, 0x49,
	0x2c, 0x4a, 0x2d, 0x48, 0xcc, 0x2c, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0,
	0xa4, 0xb8, 0xd2, 0xf3, 0xd3, 0xf3, 0x21, 0x62, 0x52, 0xf2, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9,
	0xfa, 0x60, 0x5e, 0x52, 0x69, 0x9a, 0x7e, 0x49, 0x66, 0x6e, 0x6a, 0x71, 0x49, 0x62, 0x6e, 0x01,
	0x44, 0x81, 0xd2, 0x04, 0x46, 0x2e, 0x3e, 0xcf, 0xbc, 0xac, 0xd2, 0xa2, 0xd4, 0x94, 0xe0, 0xd4,
	0xf4, 0xdc, 0xd4, 0xbc, 0x12, 0x21, 0x21, 0x2e, 0x96, 0x82, 0xc4, 0x92, 0x0c, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0x9e, 0x20, 0x30, 0x5b, 0x48, 0x9e, 0x8b, 0x3b, 0x27, 0xbf, 0xb8, 0x24, 0xbe, 0x20,
	0x33, 0x35, 0x39, 0xb5, 0x58, 0x82, 0x49, 0x81, 0x59, 0x83, 0x35, 0x88, 0x0b, 0x24, 0x14, 0x00,
	0x16, 0x11, 0xf2, 0xe4, 0xe2, 0xcd, 0xcc, 0x2b, 0x4e, 0x2d, 0x2a, 0x49, 0x4d, 0x89, 0x07, 0xd9,
	0x21, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xa5, 0x07, 0x71, 0x80, 0x1e, 0xcc, 0x01, 0x7a,
	0x21, 0x30, 0x07, 0x38, 0x71, 0x9c, 0xb8, 0x27, 0xcf, 0x30, 0xe1, 0xbe, 0x3c, 0x63, 0x10, 0x0f,
	0x4c, 0x2b, 0x48, 0xd2, 0x89, 0x25, 0x8a, 0xa9, 0x20, 0x29, 0x89, 0x0d, 0xac, 0xc3, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x66, 0xb5, 0xcc, 0x1f, 0xe8, 0x00, 0x00, 0x00,
}
