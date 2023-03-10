// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: src/proto/category.proto

package grpc

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

// CategoryServiceClient is the client API for CategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoryServiceClient interface {
	GetAllCategory(ctx context.Context, in *Blank, opts ...grpc.CallOption) (*CategoryResponse, error)
	CreateProtoCategory(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*Category, error)
	DeleteProtoCategory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*Category, error)
	UpdateProtoCategory(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*Category, error)
	// client streaming
	CreateProtoCategoryStream(ctx context.Context, opts ...grpc.CallOption) (CategoryService_CreateProtoCategoryStreamClient, error)
	DeleteProtoCategoryStream(ctx context.Context, opts ...grpc.CallOption) (CategoryService_DeleteProtoCategoryStreamClient, error)
	// server streaming
	CreateProtoCategoryServerStream(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (CategoryService_CreateProtoCategoryServerStreamClient, error)
	// bidirectional streaming
	CreateProtoCategoryBidirectionalStream(ctx context.Context, opts ...grpc.CallOption) (CategoryService_CreateProtoCategoryBidirectionalStreamClient, error)
}

type categoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryServiceClient(cc grpc.ClientConnInterface) CategoryServiceClient {
	return &categoryServiceClient{cc}
}

func (c *categoryServiceClient) GetAllCategory(ctx context.Context, in *Blank, opts ...grpc.CallOption) (*CategoryResponse, error) {
	out := new(CategoryResponse)
	err := c.cc.Invoke(ctx, "/pb.CategoryService/GetAllCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) CreateProtoCategory(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/pb.CategoryService/CreateProtoCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) DeleteProtoCategory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/pb.CategoryService/DeleteProtoCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) UpdateProtoCategory(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/pb.CategoryService/UpdateProtoCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) CreateProtoCategoryStream(ctx context.Context, opts ...grpc.CallOption) (CategoryService_CreateProtoCategoryStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &CategoryService_ServiceDesc.Streams[0], "/pb.CategoryService/CreateProtoCategoryStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &categoryServiceCreateProtoCategoryStreamClient{stream}
	return x, nil
}

type CategoryService_CreateProtoCategoryStreamClient interface {
	Send(*CategoryRequest) error
	CloseAndRecv() (*CategoryResponse, error)
	grpc.ClientStream
}

type categoryServiceCreateProtoCategoryStreamClient struct {
	grpc.ClientStream
}

func (x *categoryServiceCreateProtoCategoryStreamClient) Send(m *CategoryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *categoryServiceCreateProtoCategoryStreamClient) CloseAndRecv() (*CategoryResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CategoryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *categoryServiceClient) DeleteProtoCategoryStream(ctx context.Context, opts ...grpc.CallOption) (CategoryService_DeleteProtoCategoryStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &CategoryService_ServiceDesc.Streams[1], "/pb.CategoryService/DeleteProtoCategoryStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &categoryServiceDeleteProtoCategoryStreamClient{stream}
	return x, nil
}

type CategoryService_DeleteProtoCategoryStreamClient interface {
	Send(*CategoryDeleteRequest) error
	CloseAndRecv() (*CategoryResponse, error)
	grpc.ClientStream
}

type categoryServiceDeleteProtoCategoryStreamClient struct {
	grpc.ClientStream
}

func (x *categoryServiceDeleteProtoCategoryStreamClient) Send(m *CategoryDeleteRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *categoryServiceDeleteProtoCategoryStreamClient) CloseAndRecv() (*CategoryResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CategoryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *categoryServiceClient) CreateProtoCategoryServerStream(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (CategoryService_CreateProtoCategoryServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &CategoryService_ServiceDesc.Streams[2], "/pb.CategoryService/CreateProtoCategoryServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &categoryServiceCreateProtoCategoryServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CategoryService_CreateProtoCategoryServerStreamClient interface {
	Recv() (*CategoryResponse, error)
	grpc.ClientStream
}

type categoryServiceCreateProtoCategoryServerStreamClient struct {
	grpc.ClientStream
}

func (x *categoryServiceCreateProtoCategoryServerStreamClient) Recv() (*CategoryResponse, error) {
	m := new(CategoryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *categoryServiceClient) CreateProtoCategoryBidirectionalStream(ctx context.Context, opts ...grpc.CallOption) (CategoryService_CreateProtoCategoryBidirectionalStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &CategoryService_ServiceDesc.Streams[3], "/pb.CategoryService/CreateProtoCategoryBidirectionalStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &categoryServiceCreateProtoCategoryBidirectionalStreamClient{stream}
	return x, nil
}

type CategoryService_CreateProtoCategoryBidirectionalStreamClient interface {
	Send(*CategoryRequest) error
	Recv() (*CategoryResponse, error)
	grpc.ClientStream
}

type categoryServiceCreateProtoCategoryBidirectionalStreamClient struct {
	grpc.ClientStream
}

func (x *categoryServiceCreateProtoCategoryBidirectionalStreamClient) Send(m *CategoryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *categoryServiceCreateProtoCategoryBidirectionalStreamClient) Recv() (*CategoryResponse, error) {
	m := new(CategoryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CategoryServiceServer is the server API for CategoryService service.
// All implementations must embed UnimplementedCategoryServiceServer
// for forward compatibility
type CategoryServiceServer interface {
	GetAllCategory(context.Context, *Blank) (*CategoryResponse, error)
	CreateProtoCategory(context.Context, *CategoryRequest) (*Category, error)
	DeleteProtoCategory(context.Context, *Category) (*Category, error)
	UpdateProtoCategory(context.Context, *CategoryRequest) (*Category, error)
	// client streaming
	CreateProtoCategoryStream(CategoryService_CreateProtoCategoryStreamServer) error
	DeleteProtoCategoryStream(CategoryService_DeleteProtoCategoryStreamServer) error
	// server streaming
	CreateProtoCategoryServerStream(*CategoryRequest, CategoryService_CreateProtoCategoryServerStreamServer) error
	// bidirectional streaming
	CreateProtoCategoryBidirectionalStream(CategoryService_CreateProtoCategoryBidirectionalStreamServer) error
	mustEmbedUnimplementedCategoryServiceServer()
}

// UnimplementedCategoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCategoryServiceServer struct {
}

func (UnimplementedCategoryServiceServer) GetAllCategory(context.Context, *Blank) (*CategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCategory not implemented")
}
func (UnimplementedCategoryServiceServer) CreateProtoCategory(context.Context, *CategoryRequest) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProtoCategory not implemented")
}
func (UnimplementedCategoryServiceServer) DeleteProtoCategory(context.Context, *Category) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProtoCategory not implemented")
}
func (UnimplementedCategoryServiceServer) UpdateProtoCategory(context.Context, *CategoryRequest) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProtoCategory not implemented")
}
func (UnimplementedCategoryServiceServer) CreateProtoCategoryStream(CategoryService_CreateProtoCategoryStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateProtoCategoryStream not implemented")
}
func (UnimplementedCategoryServiceServer) DeleteProtoCategoryStream(CategoryService_DeleteProtoCategoryStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method DeleteProtoCategoryStream not implemented")
}
func (UnimplementedCategoryServiceServer) CreateProtoCategoryServerStream(*CategoryRequest, CategoryService_CreateProtoCategoryServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateProtoCategoryServerStream not implemented")
}
func (UnimplementedCategoryServiceServer) CreateProtoCategoryBidirectionalStream(CategoryService_CreateProtoCategoryBidirectionalStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateProtoCategoryBidirectionalStream not implemented")
}
func (UnimplementedCategoryServiceServer) mustEmbedUnimplementedCategoryServiceServer() {}

// UnsafeCategoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoryServiceServer will
// result in compilation errors.
type UnsafeCategoryServiceServer interface {
	mustEmbedUnimplementedCategoryServiceServer()
}

func RegisterCategoryServiceServer(s grpc.ServiceRegistrar, srv CategoryServiceServer) {
	s.RegisterService(&CategoryService_ServiceDesc, srv)
}

func _CategoryService_GetAllCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Blank)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).GetAllCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CategoryService/GetAllCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).GetAllCategory(ctx, req.(*Blank))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_CreateProtoCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).CreateProtoCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CategoryService/CreateProtoCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).CreateProtoCategory(ctx, req.(*CategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_DeleteProtoCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Category)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).DeleteProtoCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CategoryService/DeleteProtoCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).DeleteProtoCategory(ctx, req.(*Category))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_UpdateProtoCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).UpdateProtoCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CategoryService/UpdateProtoCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).UpdateProtoCategory(ctx, req.(*CategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_CreateProtoCategoryStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CategoryServiceServer).CreateProtoCategoryStream(&categoryServiceCreateProtoCategoryStreamServer{stream})
}

