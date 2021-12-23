// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package project

import (
	context "context"
	model "github.com/ez-deploy/protobuf/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OpsClient is the client API for Ops service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OpsClient interface {
	CreateProject(ctx context.Context, in *CreateProjectReq, opts ...grpc.CallOption) (*model.CommonResp, error)
	DeleteProject(ctx context.Context, in *DeleteProjectReq, opts ...grpc.CallOption) (*model.CommonResp, error)
	SetService(ctx context.Context, in *SetServiceReq, opts ...grpc.CallOption) (*model.CommonResp, error)
	GetService(ctx context.Context, in *GetServiceReq, opts ...grpc.CallOption) (*GetServiceResp, error)
	ListService(ctx context.Context, in *ListServiceReq, opts ...grpc.CallOption) (*ListServiceResp, error)
	DeleteService(ctx context.Context, in *DeleteServiceReq, opts ...grpc.CallOption) (*model.CommonResp, error)
	ListPods(ctx context.Context, in *ListPodsReq, opts ...grpc.CallOption) (*ListPodsResp, error)
}

type opsClient struct {
	cc grpc.ClientConnInterface
}

func NewOpsClient(cc grpc.ClientConnInterface) OpsClient {
	return &opsClient{cc}
}

func (c *opsClient) CreateProject(ctx context.Context, in *CreateProjectReq, opts ...grpc.CallOption) (*model.CommonResp, error) {
	out := new(model.CommonResp)
	err := c.cc.Invoke(ctx, "/Ops/CreateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsClient) DeleteProject(ctx context.Context, in *DeleteProjectReq, opts ...grpc.CallOption) (*model.CommonResp, error) {
	out := new(model.CommonResp)
	err := c.cc.Invoke(ctx, "/Ops/DeleteProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsClient) SetService(ctx context.Context, in *SetServiceReq, opts ...grpc.CallOption) (*model.CommonResp, error) {
	out := new(model.CommonResp)
	err := c.cc.Invoke(ctx, "/Ops/SetService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsClient) GetService(ctx context.Context, in *GetServiceReq, opts ...grpc.CallOption) (*GetServiceResp, error) {
	out := new(GetServiceResp)
	err := c.cc.Invoke(ctx, "/Ops/GetService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsClient) ListService(ctx context.Context, in *ListServiceReq, opts ...grpc.CallOption) (*ListServiceResp, error) {
	out := new(ListServiceResp)
	err := c.cc.Invoke(ctx, "/Ops/ListService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsClient) DeleteService(ctx context.Context, in *DeleteServiceReq, opts ...grpc.CallOption) (*model.CommonResp, error) {
	out := new(model.CommonResp)
	err := c.cc.Invoke(ctx, "/Ops/DeleteService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsClient) ListPods(ctx context.Context, in *ListPodsReq, opts ...grpc.CallOption) (*ListPodsResp, error) {
	out := new(ListPodsResp)
	err := c.cc.Invoke(ctx, "/Ops/ListPods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OpsServer is the server API for Ops service.
// All implementations must embed UnimplementedOpsServer
// for forward compatibility
type OpsServer interface {
	CreateProject(context.Context, *CreateProjectReq) (*model.CommonResp, error)
	DeleteProject(context.Context, *DeleteProjectReq) (*model.CommonResp, error)
	SetService(context.Context, *SetServiceReq) (*model.CommonResp, error)
	GetService(context.Context, *GetServiceReq) (*GetServiceResp, error)
	ListService(context.Context, *ListServiceReq) (*ListServiceResp, error)
	DeleteService(context.Context, *DeleteServiceReq) (*model.CommonResp, error)
	ListPods(context.Context, *ListPodsReq) (*ListPodsResp, error)
	mustEmbedUnimplementedOpsServer()
}

// UnimplementedOpsServer must be embedded to have forward compatible implementations.
type UnimplementedOpsServer struct {
}

func (UnimplementedOpsServer) CreateProject(context.Context, *CreateProjectReq) (*model.CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedOpsServer) DeleteProject(context.Context, *DeleteProjectReq) (*model.CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProject not implemented")
}
func (UnimplementedOpsServer) SetService(context.Context, *SetServiceReq) (*model.CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetService not implemented")
}
func (UnimplementedOpsServer) GetService(context.Context, *GetServiceReq) (*GetServiceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetService not implemented")
}
func (UnimplementedOpsServer) ListService(context.Context, *ListServiceReq) (*ListServiceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListService not implemented")
}
func (UnimplementedOpsServer) DeleteService(context.Context, *DeleteServiceReq) (*model.CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteService not implemented")
}
func (UnimplementedOpsServer) ListPods(context.Context, *ListPodsReq) (*ListPodsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPods not implemented")
}
func (UnimplementedOpsServer) mustEmbedUnimplementedOpsServer() {}

// UnsafeOpsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OpsServer will
// result in compilation errors.
type UnsafeOpsServer interface {
	mustEmbedUnimplementedOpsServer()
}

func RegisterOpsServer(s grpc.ServiceRegistrar, srv OpsServer) {
	s.RegisterService(&Ops_ServiceDesc, srv)
}

func _Ops_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ops/CreateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsServer).CreateProject(ctx, req.(*CreateProjectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ops_DeleteProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsServer).DeleteProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ops/DeleteProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsServer).DeleteProject(ctx, req.(*DeleteProjectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ops_SetService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetServiceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsServer).SetService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ops/SetService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsServer).SetService(ctx, req.(*SetServiceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ops_GetService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsServer).GetService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ops/GetService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsServer).GetService(ctx, req.(*GetServiceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ops_ListService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServiceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsServer).ListService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ops/ListService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsServer).ListService(ctx, req.(*ListServiceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ops_DeleteService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServiceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsServer).DeleteService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ops/DeleteService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsServer).DeleteService(ctx, req.(*DeleteServiceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ops_ListPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPodsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsServer).ListPods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ops/ListPods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsServer).ListPods(ctx, req.(*ListPodsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Ops_ServiceDesc is the grpc.ServiceDesc for Ops service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ops_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Ops",
	HandlerType: (*OpsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProject",
			Handler:    _Ops_CreateProject_Handler,
		},
		{
			MethodName: "DeleteProject",
			Handler:    _Ops_DeleteProject_Handler,
		},
		{
			MethodName: "SetService",
			Handler:    _Ops_SetService_Handler,
		},
		{
			MethodName: "GetService",
			Handler:    _Ops_GetService_Handler,
		},
		{
			MethodName: "ListService",
			Handler:    _Ops_ListService_Handler,
		},
		{
			MethodName: "DeleteService",
			Handler:    _Ops_DeleteService_Handler,
		},
		{
			MethodName: "ListPods",
			Handler:    _Ops_ListPods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "project/project.proto",
}
