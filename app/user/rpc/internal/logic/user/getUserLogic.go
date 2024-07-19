package userlogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/model"
	"rpc/pb/user"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.GetUserRequest) (*user.GetUserResponse, error) {
	u, e := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if errors.Is(e, model.ErrNotFound) || u == nil {
		return nil, errors.New("user not found")
	} else if e != nil {
		return nil, errors.Wrap(e, "find user error")
	}

	resp := &user.GetUserResponse{}
	_ = copier.Copy(&resp, u)

	return resp, nil
}
