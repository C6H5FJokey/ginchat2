// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package state

import (
	"context"

	"ginchat2/internal/svc"
	"ginchat2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStateLogic {
	return &CreateStateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateStateLogic) CreateState(req *types.CreateStateReq) error {
	// todo: add your logic here and delete this line

	return nil
}
