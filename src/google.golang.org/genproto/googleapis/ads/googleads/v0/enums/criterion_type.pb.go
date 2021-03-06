// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/criterion_type.proto

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

// Enum describing possible criterion types.
type CriterionTypeEnum_CriterionType int32

const (
	// Not specified.
	CriterionTypeEnum_UNSPECIFIED CriterionTypeEnum_CriterionType = 0
	// Used for return value only. Represents value unknown in this version.
	CriterionTypeEnum_UNKNOWN CriterionTypeEnum_CriterionType = 1
	// Keyword. e.g. 'mars cruise'.
	CriterionTypeEnum_KEYWORD CriterionTypeEnum_CriterionType = 2
	// Placement, aka Website. e.g. 'www.flowers4sale.com'
	CriterionTypeEnum_PLACEMENT CriterionTypeEnum_CriterionType = 3
	// Devices to target.
	CriterionTypeEnum_DEVICE CriterionTypeEnum_CriterionType = 6
	// Locations to target.
	CriterionTypeEnum_LOCATION CriterionTypeEnum_CriterionType = 7
	// Listing groups to target.
	CriterionTypeEnum_LISTING_GROUP CriterionTypeEnum_CriterionType = 8
	// Ad Schedule.
	CriterionTypeEnum_AD_SCHEDULE CriterionTypeEnum_CriterionType = 9
	// Age range.
	CriterionTypeEnum_AGE_RANGE CriterionTypeEnum_CriterionType = 10
	// Gender.
	CriterionTypeEnum_GENDER CriterionTypeEnum_CriterionType = 11
	// Income Range.
	CriterionTypeEnum_INCOME_RANGE CriterionTypeEnum_CriterionType = 12
	// Parental status.
	CriterionTypeEnum_PARENTAL_STATUS CriterionTypeEnum_CriterionType = 13
	// YouTube Video.
	CriterionTypeEnum_YOUTUBE_VIDEO CriterionTypeEnum_CriterionType = 14
	// YouTube Channel.
	CriterionTypeEnum_YOUTUBE_CHANNEL CriterionTypeEnum_CriterionType = 15
	// User list.
	CriterionTypeEnum_USER_LIST CriterionTypeEnum_CriterionType = 16
	// Proximity.
	CriterionTypeEnum_PROXIMITY CriterionTypeEnum_CriterionType = 17
	// A topic target on the display network (e.g. "Pets & Animals").
	CriterionTypeEnum_TOPIC CriterionTypeEnum_CriterionType = 18
	// Listing scope to target.
	CriterionTypeEnum_LISTING_SCOPE CriterionTypeEnum_CriterionType = 19
	// Language.
	CriterionTypeEnum_LANGUAGE CriterionTypeEnum_CriterionType = 20
	// IpBlock.
	CriterionTypeEnum_IP_BLOCK CriterionTypeEnum_CriterionType = 21
	// Content Label for category exclusion.
	CriterionTypeEnum_CONTENT_LABEL CriterionTypeEnum_CriterionType = 22
	// Carrier.
	CriterionTypeEnum_CARRIER CriterionTypeEnum_CriterionType = 23
	// A category the user is interested in.
	CriterionTypeEnum_USER_INTEREST CriterionTypeEnum_CriterionType = 24
)

var CriterionTypeEnum_CriterionType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "KEYWORD",
	3:  "PLACEMENT",
	6:  "DEVICE",
	7:  "LOCATION",
	8:  "LISTING_GROUP",
	9:  "AD_SCHEDULE",
	10: "AGE_RANGE",
	11: "GENDER",
	12: "INCOME_RANGE",
	13: "PARENTAL_STATUS",
	14: "YOUTUBE_VIDEO",
	15: "YOUTUBE_CHANNEL",
	16: "USER_LIST",
	17: "PROXIMITY",
	18: "TOPIC",
	19: "LISTING_SCOPE",
	20: "LANGUAGE",
	21: "IP_BLOCK",
	22: "CONTENT_LABEL",
	23: "CARRIER",
	24: "USER_INTEREST",
}
var CriterionTypeEnum_CriterionType_value = map[string]int32{
	"UNSPECIFIED":     0,
	"UNKNOWN":         1,
	"KEYWORD":         2,
	"PLACEMENT":       3,
	"DEVICE":          6,
	"LOCATION":        7,
	"LISTING_GROUP":   8,
	"AD_SCHEDULE":     9,
	"AGE_RANGE":       10,
	"GENDER":          11,
	"INCOME_RANGE":    12,
	"PARENTAL_STATUS": 13,
	"YOUTUBE_VIDEO":   14,
	"YOUTUBE_CHANNEL": 15,
	"USER_LIST":       16,
	"PROXIMITY":       17,
	"TOPIC":           18,
	"LISTING_SCOPE":   19,
	"LANGUAGE":        20,
	"IP_BLOCK":        21,
	"CONTENT_LABEL":   22,
	"CARRIER":         23,
	"USER_INTEREST":   24,
}

