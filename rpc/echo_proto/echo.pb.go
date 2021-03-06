// Code generated by protoc-gen-go. DO NOT EDIT.
// source: echo.proto

package echo_proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
type EchoReq struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoReq) Reset()         { *m = EchoReq{} }
func (m *EchoReq) String() string { return proto.CompactTextString(m) }
func (*EchoReq) ProtoMessage()    {}
func (*EchoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_08134aea513e0001, []int{0}
}

func (m *EchoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoReq.Unmarshal(m, b)
}
func (m *EchoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoReq.Marshal(b, m, deterministic)
}
func (m *EchoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoReq.Merge(m, src)
}
func (m *EchoReq) XXX_Size() int {
	return xxx_messageInfo_EchoReq.Size(m)
}
func (m *EchoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoReq.DiscardUnknown(m)
}

var xxx_messageInfo_EchoReq proto.InternalMessageInfo

func (m *EchoReq) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type EchoRes struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoRes) Reset()         { *m = EchoRes{} }
func (m *EchoRes) String() string { return proto.CompactTextString(m) }
func (*EchoRes) ProtoMessage()    {}
func (*EchoRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_08134aea513e0001, []int{1}
}

func (m *EchoRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoRes.Unmarshal(m, b)
}
func (m *EchoRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoRes.Marshal(b, m, deterministic)
}
func (m *EchoRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoRes.Merge(m, src)
}
func (m *EchoRes) XXX_Size() int {
	return xxx_messageInfo_EchoRes.Size(m)
}
func (m *EchoRes) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoRes.DiscardUnknown(m)
}

var xxx_messageInfo_EchoRes proto.InternalMessageInfo

func (m *EchoRes) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoReq)(nil), "echo_proto.EchoReq")
	proto.RegisterType((*EchoRes)(nil), "echo_proto.EchoRes")
}

func init() { proto.RegisterFile("echo.proto", fileDescriptor_08134aea513e0001) }

var fileDescriptor_08134aea513e0001 = []byte{
	// 102 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4d, 0xce, 0xc8,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x02, 0xb3, 0xe3, 0xc1, 0x6c, 0x25, 0x69, 0x2e, 0x76,
	0xd7, 0xe4, 0x8c, 0xfc, 0xa0, 0xd4, 0x42, 0x21, 0x01, 0x2e, 0xe6, 0xdc, 0xe2, 0x74, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xce, 0x20, 0x10, 0x13, 0x21, 0x59, 0x8c, 0x29, 0x69, 0x64, 0xc1, 0xc5, 0x02,
	0x92, 0x14, 0x32, 0xe0, 0x62, 0x01, 0x99, 0x27, 0x24, 0xac, 0x87, 0x30, 0x56, 0x0f, 0x6a, 0xa6,
	0x14, 0x16, 0xc1, 0xe2, 0x24, 0x36, 0x30, 0xd7, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x57,
	0xcd, 0x73, 0x94, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EchoClient is the client API for Echo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoClient interface {
	Echo(ctx context.Context, in *EchoReq, opts ...grpc.CallOption) (*EchoRes, error)
}

type echoClient struct {
	cc *grpc.ClientConn
}

func NewEchoClient(cc *grpc.ClientConn) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) Echo(ctx context.Context, in *EchoReq, opts ...grpc.CallOption) (*EchoRes, error) {
	out := new(EchoRes)
	err := c.cc.Invoke(ctx, "/echo_proto.Echo/echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServer is the server API for Echo service.
type EchoServer interface {
	Echo(context.Context, *EchoReq) (*EchoRes, error)
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo_proto.Echo/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Echo(ctx, req.(*EchoReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echo_proto.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "echo",
			Handler:    _Echo_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "echo.proto",
}
