// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: shop/pkg/shopproto/shop.proto

package shopproto

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

// ShopServiceClient is the client API for ShopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShopServiceClient interface {
	CreateShop(ctx context.Context, in *CreateShopRequest, opts ...grpc.CallOption) (*ShopResponse, error)
	GetAllShops(ctx context.Context, in *GetAllShopsRequest, opts ...grpc.CallOption) (*AllShopsResponse, error)
	GetShop(ctx context.Context, in *GetShopRequest, opts ...grpc.CallOption) (*ShopResponse, error)
	EditShop(ctx context.Context, in *EditShopRequest, opts ...grpc.CallOption) (*ShopResponse, error)
	DeleteShop(ctx context.Context, in *DeleteShopRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	CreateShopItem(ctx context.Context, in *CreateShopItemRequest, opts ...grpc.CallOption) (*ShopItemResponse, error)
	GetAllShopItems(ctx context.Context, in *GetAllShopItemsRequest, opts ...grpc.CallOption) (*AllShopItemsResponse, error)
	GetShopItem(ctx context.Context, in *GetShopItemRequest, opts ...grpc.CallOption) (*ShopItemResponse, error)
	EditShopItem(ctx context.Context, in *EditShopItemRequest, opts ...grpc.CallOption) (*ShopItemResponse, error)
	DeleteShopItem(ctx context.Context, in *DeleteShopItemRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	AddBankAccount(ctx context.Context, in *AddBankAccountRequest, opts ...grpc.CallOption) (*AddBankAccountResponse, error)
}

type shopServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShopServiceClient(cc grpc.ClientConnInterface) ShopServiceClient {
	return &shopServiceClient{cc}
}

func (c *shopServiceClient) CreateShop(ctx context.Context, in *CreateShopRequest, opts ...grpc.CallOption) (*ShopResponse, error) {
	out := new(ShopResponse)
	err := c.cc.Invoke(ctx, "/shop.ShopService/CreateShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) GetAllShops(ctx context.Context, in *GetAllShopsRequest, opts ...grpc.CallOption) (*AllShopsResponse, error) {
	out := new(AllShopsResponse)
	err := c.cc.Invoke(ctx, "/shop.ShopService/GetAllShops", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) GetShop(ctx context.Context, in *GetShopRequest, opts ...grpc.CallOption) (*ShopResponse, error) {
	out := new(ShopResponse)
	err := c.cc.Invoke(ctx, "/shop.ShopService/GetShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) EditShop(ctx context.Context, in *EditShopRequest, opts ...grpc.CallOption) (*ShopResponse, error) {
	out := new(ShopResponse)
	err := c.cc.Invoke(ctx, "/shop.ShopService/EditShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) DeleteShop(ctx context.Context, in *DeleteShopRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/shop.ShopService/DeleteShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) CreateShopItem(ctx context.Context, in *CreateShopItemRequest, opts ...grpc.CallOption) (*ShopItemResponse, error) {
	out := new(ShopItemResponse)
	err := c.cc.Invoke(ctx, "/shop.ShopService/CreateShopItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) GetAllShopItems(ctx context.Context, in *GetAllShopItemsRequest, opts ...grpc.CallOption) (*AllShopItemsResponse, error) {
	out := new(AllShopItemsResponse)
	err := c.cc.Invoke(ctx, "/shop.ShopService/GetAllShopItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) GetShopItem(ctx context.Context, in *GetShopItemRequest, opts ...grpc.CallOption) (*ShopItemResponse, error) {
	out := new(ShopItemResponse)
	err := c.cc.Invoke(ctx, "/shop.ShopService/GetShopItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) EditShopItem(ctx context.Context, in *EditShopItemRequest, opts ...grpc.CallOption) (*ShopItemResponse, error) {
	out := new(ShopItemResponse)
	err := c.cc.Invoke(ctx, "/shop.ShopService/EditShopItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) DeleteShopItem(ctx context.Context, in *DeleteShopItemRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/shop.ShopService/DeleteShopItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) AddBankAccount(ctx context.Context, in *AddBankAccountRequest, opts ...grpc.CallOption) (*AddBankAccountResponse, error) {
	out := new(AddBankAccountResponse)
	err := c.cc.Invoke(ctx, "/shop.ShopService/AddBankAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopServiceServer is the server API for ShopService service.
// All implementations must embed UnimplementedShopServiceServer
// for forward compatibility
type ShopServiceServer interface {
	CreateShop(context.Context, *CreateShopRequest) (*ShopResponse, error)
	GetAllShops(context.Context, *GetAllShopsRequest) (*AllShopsResponse, error)
	GetShop(context.Context, *GetShopRequest) (*ShopResponse, error)
	EditShop(context.Context, *EditShopRequest) (*ShopResponse, error)
	DeleteShop(context.Context, *DeleteShopRequest) (*DeleteResponse, error)
	CreateShopItem(context.Context, *CreateShopItemRequest) (*ShopItemResponse, error)
	GetAllShopItems(context.Context, *GetAllShopItemsRequest) (*AllShopItemsResponse, error)
	GetShopItem(context.Context, *GetShopItemRequest) (*ShopItemResponse, error)
	EditShopItem(context.Context, *EditShopItemRequest) (*ShopItemResponse, error)
	DeleteShopItem(context.Context, *DeleteShopItemRequest) (*DeleteResponse, error)
	AddBankAccount(context.Context, *AddBankAccountRequest) (*AddBankAccountResponse, error)
	mustEmbedUnimplementedShopServiceServer()
}

// UnimplementedShopServiceServer must be embedded to have forward compatible implementations.
type UnimplementedShopServiceServer struct {
}

func (UnimplementedShopServiceServer) CreateShop(context.Context, *CreateShopRequest) (*ShopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShop not implemented")
}
func (UnimplementedShopServiceServer) GetAllShops(context.Context, *GetAllShopsRequest) (*AllShopsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllShops not implemented")
}
func (UnimplementedShopServiceServer) GetShop(context.Context, *GetShopRequest) (*ShopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShop not implemented")
}
func (UnimplementedShopServiceServer) EditShop(context.Context, *EditShopRequest) (*ShopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditShop not implemented")
}
func (UnimplementedShopServiceServer) DeleteShop(context.Context, *DeleteShopRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteShop not implemented")
}
func (UnimplementedShopServiceServer) CreateShopItem(context.Context, *CreateShopItemRequest) (*ShopItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShopItem not implemented")
}
func (UnimplementedShopServiceServer) GetAllShopItems(context.Context, *GetAllShopItemsRequest) (*AllShopItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllShopItems not implemented")
}
func (UnimplementedShopServiceServer) GetShopItem(context.Context, *GetShopItemRequest) (*ShopItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShopItem not implemented")
}
func (UnimplementedShopServiceServer) EditShopItem(context.Context, *EditShopItemRequest) (*ShopItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditShopItem not implemented")
}
func (UnimplementedShopServiceServer) DeleteShopItem(context.Context, *DeleteShopItemRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteShopItem not implemented")
}
func (UnimplementedShopServiceServer) AddBankAccount(context.Context, *AddBankAccountRequest) (*AddBankAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBankAccount not implemented")
}
func (UnimplementedShopServiceServer) mustEmbedUnimplementedShopServiceServer() {}

// UnsafeShopServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShopServiceServer will
// result in compilation errors.
type UnsafeShopServiceServer interface {
	mustEmbedUnimplementedShopServiceServer()
}

func RegisterShopServiceServer(s grpc.ServiceRegistrar, srv ShopServiceServer) {
	s.RegisterService(&ShopService_ServiceDesc, srv)
}

func _ShopService_CreateShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).CreateShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.ShopService/CreateShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).CreateShop(ctx, req.(*CreateShopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_GetAllShops_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllShopsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetAllShops(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.ShopService/GetAllShops",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetAllShops(ctx, req.(*GetAllShopsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_GetShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.ShopService/GetShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetShop(ctx, req.(*GetShopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_EditShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditShopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).EditShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.ShopService/EditShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).EditShop(ctx, req.(*EditShopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_DeleteShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteShopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).DeleteShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.ShopService/DeleteShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).DeleteShop(ctx, req.(*DeleteShopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_CreateShopItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShopItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).CreateShopItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.ShopService/CreateShopItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).CreateShopItem(ctx, req.(*CreateShopItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_GetAllShopItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllShopItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetAllShopItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.ShopService/GetAllShopItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetAllShopItems(ctx, req.(*GetAllShopItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_GetShopItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShopItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetShopItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.ShopService/GetShopItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetShopItem(ctx, req.(*GetShopItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_EditShopItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditShopItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).EditShopItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.ShopService/EditShopItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).EditShopItem(ctx, req.(*EditShopItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_DeleteShopItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteShopItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).DeleteShopItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.ShopService/DeleteShopItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).DeleteShopItem(ctx, req.(*DeleteShopItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_AddBankAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBankAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).AddBankAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.ShopService/AddBankAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).AddBankAccount(ctx, req.(*AddBankAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShopService_ServiceDesc is the grpc.ServiceDesc for ShopService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShopService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shop.ShopService",
	HandlerType: (*ShopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateShop",
			Handler:    _ShopService_CreateShop_Handler,
		},
		{
			MethodName: "GetAllShops",
			Handler:    _ShopService_GetAllShops_Handler,
		},
		{
			MethodName: "GetShop",
			Handler:    _ShopService_GetShop_Handler,
		},
		{
			MethodName: "EditShop",
			Handler:    _ShopService_EditShop_Handler,
		},
		{
			MethodName: "DeleteShop",
			Handler:    _ShopService_DeleteShop_Handler,
		},
		{
			MethodName: "CreateShopItem",
			Handler:    _ShopService_CreateShopItem_Handler,
		},
		{
			MethodName: "GetAllShopItems",
			Handler:    _ShopService_GetAllShopItems_Handler,
		},
		{
			MethodName: "GetShopItem",
			Handler:    _ShopService_GetShopItem_Handler,
		},
		{
			MethodName: "EditShopItem",
			Handler:    _ShopService_EditShopItem_Handler,
		},
		{
			MethodName: "DeleteShopItem",
			Handler:    _ShopService_DeleteShopItem_Handler,
		},
		{
			MethodName: "AddBankAccount",
			Handler:    _ShopService_AddBankAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop/pkg/shopproto/shop.proto",
}
