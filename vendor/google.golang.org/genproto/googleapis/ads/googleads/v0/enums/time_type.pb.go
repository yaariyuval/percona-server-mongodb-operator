// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/time_type.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

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

// The possible time types used by certain resources as an alternative to
// absolute timestamps.
type TimeTypeEnum_TimeType int32

const (
	// Not specified.
	TimeTypeEnum_UNSPECIFIED TimeTypeEnum_TimeType = 0
	// Used for return value only. Represents value unknown in this version.
	TimeTypeEnum_UNKNOWN TimeTypeEnum_TimeType = 1
	// As soon as possible.
	TimeTypeEnum_NOW TimeTypeEnum_TimeType = 2
	// An infinite point in the future.
	TimeTypeEnum_FOREVER TimeTypeEnum_TimeType = 3
)

var TimeTypeEnum_TimeType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "NOW",
	3: "FOREVER",
}
var TimeTypeEnum_TimeType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"NOW":         2,
	"FOREVER":     3,
}

func (x TimeTypeEnum_TimeType) String() string {
	return proto.EnumName(TimeTypeEnum_TimeType_name, int32(x))
}
func (TimeTypeEnum_TimeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_time_type_4019044a86ac0043, []int{0, 0}
}

// Message describing time types.
type TimeTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeTypeEnum) Reset()         { *m = TimeTypeEnum{} }
func (m *TimeTypeEnum) String() string { return proto.CompactTextString(m) }
func (*TimeTypeEnum) ProtoMessage()    {}
func (*TimeTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_time_type_4019044a86ac0043, []int{0}
}
func (m *TimeTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeTypeEnum.Unmarshal(m, b)
}
func (m *TimeTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeTypeEnum.Marshal(b, m, deterministic)
}
func (dst *TimeTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeTypeEnum.Merge(dst, src)
}
func (m *TimeTypeEnum) XXX_Size() int {
	return xxx_messageInfo_TimeTypeEnum.Size(m)
}
func (m *TimeTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_TimeTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TimeTypeEnum)(nil), "google.ads.googleads.v0.enums.TimeTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.TimeTypeEnum_TimeType", TimeTypeEnum_TimeType_name, TimeTypeEnum_TimeType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/time_type.proto", fileDescriptor_time_type_4019044a86ac0043)
}

var fileDescriptor_time_type_4019044a86ac0043 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4d, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc,
	0xd2, 0xdc, 0x62, 0xfd, 0x92, 0xcc, 0xdc, 0xd4, 0xf8, 0x92, 0xca, 0x82, 0x54, 0xbd, 0x82, 0xa2,
	0xfc, 0x92, 0x7c, 0x21, 0x59, 0x88, 0x1a, 0xbd, 0xc4, 0x94, 0x62, 0x3d, 0xb8, 0x72, 0xbd, 0x32,
	0x03, 0x3d, 0xb0, 0x72, 0x25, 0x3f, 0x2e, 0x9e, 0x90, 0xcc, 0xdc, 0xd4, 0x90, 0xca, 0x82, 0x54,
	0xd7, 0xbc, 0xd2, 0x5c, 0x25, 0x3b, 0x2e, 0x0e, 0x18, 0x5f, 0x88, 0x9f, 0x8b, 0x3b, 0xd4, 0x2f,
	0x38, 0xc0, 0xd5, 0xd9, 0xd3, 0xcd, 0xd3, 0xd5, 0x45, 0x80, 0x41, 0x88, 0x9b, 0x8b, 0x3d, 0xd4,
	0xcf, 0xdb, 0xcf, 0x3f, 0xdc, 0x4f, 0x80, 0x51, 0x88, 0x9d, 0x8b, 0xd9, 0xcf, 0x3f, 0x5c, 0x80,
	0x09, 0x24, 0xea, 0xe6, 0x1f, 0xe4, 0x1a, 0xe6, 0x1a, 0x24, 0xc0, 0xec, 0xf4, 0x88, 0x91, 0x4b,
	0x31, 0x39, 0x3f, 0x57, 0x0f, 0xaf, 0xad, 0x4e, 0xbc, 0x30, 0x3b, 0x02, 0x40, 0x6e, 0x0c, 0x60,
	0x8c, 0x72, 0x82, 0xaa, 0x4f, 0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0xcb, 0x2f, 0x4a, 0xd7, 0x4f,
	0x4f, 0xcd, 0x03, 0xfb, 0x00, 0xe6, 0xc9, 0x82, 0xcc, 0x62, 0x1c, 0x7e, 0xb6, 0x06, 0x93, 0x8b,
	0x98, 0x98, 0xdd, 0x1d, 0x1d, 0x57, 0x31, 0xc9, 0xba, 0x43, 0x8c, 0x72, 0x4c, 0x29, 0xd6, 0x83,
	0x30, 0x41, 0xac, 0x30, 0x03, 0x3d, 0x90, 0xff, 0x8a, 0x4f, 0xc1, 0xe4, 0x63, 0x1c, 0x53, 0x8a,
	0x63, 0xe0, 0xf2, 0x31, 0x61, 0x06, 0x31, 0x60, 0xf9, 0x57, 0x4c, 0x8a, 0x10, 0x41, 0x2b, 0x2b,
	0xc7, 0x94, 0x62, 0x2b, 0x2b, 0xb8, 0x0a, 0x2b, 0xab, 0x30, 0x03, 0x2b, 0x2b, 0xb0, 0x9a, 0x24,
	0x36, 0xb0, 0xc3, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x62, 0x1a, 0x32, 0x58, 0x8b, 0x01,
	0x00, 0x00,
}
