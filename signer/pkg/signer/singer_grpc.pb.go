// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: singer.proto

package signer

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
	Signer_GetPrivateKey_FullMethodName    = "/signer.Signer/getPrivateKey"
	Signer_SendOwnSignature_FullMethodName = "/signer.Signer/sendOwnSignature"
)

// SignerClient is the client API for Signer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SignerClient interface {
	GetPrivateKey(ctx context.Context, in *GetIBEPrivatekeyRequest, opts ...grpc.CallOption) (*GetIBEPrivatekeyResponse, error)
	SendOwnSignature(ctx context.Context, in *SendSignature, opts ...grpc.CallOption) (*SendSignatureResponse, error)
}

type signerClient struct {
	cc grpc.ClientConnInterface
}

func NewSignerClient(cc grpc.ClientConnInterface) SignerClient {
	return &signerClient{cc}
}

func (c *signerClient) GetPrivateKey(ctx context.Context, in *GetIBEPrivatekeyRequest, opts ...grpc.CallOption) (*GetIBEPrivatekeyResponse, error) {
	out := new(GetIBEPrivatekeyResponse)
	err := c.cc.Invoke(ctx, Signer_GetPrivateKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signerClient) SendOwnSignature(ctx context.Context, in *SendSignature, opts ...grpc.CallOption) (*SendSignatureResponse, error) {
	out := new(SendSignatureResponse)
	err := c.cc.Invoke(ctx, Signer_SendOwnSignature_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignerServer is the server API for Signer service.
// All implementations must embed UnimplementedSignerServer
// for forward compatibility
type SignerServer interface {
	GetPrivateKey(context.Context, *GetIBEPrivatekeyRequest) (*GetIBEPrivatekeyResponse, error)
	SendOwnSignature(context.Context, *SendSignature) (*SendSignatureResponse, error)
	mustEmbedUnimplementedSignerServer()
}

// UnimplementedSignerServer must be embedded to have forward compatible implementations.
type UnimplementedSignerServer struct {
}

func (UnimplementedSignerServer) GetPrivateKey(context.Context, *GetIBEPrivatekeyRequest) (*GetIBEPrivatekeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrivateKey not implemented")
}
func (UnimplementedSignerServer) SendOwnSignature(context.Context, *SendSignature) (*SendSignatureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOwnSignature not implemented")
}
func (UnimplementedSignerServer) mustEmbedUnimplementedSignerServer() {}

// UnsafeSignerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SignerServer will
// result in compilation errors.
type UnsafeSignerServer interface {
	mustEmbedUnimplementedSignerServer()
}

func RegisterSignerServer(s grpc.ServiceRegistrar, srv SignerServer) {
	s.RegisterService(&Signer_ServiceDesc, srv)
}

func _Signer_GetPrivateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIBEPrivatekeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignerServer).GetPrivateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Signer_GetPrivateKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignerServer).GetPrivateKey(ctx, req.(*GetIBEPrivatekeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signer_SendOwnSignature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSignature)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignerServer).SendOwnSignature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Signer_SendOwnSignature_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignerServer).SendOwnSignature(ctx, req.(*SendSignature))
	}
	return interceptor(ctx, in, info, handler)
}

// Signer_ServiceDesc is the grpc.ServiceDesc for Signer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Signer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "signer.Signer",
	HandlerType: (*SignerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getPrivateKey",
			Handler:    _Signer_GetPrivateKey_Handler,
		},
		{
			MethodName: "sendOwnSignature",
			Handler:    _Signer_SendOwnSignature_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "singer.proto",
}
