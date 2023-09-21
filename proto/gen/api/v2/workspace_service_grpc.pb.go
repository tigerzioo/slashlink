// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/v2/workspace_service.proto

package apiv2

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
	WorkspaceService_GetWorkspaceSetting_FullMethodName    = "/slash.api.v2.WorkspaceService/GetWorkspaceSetting"
	WorkspaceService_UpdateWorkspaceSetting_FullMethodName = "/slash.api.v2.WorkspaceService/UpdateWorkspaceSetting"
)

// WorkspaceServiceClient is the client API for WorkspaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkspaceServiceClient interface {
	GetWorkspaceSetting(ctx context.Context, in *GetWorkspaceSettingRequest, opts ...grpc.CallOption) (*GetWorkspaceSettingResponse, error)
	UpdateWorkspaceSetting(ctx context.Context, in *UpdateWorkspaceSettingRequest, opts ...grpc.CallOption) (*UpdateWorkspaceSettingResponse, error)
}

type workspaceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkspaceServiceClient(cc grpc.ClientConnInterface) WorkspaceServiceClient {
	return &workspaceServiceClient{cc}
}

func (c *workspaceServiceClient) GetWorkspaceSetting(ctx context.Context, in *GetWorkspaceSettingRequest, opts ...grpc.CallOption) (*GetWorkspaceSettingResponse, error) {
	out := new(GetWorkspaceSettingResponse)
	err := c.cc.Invoke(ctx, WorkspaceService_GetWorkspaceSetting_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) UpdateWorkspaceSetting(ctx context.Context, in *UpdateWorkspaceSettingRequest, opts ...grpc.CallOption) (*UpdateWorkspaceSettingResponse, error) {
	out := new(UpdateWorkspaceSettingResponse)
	err := c.cc.Invoke(ctx, WorkspaceService_UpdateWorkspaceSetting_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkspaceServiceServer is the server API for WorkspaceService service.
// All implementations must embed UnimplementedWorkspaceServiceServer
// for forward compatibility
type WorkspaceServiceServer interface {
	GetWorkspaceSetting(context.Context, *GetWorkspaceSettingRequest) (*GetWorkspaceSettingResponse, error)
	UpdateWorkspaceSetting(context.Context, *UpdateWorkspaceSettingRequest) (*UpdateWorkspaceSettingResponse, error)
	mustEmbedUnimplementedWorkspaceServiceServer()
}

// UnimplementedWorkspaceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkspaceServiceServer struct {
}

func (UnimplementedWorkspaceServiceServer) GetWorkspaceSetting(context.Context, *GetWorkspaceSettingRequest) (*GetWorkspaceSettingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkspaceSetting not implemented")
}
func (UnimplementedWorkspaceServiceServer) UpdateWorkspaceSetting(context.Context, *UpdateWorkspaceSettingRequest) (*UpdateWorkspaceSettingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorkspaceSetting not implemented")
}
func (UnimplementedWorkspaceServiceServer) mustEmbedUnimplementedWorkspaceServiceServer() {}

// UnsafeWorkspaceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkspaceServiceServer will
// result in compilation errors.
type UnsafeWorkspaceServiceServer interface {
	mustEmbedUnimplementedWorkspaceServiceServer()
}

func RegisterWorkspaceServiceServer(s grpc.ServiceRegistrar, srv WorkspaceServiceServer) {
	s.RegisterService(&WorkspaceService_ServiceDesc, srv)
}

func _WorkspaceService_GetWorkspaceSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkspaceSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).GetWorkspaceSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceService_GetWorkspaceSetting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).GetWorkspaceSetting(ctx, req.(*GetWorkspaceSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_UpdateWorkspaceSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkspaceSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).UpdateWorkspaceSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceService_UpdateWorkspaceSetting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).UpdateWorkspaceSetting(ctx, req.(*UpdateWorkspaceSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkspaceService_ServiceDesc is the grpc.ServiceDesc for WorkspaceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkspaceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "slash.api.v2.WorkspaceService",
	HandlerType: (*WorkspaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWorkspaceSetting",
			Handler:    _WorkspaceService_GetWorkspaceSetting_Handler,
		},
		{
			MethodName: "UpdateWorkspaceSetting",
			Handler:    _WorkspaceService_UpdateWorkspaceSetting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v2/workspace_service.proto",
}
