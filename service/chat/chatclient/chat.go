// Code generated by goctl. DO NOT EDIT.
// Source: chat.proto

package chatclient

import (
	"context"

	"train-tiktok/service/chat/types/chat"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CharActionResp      = chat.CharActionResp
	ChatActionReq       = chat.ChatActionReq
	ChatLastMessageReq  = chat.ChatLastMessageReq
	ChatLastMessageResp = chat.ChatLastMessageResp
	ChatMessageReq      = chat.ChatMessageReq
	ChatMessageResp     = chat.ChatMessageResp
	Message             = chat.Message
	Resp                = chat.Resp

	Chat interface {
		ChatAction(ctx context.Context, in *ChatActionReq, opts ...grpc.CallOption) (*CharActionResp, error)
		ChatMessage(ctx context.Context, in *ChatMessageReq, opts ...grpc.CallOption) (*ChatMessageResp, error)
		ChatLastMessage(ctx context.Context, in *ChatLastMessageReq, opts ...grpc.CallOption) (*ChatLastMessageResp, error)
	}

	defaultChat struct {
		cli zrpc.Client
	}
)

func NewChat(cli zrpc.Client) Chat {
	return &defaultChat{
		cli: cli,
	}
}

func (m *defaultChat) ChatAction(ctx context.Context, in *ChatActionReq, opts ...grpc.CallOption) (*CharActionResp, error) {
	client := chat.NewChatClient(m.cli.Conn())
	return client.ChatAction(ctx, in, opts...)
}

func (m *defaultChat) ChatMessage(ctx context.Context, in *ChatMessageReq, opts ...grpc.CallOption) (*ChatMessageResp, error) {
	client := chat.NewChatClient(m.cli.Conn())
	return client.ChatMessage(ctx, in, opts...)
}

func (m *defaultChat) ChatLastMessage(ctx context.Context, in *ChatLastMessageReq, opts ...grpc.CallOption) (*ChatLastMessageResp, error) {
	client := chat.NewChatClient(m.cli.Conn())
	return client.ChatLastMessage(ctx, in, opts...)
}
