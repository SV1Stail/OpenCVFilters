// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: service.proto

package gen

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
	Service_AddFiltersAndChannels_FullMethodName = "/opencvfilters.Service/AddFiltersAndChannels"
	Service_FindContours_FullMethodName          = "/opencvfilters.Service/FindContours"
	Service_FindP_FullMethodName                 = "/opencvfilters.Service/FindP"
	Service_FindS_FullMethodName                 = "/opencvfilters.Service/FindS"
	Service_FindAll_FullMethodName               = "/opencvfilters.Service/FindAll"
)

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	AddFiltersAndChannels(ctx context.Context, in *ImageReq, opts ...grpc.CallOption) (*FiltersAndChannelsResp, error)
	FindContours(ctx context.Context, in *ImageReq, opts ...grpc.CallOption) (*FindContoursResp, error)
	FindP(ctx context.Context, in *ImageReq, opts ...grpc.CallOption) (*NumericalResp, error)
	FindS(ctx context.Context, in *ImageReq, opts ...grpc.CallOption) (*NumericalResp, error)
	FindAll(ctx context.Context, in *ImageReq, opts ...grpc.CallOption) (*AllResp, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) AddFiltersAndChannels(ctx context.Context, in *ImageReq, opts ...grpc.CallOption) (*FiltersAndChannelsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FiltersAndChannelsResp)
	err := c.cc.Invoke(ctx, Service_AddFiltersAndChannels_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) FindContours(ctx context.Context, in *ImageReq, opts ...grpc.CallOption) (*FindContoursResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindContoursResp)
	err := c.cc.Invoke(ctx, Service_FindContours_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) FindP(ctx context.Context, in *ImageReq, opts ...grpc.CallOption) (*NumericalResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NumericalResp)
	err := c.cc.Invoke(ctx, Service_FindP_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) FindS(ctx context.Context, in *ImageReq, opts ...grpc.CallOption) (*NumericalResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NumericalResp)
	err := c.cc.Invoke(ctx, Service_FindS_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) FindAll(ctx context.Context, in *ImageReq, opts ...grpc.CallOption) (*AllResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AllResp)
	err := c.cc.Invoke(ctx, Service_FindAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility.
type ServiceServer interface {
	AddFiltersAndChannels(context.Context, *ImageReq) (*FiltersAndChannelsResp, error)
	FindContours(context.Context, *ImageReq) (*FindContoursResp, error)
	FindP(context.Context, *ImageReq) (*NumericalResp, error)
	FindS(context.Context, *ImageReq) (*NumericalResp, error)
	FindAll(context.Context, *ImageReq) (*AllResp, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedServiceServer struct{}

func (UnimplementedServiceServer) AddFiltersAndChannels(context.Context, *ImageReq) (*FiltersAndChannelsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFiltersAndChannels not implemented")
}
func (UnimplementedServiceServer) FindContours(context.Context, *ImageReq) (*FindContoursResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindContours not implemented")
}
func (UnimplementedServiceServer) FindP(context.Context, *ImageReq) (*NumericalResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindP not implemented")
}
func (UnimplementedServiceServer) FindS(context.Context, *ImageReq) (*NumericalResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindS not implemented")
}
func (UnimplementedServiceServer) FindAll(context.Context, *ImageReq) (*AllResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}
func (UnimplementedServiceServer) testEmbeddedByValue()                 {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	// If the following call pancis, it indicates UnimplementedServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_AddFiltersAndChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).AddFiltersAndChannels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_AddFiltersAndChannels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).AddFiltersAndChannels(ctx, req.(*ImageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_FindContours_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).FindContours(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_FindContours_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).FindContours(ctx, req.(*ImageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_FindP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).FindP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_FindP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).FindP(ctx, req.(*ImageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_FindS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).FindS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_FindS_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).FindS(ctx, req.(*ImageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_FindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).FindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_FindAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).FindAll(ctx, req.(*ImageReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "opencvfilters.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFiltersAndChannels",
			Handler:    _Service_AddFiltersAndChannels_Handler,
		},
		{
			MethodName: "FindContours",
			Handler:    _Service_FindContours_Handler,
		},
		{
			MethodName: "FindP",
			Handler:    _Service_FindP_Handler,
		},
		{
			MethodName: "FindS",
			Handler:    _Service_FindS_Handler,
		},
		{
			MethodName: "FindAll",
			Handler:    _Service_FindAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
