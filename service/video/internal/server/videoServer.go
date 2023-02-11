// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package server

import (
	"context"

	"train-tiktok/service/video/internal/logic"
	"train-tiktok/service/video/internal/svc"
	"train-tiktok/service/video/types/video"
)

type VideoServer struct {
	svcCtx *svc.ServiceContext
	video.UnimplementedVideoServer
}

func NewVideoServer(svcCtx *svc.ServiceContext) *VideoServer {
	return &VideoServer{
		svcCtx: svcCtx,
	}
}

func (s *VideoServer) Publish(ctx context.Context, in *video.PublishReq) (*video.PublishResp, error) {
	l := logic.NewPublishLogic(ctx, s.svcCtx)
	return l.Publish(in)
}

func (s *VideoServer) Feed(ctx context.Context, in *video.FeedReq) (*video.FeedResp, error) {
	l := logic.NewFeedLogic(ctx, s.svcCtx)
	return l.Feed(in)
}

func (s *VideoServer) CommentAction(ctx context.Context, in *video.CommentActionReq) (*video.CommentActionResp, error) {
	l := logic.NewCommentActionLogic(ctx, s.svcCtx)
	return l.CommentAction(in)
}

func (s *VideoServer) CommentList(ctx context.Context, in *video.CommentListReq) (*video.CommentListResp, error) {
	l := logic.NewCommentListLogic(ctx, s.svcCtx)
	return l.CommentList(in)
}

func (s *VideoServer) FavoriteAction(ctx context.Context, in *video.FavoriteActionReq) (*video.FavoriteActionResp, error) {
	l := logic.NewFavoriteActionLogic(ctx, s.svcCtx)
	return l.FavoriteAction(in)
}

func (s *VideoServer) FavoriteList(ctx context.Context, in *video.FavoriteListReq) (*video.FavoriteListResp, error) {
	l := logic.NewFavoriteListLogic(ctx, s.svcCtx)
	return l.FavoriteList(in)
}

func (s *VideoServer) FavoriteCount(ctx context.Context, in *video.FavoriteCountReq) (*video.FavoriteCountResp, error) {
	l := logic.NewFavoriteCountLogic(ctx, s.svcCtx)
	return l.FavoriteCount(in)
}

func (s *VideoServer) CommentCount(ctx context.Context, in *video.CommentCountReq) (*video.CommentCountResp, error) {
	l := logic.NewCommentCountLogic(ctx, s.svcCtx)
	return l.CommentCount(in)
}
