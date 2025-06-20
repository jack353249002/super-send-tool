// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.9.0
// source: conf.proto

package proto

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ConfService_SetConf_FullMethodName     = "/super_send.ConfService/SetConf"
	ConfService_DelConf_FullMethodName     = "/super_send.ConfService/DelConf"
	ConfService_GetConfList_FullMethodName = "/super_send.ConfService/GetConfList"
	ConfService_Reload_FullMethodName      = "/super_send.ConfService/Reload"
)

// ConfServiceClient is the client API for ConfService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfServiceClient interface {
	// 设置配置信息
	SetConf(ctx context.Context, in *SetConfRequest, opts ...grpc.CallOption) (*ConfResponse, error)
	// 删除配置信息
	DelConf(ctx context.Context, in *DelConfRequest, opts ...grpc.CallOption) (*ConfResponse, error)
	// 获取配置列表
	GetConfList(ctx context.Context, in *GetConfListRequest, opts ...grpc.CallOption) (*ConfListResponse, error)
	// 重新加载配置
	Reload(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ConfResponse, error)
}

type confServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConfServiceClient(cc grpc.ClientConnInterface) ConfServiceClient {
	return &confServiceClient{cc}
}

func (c *confServiceClient) SetConf(ctx context.Context, in *SetConfRequest, opts ...grpc.CallOption) (*ConfResponse, error) {
	out := new(ConfResponse)
	err := c.cc.Invoke(ctx, ConfService_SetConf_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *confServiceClient) DelConf(ctx context.Context, in *DelConfRequest, opts ...grpc.CallOption) (*ConfResponse, error) {
	out := new(ConfResponse)
	err := c.cc.Invoke(ctx, ConfService_DelConf_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *confServiceClient) GetConfList(ctx context.Context, in *GetConfListRequest, opts ...grpc.CallOption) (*ConfListResponse, error) {
	out := new(ConfListResponse)
	err := c.cc.Invoke(ctx, ConfService_GetConfList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *confServiceClient) Reload(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ConfResponse, error) {
	out := new(ConfResponse)
	err := c.cc.Invoke(ctx, ConfService_Reload_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfServiceServer is the server API for ConfService service.
// All implementations must embed UnimplementedConfServiceServer
// for forward compatibility
type ConfServiceServer interface {
	// 设置配置信息
	SetConf(context.Context, *SetConfRequest) (*ConfResponse, error)
	// 删除配置信息
	DelConf(context.Context, *DelConfRequest) (*ConfResponse, error)
	// 获取配置列表
	GetConfList(context.Context, *GetConfListRequest) (*ConfListResponse, error)
	// 重新加载配置
	Reload(context.Context, *empty.Empty) (*ConfResponse, error)
	mustEmbedUnimplementedConfServiceServer()
}

// UnimplementedConfServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConfServiceServer struct {
}

func (UnimplementedConfServiceServer) SetConf(context.Context, *SetConfRequest) (*ConfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConf not implemented")
}
func (UnimplementedConfServiceServer) DelConf(context.Context, *DelConfRequest) (*ConfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelConf not implemented")
}
func (UnimplementedConfServiceServer) GetConfList(context.Context, *GetConfListRequest) (*ConfListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfList not implemented")
}
func (UnimplementedConfServiceServer) Reload(context.Context, *empty.Empty) (*ConfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reload not implemented")
}
func (UnimplementedConfServiceServer) mustEmbedUnimplementedConfServiceServer() {}

// UnsafeConfServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfServiceServer will
// result in compilation errors.
type UnsafeConfServiceServer interface {
	mustEmbedUnimplementedConfServiceServer()
}

func RegisterConfServiceServer(s grpc.ServiceRegistrar, srv ConfServiceServer) {
	s.RegisterService(&ConfService_ServiceDesc, srv)
}

func _ConfService_SetConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetConfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfServiceServer).SetConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfService_SetConf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfServiceServer).SetConf(ctx, req.(*SetConfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfService_DelConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelConfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfServiceServer).DelConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfService_DelConf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfServiceServer).DelConf(ctx, req.(*DelConfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfService_GetConfList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfServiceServer).GetConfList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfService_GetConfList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfServiceServer).GetConfList(ctx, req.(*GetConfListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfService_Reload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfServiceServer).Reload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfService_Reload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfServiceServer).Reload(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ConfService_ServiceDesc is the grpc.ServiceDesc for ConfService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConfService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "super_send.ConfService",
	HandlerType: (*ConfServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetConf",
			Handler:    _ConfService_SetConf_Handler,
		},
		{
			MethodName: "DelConf",
			Handler:    _ConfService_DelConf_Handler,
		},
		{
			MethodName: "GetConfList",
			Handler:    _ConfService_GetConfList_Handler,
		},
		{
			MethodName: "Reload",
			Handler:    _ConfService_Reload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "conf.proto",
}
