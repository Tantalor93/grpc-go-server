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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DNSForwardingDecisionServiceClient is the client API for DNSForwardingDecisionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DNSForwardingDecisionServiceClient interface {
	// EvaluateCustomDNSResolution evaluates custom DNS resolution
	EvaluateCustomDNSResolution(ctx context.Context, in *EvaluationRequest, opts ...grpc.CallOption) (*EvaluationResponse, error)
}

type dNSForwardingDecisionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDNSForwardingDecisionServiceClient(cc grpc.ClientConnInterface) DNSForwardingDecisionServiceClient {
	return &dNSForwardingDecisionServiceClient{cc}
}

func (c *dNSForwardingDecisionServiceClient) EvaluateCustomDNSResolution(ctx context.Context, in *EvaluationRequest, opts ...grpc.CallOption) (*EvaluationResponse, error) {
	out := new(EvaluationResponse)
	err := c.cc.Invoke(ctx, "/dnsforwardingdecision.v1.DNSForwardingDecisionService/EvaluateCustomDNSResolution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DNSForwardingDecisionServiceServer is the server API for DNSForwardingDecisionService service.
// All implementations must embed UnimplementedDNSForwardingDecisionServiceServer
// for forward compatibility
type DNSForwardingDecisionServiceServer interface {
	// EvaluateCustomDNSResolution evaluates custom DNS resolution
	EvaluateCustomDNSResolution(context.Context, *EvaluationRequest) (*EvaluationResponse, error)
	mustEmbedUnimplementedDNSForwardingDecisionServiceServer()
}

// UnimplementedDNSForwardingDecisionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDNSForwardingDecisionServiceServer struct {
}

func (UnimplementedDNSForwardingDecisionServiceServer) EvaluateCustomDNSResolution(context.Context, *EvaluationRequest) (*EvaluationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EvaluateCustomDNSResolution not implemented")
}
func (UnimplementedDNSForwardingDecisionServiceServer) mustEmbedUnimplementedDNSForwardingDecisionServiceServer() {
}

// UnsafeDNSForwardingDecisionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DNSForwardingDecisionServiceServer will
// result in compilation errors.
type UnsafeDNSForwardingDecisionServiceServer interface {
	mustEmbedUnimplementedDNSForwardingDecisionServiceServer()
}

func RegisterDNSForwardingDecisionServiceServer(s grpc.ServiceRegistrar, srv DNSForwardingDecisionServiceServer) {
	s.RegisterService(&DNSForwardingDecisionService_ServiceDesc, srv)
}

func _DNSForwardingDecisionService_EvaluateCustomDNSResolution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvaluationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSForwardingDecisionServiceServer).EvaluateCustomDNSResolution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dnsforwardingdecision.v1.DNSForwardingDecisionService/EvaluateCustomDNSResolution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSForwardingDecisionServiceServer).EvaluateCustomDNSResolution(ctx, req.(*EvaluationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DNSForwardingDecisionService_ServiceDesc is the grpc.ServiceDesc for DNSForwardingDecisionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DNSForwardingDecisionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dnsforwardingdecision.v1.DNSForwardingDecisionService",
	HandlerType: (*DNSForwardingDecisionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EvaluateCustomDNSResolution",
			Handler:    _DNSForwardingDecisionService_EvaluateCustomDNSResolution_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/dfds.proto",
}
