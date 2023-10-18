// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.0--rc3
// source: stock.proto

// 指定包名

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	StockService_GetStockById_FullMethodName = "/pb.StockService/GetStockById"
)

// StockServiceClient is the client API for StockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StockServiceClient interface {
	// GetStockById 持续获取股票信息
	GetStockById(ctx context.Context, in *StockReq, opts ...grpc.CallOption) (StockService_GetStockByIdClient, error)
}

type stockServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStockServiceClient(cc grpc.ClientConnInterface) StockServiceClient {
	return &stockServiceClient{cc}
}

func (c *stockServiceClient) GetStockById(ctx context.Context, in *StockReq, opts ...grpc.CallOption) (StockService_GetStockByIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &StockService_ServiceDesc.Streams[0], StockService_GetStockById_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &stockServiceGetStockByIdClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StockService_GetStockByIdClient interface {
	Recv() (*StockResp, error)
	grpc.ClientStream
}

type stockServiceGetStockByIdClient struct {
	grpc.ClientStream
}

func (x *stockServiceGetStockByIdClient) Recv() (*StockResp, error) {
	m := new(StockResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StockServiceServer is the server API for StockService service.
// All implementations must embed UnimplementedStockServiceServer
// for forward compatibility
type StockServiceServer interface {
	// GetStockById 持续获取股票信息
	GetStockById(*StockReq, StockService_GetStockByIdServer) error
	mustEmbedUnimplementedStockServiceServer()
}

// UnimplementedStockServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStockServiceServer struct {
}

func (UnimplementedStockServiceServer) GetStockById(*StockReq, StockService_GetStockByIdServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStockById not implemented")
}
func (UnimplementedStockServiceServer) mustEmbedUnimplementedStockServiceServer() {}

// UnsafeStockServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StockServiceServer will
// result in compilation errors.
type UnsafeStockServiceServer interface {
	mustEmbedUnimplementedStockServiceServer()
}

func RegisterStockServiceServer(s grpc.ServiceRegistrar, srv StockServiceServer) {
	s.RegisterService(&StockService_ServiceDesc, srv)
}

func _StockService_GetStockById_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StockReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StockServiceServer).GetStockById(m, &stockServiceGetStockByIdServer{stream})
}

type StockService_GetStockByIdServer interface {
	Send(*StockResp) error
	grpc.ServerStream
}

type stockServiceGetStockByIdServer struct {
	grpc.ServerStream
}

func (x *stockServiceGetStockByIdServer) Send(m *StockResp) error {
	return x.ServerStream.SendMsg(m)
}

// StockService_ServiceDesc is the grpc.ServiceDesc for StockService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StockService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.StockService",
	HandlerType: (*StockServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStockById",
			Handler:       _StockService_GetStockById_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "stock.proto",
}
