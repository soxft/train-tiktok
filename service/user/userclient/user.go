// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"

	"train-tiktok/service/user/types/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	FollowListReq    = user.FollowListReq
	FollowListResp   = user.FollowListResp
	FollowerListReq  = user.FollowerListReq
	FollowerListResp = user.FollowerListResp
	FriendListReq    = user.FriendListReq
	FriendListResp   = user.FriendListResp
	IsFriendReq      = user.IsFriendReq
	IsFriendResp     = user.IsFriendResp
	RelationActReq   = user.RelationActReq
	RelationActResp  = user.RelationActResp
	Resp             = user.Resp
	UserReq          = user.UserReq
	UserResp         = user.UserResp

	User interface {
		User(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserResp, error)
		RelationAct(ctx context.Context, in *RelationActReq, opts ...grpc.CallOption) (*RelationActResp, error)
		FollowList(ctx context.Context, in *FollowListReq, opts ...grpc.CallOption) (*FollowListResp, error)
		FollowerList(ctx context.Context, in *FollowerListReq, opts ...grpc.CallOption) (*FollowerListResp, error)
		FriendList(ctx context.Context, in *FriendListReq, opts ...grpc.CallOption) (*FriendListResp, error)
		IsFriend(ctx context.Context, in *IsFriendReq, opts ...grpc.CallOption) (*IsFriendResp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) User(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.User(ctx, in, opts...)
}

func (m *defaultUser) RelationAct(ctx context.Context, in *RelationActReq, opts ...grpc.CallOption) (*RelationActResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.RelationAct(ctx, in, opts...)
}

func (m *defaultUser) FollowList(ctx context.Context, in *FollowListReq, opts ...grpc.CallOption) (*FollowListResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.FollowList(ctx, in, opts...)
}

func (m *defaultUser) FollowerList(ctx context.Context, in *FollowerListReq, opts ...grpc.CallOption) (*FollowerListResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.FollowerList(ctx, in, opts...)
}

func (m *defaultUser) FriendList(ctx context.Context, in *FriendListReq, opts ...grpc.CallOption) (*FriendListResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.FriendList(ctx, in, opts...)
}

func (m *defaultUser) IsFriend(ctx context.Context, in *IsFriendReq, opts ...grpc.CallOption) (*IsFriendResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.IsFriend(ctx, in, opts...)
}
