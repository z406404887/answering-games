// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gatelogin.proto

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// gw请求注册到login
type GW2L_ReqRegist struct {
	Account          *string `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
	Passwd           *string `protobuf:"bytes,2,opt,name=passwd" json:"passwd,omitempty"`
	Host             *IpHost `protobuf:"bytes,3,opt,name=host" json:"host,omitempty"`
	Name             *string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GW2L_ReqRegist) Reset()                    { *m = GW2L_ReqRegist{} }
func (m *GW2L_ReqRegist) String() string            { return proto.CompactTextString(m) }
func (*GW2L_ReqRegist) ProtoMessage()               {}
func (*GW2L_ReqRegist) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *GW2L_ReqRegist) GetAccount() string {
	if m != nil && m.Account != nil {
		return *m.Account
	}
	return ""
}

func (m *GW2L_ReqRegist) GetPasswd() string {
	if m != nil && m.Passwd != nil {
		return *m.Passwd
	}
	return ""
}

func (m *GW2L_ReqRegist) GetHost() *IpHost {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *GW2L_ReqRegist) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

// 注册返回
type L2GW_RetRegist struct {
	Errocde          *string `protobuf:"bytes,1,opt,name=errocde" json:"errocde,omitempty"`
	Host             *IpHost `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *L2GW_RetRegist) Reset()                    { *m = L2GW_RetRegist{} }
func (m *L2GW_RetRegist) String() string            { return proto.CompactTextString(m) }
func (*L2GW_RetRegist) ProtoMessage()               {}
func (*L2GW_RetRegist) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *L2GW_RetRegist) GetErrocde() string {
	if m != nil && m.Errocde != nil {
		return *m.Errocde
	}
	return ""
}

func (m *L2GW_RetRegist) GetHost() *IpHost {
	if m != nil {
		return m.Host
	}
	return nil
}

// 心跳
type GW2L_HeartBeat struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GW2L_HeartBeat) Reset()                    { *m = GW2L_HeartBeat{} }
func (m *GW2L_HeartBeat) String() string            { return proto.CompactTextString(m) }
func (*GW2L_HeartBeat) ProtoMessage()               {}
func (*GW2L_HeartBeat) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

//
type L2GW_HeartBeat struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *L2GW_HeartBeat) Reset()                    { *m = L2GW_HeartBeat{} }
func (m *L2GW_HeartBeat) String() string            { return proto.CompactTextString(m) }
func (*L2GW_HeartBeat) ProtoMessage()               {}
func (*L2GW_HeartBeat) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

// 玩家登陆ls成功，注册他到相应的Gw
type L2GW_ReqRegistUser struct {
	Account          *string `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
	Expire           *int64  `protobuf:"varint,2,opt,name=expire" json:"expire,omitempty"`
	Gatehost         *string `protobuf:"bytes,3,opt,name=gatehost" json:"gatehost,omitempty"`
	Sid              *int32  `protobuf:"varint,4,opt,name=sid" json:"sid,omitempty"`
	Timestamp        *int64  `protobuf:"varint,5,opt,name=timestamp" json:"timestamp,omitempty"`
	Verifykey        *string `protobuf:"bytes,6,opt,name=verifykey" json:"verifykey,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *L2GW_ReqRegistUser) Reset()                    { *m = L2GW_ReqRegistUser{} }
func (m *L2GW_ReqRegistUser) String() string            { return proto.CompactTextString(m) }
func (*L2GW_ReqRegistUser) ProtoMessage()               {}
func (*L2GW_ReqRegistUser) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *L2GW_ReqRegistUser) GetAccount() string {
	if m != nil && m.Account != nil {
		return *m.Account
	}
	return ""
}

func (m *L2GW_ReqRegistUser) GetExpire() int64 {
	if m != nil && m.Expire != nil {
		return *m.Expire
	}
	return 0
}

func (m *L2GW_ReqRegistUser) GetGatehost() string {
	if m != nil && m.Gatehost != nil {
		return *m.Gatehost
	}
	return ""
}

func (m *L2GW_ReqRegistUser) GetSid() int32 {
	if m != nil && m.Sid != nil {
		return *m.Sid
	}
	return 0
}

func (m *L2GW_ReqRegistUser) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *L2GW_ReqRegistUser) GetVerifykey() string {
	if m != nil && m.Verifykey != nil {
		return *m.Verifykey
	}
	return ""
}

// 注册返回
type GW2L_RegistUserRet struct {
	Account          *string `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
	Gatehost         *string `protobuf:"bytes,2,opt,name=gatehost" json:"gatehost,omitempty"`
	Errcode          *string `protobuf:"bytes,3,opt,name=errcode" json:"errcode,omitempty"`
	Sid              *int32  `protobuf:"varint,4,opt,name=sid" json:"sid,omitempty"`
	Verifykey        *string `protobuf:"bytes,5,opt,name=verifykey" json:"verifykey,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GW2L_RegistUserRet) Reset()                    { *m = GW2L_RegistUserRet{} }
func (m *GW2L_RegistUserRet) String() string            { return proto.CompactTextString(m) }
func (*GW2L_RegistUserRet) ProtoMessage()               {}
func (*GW2L_RegistUserRet) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *GW2L_RegistUserRet) GetAccount() string {
	if m != nil && m.Account != nil {
		return *m.Account
	}
	return ""
}