func (x CriterionTypeEnum_CriterionType) String() string {
	return proto.EnumName(CriterionTypeEnum_CriterionType_name, int32(x))
}
func (CriterionTypeEnum_CriterionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_criterion_type_f43400bfbfc8f5e1, []int{0, 0}
}

// The possible types of a criterion.
type CriterionTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CriterionTypeEnum) Reset()         { *m = CriterionTypeEnum{} }
func (m *CriterionTypeEnum) String() string { return proto.CompactTextString(m) }
func (*CriterionTypeEnum) ProtoMessage()    {}
func (*CriterionTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_criterion_type_f43400bfbfc8f5e1, []int{0}
}
func (m *CriterionTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CriterionTypeEnum.Unmarshal(m, b)
}
func (m *CriterionTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CriterionTypeEnum.Marshal(b, m, deterministic)
}
func (dst *CriterionTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CriterionTypeEnum.Merge(dst, src)
}
func (m *CriterionTypeEnum) XXX_Size() int {
	return xxx_messageInfo_CriterionTypeEnum.Size(m)
}
func (m *CriterionTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CriterionTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CriterionTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CriterionTypeEnum)(nil), "google.ads.googleads.v0.enums.CriterionTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.CriterionTypeEnum_CriterionType", CriterionTypeEnum_CriterionType_name, CriterionTypeEnum_CriterionType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/criterion_type.proto", fileDescriptor_criterion_type_f43400bfbfc8f5e1)
}

