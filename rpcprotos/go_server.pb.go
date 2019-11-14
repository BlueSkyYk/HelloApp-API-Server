// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go_server.proto

package rpcprotos

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

//用户登录
type LoginParam struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginParam) Reset()         { *m = LoginParam{} }
func (m *LoginParam) String() string { return proto.CompactTextString(m) }
func (*LoginParam) ProtoMessage()    {}
func (*LoginParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d212c4d347d3e7b, []int{0}
}

func (m *LoginParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginParam.Unmarshal(m, b)
}
func (m *LoginParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginParam.Marshal(b, m, deterministic)
}
func (m *LoginParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginParam.Merge(m, src)
}
func (m *LoginParam) XXX_Size() int {
	return xxx_messageInfo_LoginParam.Size(m)
}
func (m *LoginParam) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginParam.DiscardUnknown(m)
}

var xxx_messageInfo_LoginParam proto.InternalMessageInfo

func (m *LoginParam) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *LoginParam) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

//ServerToClientMsgResult
type ServerToClientMsgResultParam struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	MsgId                int32    `protobuf:"varint,2,opt,name=msgId,proto3" json:"msgId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerToClientMsgResultParam) Reset()         { *m = ServerToClientMsgResultParam{} }
func (m *ServerToClientMsgResultParam) String() string { return proto.CompactTextString(m) }
func (*ServerToClientMsgResultParam) ProtoMessage()    {}
func (*ServerToClientMsgResultParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d212c4d347d3e7b, []int{1}
}

func (m *ServerToClientMsgResultParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerToClientMsgResultParam.Unmarshal(m, b)
}
func (m *ServerToClientMsgResultParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerToClientMsgResultParam.Marshal(b, m, deterministic)
}
func (m *ServerToClientMsgResultParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerToClientMsgResultParam.Merge(m, src)
}
func (m *ServerToClientMsgResultParam) XXX_Size() int {
	return xxx_messageInfo_ServerToClientMsgResultParam.Size(m)
}
func (m *ServerToClientMsgResultParam) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerToClientMsgResultParam.DiscardUnknown(m)
}

var xxx_messageInfo_ServerToClientMsgResultParam proto.InternalMessageInfo

func (m *ServerToClientMsgResultParam) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ServerToClientMsgResultParam) GetMsgId() int32 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

//JavaBaseResult
type GoBaseResult struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoBaseResult) Reset()         { *m = GoBaseResult{} }
func (m *GoBaseResult) String() string { return proto.CompactTextString(m) }
func (*GoBaseResult) ProtoMessage()    {}
func (*GoBaseResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d212c4d347d3e7b, []int{2}
}

func (m *GoBaseResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoBaseResult.Unmarshal(m, b)
}
func (m *GoBaseResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoBaseResult.Marshal(b, m, deterministic)
}
func (m *GoBaseResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoBaseResult.Merge(m, src)
}
func (m *GoBaseResult) XXX_Size() int {
	return xxx_messageInfo_GoBaseResult.Size(m)
}
func (m *GoBaseResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GoBaseResult.DiscardUnknown(m)
}

var xxx_messageInfo_GoBaseResult proto.InternalMessageInfo

func (m *GoBaseResult) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GoBaseResult) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginParam)(nil), "rpcprotos.LoginParam")
	proto.RegisterType((*ServerToClientMsgResultParam)(nil), "rpcprotos.ServerToClientMsgResultParam")
	proto.RegisterType((*GoBaseResult)(nil), "rpcprotos.GoBaseResult")
}

func init() { proto.RegisterFile("go_server.proto", fileDescriptor_9d212c4d347d3e7b) }

var fileDescriptor_9d212c4d347d3e7b = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xed, 0xea, 0x5a, 0x77, 0x10, 0x84, 0xb0, 0x62, 0x59, 0x3d, 0x2c, 0xbd, 0xe8, 0x29,
	0x07, 0x05, 0x2f, 0x7a, 0xaa, 0x87, 0x5a, 0x70, 0xa1, 0x44, 0xef, 0x12, 0xb3, 0x31, 0x16, 0xdb,
	0x4e, 0x48, 0x52, 0xc5, 0xbf, 0xe3, 0x8f, 0xf0, 0xf7, 0x89, 0xc9, 0x76, 0x2d, 0x22, 0x7b, 0x9b,
	0x97, 0xe1, 0xbd, 0x79, 0x5f, 0xe0, 0x40, 0xe1, 0xa3, 0x95, 0xe6, 0x4d, 0x1a, 0xaa, 0x0d, 0x3a,
	0x24, 0x13, 0xa3, 0x85, 0x9f, 0x6c, 0x9a, 0x01, 0xdc, 0xa1, 0xaa, 0xda, 0x92, 0x1b, 0xde, 0x90,
	0x04, 0x62, 0x2e, 0x04, 0x76, 0xad, 0x4b, 0xa2, 0x79, 0x74, 0x36, 0x61, 0xbd, 0x24, 0x33, 0xd8,
	0xd3, 0xdc, 0xda, 0x77, 0x34, 0xcb, 0x64, 0xe4, 0x57, 0x6b, 0x9d, 0xde, 0xc2, 0xc9, 0xbd, 0x8f,
	0x7f, 0xc0, 0x9b, 0xba, 0x92, 0xad, 0x5b, 0x58, 0xc5, 0xa4, 0xed, 0x6a, 0x17, 0x52, 0x09, 0xec,
	0x08, 0x5c, 0x4a, 0x1f, 0x39, 0x66, 0x7e, 0x26, 0x53, 0x18, 0x37, 0x56, 0x15, 0x21, 0x6c, 0xcc,
	0x82, 0x48, 0xaf, 0x61, 0x3f, 0xc7, 0x8c, 0x5b, 0x19, 0xec, 0xff, 0x3a, 0x13, 0x88, 0x1b, 0x69,
	0x2d, 0x57, 0x72, 0x55, 0xa4, 0x97, 0xe7, 0x5f, 0x11, 0x4c, 0x73, 0x64, 0x5a, 0x84, 0x36, 0x45,
	0xeb, 0xa4, 0x79, 0xe6, 0x42, 0x92, 0x2b, 0x88, 0x8b, 0x85, 0xc7, 0x24, 0x87, 0x74, 0xcd, 0x4e,
	0x7f, 0xc1, 0x67, 0x47, 0x83, 0xe7, 0x61, 0x83, 0x74, 0x8b, 0x70, 0x38, 0xfe, 0x43, 0x17, 0xce,
	0xad, 0x2a, 0x9e, 0x0e, 0x9c, 0x9b, 0x7e, 0x61, 0xc3, 0x89, 0xec, 0x12, 0xe6, 0x02, 0x1b, 0x2a,
	0xba, 0xca, 0x7d, 0xbc, 0xd2, 0x17, 0x59, 0xd7, 0xc8, 0xb5, 0xae, 0x1a, 0xaa, 0xb0, 0xf7, 0x64,
	0xe0, 0xc9, 0xca, 0x9f, 0xb9, 0x8c, 0x3e, 0x47, 0xdb, 0x39, 0x2b, 0x9f, 0x76, 0xfd, 0xea, 0xe2,
	0x3b, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x49, 0xe0, 0xb8, 0xe1, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GoRpcServerInterfaceClient is the client API for GoRpcServerInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GoRpcServerInterfaceClient interface {
	//登录IM
	IMLogin(ctx context.Context, in *LoginParam, opts ...grpc.CallOption) (*GoBaseResult, error)
	//Server To Client 消息发送结果回调
	ServerToClientMessageResult(ctx context.Context, in *ServerToClientMsgResultParam, opts ...grpc.CallOption) (*GoBaseResult, error)
}

type goRpcServerInterfaceClient struct {
	cc *grpc.ClientConn
}

func NewGoRpcServerInterfaceClient(cc *grpc.ClientConn) GoRpcServerInterfaceClient {
	return &goRpcServerInterfaceClient{cc}
}

func (c *goRpcServerInterfaceClient) IMLogin(ctx context.Context, in *LoginParam, opts ...grpc.CallOption) (*GoBaseResult, error) {
	out := new(GoBaseResult)
	err := c.cc.Invoke(ctx, "/rpcprotos.GoRpcServerInterface/IMLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goRpcServerInterfaceClient) ServerToClientMessageResult(ctx context.Context, in *ServerToClientMsgResultParam, opts ...grpc.CallOption) (*GoBaseResult, error) {
	out := new(GoBaseResult)
	err := c.cc.Invoke(ctx, "/rpcprotos.GoRpcServerInterface/ServerToClientMessageResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoRpcServerInterfaceServer is the server API for GoRpcServerInterface service.
type GoRpcServerInterfaceServer interface {
	//登录IM
	IMLogin(context.Context, *LoginParam) (*GoBaseResult, error)
	//Server To Client 消息发送结果回调
	ServerToClientMessageResult(context.Context, *ServerToClientMsgResultParam) (*GoBaseResult, error)
}

// UnimplementedGoRpcServerInterfaceServer can be embedded to have forward compatible implementations.
type UnimplementedGoRpcServerInterfaceServer struct {
}

func (*UnimplementedGoRpcServerInterfaceServer) IMLogin(ctx context.Context, req *LoginParam) (*GoBaseResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IMLogin not implemented")
}
func (*UnimplementedGoRpcServerInterfaceServer) ServerToClientMessageResult(ctx context.Context, req *ServerToClientMsgResultParam) (*GoBaseResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServerToClientMessageResult not implemented")
}

func RegisterGoRpcServerInterfaceServer(s *grpc.Server, srv GoRpcServerInterfaceServer) {
	s.RegisterService(&_GoRpcServerInterface_serviceDesc, srv)
}

func _GoRpcServerInterface_IMLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoRpcServerInterfaceServer).IMLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcprotos.GoRpcServerInterface/IMLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoRpcServerInterfaceServer).IMLogin(ctx, req.(*LoginParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoRpcServerInterface_ServerToClientMessageResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerToClientMsgResultParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoRpcServerInterfaceServer).ServerToClientMessageResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcprotos.GoRpcServerInterface/ServerToClientMessageResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoRpcServerInterfaceServer).ServerToClientMessageResult(ctx, req.(*ServerToClientMsgResultParam))
	}
	return interceptor(ctx, in, info, handler)
}

var _GoRpcServerInterface_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpcprotos.GoRpcServerInterface",
	HandlerType: (*GoRpcServerInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IMLogin",
			Handler:    _GoRpcServerInterface_IMLogin_Handler,
		},
		{
			MethodName: "ServerToClientMessageResult",
			Handler:    _GoRpcServerInterface_ServerToClientMessageResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go_server.proto",
}