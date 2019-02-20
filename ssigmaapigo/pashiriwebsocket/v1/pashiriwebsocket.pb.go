// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ssigmaapi/pashiriwebsocket/v1/pashiriwebsocket.proto

package pashiriwebsocket

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	market "github.com/sansigma/ssigmaapigo/type/market"
	trade "github.com/sansigma/ssigmaapigo/type/trade"
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

func init() {
	proto.RegisterFile("ssigmaapi/pashiriwebsocket/v1/pashiriwebsocket.proto", fileDescriptor_0b61436759cfed1c)
}

var fileDescriptor_0b61436759cfed1c = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x29, 0x2e, 0xce, 0x4c,
	0xcf, 0x4d, 0x4c, 0x2c, 0xc8, 0xd4, 0x2f, 0x48, 0x2c, 0xce, 0xc8, 0x2c, 0xca, 0x2c, 0x4f, 0x4d,
	0x2a, 0xce, 0x4f, 0xce, 0x4e, 0x2d, 0xd1, 0x2f, 0x33, 0xc4, 0x10, 0xd3, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x92, 0x85, 0xeb, 0xd2, 0xc3, 0x50, 0x51, 0x66, 0x28, 0x25, 0x8d, 0x30, 0xb4, 0xa4,
	0xb2, 0x20, 0x55, 0x3f, 0x37, 0xb1, 0x08, 0xae, 0x57, 0x4a, 0x0a, 0x4d, 0xb2, 0xa4, 0x28, 0x31,
	0x25, 0x15, 0x22, 0x67, 0x14, 0xc3, 0x25, 0x19, 0x00, 0x31, 0x2f, 0x1c, 0x66, 0x5e, 0x70, 0x6a,
	0x51, 0x59, 0x66, 0x72, 0x6a, 0x98, 0xa1, 0x90, 0x3d, 0x17, 0xa7, 0x7b, 0x6a, 0x49, 0x08, 0x48,
	0x79, 0xb1, 0x90, 0x98, 0x1e, 0xc2, 0x09, 0x20, 0x63, 0xf4, 0x7c, 0xc1, 0x76, 0x48, 0x61, 0x88,
	0x43, 0xd4, 0x2b, 0x31, 0x18, 0x30, 0x3a, 0xb9, 0x45, 0xb9, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26,
	0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x17, 0x27, 0xe6, 0x81, 0x55, 0xea, 0xc3, 0x35, 0xa4, 0xe7, 0x63,
	0x0b, 0x03, 0x6b, 0x74, 0xb1, 0x24, 0x36, 0xb0, 0x63, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xf9, 0xa8, 0xf6, 0xc2, 0x3c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PashiriWebsocketServiceV1Client is the client API for PashiriWebsocketServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PashiriWebsocketServiceV1Client interface {
	GetTrades(ctx context.Context, in *market.Market, opts ...grpc.CallOption) (PashiriWebsocketServiceV1_GetTradesClient, error)
}

type pashiriWebsocketServiceV1Client struct {
	cc *grpc.ClientConn
}

func NewPashiriWebsocketServiceV1Client(cc *grpc.ClientConn) PashiriWebsocketServiceV1Client {
	return &pashiriWebsocketServiceV1Client{cc}
}

func (c *pashiriWebsocketServiceV1Client) GetTrades(ctx context.Context, in *market.Market, opts ...grpc.CallOption) (PashiriWebsocketServiceV1_GetTradesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PashiriWebsocketServiceV1_serviceDesc.Streams[0], "/ssigmaapi.pashiriwebsocket.v1.PashiriWebsocketServiceV1/GetTrades", opts...)
	if err != nil {
		return nil, err
	}
	x := &pashiriWebsocketServiceV1GetTradesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PashiriWebsocketServiceV1_GetTradesClient interface {
	Recv() (*trade.Trades, error)
	grpc.ClientStream
}

type pashiriWebsocketServiceV1GetTradesClient struct {
	grpc.ClientStream
}

func (x *pashiriWebsocketServiceV1GetTradesClient) Recv() (*trade.Trades, error) {
	m := new(trade.Trades)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PashiriWebsocketServiceV1Server is the server API for PashiriWebsocketServiceV1 service.
type PashiriWebsocketServiceV1Server interface {
	GetTrades(*market.Market, PashiriWebsocketServiceV1_GetTradesServer) error
}

func RegisterPashiriWebsocketServiceV1Server(s *grpc.Server, srv PashiriWebsocketServiceV1Server) {
	s.RegisterService(&_PashiriWebsocketServiceV1_serviceDesc, srv)
}

func _PashiriWebsocketServiceV1_GetTrades_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(market.Market)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PashiriWebsocketServiceV1Server).GetTrades(m, &pashiriWebsocketServiceV1GetTradesServer{stream})
}

type PashiriWebsocketServiceV1_GetTradesServer interface {
	Send(*trade.Trades) error
	grpc.ServerStream
}

type pashiriWebsocketServiceV1GetTradesServer struct {
	grpc.ServerStream
}

func (x *pashiriWebsocketServiceV1GetTradesServer) Send(m *trade.Trades) error {
	return x.ServerStream.SendMsg(m)
}

var _PashiriWebsocketServiceV1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ssigmaapi.pashiriwebsocket.v1.PashiriWebsocketServiceV1",
	HandlerType: (*PashiriWebsocketServiceV1Server)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTrades",
			Handler:       _PashiriWebsocketServiceV1_GetTrades_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ssigmaapi/pashiriwebsocket/v1/pashiriwebsocket.proto",
}
