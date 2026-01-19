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

type EditPersonaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditPersonaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditPersonaLogic {
	return &EditPersonaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditPersonaLogic) EditPersona(req *types.EditPersonaReq) error {
	userId, err := strconv.ParseInt(l.ctx.Value("userId").(string), 10, 64)
	if err == models.ErrNotFound {
		return errors.New(5002, "不存在该用户")
	}
	persona, err := l.svcCtx.PersonaModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return err
	}
	if persona.UserId != userId {
		return errors.New(2002, "persona与用户不匹配!")
	}
	if err := l.svcCtx.PersonaModel.Update(l.ctx, &models.Persona{
		Id:     req.Id,
		UserId: userId,
		Name:   req.Name,
		Prompt: req.Prompt,
	}); err != nil {
		logx.Info(err)
		return errors.New(2002, "存在同名persona!")
	}
	return nil
}