type CategoryService_CreateProtoCategoryStreamServer interface {
	SendAndClose(*CategoryResponse) error
	Recv() (*CategoryRequest, error)
	grpc.ServerStream
}

type categoryServiceCreateProtoCategoryStreamServer struct {
	grpc.ServerStream
}

func (x *categoryServiceCreateProtoCategoryStreamServer) SendAndClose(m *CategoryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *categoryServiceCreateProtoCategoryStreamServer) Recv() (*CategoryRequest, error) {
	m := new(CategoryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CategoryService_DeleteProtoCategoryStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CategoryServiceServer).DeleteProtoCategoryStream(&categoryServiceDeleteProtoCategoryStreamServer{stream})
}

type CategoryService_DeleteProtoCategoryStreamServer interface {
	SendAndClose(*CategoryResponse) error
	Recv() (*CategoryDeleteRequest, error)
	grpc.ServerStream
}

type categoryServiceDeleteProtoCategoryStreamServer struct {
	grpc.ServerStream
}

func (x *categoryServiceDeleteProtoCategoryStreamServer) SendAndClose(m *CategoryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *categoryServiceDeleteProtoCategoryStreamServer) Recv() (*CategoryDeleteRequest, error) {
	m := new(CategoryDeleteRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CategoryService_CreateProtoCategoryServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CategoryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CategoryServiceServer).CreateProtoCategoryServerStream(m, &categoryServiceCreateProtoCategoryServerStreamServer{stream})
}

