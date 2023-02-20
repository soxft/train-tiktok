// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: video.proto

package video

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

// VideoClient is the client API for Video service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoClient interface {
	Publish(ctx context.Context, in *PublishReq, opts ...grpc.CallOption) (*PublishResp, error)
	PublishList(ctx context.Context, in *PublishListReq, opts ...grpc.CallOption) (*PublishListResp, error)
	Feed(ctx context.Context, in *FeedReq, opts ...grpc.CallOption) (*FeedResp, error)
	CommentAction(ctx context.Context, in *CommentActionReq, opts ...grpc.CallOption) (*CommentActionResp, error)
	CommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListResp, error)
	FavoriteAction(ctx context.Context, in *FavoriteActionReq, opts ...grpc.CallOption) (*FavoriteActionResp, error)
	FavoriteList(ctx context.Context, in *FavoriteListReq, opts ...grpc.CallOption) (*FavoriteListResp, error)
	FavoriteCount(ctx context.Context, in *FavoriteCountReq, opts ...grpc.CallOption) (*FavoriteCountResp, error)
	CommentCount(ctx context.Context, in *CommentCountReq, opts ...grpc.CallOption) (*CommentCountResp, error)
	IsFavorite(ctx context.Context, in *IsFavoriteReq, opts ...grpc.CallOption) (*IsFavoriteResp, error)
	WorkCount(ctx context.Context, in *WorkCountReq, opts ...grpc.CallOption) (*WorkCountResp, error)
	FavoritedCount(ctx context.Context, in *FavoritedCountReq, opts ...grpc.CallOption) (*FavoritedCountResp, error)
	UserFavoriteCount(ctx context.Context, in *UserFavoriteCountReq, opts ...grpc.CallOption) (*UserFavoriteCountResp, error)
}

type videoClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoClient(cc grpc.ClientConnInterface) VideoClient {
	return &videoClient{cc}
}

