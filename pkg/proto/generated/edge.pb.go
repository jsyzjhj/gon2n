// Code generated by protoc-gen-go. DO NOT EDIT.
// source: edge.proto

package gon2n

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type EdgeManagerCreateArgs struct {
	AllowP2P             bool     `protobuf:"varint,1,opt,name=AllowP2P,proto3" json:"AllowP2P,omitempty"`
	AllowRouting         bool     `protobuf:"varint,2,opt,name=AllowRouting,proto3" json:"AllowRouting,omitempty"`
	CommunityName        string   `protobuf:"bytes,3,opt,name=CommunityName,proto3" json:"CommunityName,omitempty"`
	DisablePMTUDiscovery bool     `protobuf:"varint,4,opt,name=DisablePMTUDiscovery,proto3" json:"DisablePMTUDiscovery,omitempty"`
	DisableMulticast     bool     `protobuf:"varint,5,opt,name=DisableMulticast,proto3" json:"DisableMulticast,omitempty"`
	DynamicIPMode        bool     `protobuf:"varint,6,opt,name=DynamicIPMode,proto3" json:"DynamicIPMode,omitempty"`
	EncryptionKey        string   `protobuf:"bytes,7,opt,name=EncryptionKey,proto3" json:"EncryptionKey,omitempty"`
	LocalPort            int64    `protobuf:"varint,8,opt,name=LocalPort,proto3" json:"LocalPort,omitempty"`
	ManagementPort       int64    `protobuf:"varint,9,opt,name=ManagementPort,proto3" json:"ManagementPort,omitempty"`
	RegisterInterval     int64    `protobuf:"varint,10,opt,name=RegisterInterval,proto3" json:"RegisterInterval,omitempty"`
	RegisterTTL          int64    `protobuf:"varint,11,opt,name=RegisterTTL,proto3" json:"RegisterTTL,omitempty"`
	SupernodeHostPort    string   `protobuf:"bytes,12,opt,name=SupernodeHostPort,proto3" json:"SupernodeHostPort,omitempty"`
	TypeOfService        int64    `protobuf:"varint,13,opt,name=TypeOfService,proto3" json:"TypeOfService,omitempty"`
	EncryptionMethod     int64    `protobuf:"varint,14,opt,name=EncryptionMethod,proto3" json:"EncryptionMethod,omitempty"`
	DeviceName           string   `protobuf:"bytes,15,opt,name=DeviceName,proto3" json:"DeviceName,omitempty"`
	AddressMode          string   `protobuf:"bytes,16,opt,name=AddressMode,proto3" json:"AddressMode,omitempty"`
	DeviceIP             string   `protobuf:"bytes,17,opt,name=DeviceIP,proto3" json:"DeviceIP,omitempty"`
	DeviceNetmask        string   `protobuf:"bytes,18,opt,name=DeviceNetmask,proto3" json:"DeviceNetmask,omitempty"`
	DeviceMACAddress     string   `protobuf:"bytes,19,opt,name=DeviceMACAddress,proto3" json:"DeviceMACAddress,omitempty"`
	MTU                  int64    `protobuf:"varint,20,opt,name=MTU,proto3" json:"MTU,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EdgeManagerCreateArgs) Reset()         { *m = EdgeManagerCreateArgs{} }
func (m *EdgeManagerCreateArgs) String() string { return proto.CompactTextString(m) }
func (*EdgeManagerCreateArgs) ProtoMessage()    {}
func (*EdgeManagerCreateArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_cab1176173a95651, []int{0}
}

func (m *EdgeManagerCreateArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EdgeManagerCreateArgs.Unmarshal(m, b)
}
func (m *EdgeManagerCreateArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EdgeManagerCreateArgs.Marshal(b, m, deterministic)
}
func (m *EdgeManagerCreateArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EdgeManagerCreateArgs.Merge(m, src)
}
func (m *EdgeManagerCreateArgs) XXX_Size() int {
	return xxx_messageInfo_EdgeManagerCreateArgs.Size(m)
}
func (m *EdgeManagerCreateArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_EdgeManagerCreateArgs.DiscardUnknown(m)
}

var xxx_messageInfo_EdgeManagerCreateArgs proto.InternalMessageInfo

func (m *EdgeManagerCreateArgs) GetAllowP2P() bool {
	if m != nil {
		return m.AllowP2P
	}
	return false
}

func (m *EdgeManagerCreateArgs) GetAllowRouting() bool {
	if m != nil {
		return m.AllowRouting
	}
	return false
}

func (m *EdgeManagerCreateArgs) GetCommunityName() string {
	if m != nil {
		return m.CommunityName
	}
	return ""
}

func (m *EdgeManagerCreateArgs) GetDisablePMTUDiscovery() bool {
	if m != nil {
		return m.DisablePMTUDiscovery
	}
	return false
}

func (m *EdgeManagerCreateArgs) GetDisableMulticast() bool {
	if m != nil {
		return m.DisableMulticast
	}
	return false
}

func (m *EdgeManagerCreateArgs) GetDynamicIPMode() bool {
	if m != nil {
		return m.DynamicIPMode
	}
	return false
}

func (m *EdgeManagerCreateArgs) GetEncryptionKey() string {
	if m != nil {
		return m.EncryptionKey
	}
	return ""
}

func (m *EdgeManagerCreateArgs) GetLocalPort() int64 {
	if m != nil {
		return m.LocalPort
	}
	return 0
}

func (m *EdgeManagerCreateArgs) GetManagementPort() int64 {
	if m != nil {
		return m.ManagementPort
	}
	return 0
}

func (m *EdgeManagerCreateArgs) GetRegisterInterval() int64 {
	if m != nil {
		return m.RegisterInterval
	}
	return 0
}

func (m *EdgeManagerCreateArgs) GetRegisterTTL() int64 {
	if m != nil {
		return m.RegisterTTL
	}
	return 0
}

func (m *EdgeManagerCreateArgs) GetSupernodeHostPort() string {
	if m != nil {
		return m.SupernodeHostPort
	}
	return ""
}

func (m *EdgeManagerCreateArgs) GetTypeOfService() int64 {
	if m != nil {
		return m.TypeOfService
	}
	return 0
}

func (m *EdgeManagerCreateArgs) GetEncryptionMethod() int64 {
	if m != nil {
		return m.EncryptionMethod
	}
	return 0
}

func (m *EdgeManagerCreateArgs) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *EdgeManagerCreateArgs) GetAddressMode() string {
	if m != nil {
		return m.AddressMode
	}
	return ""
}

func (m *EdgeManagerCreateArgs) GetDeviceIP() string {
	if m != nil {
		return m.DeviceIP
	}
	return ""
}

func (m *EdgeManagerCreateArgs) GetDeviceNetmask() string {
	if m != nil {
		return m.DeviceNetmask
	}
	return ""
}

func (m *EdgeManagerCreateArgs) GetDeviceMACAddress() string {
	if m != nil {
		return m.DeviceMACAddress
	}
	return ""
}

func (m *EdgeManagerCreateArgs) GetMTU() int64 {
	if m != nil {
		return m.MTU
	}
	return 0
}

type EdgeManagerCreateReply struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EdgeManagerCreateReply) Reset()         { *m = EdgeManagerCreateReply{} }
func (m *EdgeManagerCreateReply) String() string { return proto.CompactTextString(m) }
func (*EdgeManagerCreateReply) ProtoMessage()    {}
func (*EdgeManagerCreateReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_cab1176173a95651, []int{1}
}

func (m *EdgeManagerCreateReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EdgeManagerCreateReply.Unmarshal(m, b)
}
func (m *EdgeManagerCreateReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EdgeManagerCreateReply.Marshal(b, m, deterministic)
}
func (m *EdgeManagerCreateReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EdgeManagerCreateReply.Merge(m, src)
}
func (m *EdgeManagerCreateReply) XXX_Size() int {
	return xxx_messageInfo_EdgeManagerCreateReply.Size(m)
}
func (m *EdgeManagerCreateReply) XXX_DiscardUnknown() {
	xxx_messageInfo_EdgeManagerCreateReply.DiscardUnknown(m)
}

var xxx_messageInfo_EdgeManagerCreateReply proto.InternalMessageInfo

func (m *EdgeManagerCreateReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*EdgeManagerCreateArgs)(nil), "gon2n.EdgeManagerCreateArgs")
	proto.RegisterType((*EdgeManagerCreateReply)(nil), "gon2n.EdgeManagerCreateReply")
}

func init() { proto.RegisterFile("edge.proto", fileDescriptor_cab1176173a95651) }

var fileDescriptor_cab1176173a95651 = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0x80, 0x69, 0x4b, 0x4b, 0x7b, 0xdd, 0x4a, 0x67, 0x06, 0xb2, 0xa6, 0x81, 0xaa, 0x0a, 0xa1,
	0x0a, 0xa1, 0x3e, 0x94, 0x5f, 0x50, 0xb5, 0x13, 0x44, 0x2c, 0x10, 0x65, 0x19, 0xef, 0x5e, 0x72,
	0x84, 0x88, 0xc4, 0x8e, 0x6c, 0xb7, 0x28, 0xbf, 0x1e, 0xe4, 0xcb, 0xd8, 0x9a, 0x65, 0xbc, 0xc5,
	0xdf, 0x7d, 0x39, 0xdf, 0xd9, 0x3e, 0x00, 0x4c, 0x52, 0x5c, 0x96, 0x5a, 0x59, 0xc5, 0xfa, 0xa9,
	0x92, 0x2b, 0x39, 0xff, 0xd3, 0x87, 0x97, 0x17, 0x49, 0x8a, 0xbe, 0x90, 0x22, 0x45, 0xbd, 0xd1,
	0x28, 0x2c, 0xae, 0x75, 0x6a, 0xd8, 0x19, 0x0c, 0xd7, 0x79, 0xae, 0x7e, 0x07, 0xab, 0x80, 0x77,
	0x66, 0x9d, 0xc5, 0x30, 0xbc, 0x5b, 0xb3, 0x39, 0x1c, 0xd1, 0x77, 0xa8, 0x76, 0x36, 0x93, 0x29,
	0xef, 0x52, 0xbc, 0xc1, 0xd8, 0x5b, 0x38, 0xde, 0xa8, 0xa2, 0xd8, 0xc9, 0xcc, 0x56, 0x5f, 0x45,
	0x81, 0xbc, 0x37, 0xeb, 0x2c, 0x46, 0x61, 0x13, 0xb2, 0x15, 0x9c, 0x6e, 0x33, 0x23, 0x6e, 0x72,
	0x0c, 0xfc, 0xe8, 0x7a, 0x9b, 0x99, 0x58, 0xed, 0x51, 0x57, 0xfc, 0x29, 0x65, 0x7c, 0x34, 0xc6,
	0xde, 0xc3, 0xf4, 0x96, 0xfb, 0xbb, 0xdc, 0x66, 0xb1, 0x30, 0x96, 0xf7, 0xc9, 0x6f, 0x71, 0x57,
	0xc5, 0xb6, 0x92, 0xa2, 0xc8, 0x62, 0x2f, 0xf0, 0x55, 0x82, 0x7c, 0x40, 0x62, 0x13, 0x3a, 0xeb,
	0x42, 0xc6, 0xba, 0x2a, 0x6d, 0xa6, 0xe4, 0x17, 0xac, 0xf8, 0xb3, 0xba, 0xd6, 0x06, 0x64, 0xe7,
	0x30, 0xba, 0x54, 0xb1, 0xc8, 0x03, 0xa5, 0x2d, 0x1f, 0xce, 0x3a, 0x8b, 0x5e, 0x78, 0x0f, 0xd8,
	0x3b, 0x98, 0xd4, 0x87, 0x58, 0xa0, 0xb4, 0xa4, 0x8c, 0x48, 0x79, 0x40, 0x5d, 0xf5, 0x21, 0xa6,
	0x99, 0xb1, 0xa8, 0x3d, 0x69, 0x51, 0xef, 0x45, 0xce, 0x81, 0xcc, 0x16, 0x67, 0x33, 0x18, 0xff,
	0x63, 0x51, 0x74, 0xc9, 0xc7, 0xa4, 0x1d, 0x22, 0xf6, 0x01, 0x4e, 0xae, 0x76, 0x25, 0x6a, 0xa9,
	0x12, 0xfc, 0xac, 0x4c, 0xbd, 0xf1, 0x11, 0x55, 0xdf, 0x0e, 0xb8, 0x3e, 0xa3, 0xaa, 0xc4, 0x6f,
	0x3f, 0xae, 0x50, 0xef, 0xb3, 0x18, 0xf9, 0x31, 0x65, 0x6c, 0x42, 0x57, 0xe1, 0x7d, 0xe3, 0x3e,
	0xda, 0x9f, 0x2a, 0xe1, 0x93, 0xba, 0xc2, 0x87, 0x9c, 0xbd, 0x01, 0xd8, 0xa2, 0xfb, 0x8b, 0xae,
	0xf8, 0x39, 0x6d, 0x7c, 0x40, 0x5c, 0x07, 0xeb, 0x24, 0xd1, 0x68, 0x0c, 0x9d, 0xfe, 0x94, 0x84,
	0x43, 0xe4, 0xde, 0x59, 0xed, 0x7b, 0x01, 0x3f, 0xa1, 0xf0, 0xdd, 0x9a, 0x6e, 0xaf, 0xce, 0x85,
	0xb6, 0x10, 0xe6, 0x17, 0x67, 0xf5, 0xbd, 0x34, 0x20, 0xbd, 0x07, 0x02, 0xfe, 0x7a, 0x73, 0x9b,
	0x99, 0xbf, 0x20, 0xb1, 0xc5, 0xd9, 0x14, 0x7a, 0x7e, 0x74, 0xcd, 0x4f, 0xa9, 0x1d, 0xf7, 0x39,
	0x5f, 0xc0, 0xab, 0xd6, 0x00, 0x84, 0x58, 0xe6, 0x15, 0x9b, 0x40, 0xd7, 0x4b, 0xe8, 0xed, 0x8f,
	0xc2, 0xae, 0x97, 0xac, 0xbe, 0xc3, 0xf8, 0xc0, 0x64, 0x9f, 0x60, 0x50, 0xdb, 0xec, 0x7c, 0x49,
	0xc3, 0xb4, 0x7c, 0x74, 0x90, 0xce, 0x5e, 0xff, 0x2f, 0x4a, 0xbb, 0xcc, 0x9f, 0xdc, 0x0c, 0x68,
	0x22, 0x3f, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x97, 0x7f, 0xa4, 0xa7, 0x9f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EdgeManagerClient is the client API for EdgeManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EdgeManagerClient interface {
	Create(ctx context.Context, in *EdgeManagerCreateArgs, opts ...grpc.CallOption) (*EdgeManagerCreateReply, error)
}

type edgeManagerClient struct {
	cc *grpc.ClientConn
}

func NewEdgeManagerClient(cc *grpc.ClientConn) EdgeManagerClient {
	return &edgeManagerClient{cc}
}

func (c *edgeManagerClient) Create(ctx context.Context, in *EdgeManagerCreateArgs, opts ...grpc.CallOption) (*EdgeManagerCreateReply, error) {
	out := new(EdgeManagerCreateReply)
	err := c.cc.Invoke(ctx, "/gon2n.EdgeManager/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EdgeManagerServer is the server API for EdgeManager service.
type EdgeManagerServer interface {
	Create(context.Context, *EdgeManagerCreateArgs) (*EdgeManagerCreateReply, error)
}

// UnimplementedEdgeManagerServer can be embedded to have forward compatible implementations.
type UnimplementedEdgeManagerServer struct {
}

func (*UnimplementedEdgeManagerServer) Create(ctx context.Context, req *EdgeManagerCreateArgs) (*EdgeManagerCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func RegisterEdgeManagerServer(s *grpc.Server, srv EdgeManagerServer) {
	s.RegisterService(&_EdgeManager_serviceDesc, srv)
}

func _EdgeManager_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EdgeManagerCreateArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EdgeManagerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gon2n.EdgeManager/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EdgeManagerServer).Create(ctx, req.(*EdgeManagerCreateArgs))
	}
	return interceptor(ctx, in, info, handler)
}

var _EdgeManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gon2n.EdgeManager",
	HandlerType: (*EdgeManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _EdgeManager_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "edge.proto",
}