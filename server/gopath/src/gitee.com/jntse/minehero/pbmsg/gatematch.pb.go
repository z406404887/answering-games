// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gatematch.proto

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// gw请求注册到match
type GW2MS_ReqRegist struct {
	Account          *string `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
	Passwd           *string `protobuf:"bytes,2,opt,name=passwd" json:"passwd,omitempty"`
	Agentname        *string `protobuf:"bytes,3,opt,name=agentname" json:"agentname,omitempty"`
	Host             *IpHost `protobuf:"bytes,4,opt,name=host" json:"host,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GW2MS_ReqRegist) Reset()                    { *m = GW2MS_ReqRegist{} }
func (m *GW2MS_ReqRegist) String() string            { return proto.CompactTextString(m) }
func (*GW2MS_ReqRegist) ProtoMessage()               {}
func (*GW2MS_ReqRegist) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *GW2MS_ReqRegist) GetAccount() string {
	if m != nil && m.Account != nil {
		return *m.Account
	}
	return ""
}

func (m *GW2MS_ReqRegist) GetPasswd() string {
	if m != nil && m.Passwd != nil {
		return *m.Passwd
	}
	return ""
}

func (m *GW2MS_ReqRegist) GetAgentname() string {
	if m != nil && m.Agentname != nil {
		return *m.Agentname
	}
	return ""
}

func (m *GW2MS_ReqRegist) GetHost() *IpHost {
	if m != nil {
		return m.Host
	}
	return nil
}

