// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package state

import (
	"context"

	"ginchat2/internal/svc"
	"ginchat2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserStateByGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserStateByGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserStateByGroupLogic {
	return &GetUserStateByGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserStateByGroupLogic) GetUserStateByGroup(req *types.GetUserStateByGroupReq) (resp []types.GetStateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
