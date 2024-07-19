// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package user

import (
	"context"

	"rpc/pb/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateUserRequest  = user.CreateUserRequest
	CreateUserResponse = user.CreateUserResponse
	GetUserRequest     = user.GetUserRequest
	GetUserResponse    = user.GetUserResponse
	Request            = user.Request
	Response           = user.Response
	UpdateUserRequest  = user.UpdateUserRequest
	UpdateUserResponse = user.UpdateUserResponse

	User interface {
		GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
		CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
		UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
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

func (m *defaultUser) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}

func (m *defaultUser) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.CreateUser(ctx, in, opts...)
}

func (m *defaultUser) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UpdateUser(ctx, in, opts...)
}