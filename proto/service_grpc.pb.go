// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// GreetingClient is the client API for Greeting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetingClient interface {
	SayHello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type greetingClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetingClient(cc grpc.ClientConnInterface) GreetingClient {
	return &greetingClient{cc}
}

func (c *greetingClient) SayHello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Greeting/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetingServer is the server API for Greeting service.
// All implementations must embed UnimplementedGreetingServer
// for forward compatibility
type GreetingServer interface {
	SayHello(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedGreetingServer()
}

// UnimplementedGreetingServer must be embedded to have forward compatible implementations.
type UnimplementedGreetingServer struct {
}

func (UnimplementedGreetingServer) SayHello(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreetingServer) mustEmbedUnimplementedGreetingServer() {}

// UnsafeGreetingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetingServer will
// result in compilation errors.
type UnsafeGreetingServer interface {
	mustEmbedUnimplementedGreetingServer()
}

func RegisterGreetingServer(s *grpc.Server, srv GreetingServer) {
	s.RegisterService(&_Greeting_serviceDesc, srv)
}

func _Greeting_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetingServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Greeting/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetingServer).SayHello(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeting_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Greeting",
	HandlerType: (*GreetingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeting_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service.proto",
}