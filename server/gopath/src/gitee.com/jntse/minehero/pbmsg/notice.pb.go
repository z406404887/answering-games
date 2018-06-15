// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notice.proto

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 公告类型
type NoticeType int32

const (
	NoticeType_Suspension NoticeType = 1
	NoticeType_Marquee    NoticeType = 2
)

var NoticeType_name = map[int32]string{
	1: "Suspension",
	2: "Marquee",
}
var NoticeType_value = map[string]int32{
	"Suspension": 1,
	"Marquee":    2,
}

func (x NoticeType) Enum() *NoticeType {
	p := new(NoticeType)
	*p = x
	return p
}
func (x NoticeType) String() string {
	return proto.EnumName(NoticeType_name, int32(x))
}
func (x *NoticeType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NoticeType_value, data, "NoticeType")
	if err != nil {
		return err
	}
	*x = NoticeType(value)
	return nil
}
func (NoticeType) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

// 个人提示信息
type GW2C_MsgNotify struct {
	Userid           *uint64 `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
	Text             *string `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GW2C_MsgNotify) Reset()                    { *m = GW2C_MsgNotify{} }
func (m *GW2C_MsgNotify) String() string            { return proto.CompactTextString(m) }
func (*GW2C_MsgNotify) ProtoMessage()               {}
func (*GW2C_MsgNotify) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *GW2C_MsgNotify) GetUserid() uint64 {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return 0
}

func (m *GW2C_MsgNotify) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

// 公告提示信息
type GW2C_MsgNotice struct {
	Userid           *uint64 `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
	Face             *string `protobuf:"bytes,2,opt,name=face" json:"face,omitempty"`
	Name             *string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Type             *int32  `protobuf:"varint,4,opt,name=type" json:"type,omitempty"`
	Text             *string `protobuf:"bytes,5,opt,name=text" json:"text,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GW2C_MsgNotice) Reset()                    { *m = GW2C_MsgNotice{} }
func (m *GW2C_MsgNotice) String() string            { return proto.CompactTextString(m) }
func (*GW2C_MsgNotice) ProtoMessage()               {}
func (*GW2C_MsgNotice) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

func (m *GW2C_MsgNotice) GetUserid() uint64 {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return 0
}

func (m *GW2C_MsgNotice) GetFace() string {
	if m != nil && m.Face != nil {
		return *m.Face
	}
	return ""
}

func (m *GW2C_MsgNotice) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *GW2C_MsgNotice) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *GW2C_MsgNotice) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

// 转发
type GW2MS_MsgNotice struct {
	Notice           *GW2C_MsgNotice `protobuf:"bytes,1,opt,name=notice" json:"notice,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *GW2MS_MsgNotice) Reset()                    { *m = GW2MS_MsgNotice{} }
func (m *GW2MS_MsgNotice) String() string            { return proto.CompactTextString(m) }
func (*GW2MS_MsgNotice) ProtoMessage()               {}
func (*GW2MS_MsgNotice) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

func (m *GW2MS_MsgNotice) GetNotice() *GW2C_MsgNotice {
	if m != nil {
		return m.Notice
	}
	return nil
}

// 转发
type RS2MS_MsgNotice struct {
	Notice           *GW2C_MsgNotice `protobuf:"bytes,1,opt,name=notice" json:"notice,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *RS2MS_MsgNotice) Reset()                    { *m = RS2MS_MsgNotice{} }
func (m *RS2MS_MsgNotice) String() string            { return proto.CompactTextString(m) }
func (*RS2MS_MsgNotice) ProtoMessage()               {}
func (*RS2MS_MsgNotice) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{3} }

func (m *RS2MS_MsgNotice) GetNotice() *GW2C_MsgNotice {
	if m != nil {
		return m.Notice
	}
	return nil
}

// 分发
type MS2GW_MsgNotice struct {
	Notice           *GW2C_MsgNotice `protobuf:"bytes,1,opt,name=notice" json:"notice,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *MS2GW_MsgNotice) Reset()                    { *m = MS2GW_MsgNotice{} }
func (m *MS2GW_MsgNotice) String() string            { return proto.CompactTextString(m) }
func (*MS2GW_MsgNotice) ProtoMessage()               {}
func (*MS2GW_MsgNotice) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{4} }

func (m *MS2GW_MsgNotice) GetNotice() *GW2C_MsgNotice {
	if m != nil {
		return m.Notice
	}
	return nil
}

func init() {
	proto.RegisterType((*GW2C_MsgNotify)(nil), "msg.GW2C_MsgNotify")
	proto.RegisterType((*GW2C_MsgNotice)(nil), "msg.GW2C_MsgNotice")
	proto.RegisterType((*GW2MS_MsgNotice)(nil), "msg.GW2MS_MsgNotice")
	proto.RegisterType((*RS2MS_MsgNotice)(nil), "msg.RS2MS_MsgNotice")
	proto.RegisterType((*MS2GW_MsgNotice)(nil), "msg.MS2GW_MsgNotice")
	proto.RegisterEnum("msg.NoticeType", NoticeType_name, NoticeType_value)
}

func init() { proto.RegisterFile("notice.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0xcb, 0x2f, 0xc9,
	0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x2d, 0x4e, 0x57, 0xd2, 0xe3,
	0xe2, 0x73, 0x0f, 0x37, 0x72, 0x8e, 0xf7, 0x2d, 0x4e, 0xf7, 0xcb, 0x2f, 0xc9, 0x4c, 0xab, 0x14,
	0xe2, 0xe3, 0x62, 0x2b, 0x2d, 0x4e, 0x2d, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x11,
	0xe2, 0xe1, 0x62, 0x29, 0x49, 0xad, 0x28, 0x91, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x54, 0x8a, 0x40,
	0x55, 0x9f, 0x9c, 0x8a, 0x4d, 0x7d, 0x5a, 0x62, 0x72, 0x2a, 0x44, 0x3d, 0x88, 0x97, 0x97, 0x98,
	0x9b, 0x2a, 0xc1, 0x0c, 0xe3, 0x95, 0x54, 0x16, 0xa4, 0x4a, 0xb0, 0x28, 0x30, 0x6a, 0xb0, 0xc2,
	0x4d, 0x66, 0x05, 0x9b, 0x6c, 0xc6, 0xc5, 0xef, 0x1e, 0x6e, 0xe4, 0x1b, 0x8c, 0x64, 0xb4, 0x32,
	0x17, 0x1b, 0xc4, 0xc5, 0x60, 0xa3, 0xb9, 0x8d, 0x84, 0xf5, 0x72, 0x8b, 0xd3, 0xf5, 0x50, 0xed,
	0x07, 0xe9, 0x0b, 0x0a, 0x26, 0x4f, 0x9f, 0x6f, 0xb0, 0x91, 0x7b, 0x38, 0x89, 0xfa, 0xb4, 0x34,
	0xb9, 0xb8, 0x20, 0xac, 0x90, 0xca, 0x02, 0x90, 0xef, 0xb9, 0x82, 0x4b, 0x8b, 0x0b, 0x52, 0xf3,
	0x8a, 0x33, 0xf3, 0xf3, 0x04, 0x18, 0x85, 0xb8, 0xb9, 0xd8, 0x7d, 0x13, 0x8b, 0x0a, 0x4b, 0x53,
	0x53, 0x05, 0x98, 0x00, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xa8, 0xab, 0x8d, 0x70, 0x01, 0x00,
	0x00,
}