func (c *videoClient) Publish(ctx context.Context, in *PublishReq, opts ...grpc.CallOption) (*PublishResp, error) {
	out := new(PublishResp)
	err := c.cc.Invoke(ctx, "/video.video/publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) PublishList(ctx context.Context, in *PublishListReq, opts ...grpc.CallOption) (*PublishListResp, error) {
	out := new(PublishListResp)
	err := c.cc.Invoke(ctx, "/video.video/publishList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) Feed(ctx context.Context, in *FeedReq, opts ...grpc.CallOption) (*FeedResp, error) {
	out := new(FeedResp)
	err := c.cc.Invoke(ctx, "/video.video/feed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) CommentAction(ctx context.Context, in *CommentActionReq, opts ...grpc.CallOption) (*CommentActionResp, error) {
	out := new(CommentActionResp)
	err := c.cc.Invoke(ctx, "/video.video/commentAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) CommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListResp, error) {
	out := new(CommentListResp)
	err := c.cc.Invoke(ctx, "/video.video/commentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) FavoriteAction(ctx context.Context, in *FavoriteActionReq, opts ...grpc.CallOption) (*FavoriteActionResp, error) {
	out := new(FavoriteActionResp)
	err := c.cc.Invoke(ctx, "/video.video/favoriteAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) FavoriteList(ctx context.Context, in *FavoriteListReq, opts ...grpc.CallOption) (*FavoriteListResp, error) {
	out := new(FavoriteListResp)
	err := c.cc.Invoke(ctx, "/video.video/favoriteList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) FavoriteCount(ctx context.Context, in *FavoriteCountReq, opts ...grpc.CallOption) (*FavoriteCountResp, error) {
	out := new(FavoriteCountResp)
	err := c.cc.Invoke(ctx, "/video.video/favoriteCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) CommentCount(ctx context.Context, in *CommentCountReq, opts ...grpc.CallOption) (*CommentCountResp, error) {
	out := new(CommentCountResp)
	err := c.cc.Invoke(ctx, "/video.video/commentCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) IsFavorite(ctx context.Context, in *IsFavoriteReq, opts ...grpc.CallOption) (*IsFavoriteResp, error) {
	out := new(IsFavoriteResp)
	err := c.cc.Invoke(ctx, "/video.video/isFavorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) WorkCount(ctx context.Context, in *WorkCountReq, opts ...grpc.CallOption) (*WorkCountResp, error) {
	out := new(WorkCountResp)
	err := c.cc.Invoke(ctx, "/video.video/workCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) FavoritedCount(ctx context.Context, in *FavoritedCountReq, opts ...grpc.CallOption) (*FavoritedCountResp, error) {
	out := new(FavoritedCountResp)
	err := c.cc.Invoke(ctx, "/video.video/favoritedCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) UserFavoriteCount(ctx context.Context, in *UserFavoriteCountReq, opts ...grpc.CallOption) (*UserFavoriteCountResp, error) {
	out := new(UserFavoriteCountResp)
	err := c.cc.Invoke(ctx, "/video.video/userFavoriteCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoServer is the server API for Video service.
// All implementations must embed UnimplementedVideoServer
// for forward compatibility
type VideoServer interface {
	Publish(context.Context, *PublishReq) (*PublishResp, error)
	PublishList(context.Context, *PublishListReq) (*PublishListResp, error)
	Feed(context.Context, *FeedReq) (*FeedResp, error)
	CommentAction(context.Context, *CommentActionReq) (*CommentActionResp, error)
	CommentList(context.Context, *CommentListReq) (*CommentListResp, error)
	FavoriteAction(context.Context, *FavoriteActionReq) (*FavoriteActionResp, error)
	FavoriteList(context.Context, *FavoriteListReq) (*FavoriteListResp, error)
	FavoriteCount(context.Context, *FavoriteCountReq) (*FavoriteCountResp, error)
	CommentCount(context.Context, *CommentCountReq) (*CommentCountResp, error)
	IsFavorite(context.Context, *IsFavoriteReq) (*IsFavoriteResp, error)
	WorkCount(context.Context, *WorkCountReq) (*WorkCountResp, error)
	FavoritedCount(context.Context, *FavoritedCountReq) (*FavoritedCountResp, error)
	UserFavoriteCount(context.Context, *UserFavoriteCountReq) (*UserFavoriteCountResp, error)
	mustEmbedUnimplementedVideoServer()
}

// UnimplementedVideoServer must be embedded to have forward compatible implementations.
type UnimplementedVideoServer struct {
}

func (UnimplementedVideoServer) Publish(context.Context, *PublishReq) (*PublishResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedVideoServer) PublishList(context.Context, *PublishListReq) (*PublishListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishList not implemented")
}
func (UnimplementedVideoServer) Feed(context.Context, *FeedReq) (*FeedResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Feed not implemented")
}
func (UnimplementedVideoServer) CommentAction(context.Context, *CommentActionReq) (*CommentActionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentAction not implemented")
}
func (UnimplementedVideoServer) CommentList(context.Context, *CommentListReq) (*CommentListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentList not implemented")
}
func (UnimplementedVideoServer) FavoriteAction(context.Context, *FavoriteActionReq) (*FavoriteActionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteAction not implemented")
}
func (UnimplementedVideoServer) FavoriteList(context.Context, *FavoriteListReq) (*FavoriteListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteList not implemented")
}
func (UnimplementedVideoServer) FavoriteCount(context.Context, *FavoriteCountReq) (*FavoriteCountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteCount not implemented")
}
func (UnimplementedVideoServer) CommentCount(context.Context, *CommentCountReq) (*CommentCountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentCount not implemented")
}
func (UnimplementedVideoServer) IsFavorite(context.Context, *IsFavoriteReq) (*IsFavoriteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsFavorite not implemented")
}
func (UnimplementedVideoServer) WorkCount(context.Context, *WorkCountReq) (*WorkCountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WorkCount not implemented")
}
func (UnimplementedVideoServer) FavoritedCount(context.Context, *FavoritedCountReq) (*FavoritedCountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoritedCount not implemented")
}
func (UnimplementedVideoServer) UserFavoriteCount(context.Context, *UserFavoriteCountReq) (*UserFavoriteCountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserFavoriteCount not implemented")
}
func (UnimplementedVideoServer) mustEmbedUnimplementedVideoServer() {}

// UnsafeVideoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoServer will
// result in compilation errors.
type UnsafeVideoServer interface {
	mustEmbedUnimplementedVideoServer()
}

func RegisterVideoServer(s grpc.ServiceRegistrar, srv VideoServer) {
	s.RegisterService(&Video_ServiceDesc, srv)
}

func _Video_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video.video/publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).Publish(ctx, req.(*PublishReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_PublishList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).PublishList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video.video/publishList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).PublishList(ctx, req.(*PublishListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_Feed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).Feed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video.video/feed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).Feed(ctx, req.(*FeedReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_CommentAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentActionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).CommentAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video.video/commentAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).CommentAction(ctx, req.(*CommentActionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_CommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).CommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video.video/commentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).CommentList(ctx, req.(*CommentListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_FavoriteAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteActionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).FavoriteAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video.video/favoriteAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).FavoriteAction(ctx, req.(*FavoriteActionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_FavoriteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).FavoriteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video.video/favoriteList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).FavoriteList(ctx, req.(*FavoriteListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_FavoriteCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).FavoriteCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video.video/favoriteCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).FavoriteCount(ctx, req.(*FavoriteCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_CommentCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).CommentCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video.video/commentCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).CommentCount(ctx, req.(*CommentCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_IsFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsFavoriteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).IsFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video.video/isFavorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).IsFavorite(ctx, req.(*IsFavoriteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_WorkCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).WorkCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video.video/workCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).WorkCount(ctx, req.(*WorkCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_FavoritedCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoritedCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).FavoritedCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video.video/favoritedCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).FavoritedCount(ctx, req.(*FavoritedCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_UserFavoriteCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFavoriteCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).UserFavoriteCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video.video/userFavoriteCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).UserFavoriteCount(ctx, req.(*UserFavoriteCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Video_ServiceDesc is the grpc.ServiceDesc for Video service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Video_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "video.video",
	HandlerType: (*VideoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "publish",
			Handler:    _Video_Publish_Handler,
		},
		{
			MethodName: "publishList",
			Handler:    _Video_PublishList_Handler,
		},
		{
			MethodName: "feed",
			Handler:    _Video_Feed_Handler,
		},
		{
			MethodName: "commentAction",
			Handler:    _Video_CommentAction_Handler,
		},
		{
			MethodName: "commentList",
			Handler:    _Video_CommentList_Handler,
		},
		{
			MethodName: "favoriteAction",
			Handler:    _Video_FavoriteAction_Handler,
		},
		{
			MethodName: "favoriteList",
			Handler:    _Video_FavoriteList_Handler,
		},
		{
			MethodName: "favoriteCount",
			Handler:    _Video_FavoriteCount_Handler,
		},
		{
			MethodName: "commentCount",
			Handler:    _Video_CommentCount_Handler,
		},
		{
			MethodName: "isFavorite",
			Handler:    _Video_IsFavorite_Handler,
		},
		{
			MethodName: "workCount",
			Handler:    _Video_WorkCount_Handler,
		},
		{
			MethodName: "favoritedCount",
			Handler:    _Video_FavoritedCount_Handler,
		},
		{
			MethodName: "userFavoriteCount",
			Handler:    _Video_UserFavoriteCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "video.proto",
}
