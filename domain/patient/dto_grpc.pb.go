// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: domain/patient/proto/dto.proto

package patient

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PatientProtoServiceClient is the client API for PatientProtoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PatientProtoServiceClient interface {
	Create(ctx context.Context, in *PatientProto, opts ...grpc.CallOption) (*emptypb.Empty, error)
	List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PatientListProto, error)
	Find(ctx context.Context, in *PatientFilterProto, opts ...grpc.CallOption) (*PatientListProto, error)
}

type patientProtoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPatientProtoServiceClient(cc grpc.ClientConnInterface) PatientProtoServiceClient {
	return &patientProtoServiceClient{cc}
}

func (c *patientProtoServiceClient) Create(ctx context.Context, in *PatientProto, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/patient.PatientProtoService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientProtoServiceClient) List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PatientListProto, error) {
	out := new(PatientListProto)
	err := c.cc.Invoke(ctx, "/patient.PatientProtoService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientProtoServiceClient) Find(ctx context.Context, in *PatientFilterProto, opts ...grpc.CallOption) (*PatientListProto, error) {
	out := new(PatientListProto)
	err := c.cc.Invoke(ctx, "/patient.PatientProtoService/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PatientProtoServiceServer is the server API for PatientProtoService service.
// All implementations must embed UnimplementedPatientProtoServiceServer
// for forward compatibility
type PatientProtoServiceServer interface {
	Create(context.Context, *PatientProto) (*emptypb.Empty, error)
	List(context.Context, *emptypb.Empty) (*PatientListProto, error)
	Find(context.Context, *PatientFilterProto) (*PatientListProto, error)
	mustEmbedUnimplementedPatientProtoServiceServer()
}

// UnimplementedPatientProtoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPatientProtoServiceServer struct {
}

func (UnimplementedPatientProtoServiceServer) Create(context.Context, *PatientProto) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPatientProtoServiceServer) List(context.Context, *emptypb.Empty) (*PatientListProto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPatientProtoServiceServer) Find(context.Context, *PatientFilterProto) (*PatientListProto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedPatientProtoServiceServer) mustEmbedUnimplementedPatientProtoServiceServer() {}

// UnsafePatientProtoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PatientProtoServiceServer will
// result in compilation errors.
type UnsafePatientProtoServiceServer interface {
	mustEmbedUnimplementedPatientProtoServiceServer()
}

func RegisterPatientProtoServiceServer(s grpc.ServiceRegistrar, srv PatientProtoServiceServer) {
	s.RegisterService(&PatientProtoService_ServiceDesc, srv)
}

func _PatientProtoService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatientProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientProtoServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/patient.PatientProtoService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientProtoServiceServer).Create(ctx, req.(*PatientProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientProtoService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientProtoServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/patient.PatientProtoService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientProtoServiceServer).List(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientProtoService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatientFilterProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientProtoServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/patient.PatientProtoService/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientProtoServiceServer).Find(ctx, req.(*PatientFilterProto))
	}
	return interceptor(ctx, in, info, handler)
}

// PatientProtoService_ServiceDesc is the grpc.ServiceDesc for PatientProtoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PatientProtoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "patient.PatientProtoService",
	HandlerType: (*PatientProtoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PatientProtoService_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _PatientProtoService_List_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _PatientProtoService_Find_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "domain/patient/proto/dto.proto",
}
