// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.1
// source: lead.proto

package leadpb

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
	LeadService_ReceiveLead_FullMethodName = "/leadpb.LeadService/ReceiveLead"
)

// LeadServiceClient is the client API for LeadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LeadServiceClient interface {
	ReceiveLead(ctx context.Context, in *LeadRequest, opts ...grpc.CallOption) (*LeadResponse, error)
}

type leadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLeadServiceClient(cc grpc.ClientConnInterface) LeadServiceClient {
	return &leadServiceClient{cc}
}

func (c *leadServiceClient) ReceiveLead(ctx context.Context, in *LeadRequest, opts ...grpc.CallOption) (*LeadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LeadResponse)
	err := c.cc.Invoke(ctx, LeadService_ReceiveLead_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LeadServiceServer is the server API for LeadService service.
// All implementations must embed UnimplementedLeadServiceServer
// for forward compatibility.
type LeadServiceServer interface {
	ReceiveLead(context.Context, *LeadRequest) (*LeadResponse, error)
	mustEmbedUnimplementedLeadServiceServer()
}

// UnimplementedLeadServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLeadServiceServer struct{}

func (UnimplementedLeadServiceServer) ReceiveLead(context.Context, *LeadRequest) (*LeadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveLead not implemented")
}
func (UnimplementedLeadServiceServer) mustEmbedUnimplementedLeadServiceServer() {}
func (UnimplementedLeadServiceServer) testEmbeddedByValue()                     {}

// UnsafeLeadServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LeadServiceServer will
// result in compilation errors.
type UnsafeLeadServiceServer interface {
	mustEmbedUnimplementedLeadServiceServer()
}

func RegisterLeadServiceServer(s grpc.ServiceRegistrar, srv LeadServiceServer) {
	// If the following call pancis, it indicates UnimplementedLeadServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LeadService_ServiceDesc, srv)
}

func _LeadService_ReceiveLead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeadServiceServer).ReceiveLead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LeadService_ReceiveLead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeadServiceServer).ReceiveLead(ctx, req.(*LeadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LeadService_ServiceDesc is the grpc.ServiceDesc for LeadService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LeadService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "leadpb.LeadService",
	HandlerType: (*LeadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReceiveLead",
			Handler:    _LeadService_ReceiveLead_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lead.proto",
}
