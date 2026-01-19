// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package state

import (
	"context"

	"ginchat2/internal/svc"
	"ginchat2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllStateByGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllStateByGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllStateByGroupLogic {
	return &GetAllStateByGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllStateByGroupLogic) GetAllStateByGroup(req *types.GetStateByGroupReq) (resp []types.GetUserStateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
