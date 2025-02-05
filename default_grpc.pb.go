// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: default.proto

package hedgeproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Hedge_Send_FullMethodName      = "/hedgeproto.Hedge/Send"
	Hedge_Broadcast_FullMethodName = "/hedgeproto.Hedge/Broadcast"
	Hedge_SoSWrite_FullMethodName  = "/hedgeproto.Hedge/SoSWrite"
	Hedge_SoSRead_FullMethodName   = "/hedgeproto.Hedge/SoSRead"
	Hedge_SoSClose_FullMethodName  = "/hedgeproto.Hedge/SoSClose"
)

// HedgeClient is the client API for Hedge service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HedgeClient interface {
	Send(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Payload, Payload], error)
	Broadcast(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Payload, Payload], error)
	SoSWrite(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Payload, Payload], error)
	SoSRead(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Payload, Payload], error)
	SoSClose(ctx context.Context, in *Payload, opts ...grpc.CallOption) (*Payload, error)
}

type hedgeClient struct {
	cc grpc.ClientConnInterface
}

func NewHedgeClient(cc grpc.ClientConnInterface) HedgeClient {
	return &hedgeClient{cc}
}

func (c *hedgeClient) Send(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Payload, Payload], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Hedge_ServiceDesc.Streams[0], Hedge_Send_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Payload, Payload]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Hedge_SendClient = grpc.BidiStreamingClient[Payload, Payload]

func (c *hedgeClient) Broadcast(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Payload, Payload], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Hedge_ServiceDesc.Streams[1], Hedge_Broadcast_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Payload, Payload]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Hedge_BroadcastClient = grpc.BidiStreamingClient[Payload, Payload]

func (c *hedgeClient) SoSWrite(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Payload, Payload], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Hedge_ServiceDesc.Streams[2], Hedge_SoSWrite_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Payload, Payload]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Hedge_SoSWriteClient = grpc.BidiStreamingClient[Payload, Payload]

func (c *hedgeClient) SoSRead(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Payload, Payload], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Hedge_ServiceDesc.Streams[3], Hedge_SoSRead_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Payload, Payload]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Hedge_SoSReadClient = grpc.BidiStreamingClient[Payload, Payload]

func (c *hedgeClient) SoSClose(ctx context.Context, in *Payload, opts ...grpc.CallOption) (*Payload, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Payload)
	err := c.cc.Invoke(ctx, Hedge_SoSClose_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HedgeServer is the server API for Hedge service.
// All implementations must embed UnimplementedHedgeServer
// for forward compatibility.
type HedgeServer interface {
	Send(grpc.BidiStreamingServer[Payload, Payload]) error
	Broadcast(grpc.BidiStreamingServer[Payload, Payload]) error
	SoSWrite(grpc.BidiStreamingServer[Payload, Payload]) error
	SoSRead(grpc.BidiStreamingServer[Payload, Payload]) error
	SoSClose(context.Context, *Payload) (*Payload, error)
	mustEmbedUnimplementedHedgeServer()
}

// UnimplementedHedgeServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHedgeServer struct{}

func (UnimplementedHedgeServer) Send(grpc.BidiStreamingServer[Payload, Payload]) error {
	return status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedHedgeServer) Broadcast(grpc.BidiStreamingServer[Payload, Payload]) error {
	return status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedHedgeServer) SoSWrite(grpc.BidiStreamingServer[Payload, Payload]) error {
	return status.Errorf(codes.Unimplemented, "method SoSWrite not implemented")
}
func (UnimplementedHedgeServer) SoSRead(grpc.BidiStreamingServer[Payload, Payload]) error {
	return status.Errorf(codes.Unimplemented, "method SoSRead not implemented")
}
func (UnimplementedHedgeServer) SoSClose(context.Context, *Payload) (*Payload, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SoSClose not implemented")
}
func (UnimplementedHedgeServer) mustEmbedUnimplementedHedgeServer() {}
func (UnimplementedHedgeServer) testEmbeddedByValue()               {}

// UnsafeHedgeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HedgeServer will
// result in compilation errors.
type UnsafeHedgeServer interface {
	mustEmbedUnimplementedHedgeServer()
}

func RegisterHedgeServer(s grpc.ServiceRegistrar, srv HedgeServer) {
	// If the following call pancis, it indicates UnimplementedHedgeServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Hedge_ServiceDesc, srv)
}

func _Hedge_Send_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HedgeServer).Send(&grpc.GenericServerStream[Payload, Payload]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Hedge_SendServer = grpc.BidiStreamingServer[Payload, Payload]

func _Hedge_Broadcast_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HedgeServer).Broadcast(&grpc.GenericServerStream[Payload, Payload]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Hedge_BroadcastServer = grpc.BidiStreamingServer[Payload, Payload]

func _Hedge_SoSWrite_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HedgeServer).SoSWrite(&grpc.GenericServerStream[Payload, Payload]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Hedge_SoSWriteServer = grpc.BidiStreamingServer[Payload, Payload]

func _Hedge_SoSRead_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HedgeServer).SoSRead(&grpc.GenericServerStream[Payload, Payload]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Hedge_SoSReadServer = grpc.BidiStreamingServer[Payload, Payload]

func _Hedge_SoSClose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Payload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HedgeServer).SoSClose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hedge_SoSClose_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HedgeServer).SoSClose(ctx, req.(*Payload))
	}
	return interceptor(ctx, in, info, handler)
}

// Hedge_ServiceDesc is the grpc.ServiceDesc for Hedge service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hedge_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hedgeproto.Hedge",
	HandlerType: (*HedgeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SoSClose",
			Handler:    _Hedge_SoSClose_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Send",
			Handler:       _Hedge_Send_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Broadcast",
			Handler:       _Hedge_Broadcast_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SoSWrite",
			Handler:       _Hedge_SoSWrite_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SoSRead",
			Handler:       _Hedge_SoSRead_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "default.proto",
}
