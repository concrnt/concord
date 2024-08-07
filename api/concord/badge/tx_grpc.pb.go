// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: concord/badge/tx.proto

package badge

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
	Msg_UpdateParams_FullMethodName = "/concord.badge.Msg/UpdateParams"
	Msg_CreateSeries_FullMethodName = "/concord.badge.Msg/CreateSeries"
	Msg_MintBadge_FullMethodName    = "/concord.badge.Msg/MintBadge"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	CreateSeries(ctx context.Context, in *MsgCreateSeries, opts ...grpc.CallOption) (*MsgCreateSeriesResponse, error)
	MintBadge(ctx context.Context, in *MsgMintBadge, opts ...grpc.CallOption) (*MsgMintBadgeResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateSeries(ctx context.Context, in *MsgCreateSeries, opts ...grpc.CallOption) (*MsgCreateSeriesResponse, error) {
	out := new(MsgCreateSeriesResponse)
	err := c.cc.Invoke(ctx, Msg_CreateSeries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MintBadge(ctx context.Context, in *MsgMintBadge, opts ...grpc.CallOption) (*MsgMintBadgeResponse, error) {
	out := new(MsgMintBadgeResponse)
	err := c.cc.Invoke(ctx, Msg_MintBadge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	CreateSeries(context.Context, *MsgCreateSeries) (*MsgCreateSeriesResponse, error)
	MintBadge(context.Context, *MsgMintBadge) (*MsgMintBadgeResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) CreateSeries(context.Context, *MsgCreateSeries) (*MsgCreateSeriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSeries not implemented")
}
func (UnimplementedMsgServer) MintBadge(context.Context, *MsgMintBadge) (*MsgMintBadgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintBadge not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateSeries)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateSeries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateSeries(ctx, req.(*MsgCreateSeries))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MintBadge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMintBadge)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MintBadge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_MintBadge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MintBadge(ctx, req.(*MsgMintBadge))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "concord.badge.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "CreateSeries",
			Handler:    _Msg_CreateSeries_Handler,
		},
		{
			MethodName: "MintBadge",
			Handler:    _Msg_MintBadge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "concord/badge/tx.proto",
}