// 注册返回
type MS2GW_RetRegist struct {
	Errcode          *string `protobuf:"bytes,1,opt,name=errcode" json:"errcode,omitempty"`
	Host             *IpHost `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MS2GW_RetRegist) Reset()                    { *m = MS2GW_RetRegist{} }
func (m *MS2GW_RetRegist) String() string            { return proto.CompactTextString(m) }
func (*MS2GW_RetRegist) ProtoMessage()               {}
func (*MS2GW_RetRegist) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *MS2GW_RetRegist) GetErrcode() string {
	if m != nil && m.Errcode != nil {
		return *m.Errcode
	}
	return ""
}

func (m *MS2GW_RetRegist) GetHost() *IpHost {
	if m != nil {
		return m.Host
	}
	return nil
}

// 心跳
type GW2MS_HeartBeat struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GW2MS_HeartBeat) Reset()                    { *m = GW2MS_HeartBeat{} }
func (m *GW2MS_HeartBeat) String() string            { return proto.CompactTextString(m) }
func (*GW2MS_HeartBeat) ProtoMessage()               {}
func (*GW2MS_HeartBeat) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

//
type MS2GW_HeartBeat struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *MS2GW_HeartBeat) Reset()                    { *m = MS2GW_HeartBeat{} }
func (m *MS2GW_HeartBeat) String() string            { return proto.CompactTextString(m) }
func (*MS2GW_HeartBeat) ProtoMessage()               {}
func (*MS2GW_HeartBeat) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

// 请求创建游戏房间
type GW2MS_ReqCreateRoom struct {
	Userid           *uint64 `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
	Gamekind         *int32  `protobuf:"varint,2,opt,name=gamekind" json:"gamekind,omitempty"`
	Gridnum          *int32  `protobuf:"varint,3,opt,name=gridnum" json:"gridnum,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GW2MS_ReqCreateRoom) Reset()                    { *m = GW2MS_ReqCreateRoom{} }
func (m *GW2MS_ReqCreateRoom) String() string            { return proto.CompactTextString(m) }
func (*GW2MS_ReqCreateRoom) ProtoMessage()               {}
func (*GW2MS_ReqCreateRoom) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{4} }

func (m *GW2MS_ReqCreateRoom) GetUserid() uint64 {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return 0
}

func (m *GW2MS_ReqCreateRoom) GetGamekind() int32 {
	if m != nil && m.Gamekind != nil {
		return *m.Gamekind
	}
	return 0
}

func (m *GW2MS_ReqCreateRoom) GetGridnum() int32 {
	if m != nil && m.Gridnum != nil {
		return *m.Gridnum
	}
	return 0
}

//
type MS2GW_RetCreateRoom struct {
	Userid           *uint64 `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
	Roomid           *int64  `protobuf:"varint,2,opt,name=roomid" json:"roomid,omitempty"`
	Errcode          *string `protobuf:"bytes,3,opt,name=errcode" json:"errcode,omitempty"`
	Roomagent        *string `protobuf:"bytes,4,opt,name=roomagent" json:"roomagent,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MS2GW_RetCreateRoom) Reset()                    { *m = MS2GW_RetCreateRoom{} }
func (m *MS2GW_RetCreateRoom) String() string            { return proto.CompactTextString(m) }
func (*MS2GW_RetCreateRoom) ProtoMessage()               {}
func (*MS2GW_RetCreateRoom) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{5} }

func (m *MS2GW_RetCreateRoom) GetUserid() uint64 {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return 0
}

func (m *MS2GW_RetCreateRoom) GetRoomid() int64 {
	if m != nil && m.Roomid != nil {
		return *m.Roomid
	}
	return 0
}

func (m *MS2GW_RetCreateRoom) GetErrcode() string {
	if m != nil && m.Errcode != nil {
		return *m.Errcode
	}
	return ""
}

func (m *MS2GW_RetCreateRoom) GetRoomagent() string {
	if m != nil && m.Roomagent != nil {
		return *m.Roomagent
	}
	return ""
}

func init() {
	proto.RegisterType((*GW2MS_ReqRegist)(nil), "msg.GW2MS_ReqRegist")
	proto.RegisterType((*MS2GW_RetRegist)(nil), "msg.MS2GW_RetRegist")
	proto.RegisterType((*GW2MS_HeartBeat)(nil), "msg.GW2MS_HeartBeat")
	proto.RegisterType((*MS2GW_HeartBeat)(nil), "msg.MS2GW_HeartBeat")
	proto.RegisterType((*GW2MS_ReqCreateRoom)(nil), "msg.GW2MS_ReqCreateRoom")
	proto.RegisterType((*MS2GW_RetCreateRoom)(nil), "msg.MS2GW_RetCreateRoom")
}

func init() { proto.RegisterFile("gatematch.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x14, 0x84, 0x95, 0x26, 0x2d, 0xe4, 0x15, 0x35, 0x34, 0x5d, 0x02, 0x53, 0x95, 0xa9, 0x53, 0x86,
	0xec, 0x2c, 0x30, 0x34, 0x0c, 0x5d, 0xd2, 0xa1, 0x62, 0x2a, 0x56, 0xfc, 0x70, 0x2d, 0x64, 0x3b,
	0xd8, 0x2f, 0xe2, 0xef, 0xa3, 0x38, 0xa4, 0x45, 0x42, 0xea, 0xe8, 0x93, 0xdf, 0xdd, 0x7d, 0x07,
	0x89, 0x60, 0x84, 0x8a, 0x51, 0x73, 0x2a, 0x5a, 0x6b, 0xc8, 0xa4, 0xa1, 0x72, 0xe2, 0xf1, 0x8e,
	0xe3, 0x87, 0xd4, 0x38, 0x48, 0xf9, 0x3b, 0x24, 0xdb, 0x43, 0xb9, 0xdb, 0x1f, 0x6b, 0xfc, 0xaa,
	0x51, 0x48, 0x47, 0x69, 0x02, 0x37, 0xac, 0x69, 0x4c, 0xa7, 0x29, 0x0b, 0xd6, 0xc1, 0x26, 0x4e,
	0x17, 0x30, 0x6b, 0x99, 0x73, 0xdf, 0x3c, 0x9b, 0xf8, 0xf7, 0x12, 0x62, 0x26, 0x50, 0x93, 0x66,
	0x0a, 0xb3, 0xd0, 0x4b, 0x0f, 0x10, 0x9d, 0x8c, 0xa3, 0x2c, 0x5a, 0x07, 0x9b, 0x79, 0x39, 0x2f,
	0x94, 0x13, 0xc5, 0x6b, 0x5b, 0x19, 0x47, 0xf9, 0x13, 0x24, 0xbb, 0x7d, 0xb9, 0x3d, 0x1c, 0x6b,
	0xa4, 0x4b, 0x02, 0x5a, 0xdb, 0x18, 0x8e, 0xbf, 0x09, 0xe3, 0xf9, 0xe4, 0xff, 0xf9, 0x72, 0x2c,
	0x58, 0x21, 0xb3, 0xf4, 0x8c, 0xcc, 0x4b, 0x83, 0xe3, 0x45, 0xaa, 0x60, 0x75, 0xc6, 0x78, 0xb1,
	0xc8, 0x08, 0x6b, 0x63, 0x54, 0xdf, 0xbc, 0x73, 0x68, 0x25, 0xf7, 0x39, 0x51, 0x7a, 0x0f, 0xb7,
	0x82, 0x29, 0xfc, 0x94, 0x7a, 0x60, 0x99, 0xf6, 0x55, 0x84, 0x95, 0x5c, 0x77, 0xca, 0x93, 0x4c,
	0xf3, 0x37, 0x58, 0x9d, 0xeb, 0x5e, 0x71, 0x5a, 0xc0, 0xcc, 0x1a, 0xa3, 0xe4, 0xe0, 0x13, 0xfe,
	0x45, 0x0a, 0xc7, 0x91, 0xfa, 0x0f, 0x7e, 0x28, 0x3f, 0x4b, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff,
	0xaa, 0xee, 0xce, 0x1e, 0x90, 0x01, 0x00, 0x00,
}
