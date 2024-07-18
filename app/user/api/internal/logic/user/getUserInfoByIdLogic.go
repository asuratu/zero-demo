package user

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewGetUserInfoByIdLogic 根据id获取用户详情
func NewGetUserInfoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByIdLogic {
	return &GetUserInfoByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoByIdLogic) GetUserInfoById(req *types.UserInfoByIdReq) (resp *types.SimpleUserInfoReply, err error) {
	logx.Info("GetUserInfoByIdLogic")
	resp = &types.SimpleUserInfoReply{
		Id:   req.Id,
		Name: "test",
	}
	return
}
