// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package persona

import (
	"context"
	"strconv"

	"ginchat2/internal/svc"
	"ginchat2/internal/types"
	"ginchat2/models"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
)

type GetPersonaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPersonaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPersonaLogic {
	return &GetPersonaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPersonaLogic) GetPersona(req *types.GetPersonaReq) (resp *types.GetPersonaResp, err error) {
	userId, err := strconv.ParseInt(l.ctx.Value("userId").(string), 10, 64)
	if err == models.ErrNotFound {
		return nil, errors.New(5002, "不存在该用户")
	}
	persona, err := l.svcCtx.PersonaModel.FindOneByUserAndName(l.ctx, userId, req.Name)
	resp = &types.GetPersonaResp{
		Id:     persona.Id,
		Name:   persona.Name,
		Prompt: persona.Prompt,
	}
	return
}
