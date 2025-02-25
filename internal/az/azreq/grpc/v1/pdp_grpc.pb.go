// Copyright 2025 Nitro Agility S.r.l.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: pdp.proto

package v1

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
	V1PDPService_AuthorizationCheck_FullMethodName = "/policydecisionpoint.V1PDPService/AuthorizationCheck"
)

// V1PDPServiceClient is the client API for V1PDPService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// V1PDPService	is the service for the Policy Decision Point.
type V1PDPServiceClient interface {
	AuthorizationCheck(ctx context.Context, in *AuthorizationCheckRequest, opts ...grpc.CallOption) (*AuthorizationCheckResponse, error)
}

type v1PDPServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewV1PDPServiceClient(cc grpc.ClientConnInterface) V1PDPServiceClient {
	return &v1PDPServiceClient{cc}
}

func (c *v1PDPServiceClient) AuthorizationCheck(ctx context.Context, in *AuthorizationCheckRequest, opts ...grpc.CallOption) (*AuthorizationCheckResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthorizationCheckResponse)
	err := c.cc.Invoke(ctx, V1PDPService_AuthorizationCheck_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// V1PDPServiceServer is the server API for V1PDPService service.
// All implementations should embed UnimplementedV1PDPServiceServer
// for forward compatibility.
//
// V1PDPService	is the service for the Policy Decision Point.
type V1PDPServiceServer interface {
	AuthorizationCheck(context.Context, *AuthorizationCheckRequest) (*AuthorizationCheckResponse, error)
}

// UnimplementedV1PDPServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedV1PDPServiceServer struct{}

func (UnimplementedV1PDPServiceServer) AuthorizationCheck(context.Context, *AuthorizationCheckRequest) (*AuthorizationCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizationCheck not implemented")
}
func (UnimplementedV1PDPServiceServer) testEmbeddedByValue() {}

// UnsafeV1PDPServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to V1PDPServiceServer will
// result in compilation errors.
type UnsafeV1PDPServiceServer interface {
	mustEmbedUnimplementedV1PDPServiceServer()
}

func RegisterV1PDPServiceServer(s grpc.ServiceRegistrar, srv V1PDPServiceServer) {
	// If the following call pancis, it indicates UnimplementedV1PDPServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&V1PDPService_ServiceDesc, srv)
}

func _V1PDPService_AuthorizationCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizationCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(V1PDPServiceServer).AuthorizationCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: V1PDPService_AuthorizationCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(V1PDPServiceServer).AuthorizationCheck(ctx, req.(*AuthorizationCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// V1PDPService_ServiceDesc is the grpc.ServiceDesc for V1PDPService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var V1PDPService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "policydecisionpoint.V1PDPService",
	HandlerType: (*V1PDPServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthorizationCheck",
			Handler:    _V1PDPService_AuthorizationCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pdp.proto",
}
