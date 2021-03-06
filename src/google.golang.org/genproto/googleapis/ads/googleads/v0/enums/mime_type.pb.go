// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/mime_type.proto

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

// The mime type
type MimeTypeEnum_MimeType int32

const (
	// The mime type has not been specified.
	MimeTypeEnum_UNSPECIFIED MimeTypeEnum_MimeType = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	MimeTypeEnum_UNKNOWN MimeTypeEnum_MimeType = 1
	// MIME type of image/jpeg.
	MimeTypeEnum_IMAGE_JPEG MimeTypeEnum_MimeType = 2
	// MIME type of image/gif.
	MimeTypeEnum_IMAGE_GIF MimeTypeEnum_MimeType = 3
	// MIME type of image/png.
	MimeTypeEnum_IMAGE_PNG MimeTypeEnum_MimeType = 4
	// MIME type of application/x-shockwave-flash.
	MimeTypeEnum_FLASH MimeTypeEnum_MimeType = 5
	// MIME type of text/html.
	MimeTypeEnum_TEXT_HTML MimeTypeEnum_MimeType = 6
	// MIME type of application/pdf.
	MimeTypeEnum_PDF MimeTypeEnum_MimeType = 7
	// MIME type of application/msword.
	MimeTypeEnum_MSWORD MimeTypeEnum_MimeType = 8
	// MIME type of application/vnd.ms-excel.
	MimeTypeEnum_MSEXCEL MimeTypeEnum_MimeType = 9
	// MIME type of application/rtf.
	MimeTypeEnum_RTF MimeTypeEnum_MimeType = 10
	// MIME type of audio/wav.
	MimeTypeEnum_AUDIO_WAV MimeTypeEnum_MimeType = 11
	// MIME type of audio/mp3.
	MimeTypeEnum_AUDIO_MP3 MimeTypeEnum_MimeType = 12
	// MIME type of application/x-html5-ad-zip.
	MimeTypeEnum_HTML5_AD_ZIP MimeTypeEnum_MimeType = 13
)

var MimeTypeEnum_MimeType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "IMAGE_JPEG",
	3:  "IMAGE_GIF",
	4:  "IMAGE_PNG",
	5:  "FLASH",
	6:  "TEXT_HTML",
	7:  "PDF",
	8:  "MSWORD",
	9:  "MSEXCEL",
	10: "RTF",
	11: "AUDIO_WAV",
	12: "AUDIO_MP3",
	13: "HTML5_AD_ZIP",
}
var MimeTypeEnum_MimeType_value = map[string]int32{
	"UNSPECIFIED":  0,
	"UNKNOWN":      1,
	"IMAGE_JPEG":   2,
	"IMAGE_GIF":    3,
	"IMAGE_PNG":    4,
	"FLASH":        5,
	"TEXT_HTML":    6,
	"PDF":          7,
	"MSWORD":       8,
	"MSEXCEL":      9,
	"RTF":          10,
	"AUDIO_WAV":    11,
	"AUDIO_MP3":    12,
	"HTML5_AD_ZIP": 13,
}

func (x MimeTypeEnum_MimeType) String() string {
	return proto.EnumName(MimeTypeEnum_MimeType_name, int32(x))
}
func (MimeTypeEnum_MimeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_mime_type_488ab8ea3bec433a, []int{0, 0}
}

// Container for enum describing the mime types.
type MimeTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MimeTypeEnum) Reset()         { *m = MimeTypeEnum{} }
func (m *MimeTypeEnum) String() string { return proto.CompactTextString(m) }
func (*MimeTypeEnum) ProtoMessage()    {}
func (*MimeTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_mime_type_488ab8ea3bec433a, []int{0}
}
func (m *MimeTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MimeTypeEnum.Unmarshal(m, b)
}
func (m *MimeTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MimeTypeEnum.Marshal(b, m, deterministic)
}
func (dst *MimeTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MimeTypeEnum.Merge(dst, src)
}
func (m *MimeTypeEnum) XXX_Size() int {
	return xxx_messageInfo_MimeTypeEnum.Size(m)
}
func (m *MimeTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_MimeTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_MimeTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MimeTypeEnum)(nil), "google.ads.googleads.v0.enums.MimeTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.MimeTypeEnum_MimeType", MimeTypeEnum_MimeType_name, MimeTypeEnum_MimeType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/mime_type.proto", fileDescriptor_mime_type_488ab8ea3bec433a)
}