var fileDescriptor_criterion_type_f43400bfbfc8f5e1 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcd, 0x6e, 0xda, 0x40,
	0x10, 0x2e, 0xa0, 0x92, 0xb0, 0x40, 0x19, 0x36, 0xfd, 0xbb, 0xe4, 0xd0, 0x3c, 0x80, 0x8d, 0xda,
	0x63, 0x4f, 0xeb, 0xf5, 0xd4, 0x59, 0x61, 0x76, 0xad, 0xf5, 0x9a, 0x94, 0x0a, 0xc9, 0xa2, 0xc1,
	0xb2, 0x90, 0x82, 0x8d, 0x30, 0x89, 0x94, 0xd7, 0xe8, 0x23, 0xf4, 0xd8, 0xd7, 0xe8, 0xad, 0xcf,
	0xd4, 0x43, 0xb5, 0xa6, 0x8e, 0x9a, 0x43, 0x73, 0xb1, 0xbe, 0xdd, 0x6f, 0xbe, 0xf9, 0x76, 0x3c,
	0x1f, 0x79, 0x9f, 0x97, 0x65, 0x7e, 0x93, 0xb9, 0xab, 0x75, 0xe5, 0x1e, 0xa1, 0x45, 0x77, 0x13,
	0x37, 0x2b, 0x6e, 0xb7, 0x95, 0x7b, 0xbd, 0xdf, 0x1c, 0xb2, 0xfd, 0xa6, 0x2c, 0xd2, 0xc3, 0xfd,
	0x2e, 0x73, 0x76, 0xfb, 0xf2, 0x50, 0xd2, 0xf3, 0x63, 0xa1, 0xb3, 0x5a, 0x57, 0xce, 0x83, 0xc6,
	0xb9, 0x9b, 0x38, 0xb5, 0xe6, 0xe2, 0x5b, 0x87, 0x8c, 0x79, 0xa3, 0x33, 0xf7, 0xbb, 0x0c, 0x8b,
	0xdb, 0xed, 0xc5, 0xef, 0x36, 0x19, 0x3e, 0xba, 0xa5, 0x23, 0xd2, 0x4f, 0x64, 0x1c, 0x21, 0x17,
	0x9f, 0x04, 0xfa, 0xf0, 0x8c, 0xf6, 0xc9, 0x49, 0x22, 0xa7, 0x52, 0x5d, 0x49, 0x68, 0xd9, 0xc3,
	0x14, 0x17, 0x57, 0x4a, 0xfb, 0xd0, 0xa6, 0x43, 0xd2, 0x8b, 0x42, 0xc6, 0x71, 0x86, 0xd2, 0x40,
	0x87, 0x12, 0xd2, 0xf5, 0x71, 0x2e, 0x38, 0x42, 0x97, 0x0e, 0xc8, 0x69, 0xa8, 0x38, 0x33, 0x42,
	0x49, 0x38, 0xa1, 0x63, 0x32, 0x0c, 0x45, 0x6c, 0x84, 0x0c, 0xd2, 0x40, 0xab, 0x24, 0x82, 0x53,
	0x6b, 0xc3, 0xfc, 0x34, 0xe6, 0x97, 0xe8, 0x27, 0x21, 0x42, 0xcf, 0x36, 0x63, 0x01, 0xa6, 0x9a,
	0xc9, 0x00, 0x81, 0xd8, 0x66, 0x01, 0x4a, 0x1f, 0x35, 0xf4, 0x29, 0x90, 0x81, 0x90, 0x5c, 0xcd,
	0x1a, 0x76, 0x40, 0xcf, 0xc8, 0x28, 0x62, 0x1a, 0xa5, 0x61, 0x61, 0x1a, 0x1b, 0x66, 0x92, 0x18,
	0x86, 0xd6, 0x65, 0xa1, 0x12, 0x93, 0x78, 0x98, 0xce, 0x85, 0x8f, 0x0a, 0x5e, 0xd8, 0xba, 0xe6,
	0x8a, 0x5f, 0x32, 0x29, 0x31, 0x84, 0x91, 0x75, 0x4a, 0x62, 0xd4, 0xa9, 0x7d, 0x12, 0x40, 0x3d,
	0x85, 0x56, 0x9f, 0xc5, 0x4c, 0x98, 0x05, 0x8c, 0x69, 0x8f, 0x3c, 0x37, 0x2a, 0x12, 0x1c, 0xe8,
	0xbf, 0xcf, 0x8e, 0xb9, 0x8a, 0x10, 0xce, 0xea, 0xb9, 0x98, 0x0c, 0x12, 0x16, 0x20, 0xbc, 0xb4,
	0x27, 0x11, 0xa5, 0x5e, 0xa8, 0xf8, 0x14, 0x5e, 0xd9, 0x72, 0xae, 0xa4, 0x41, 0x69, 0xd2, 0x90,
	0x79, 0x18, 0xc2, 0x6b, 0xfb, 0xbb, 0x38, 0xd3, 0x5a, 0xa0, 0x86, 0x37, 0x96, 0xaf, 0x7d, 0x85,
	0x34, 0xa8, 0x31, 0x36, 0xf0, 0xd6, 0xfb, 0xd9, 0x22, 0xef, 0xae, 0xcb, 0xad, 0xf3, 0xe4, 0xea,
	0x3c, 0xfa, 0x68, 0x43, 0x91, 0xdd, 0x76, 0xd4, 0xfa, 0xe2, 0xfd, 0x15, 0xe5, 0xe5, 0xcd, 0xaa,
	0xc8, 0x9d, 0x72, 0x9f, 0xbb, 0x79, 0x56, 0xd4, 0x59, 0x68, 0x32, 0xb3, 0xdb, 0x54, 0xff, 0x89,
	0xd0, 0xc7, 0xfa, 0xfb, 0xbd, 0xdd, 0x09, 0x18, 0xfb, 0xd1, 0x3e, 0x0f, 0x8e, 0xad, 0xd8, 0xba,
	0x72, 0x8e, 0xd0, 0xa2, 0xf9, 0xc4, 0xb1, 0x19, 0xa9, 0x7e, 0x35, 0xfc, 0x92, 0xad, 0xab, 0xe5,
	0x03, 0xbf, 0x9c, 0x4f, 0x96, 0x35, 0xff, 0xb5, 0x5b, 0x9b, 0x7e, 0xf8, 0x13, 0x00, 0x00, 0xff,
	0xff, 0xd0, 0xd6, 0x60, 0x96, 0xb6, 0x02, 0x00, 0x00,
}