func (m *GW2L_RegistUserRet) GetGatehost() string {
	if m != nil && m.Gatehost != nil {
		return *m.Gatehost
	}
	return ""
}

func (m *GW2L_RegistUserRet) GetErrcode() string {
	if m != nil && m.Errcode != nil {
		return *m.Errcode
	}
	return ""
}

func (m *GW2L_RegistUserRet) GetSid() int32 {
	if m != nil && m.Sid != nil {
		return *m.Sid
	}
	return 0
}

func (m *GW2L_RegistUserRet) GetVerifykey() string {
	if m != nil && m.Verifykey != nil {
		return *m.Verifykey
	}
	return ""
}

func init() {
	proto.RegisterType((*GW2L_ReqRegist)(nil), "msg.GW2L_ReqRegist")
	proto.RegisterType((*L2GW_RetRegist)(nil), "msg.L2GW_RetRegist")
	proto.RegisterType((*GW2L_HeartBeat)(nil), "msg.GW2L_HeartBeat")
	proto.RegisterType((*L2GW_HeartBeat)(nil), "msg.L2GW_HeartBeat")
	proto.RegisterType((*L2GW_ReqRegistUser)(nil), "msg.L2GW_ReqRegistUser")
	proto.RegisterType((*GW2L_RegistUserRet)(nil), "msg.GW2L_RegistUserRet")
}

func init() { proto.RegisterFile("gatelogin.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x3f, 0x4f, 0xc3, 0x30,
	0x14, 0xc4, 0x95, 0x26, 0x2d, 0xe4, 0xa5, 0x4a, 0x5b, 0x4f, 0x86, 0xa9, 0xca, 0xd4, 0x29, 0x43,
	0x66, 0x26, 0x96, 0x16, 0xa9, 0x53, 0x24, 0x54, 0x89, 0xa5, 0xb2, 0x92, 0xd7, 0xd4, 0x82, 0xc4,
	0xc1, 0x7e, 0xfc, 0x29, 0x9f, 0x1e, 0xc5, 0x71, 0x09, 0x12, 0x65, 0xf4, 0xc9, 0xef, 0x7e, 0x77,
	0x07, 0xb3, 0x4a, 0x10, 0xbe, 0xa8, 0x4a, 0x36, 0x69, 0xab, 0x15, 0x29, 0xe6, 0xd7, 0xa6, 0xba,
	0x9d, 0x96, 0x78, 0x90, 0x0d, 0xf6, 0x52, 0xf2, 0x04, 0xf1, 0x7a, 0x97, 0x6d, 0xf7, 0x39, 0xbe,
	0xe6, 0x58, 0x49, 0x43, 0x6c, 0x06, 0x57, 0xa2, 0x28, 0xd4, 0x5b, 0x43, 0xdc, 0x5b, 0x7a, 0xab,
	0x90, 0xc5, 0x30, 0x69, 0x85, 0x31, 0x1f, 0x25, 0x1f, 0xd9, 0xf7, 0x0d, 0x04, 0x47, 0x65, 0x88,
	0xfb, 0x4b, 0x6f, 0x15, 0x65, 0x51, 0x5a, 0x9b, 0x2a, 0x7d, 0x68, 0x37, 0xca, 0x10, 0x9b, 0x42,
	0xd0, 0x88, 0x1a, 0x79, 0xd0, 0x7d, 0x4c, 0xee, 0x20, 0xde, 0x66, 0xeb, 0xdd, 0x3e, 0x47, 0x1a,
	0xbc, 0x51, 0x6b, 0x55, 0x94, 0xe8, 0xbc, 0xcf, 0x5e, 0xa3, 0x3f, 0x5e, 0xc9, 0xdc, 0x25, 0xdb,
	0xa0, 0xd0, 0x74, 0x8f, 0xc2, 0x2a, 0xd6, 0x6f, 0x50, 0xbe, 0x80, 0x39, 0x82, 0x4b, 0xff, 0x68,
	0x50, 0x5f, 0x6c, 0x80, 0x9f, 0xad, 0xd4, 0x68, 0x39, 0x3e, 0x9b, 0xc3, 0x75, 0x37, 0xcd, 0x4f,
	0x8b, 0x90, 0x45, 0xe0, 0x1b, 0x59, 0xda, 0xdc, 0x63, 0xb6, 0x80, 0x90, 0x64, 0x8d, 0x86, 0x44,
	0xdd, 0xf2, 0xb1, 0xbd, 0x58, 0x40, 0xf8, 0x8e, 0x5a, 0x1e, 0x4e, 0xcf, 0x78, 0xe2, 0x13, 0xdb,
	0xee, 0x08, 0xcc, 0x2d, 0x77, 0x06, 0xe7, 0x78, 0x61, 0xbd, 0xdf, 0xac, 0x7e, 0xbf, 0x7e, 0x84,
	0x42, 0x95, 0xf8, 0x0f, 0x7c, 0x20, 0x75, 0xf0, 0xf0, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xb4,
	0x5d, 0x13, 0xc8, 0x01, 0x00, 0x00,
}