var fileDescriptor_mime_type_488ab8ea3bec433a = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcd, 0x6e, 0xaa, 0x40,
	0x1c, 0xc5, 0x2f, 0x7a, 0xfd, 0x1a, 0xf5, 0xde, 0x7f, 0x66, 0xef, 0xa2, 0xee, 0x3b, 0x90, 0x98,
	0xae, 0xba, 0x1a, 0x65, 0x40, 0x5a, 0xc1, 0x89, 0xe0, 0x47, 0x0c, 0x09, 0xb1, 0x85, 0x10, 0x13,
	0xf9, 0x88, 0xa8, 0x89, 0xaf, 0xd3, 0x65, 0x1f, 0xa2, 0x0f, 0xd0, 0xbe, 0x42, 0x1f, 0xa6, 0x19,
	0x28, 0xba, 0x6a, 0x37, 0xe4, 0x1c, 0xce, 0x99, 0xff, 0x3f, 0xf3, 0x1b, 0x74, 0x1b, 0x26, 0x49,
	0xb8, 0x0b, 0xe4, 0x8d, 0x9f, 0xc9, 0x85, 0x14, 0xea, 0xa4, 0xc8, 0x41, 0x7c, 0x8c, 0x32, 0x39,
	0xda, 0x46, 0x81, 0x77, 0x38, 0xa7, 0x01, 0x49, 0xf7, 0xc9, 0x21, 0xc1, 0xbd, 0xa2, 0x43, 0x36,
	0x7e, 0x46, 0x2e, 0x75, 0x72, 0x52, 0x48, 0x5e, 0xef, 0x7f, 0x4a, 0xa8, 0x63, 0x6e, 0xa3, 0xc0,
	0x39, 0xa7, 0x01, 0x8b, 0x8f, 0x51, 0xff, 0x43, 0x42, 0xcd, 0xf2, 0x07, 0xfe, 0x8f, 0xda, 0x73,
	0xcb, 0xe6, 0x6c, 0x64, 0x68, 0x06, 0x53, 0xe1, 0x0f, 0x6e, 0xa3, 0xc6, 0xdc, 0x7a, 0xb4, 0xa6,
	0x4b, 0x0b, 0x24, 0xfc, 0x0f, 0x21, 0xc3, 0xa4, 0x3a, 0xf3, 0x1e, 0x38, 0xd3, 0xa1, 0x82, 0xbb,
	0xa8, 0x55, 0x78, 0xdd, 0xd0, 0xa0, 0x7a, 0xb5, 0xdc, 0xd2, 0xe1, 0x2f, 0x6e, 0xa1, 0x9a, 0x36,
	0xa1, 0xf6, 0x18, 0x6a, 0x22, 0x71, 0xd8, 0xca, 0xf1, 0xc6, 0x8e, 0x39, 0x81, 0x3a, 0x6e, 0xa0,
	0x2a, 0x57, 0x35, 0x68, 0x60, 0x84, 0xea, 0xa6, 0xbd, 0x9c, 0xce, 0x54, 0x68, 0x8a, 0x4d, 0xa6,
	0xcd, 0x56, 0x23, 0x36, 0x81, 0x96, 0x68, 0xcc, 0x1c, 0x0d, 0x90, 0x38, 0x49, 0xe7, 0xaa, 0x31,
	0xf5, 0x96, 0x74, 0x01, 0xed, 0xab, 0x35, 0xf9, 0x00, 0x3a, 0x18, 0x50, 0x47, 0x8c, 0xbc, 0xf3,
	0xa8, 0xea, 0xad, 0x0d, 0x0e, 0xdd, 0xe1, 0x9b, 0x84, 0x6e, 0x9e, 0x93, 0x88, 0xfc, 0x0a, 0x61,
	0xd8, 0x2d, 0x2f, 0xcc, 0x05, 0x32, 0x2e, 0xad, 0x87, 0xdf, 0xfd, 0x30, 0xd9, 0x6d, 0xe2, 0x90,
	0x24, 0xfb, 0x50, 0x0e, 0x83, 0x38, 0x07, 0x5a, 0x32, 0x4f, 0xb7, 0xd9, 0x0f, 0x4f, 0x70, 0x9f,
	0x7f, 0x5f, 0x2a, 0x55, 0x9d, 0xd2, 0xd7, 0x4a, 0x4f, 0x2f, 0x46, 0x51, 0x3f, 0x23, 0x85, 0x14,
	0x6a, 0xa1, 0x10, 0x41, 0x3b, 0x7b, 0x2f, 0x73, 0x97, 0xfa, 0x99, 0x7b, 0xc9, 0xdd, 0x85, 0xe2,
	0xe6, 0xf9, 0x53, 0x3d, 0x5f, 0x3a, 0xf8, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x64, 0xde, 0xc3,
	0xf6, 0x01, 0x00, 0x00,
}