type CategoryService_CreateProtoCategoryServerStreamServer interface {
	Send(*CategoryResponse) error
	grpc.ServerStream
}

type categoryServiceCreateProtoCategoryServerStreamServer struct {
	grpc.ServerStream
}

func (x *categoryServiceCreateProtoCategoryServerStreamServer) Send(m *CategoryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CategoryService_CreateProtoCategoryBidirectionalStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CategoryServiceServer).CreateProtoCategoryBidirectionalStream(&categoryServiceCreateProtoCategoryBidirectionalStreamServer{stream})
}

type CategoryService_CreateProtoCategoryBidirectionalStreamServer interface {
	Send(*CategoryResponse) error
	Recv() (*CategoryRequest, error)
	grpc.ServerStream
}

type categoryServiceCreateProtoCategoryBidirectionalStreamServer struct {
	grpc.ServerStream
}

func (x *categoryServiceCreateProtoCategoryBidirectionalStreamServer) Send(m *CategoryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *categoryServiceCreateProtoCategoryBidirectionalStreamServer) Recv() (*CategoryRequest, error) {
	m := new(CategoryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CategoryService_ServiceDesc is the grpc.ServiceDesc for CategoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CategoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CategoryService",
	HandlerType: (*CategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllCategory",
			Handler:    _CategoryService_GetAllCategory_Handler,
		},
		{
			MethodName: "CreateProtoCategory",
			Handler:    _CategoryService_CreateProtoCategory_Handler,
		},
		{
			MethodName: "DeleteProtoCategory",
			Handler:    _CategoryService_DeleteProtoCategory_Handler,
		},
		{
			MethodName: "UpdateProtoCategory",
			Handler:    _CategoryService_UpdateProtoCategory_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateProtoCategoryStream",
			Handler:       _CategoryService_CreateProtoCategoryStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DeleteProtoCategoryStream",
			Handler:       _CategoryService_DeleteProtoCategoryStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "CreateProtoCategoryServerStream",
			Handler:       _CategoryService_CreateProtoCategoryServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CreateProtoCategoryBidirectionalStream",
			Handler:       _CategoryService_CreateProtoCategoryBidirectionalStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "src/proto/category.proto",
}
