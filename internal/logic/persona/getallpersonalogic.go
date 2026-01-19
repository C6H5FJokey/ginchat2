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

type GetAllPersonaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllPersonaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllPersonaLogic {
	return &GetAllPersonaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllPersonaLogic) GetAllPersona() (resp []types.GetPersonaResp, err error) {
	userId, err := strconv.ParseInt(l.ctx.Value("userId").(string), 10, 64)
	if err == models.ErrNotFound {
		return nil, errors.New(5002, "不存在该用户")
	}
	personaList, err := l.svcCtx.PersonaModel.FindAllByUser(l.ctx, userId)
	if err != nil {
		return nil, err
	}
	resp = make([]types.GetPersonaResp, 0, len(personaList))
	for _, persona := range personaList {
		resp = append(resp, types.GetPersonaResp{
			Id:     persona.Id,
			Name:   persona.Name,
			Prompt: persona.Prompt,
		})
	}
	return
}
