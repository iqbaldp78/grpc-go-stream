// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// AvgAlarmServiceClient is the client API for AvgAlarmService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AvgAlarmServiceClient interface {
	// Push a number to aggregate Returns an error when average exceeds float range
	Push(ctx context.Context, in *PushReq, opts ...grpc.CallOption) (*Empty, error)
	// Gets a stream of notifications when current average exceeds the threshold provided in request
	// Never returns an error
	GetAvgAlarmStream(ctx context.Context, in *GetAvgAlarmStreamReq, opts ...grpc.CallOption) (AvgAlarmService_GetAvgAlarmStreamClient, error)
}

type avgAlarmServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAvgAlarmServiceClient(cc grpc.ClientConnInterface) AvgAlarmServiceClient {
	return &avgAlarmServiceClient{cc}
}

func (c *avgAlarmServiceClient) Push(ctx context.Context, in *PushReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.AvgAlarmService/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *avgAlarmServiceClient) GetAvgAlarmStream(ctx context.Context, in *GetAvgAlarmStreamReq, opts ...grpc.CallOption) (AvgAlarmService_GetAvgAlarmStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &AvgAlarmService_ServiceDesc.Streams[0], "/pb.AvgAlarmService/GetAvgAlarmStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &avgAlarmServiceGetAvgAlarmStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AvgAlarmService_GetAvgAlarmStreamClient interface {
	Recv() (*GetAvgAlarmStreamRes, error)
	grpc.ClientStream
}

type avgAlarmServiceGetAvgAlarmStreamClient struct {
	grpc.ClientStream
}

func (x *avgAlarmServiceGetAvgAlarmStreamClient) Recv() (*GetAvgAlarmStreamRes, error) {
	m := new(GetAvgAlarmStreamRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AvgAlarmServiceServer is the server API for AvgAlarmService service.
// All implementations must embed UnimplementedAvgAlarmServiceServer
// for forward compatibility
type AvgAlarmServiceServer interface {
	// Push a number to aggregate Returns an error when average exceeds float range
	Push(context.Context, *PushReq) (*Empty, error)
	// Gets a stream of notifications when current average exceeds the threshold provided in request
	// Never returns an error
	GetAvgAlarmStream(*GetAvgAlarmStreamReq, AvgAlarmService_GetAvgAlarmStreamServer) error
	mustEmbedUnimplementedAvgAlarmServiceServer()
}

// UnimplementedAvgAlarmServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAvgAlarmServiceServer struct {
}

func (UnimplementedAvgAlarmServiceServer) Push(context.Context, *PushReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (UnimplementedAvgAlarmServiceServer) GetAvgAlarmStream(*GetAvgAlarmStreamReq, AvgAlarmService_GetAvgAlarmStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAvgAlarmStream not implemented")
}
func (UnimplementedAvgAlarmServiceServer) mustEmbedUnimplementedAvgAlarmServiceServer() {}

// UnsafeAvgAlarmServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AvgAlarmServiceServer will
// result in compilation errors.
type UnsafeAvgAlarmServiceServer interface {
	mustEmbedUnimplementedAvgAlarmServiceServer()
}

func RegisterAvgAlarmServiceServer(s grpc.ServiceRegistrar, srv AvgAlarmServiceServer) {
	s.RegisterService(&AvgAlarmService_ServiceDesc, srv)
}

func _AvgAlarmService_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvgAlarmServiceServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AvgAlarmService/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvgAlarmServiceServer).Push(ctx, req.(*PushReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AvgAlarmService_GetAvgAlarmStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAvgAlarmStreamReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AvgAlarmServiceServer).GetAvgAlarmStream(m, &avgAlarmServiceGetAvgAlarmStreamServer{stream})
}

type AvgAlarmService_GetAvgAlarmStreamServer interface {
	Send(*GetAvgAlarmStreamRes) error
	grpc.ServerStream
}

type avgAlarmServiceGetAvgAlarmStreamServer struct {
	grpc.ServerStream
}

func (x *avgAlarmServiceGetAvgAlarmStreamServer) Send(m *GetAvgAlarmStreamRes) error {
	return x.ServerStream.SendMsg(m)
}

// AvgAlarmService_ServiceDesc is the grpc.ServiceDesc for AvgAlarmService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AvgAlarmService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AvgAlarmService",
	HandlerType: (*AvgAlarmServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Push",
			Handler:    _AvgAlarmService_Push_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAvgAlarmStream",
			Handler:       _AvgAlarmService_GetAvgAlarmStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "modules/pb/pb.proto",
}