// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: api/v1/debug/debug.proto

package debug

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
	DebugService_UpsertDefaultSettings_FullMethodName      = "/openpgl.settings.v1.debug.DebugService/UpsertDefaultSettings"
	DebugService_ProduceUpsertSettingsEvent_FullMethodName = "/openpgl.settings.v1.debug.DebugService/ProduceUpsertSettingsEvent"
)

// DebugServiceClient is the client API for DebugService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DebugServiceClient interface {
	UpsertDefaultSettings(ctx context.Context, in *UpsertDefaultSettingsRequest, opts ...grpc.CallOption) (*UpsertDefaultSettingsResponse, error)
	ProduceUpsertSettingsEvent(ctx context.Context, in *ProduceUpsertSettingsEventRequest, opts ...grpc.CallOption) (*ProduceUpsertSettingsEventResponse, error)
}

type debugServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDebugServiceClient(cc grpc.ClientConnInterface) DebugServiceClient {
	return &debugServiceClient{cc}
}

func (c *debugServiceClient) UpsertDefaultSettings(ctx context.Context, in *UpsertDefaultSettingsRequest, opts ...grpc.CallOption) (*UpsertDefaultSettingsResponse, error) {
	out := new(UpsertDefaultSettingsResponse)
	err := c.cc.Invoke(ctx, DebugService_UpsertDefaultSettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugServiceClient) ProduceUpsertSettingsEvent(ctx context.Context, in *ProduceUpsertSettingsEventRequest, opts ...grpc.CallOption) (*ProduceUpsertSettingsEventResponse, error) {
	out := new(ProduceUpsertSettingsEventResponse)
	err := c.cc.Invoke(ctx, DebugService_ProduceUpsertSettingsEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DebugServiceServer is the server API for DebugService service.
// All implementations must embed UnimplementedDebugServiceServer
// for forward compatibility
type DebugServiceServer interface {
	UpsertDefaultSettings(context.Context, *UpsertDefaultSettingsRequest) (*UpsertDefaultSettingsResponse, error)
	ProduceUpsertSettingsEvent(context.Context, *ProduceUpsertSettingsEventRequest) (*ProduceUpsertSettingsEventResponse, error)
	mustEmbedUnimplementedDebugServiceServer()
}

// UnimplementedDebugServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDebugServiceServer struct {
}

func (UnimplementedDebugServiceServer) UpsertDefaultSettings(context.Context, *UpsertDefaultSettingsRequest) (*UpsertDefaultSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertDefaultSettings not implemented")
}
func (UnimplementedDebugServiceServer) ProduceUpsertSettingsEvent(context.Context, *ProduceUpsertSettingsEventRequest) (*ProduceUpsertSettingsEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProduceUpsertSettingsEvent not implemented")
}
func (UnimplementedDebugServiceServer) mustEmbedUnimplementedDebugServiceServer() {}

// UnsafeDebugServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DebugServiceServer will
// result in compilation errors.
type UnsafeDebugServiceServer interface {
	mustEmbedUnimplementedDebugServiceServer()
}

func RegisterDebugServiceServer(s grpc.ServiceRegistrar, srv DebugServiceServer) {
	s.RegisterService(&DebugService_ServiceDesc, srv)
}

func _DebugService_UpsertDefaultSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertDefaultSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugServiceServer).UpsertDefaultSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DebugService_UpsertDefaultSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugServiceServer).UpsertDefaultSettings(ctx, req.(*UpsertDefaultSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DebugService_ProduceUpsertSettingsEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProduceUpsertSettingsEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugServiceServer).ProduceUpsertSettingsEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DebugService_ProduceUpsertSettingsEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugServiceServer).ProduceUpsertSettingsEvent(ctx, req.(*ProduceUpsertSettingsEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DebugService_ServiceDesc is the grpc.ServiceDesc for DebugService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DebugService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openpgl.settings.v1.debug.DebugService",
	HandlerType: (*DebugServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertDefaultSettings",
			Handler:    _DebugService_UpsertDefaultSettings_Handler,
		},
		{
			MethodName: "ProduceUpsertSettingsEvent",
			Handler:    _DebugService_ProduceUpsertSettingsEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/debug/debug.proto",
}