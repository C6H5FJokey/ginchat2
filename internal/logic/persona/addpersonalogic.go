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

type AddPersonaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPersonaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPersonaLogic {
	return &AddPersonaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPersonaLogic) AddPersona(req *types.AddPersonaReq) error {
	userId, err := strconv.ParseInt(l.ctx.Value("userId").(string), 10, 64)
	if err == models.ErrNotFound {
		return errors.New(5002, "不存在该用户")
	}
	if _, err := l.svcCtx.PersonaModel.FindOneByUserAndName(l.ctx, userId, req.Name); err == nil {
		return errors.New(2001, "存在同名persona")
	} else if err != models.ErrNotFound {
		return errors.New(2001, err.Error())
	}
	if _, err := l.svcCtx.PersonaModel.Insert(l.ctx, &models.Persona{
		UserId: userId,
		Name:   req.Name,
		Prompt: req.Prompt,
	}); err != nil {
		return errors.New(2001, err.Error())
	}
	return nil
}
